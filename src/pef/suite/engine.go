package suite

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	"pef/client"
	"pef/datagen"
	peflag "pef/flag"

	"github.com/feiyuw/boomer"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var (
	signBufPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	fieldsBuffer = make(chan map[string]interface{}, 5000)
	hatchEvt     = make(chan struct{}, 1)
	stopEvt      = make(chan struct{}, 1)
)

func init() {
	SuiteMap.Add("engine", NewEngineSuite())
}

func clearFieldsBuffer() {
	for {
		select {
		case <-fieldsBuffer:
		default:
			return
		}
	}
}

// EngineSuite object of engine pet tester
type EngineSuite struct {
	URLWithoutSign string

	Host       string
	AppKey     string
	AppSecret  string
	AppCode    string
	EventCodeList  string
	EventCode string
	DataSource string
	RunOnce    bool
	FieldsFn   map[string]func() interface{}
}

// NewEngineSuite generate a new EngineSuite object
func NewEngineSuite() *EngineSuite {
	return &EngineSuite{FieldsFn: map[string]func() interface{}{}}
}

// Init arguments parsing and initialization
func (e *EngineSuite) Init(flagSet *flag.FlagSet, args []string) error {
	fieldsSet := peflag.ListFlags{}

	flagSet.StringVar(&e.Host, "host", "", "engine host and port, eg. 127.0.0.1:7090")
	flagSet.StringVar(&e.AppKey, "app-key", "", "key of app")
	flagSet.StringVar(&e.AppSecret, "app-secret", "", "secret of app")
	flagSet.StringVar(&e.AppCode, "app-code", "", "app code, used by credit")
	flagSet.StringVar(&e.EventCodeList, "event-codes", "", "event codes to test")
	flagSet.StringVar(&e.DataSource, "data-source", "random", "data source of test data, file path|random")
	flagSet.BoolVar(&e.RunOnce, "run-once", false, "only for file data source, if set, will not rotate at the end of file")
	flagSet.Var(&fieldsSet, "field", "random post field name, include: ip, email, phone_number, const_id, or dynamic field like: ext_amount:float_2000, ext_uid:uuid, ext_name:string_64")
	if err := flagSet.Parse(args); err != nil {
		return errors.New("command line args parsing error")
	}

	if e.Host == "" || e.AppKey == "" || e.AppSecret == "" || e.EventCodeList == "" {
		return errors.New("host, app-key, app-secret, event-codes should be set")
	}
	if e.DataSource == "random" {
		if len(fieldsSet) == 0 {
			return errors.New("on random data source, field should be set")
		}
		for _, field := range fieldsSet {
			name, handler, err := datagen.GetGenerator(field)
			if err != nil {
				return err
			}
			e.FieldsFn[name] = handler
		}
	} else {
		if _, err := os.Stat(e.DataSource); os.IsNotExist(err) {
			return errors.New("datasource does not exist")
		}
		go e.readFromDataSource()
	}
	e.URLWithoutSign = fmt.Sprintf("http://%s/ctu/event.do?appKey=%s&version=1&sign=", e.Host, e.AppKey)

	return nil
}

func (e *EngineSuite) readFromDataSource() {
	file, err := os.Open(e.DataSource)
	if err != nil {
		panic("open datasource error")
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic("close datasource error")
		}
	}()

	for {
		stopped := false
		if e.RunOnce {
			<-hatchEvt
			log.Println("start to read from data source")
		}
		scanner := bufio.NewScanner(file)
	LOOP:
		for scanner.Scan() {
			fields := make(map[string]interface{})
			line := scanner.Bytes()
			if err := json.Unmarshal(line, &fields); err != nil {
				panic("invalid data line")
			}
			select {
			case <-stopEvt:
				log.Println("got stop event")
				stopped = true
				break LOOP
			default:
				fieldsBuffer <- fields
			}
		}
		if e.RunOnce {
			if !stopped {
				<-stopEvt
			}
		}
		log.Println("read to the end of the data source")
		if _, err := file.Seek(0, 0); err != nil {
			panic("seek to the beginning of file error")
		}
	}
}

// GetTask return a boomer task of engine
func (e *EngineSuite) GetTask() *boomer.Task {
	return &boomer.Task{
		Name: "engine",
		OnStart: func() {
			if e.DataSource != "random" && e.RunOnce {
				hatchEvt <- struct{}{}
			}
		},
		OnStop: func() {
			if e.DataSource != "random" && e.RunOnce {
				stopEvt <- struct{}{}
				clearFieldsBuffer()
			}
		},
		Fn: func() {
			codes := strings.Fields(strings.ReplaceAll(e.EventCodeList,","," "))
			for _,code := range  codes{
				e.EventCode = code
				fields := e.getFields()
				start := boomer.Now()
				statusCode, err := e.callRiskEngine(fields)

				elapsed := boomer.Now() - start

				if err != nil {
					boomer.RecordFailure("/ctu/event.do", "err", elapsed, err.Error())
					return
				}

				if statusCode == fasthttp.StatusOK {
					boomer.RecordSuccess("/ctu/event.do", "200", elapsed, int64(0))
				} else {
					boomer.RecordFailure("/ctu/event.do", strconv.Itoa(statusCode), elapsed, "")
				}
			}

		},
	}
}

func (e *EngineSuite) callRiskEngine(fields map[string]interface{}) (int, error) {
	data, err := e.getData(fields)
	if err != nil {
		return 0, err
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(e.getURL(fields))
	req.Header.SetMethod("POST")
	req.Header.SetContentType("text/plain")
	req.SetBody(data)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if err := client.HTTPClient.Do(req, resp); err != nil {
		return 0, err
	}

	return resp.StatusCode(), nil
}

func (e *EngineSuite) getFields() map[string]interface{} {
	if e.DataSource == "random" {
		return e.randomFields()
	}

	return <-fieldsBuffer
}

func (e *EngineSuite) randomFields() map[string]interface{} {
	fields := make(map[string]interface{}, len(e.FieldsFn))
	for name, handler := range e.FieldsFn {
		fields[name] = handler()
	}

	return fields
}

func (e *EngineSuite) getURL(fields map[string]interface{}) string {
	return e.URLWithoutSign + e.getSign(fields)
}

func (e *EngineSuite) getSign(fields map[string]interface{}) string {
	data := signBufPool.Get().(*bytes.Buffer)
	data.Reset()
	defer signBufPool.Put(data)
	data.WriteString(e.AppSecret)
	if e.AppCode != "" {
		data.WriteString("appCode")
		data.WriteString(e.AppCode)
	}
	data.WriteString("eventCode")
	data.WriteString(e.EventCode)
	data.WriteString("flag")
	data.WriteString(e.EventCode)

	keys := make([]string, len(fields))
	idx := 0
	for key := range fields {
		keys[idx] = key
		idx++
	}
	sort.Strings(keys)

	for _, key := range keys {
		data.WriteString(key)
		value := fields[key]
		if value != nil && reflect.TypeOf(value).Name() == "string" {
			data.WriteString(value.(string))
		} else {
			valueStr, _ := json.Marshal(value)
			data.Write(valueStr)
		}
	}

	data.WriteString(e.AppSecret)

	sum := md5.Sum(data.Bytes())

	return hex.EncodeToString(sum[:])
}

func (e *EngineSuite) getData(fields map[string]interface{}) ([]byte, error) {
	var data []byte
	var err error
	dataMap := map[string]interface{}{
		"flag":      e.EventCode,
		"data":      fields,
		"eventCode": e.EventCode,
	}
	if e.AppCode != "" {
		dataMap["appCode"] = e.AppCode
	}

	data, err = json.Marshal(dataMap)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)

	return dst, nil
}

