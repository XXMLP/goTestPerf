package suite

import (
	"bufio"
	"errors"
	"flag"
	"github.com/feiyuw/boomer"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"pef/client"
	"strconv"
)

var (
	generalfieldsBuffer = make(chan map[string]interface{}, 5000)
	generalhatchEvt     = make(chan struct{}, 1)
	generalstopEvt      = make(chan struct{}, 1)
)

func init() {
	SuiteMap.Add("general", NewGeneralSuite())
}
type apiResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
type GeneralSuite struct {
	Host       string
	DataSource string
	Api        string
	RunOnce    bool
}

func generalClearFieldsBuffer() {
	for {
		select {
		case <-generalfieldsBuffer:
		default:
			return
		}
	}
}

func NewGeneralSuite() *GeneralSuite {
	return &GeneralSuite{}
}

func (general *GeneralSuite) Init(flagSet *flag.FlagSet, args []string) error {
	flagSet.StringVar(&general.Host, "host", "", "general host and port, eg. 127.0.0.1:8080")
	flagSet.StringVar(&general.Api, "api", "", "general api url, eg. /api/a")
	flagSet.StringVar(&general.DataSource, "data-source", "", "data source of test data, file path")
	flagSet.BoolVar(&general.RunOnce, "run-once", false, "only for file data source, if set, will not rotate at the end of file")
	if err := flagSet.Parse(args); err != nil {
		return errors.New("command line args parsing error")
	}

	if general.Host == "" || general.Api == "" || general.DataSource == ""{
		return errors.New("host, api, data-source should be set")
	}
	if _, err := os.Stat(general.DataSource); os.IsNotExist(err) {
			return errors.New("datasource does not exist")
		}
	go general.readFromDataSource()
	return nil
}

func (general *GeneralSuite) readFromDataSource() {
	file, err := os.Open(general.DataSource)
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
		if general.RunOnce {
			<-generalhatchEvt
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
			case <-generalstopEvt:
				log.Println("got stop event")
				stopped = true
				break LOOP
			default:
				generalfieldsBuffer <- fields
			}
		}
		if general.RunOnce {
			if !stopped {
				<-generalstopEvt
			}
		}
		log.Println("read to the end of the data source")
		if _, err := file.Seek(0, 0); err != nil {
			panic("seek to the beginning of file error")
		}
	}
}
// GetTask return a boomer task of general
func (general *GeneralSuite) GetTask() *boomer.Task {
	return &boomer.Task{
		Name: "engine",
		OnStart: func() {
			if general.RunOnce {
				generalhatchEvt <- struct{}{}
			}
		},
		OnStop: func() {
			if general.RunOnce {
				generalstopEvt <- struct{}{}
				generalClearFieldsBuffer()
			}
		},
		Fn: func() {
				start := boomer.Now()
				statusCode, body, err := general.callGeneralApi(<-generalfieldsBuffer)
				elapsed := boomer.Now() - start
				if err != nil {
					boomer.RecordFailure(general.Api, strconv.Itoa(statusCode), elapsed, err.Error())
					return
				}

				resp := apiResponse{}
				if err := json.Unmarshal(body, &resp); err != nil {
				boomer.RecordFailure(general.Api, "err", elapsed, err.Error())
					return
				}

				if statusCode == fasthttp.StatusOK && resp.Success{
					boomer.RecordSuccess(general.Api, strconv.Itoa(statusCode), elapsed, int64(0))
				} else if statusCode == fasthttp.StatusOK && !resp.Success{
					boomer.RecordFailure(general.Api, "success:"+strconv.FormatBool(resp.Success), elapsed, resp.Msg)
					return
				} else if statusCode != fasthttp.StatusOK{
					boomer.RecordFailure(general.Api, strconv.Itoa(statusCode), elapsed, string(body))
					return
				}
		},
	}
}

func (general *GeneralSuite) callGeneralApi(fields map[string]interface{}) (int, []byte, error) {
	var data []byte
	var jsonErr error
	data, jsonErr = json.Marshal(fields)
	if jsonErr != nil {
		return 0, nil, jsonErr
	}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://" + general.Host + general.Api)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(data)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if err := client.HTTPClient.Do(req, resp); err != nil {
		return 0, nil ,err
	}

	return resp.StatusCode(), resp.Body(), nil
}