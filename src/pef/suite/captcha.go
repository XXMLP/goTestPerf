package suite

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"pef/client"
	"strconv"

	"github.com/feiyuw/boomer"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

const (
	aIDChars    = "abcdefghijklmnopqrstuvwxyz0123456789"
	aIDCharsLen = len(aIDChars)

	defaultX  = "47"
	defaultY  = "222"
	defaultAC = "s_v3#X8Xnl+A7s20Z8ooYXXOcX8XYZw33XYGOzXNP87u4dmL8abV1w9D4X4Q8MXPDhv4FwXhhT7/3A9kUjx9RCmB4D7/izwHCm8XsYR8VINCiUyggaPf1XX1Tm2X3D9bHD2rXm3uZRA/vXYvN63yHPJ12QskoqG6dlH1EbBAoBHlaMvYhto0itbiuVfYXX2X/VFKbDB6BKTLgF7foGOKR+G3BGRrsXXluCfa+6dC5N4zOjXXSrf1P6njlr1AmrOwcK4GE5FP95pzrdpFrgpLEd0sQqUDygbGSGUkJoXfXXO3sXYJuCf94DvYTHdbEmz5C8EZRvV8oRTv236ns+j43i33c8jvQm1yU/6XTXXLMgHWSi/vjlM3XYFPMbShuqpoj38XngHPN4D81mKgTXXLMgHOEfM24bT3XYFPMbcZh9QWB38XngHPN8yFvtYXTXXLMgHf2I+UEhC3XYFPMbE3rpGxh38XngHPN3TYFLAfTXXLMgHfUjqpl8C3XYFPMKHP26aE238XngHP75UuZRi8TXXLMgHzRVzMbcT3XYFPMKWFYWzUk38XngHP7oeImXUyTXXLMgHAESMW6BM3XYFPMKOiRvhRU38XngHP7Cn/tIbyTXXLMgHcd4owZvC3XYFPMK64ANbSN38XngHP7acJacPyTXXLMgHc1T51XkT3XYFPMKhyS7U/s38XngHP71TLFwE3TXXLMgHynY4D/vm3XYFPMKFeGPMGp38XngHP5qKSFj0gTXXLMgHLkNL/VMC3XYFPMKIdIUL0n38XngHP50kMX8k/TXXLMgHEQeqplDM3XYFPMKvqWcvar38XngHP5L5Njrz8TXXLMgHC2WAurcC3XYFPMK1mEJzpI38XngHP5/aZURGrTXXLMgHCm8NvLcm3XYFPMKjn2uMkp38XngHP5nt3iYBITXXLMgHr/uLoq2C3XYFPMKX81ZLHF38XngHPGbBDPML5TXXLMgH0A7ZP8im3XYFPMKbk/5pIk38XngHPGMIkMgXrTXXLMgHpVQSqF4M3XYFPMK7pOS+mf38XngHPGESE2o/yTXXLMgHxyOcbJxC3XYFPMKc4ak1Na38XngHPqTVa7Jn3TXXLMgHS9Z2SsgC3XYFPMKE/z1TLC38XngHPqYrrdZogTXXLMgHdgBoPXUC3XYFPMBWDPKFyj38XngHPb4D4nhkIiXXMjFybKmraM3y7KXXX="
)

func init() {
	SuiteMap.Add("captcha", NewCaptchaSuite())
}

// json struct of response of a interface
type aResponse struct {
	Sid     string `json:"sid"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// CaptchaSuite object of captcha pet tester
type CaptchaSuite struct {
	AURLPrefix string
	V1URL      string
	TokenVerifyURL string

	Host  string
	Iface string
	Ak    string
	Jsv   string
	C     string
	Sid   string
	X     string
	Y     string
	Ac    string
}

// NewCaptchaSuite generate a new CaptchaSuite object
func NewCaptchaSuite() *CaptchaSuite {
	return &CaptchaSuite{}
}

// Init arguments parsing and initialization
func (cs *CaptchaSuite) Init(flagSet *flag.FlagSet, args []string) error {
	flagSet.StringVar(&cs.Host, "host", "", "captcha host and port, eg. 127.0.0.1:7776")
	flagSet.StringVar(&cs.Iface, "iface", "all", "test interface, a|v1|all")
	flagSet.StringVar(&cs.Ak, "ak", "", "app ID")
	flagSet.StringVar(&cs.X, "x", defaultX, "valid X")
	flagSet.StringVar(&cs.Y, "y", defaultY, "valid Y")
	flagSet.StringVar(&cs.Ac, "ac", defaultAC, "valid captcha path")
	flagSet.StringVar(&cs.Jsv, "jsv", "1.3.9.91", "js version, eg. 1.3.9.91")
	flagSet.StringVar(&cs.C, "c", "", "valid constant ID")
	flagSet.StringVar(&cs.Sid, "sid", "", "sid return by a interface, mandentary if iface is v1")

	flagSet.Parse(args)

	if cs.Host == "" || cs.Ak == "" || cs.Jsv == "" || cs.C == "" {
		return errors.New("host, ak, c should be set")
	}
	if cs.Iface != "a" && cs.Iface != "v1" && cs.Iface != "tokenVerify" && cs.Iface != "all" {
		return errors.New("iface should be one of a|v1|tokenVerify|all")
	}
	if cs.Iface == "v1" && cs.Sid == "" {
		return errors.New("sid should be set if iface is v1")
	}

	cs.AURLPrefix = fmt.Sprintf("http://%s/api/a?jsv=%s&c=%s&ak=%s&s=50&h=150&w=300&aid=", cs.Host, cs.Jsv, cs.C, cs.Ak)
	cs.V1URL = fmt.Sprintf("http://%s/api/v1", cs.Host)
	cs.TokenVerifyURL = fmt.Sprintf("http://%s/api/tokenVerify?ak=%s&token=", cs.Host, cs.Ak)

	return nil
}

// GetTask return a boomer task of captcha
func (cs *CaptchaSuite) GetTask() *boomer.Task {
	return &boomer.Task{
		Name:    "captcha",
		OnStart: func() {},
		OnStop:  func() {},
		Fn: func() {
			aid := randomAID()
			switch cs.Iface {
			case "a":
				cs.doA(aid)
				break
			case "v1":
				cs.doV1(aid, cs.Sid)
				break
			case "tokenVerify":
				cs.doTokenVerify()
				break
			case "all":
				sid, err := cs.doA(aid)
				if err == nil {
					err := cs.doV1(aid, sid)
					if err == nil{
						cs.doTokenVerify()
					}
				}
				break
			}
		},
	}
}

func (cs *CaptchaSuite) doA(aid string) (string, error) {
	start := boomer.Now()
	statusCode, body, err := client.HTTPClient.Get(nil, cs.AURLPrefix+aid)
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure("/api/a", "err", elapsed, err.Error())
		return "", err
	}
	if statusCode != fasthttp.StatusOK {
		boomer.RecordFailure("/api/a", strconv.Itoa(statusCode), elapsed, "")
		return "", errors.New(string(body))
	}

	resp := aResponse{}

	if err := json.Unmarshal(body, &resp); err != nil {
		boomer.RecordFailure("/api/a", "err", elapsed, err.Error())
		return "", err
	}

	if !resp.Success {
		boomer.RecordFailure("/api/a", "err", elapsed, resp.Msg)
		return "", errors.New(resp.Msg)
	}

	boomer.RecordSuccess("/api/a", "200", elapsed, int64(0))

	return resp.Sid, nil
}

func (cs *CaptchaSuite) doV1(aid, sid string) error {
	args := fasthttp.Args{}
	args.Set("x", "34")
	args.Set("y", "127")
	args.Set("aid", aid)
	args.Set("sid", sid)
	args.Set("jsv", cs.Jsv)
	args.Set("c", cs.C)
	args.Set("ak", cs.Ak)
	args.Set("ac", cs.Ac)

	start := boomer.Now()
	statusCode, _, err := client.HTTPClient.Post(nil, cs.V1URL, &args)
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure("/api/v1", "err", elapsed, err.Error())
		return err
	}
	if statusCode == fasthttp.StatusOK {
		boomer.RecordSuccess("/api/v1", "200", elapsed, int64(0))
		return nil
	} else {
		boomer.RecordFailure("/api/v1", strconv.Itoa(statusCode), elapsed, "")
		return errors.New(strconv.Itoa(statusCode))
	}

}

func (cs *CaptchaSuite) doTokenVerify() {
	start := boomer.Now()
	statusCode, _, err := client.HTTPClient.Get(nil, cs.TokenVerifyURL+"1")
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure("/api/tokenVerify", "err", elapsed, err.Error())
		return
	}
	if statusCode == fasthttp.StatusOK {
		boomer.RecordSuccess("/api/tokenVerify", "200", elapsed, int64(0))
	} else {
		boomer.RecordFailure("/api/tokenVerify", strconv.Itoa(statusCode), elapsed, "")
	}
}

func randomAID() string {
	aid := make([]byte, 32)
	for i := 0; i < 32; i++ {
		aid[i] = aIDChars[rand.Intn(aIDCharsLen)]
	}

	return string(aid)
}
