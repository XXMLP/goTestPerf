package suite

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"time"

	"pef/client"
	"pef/datagen"
	peflag "pef/flag"

	"github.com/dubbogo/hessian2"
	"github.com/feiyuw/boomer"
	"github.com/feiyuw/dubbo-go/common"
	"github.com/feiyuw/dubbo-go/protocol/dubbo"
	"github.com/satori/go.uuid"
)

const (
	dcURLTpl = "dubbo://%s/cn.securitystack.indicatorcenter.api.client.DataSourceClient?anyhost=true&application=indicator-center&dubbo=2.6.0&generic=false&interface=cn.securitystack.indicatorcenter.api.client.DataSourceClient&logger=slf4j&methods=processMessage&pid=1&revision=1.0.0&side=provider&timestamp=1560575961733&version=1.0.0"
	ocURLTpl = "dubbo://%s/cn.securitystack.indicatorcenter.api.client.ObtainValueClient?anyhost=true&application=indicator-center&dubbo=2.6.0&generic=false&interface=cn.securitystack.indicatorcenter.api.client.ObtainValueClient&logger=slf4j&methods=get,mget&pid=1&revision=1.0.0&side=provider&timestamp=1560575962916&version=1.0.0"

	statusOK = 200
)

func init() {
	SuiteMap.Add("indicator", NewIndicatorSuite())
	hessian.RegisterPOJO(&indicatorResult{})
}

type requestContext struct {
	CurrentTimestamp int64
	Uuid             string
	DatasourceCode   string
	DataDict         map[string]interface{}
	ExtraMap         map[uint32]string
	ExtraParam       extraParam
}

func (r requestContext) JavaClassName() string {
	return "cn.securitystack.indicatorcenter.api.pojo.IndicatorCenterRequestContext"
}

type mgetParam struct {
	CurrentTimestamp int64
	Uuid             string
	DatasourceCode   string
	DataDict         map[string]interface{}
	ExtraMap         map[uint32]string
	ExtraParam       extraParam
	IndicatorCodeSet []string
}

func (m mgetParam) JavaClassName() string {
	return "cn.securitystack.indicatorcenter.api.pojo.MgetParam"
}

type extraParam struct {
	OffLineRequest    bool
	ObtainCallRule    bool
	MergeValueRequest bool
	PolicyLabRequest  bool
	StockRequest      bool
}

func (p extraParam) JavaClassName() string {
	return "cn.securitystack.indicatorcenter.api.pojo.ExtraParam"
}

type indicatorResult struct {
	Data    map[string]interface{}
	Message string
	Code    int
}

func (r indicatorResult) JavaClassName() string {
	return "cn.securitystack.indicatorcenter.api.pojo.IndicatorResult"
}

// IndicatorSuite object of indicator(center) pet tester
type IndicatorSuite struct {
	Host             string
	Iface            string
	DataSourceCode   string
	IndicatorCodeSet peflag.ListFlags
	ExtraParam       extraParam
	ExtraMap         map[uint32]string
	FieldsFn         map[string]func() interface{}

	DcURL common.URL // URL for datasourceClient
	OcURL common.URL // URL for obtainValueClient

	stopCh chan bool
}

// NewIndicatorSuite is used generate a fresh IndicatorSuite
func NewIndicatorSuite() *IndicatorSuite {
	return &IndicatorSuite{FieldsFn: map[string]func() interface{}{}}
}

// Init arguments parsing and initialization
func (is *IndicatorSuite) Init(flagSet *flag.FlagSet, args []string) error {
	fieldsSet := peflag.ListFlags{}

	flagSet.StringVar(&is.Host, "host", "", "indicator host and port, eg. 127.0.0.1:31000")
	flagSet.StringVar(&is.Iface, "iface", "all", "test interface, processMessage|mget|all")
	flagSet.StringVar(&is.DataSourceCode, "dc", "", "datasource code")
	flagSet.Var(&is.IndicatorCodeSet, "ic", "indicator code")
	flagSet.Var(&fieldsSet, "field", "random field name, include: ip, email, phone_number, const_id, or dynamic field like: ext_amount:float_2000, ext_uid:uuid, ext_name:string_64")

	if err := flagSet.Parse(args); err != nil {
		return errors.New("command line args parsing error")
	}

	if is.Host == "" || is.DataSourceCode == "" || len(fieldsSet) == 0 {
		return errors.New("host, datasource code and field should be set")
	}

	if is.Iface != "processMessage" && is.Iface != "mget" && is.Iface != "all" {
		return errors.New("invalid iface, should be one of processMessage, mget, all")
	}

	if is.Iface != "processMessage" && len(is.IndicatorCodeSet) == 0 {
		return errors.New("indicator codeset is mandentary when iface is mget or all")
	}

	ctx := context.Background()
	dcurl, err := common.NewURL(ctx, fmt.Sprintf(dcURLTpl, is.Host))
	if err != nil {
		return err
	}
	is.DcURL = dcurl
	ocurl, err := common.NewURL(ctx, fmt.Sprintf(ocURLTpl, is.Host))
	if err != nil {
		return err
	}
	is.OcURL = ocurl

	for _, field := range fieldsSet {
		name, handler, err := datagen.GetGenerator(field)
		if err != nil {
			return err
		}
		is.FieldsFn[name] = handler
	}

	// TODO: online request support only
	is.ExtraParam = extraParam{
		OffLineRequest:    false,
		ObtainCallRule:    false,
		MergeValueRequest: false,
		PolicyLabRequest:  false,
		StockRequest:      false,
	}
	is.ExtraMap = map[uint32]string{}

	return nil
}

// GetTask return a boomer.Task that do the real job
func (is *IndicatorSuite) GetTask() *boomer.Task {
	return &boomer.Task{
		Name: "indicator",
		OnStart: func() {
			is.stopCh = make(chan bool)
		},
		OnStop: func() {
			close(is.stopCh)
		},
		Fn: func() {
			client := client.NewDubboClient()
			for {
				select {
				case <-is.stopCh:
					return
				default:
					fields := is.randomFields()
					switch is.Iface {
					case "processMessage":
						is.doProcessMessage(client, fields)
					case "mget":
						is.doMget(client, fields)
					case "all":
						is.doMget(client, fields)
						is.doProcessMessage(client, fields)
					}
				}
			}
		},
	}
}

func (is *IndicatorSuite) doProcessMessage(client *dubbo.Client, fields map[string]interface{}) {
	is.doDubboRequest(
		client,
		"processMessage",
		[]interface{}{
			requestContext{
				CurrentTimestamp: currentTimestampMills(),
				Uuid:             uuid.Must(uuid.NewV4()).String(),
				DatasourceCode:   is.DataSourceCode,
				DataDict:         fields,
				ExtraMap:         is.ExtraMap,
				ExtraParam:       is.ExtraParam,
			},
		})
}

func (is *IndicatorSuite) doMget(client *dubbo.Client, fields map[string]interface{}) {
	is.doDubboRequest(
		client,
		"mget",
		[]interface{}{
			mgetParam{
				CurrentTimestamp: currentTimestampMills(),
				Uuid:             uuid.Must(uuid.NewV4()).String(),
				DatasourceCode:   is.DataSourceCode,
				DataDict:         fields,
				ExtraMap:         is.ExtraMap,
				ExtraParam:       is.ExtraParam,
				IndicatorCodeSet: is.IndicatorCodeSet,
			},
		})
}

func (is *IndicatorSuite) doDubboRequest(client *dubbo.Client, method string, args []interface{}) {
	var err error

	resp := &indicatorResult{}
	start := boomer.Now()
	if method == "processMessage" {
		err = client.Call(is.Host, is.DcURL, method, args, resp)
	} else {
		err = client.Call(is.Host, is.OcURL, method, args, resp)
	}
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure(method, "err", elapsed, err.Error())
		return
	}

	if resp.Code != statusOK {
		boomer.RecordFailure(method, strconv.Itoa(resp.Code), elapsed, resp.Message)
		return
	}
	boomer.RecordSuccess(method, "200", elapsed, int64(0))
}

func (is *IndicatorSuite) randomFields() map[string]interface{} {
	fields := make(map[string]interface{}, len(is.FieldsFn))
	for name, handler := range is.FieldsFn {
		fields[name] = handler()
	}

	return fields
}

func currentTimestampMills() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}
