package suite

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"pef/client"
	"pef/datagen"
	"pef/mobile"
	"strconv"
	"sync"

	"github.com/dubbogo/hessian2"
	"github.com/feiyuw/boomer"
	"github.com/feiyuw/dubbo-go/common"
	"github.com/feiyuw/dubbo-go/protocol/dubbo"
	"github.com/feiyuw/xxtea"
	"github.com/golang/protobuf/proto"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

const (
	zipfileHeader = 0x04034b50 // PK\x03\x04 in little endian
	dubboURLTpl   = "dubbo://%s/cn.securitystack.ctucommon.constid.api.ConstIDService?anyhost=true&application=constid-api&default.timeout=2000&dubbo=2.5.6&generic=false&interface=cn.securitystack.ctucommon.constid.api.ConstIDService&methods=findByToken&pid=8&revision=1.0.0&side=provider&timestamp=1560575985440&version=1.0.0"
)

var (
	base64Encoder = base64.NewEncoding("S0DOZN9bBJyPV-qczRa3oYvhGlUMrdjW7m2CkE5_FuKiTQXnwe6pg8fs4HAtIL1x")
	types         = [4]string{"web", "ios", "android", "miniProgram"}

	constidSignPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func init() {
	SuiteMap.Add("constid", NewConstIDSuite())
	SuiteMap.Add("constid-dubbo", NewConstIDDubboSuite())
	hessian.RegisterPOJO(&constIdApiResult{})
	hessian.RegisterPOJO(&deviceInfo{})
}

type deviceInfoResponse struct {
	StateCode int               `json:"stateCode"`
	Message   string            `json:"message"`
}

// ConstIDSuite object of constid pet tester
type ConstIDSuite struct {
	Host      string
	Iface     string
	AppKey    string
	AppSecret string
	Type      string

	C1URL         string
	M1URL         string
	W1URL         string
	DeviceInfoURL string

	SecretKey []byte
}

// NewConstIDSuite generate a new ConstIDSuite object
func NewConstIDSuite() *ConstIDSuite {
	return &ConstIDSuite{}
}

// Init arguments parsing and initialization
func (suite *ConstIDSuite) Init(flagSet *flag.FlagSet, args []string) error {
	flagSet.StringVar(&suite.Host, "host", "", "constid host and port, eg. 127.0.0.1:7776")
	flagSet.StringVar(&suite.Iface, "iface", "token", "test interface, token|verify|all")
	flagSet.StringVar(&suite.AppKey, "app-key", "", "key of app")
	flagSet.StringVar(&suite.AppSecret, "app-secret", "", "secret of app")
	flagSet.StringVar(&suite.Type, "type", "web", "client type, web|ios|android|miniProgram|random")

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if suite.Host == "" || suite.AppKey == "" {
		return errors.New("host, app-key should be set")
	}
	if suite.Iface != "token" && suite.Iface != "verify" && suite.Iface != "all" {
		return errors.New("iface should be one of token|verify|all")
	}
	if suite.Iface != "token" && suite.AppSecret == "" {
		return errors.New("app-secret should be set if iface is not token")
	}

	if suite.Type != "web" && suite.Type != "ios" && suite.Type != "android" && suite.Type != "miniProgram" && suite.Type != "random" {
		return errors.New("type should be one of web|ioc|android|miniProgram|random")
	}

	suite.C1URL = fmt.Sprintf("http://%s/udid/c1", suite.Host)
	suite.W1URL = fmt.Sprintf("http://%s/udid/w1", suite.Host)
	suite.M1URL = fmt.Sprintf("http://%s/udid/m1?appKey=%s&version=5.5.0&sign=", suite.Host, suite.AppKey)
	suite.DeviceInfoURL = fmt.Sprintf("http://%s/udid/api/getDeviceInfo?", suite.Host)
	suite.SecretKey = generateSecretKey(suite.AppKey)

	return nil
}

// GetTask return a boomer task of constid
func (suite *ConstIDSuite) GetTask() *boomer.Task {
	return &boomer.Task{
		Name:    "constid",
		OnStart: func() {},
		OnStop:  func() {},
		Fn:      suite.doRequest,
	}
}

func (suite *ConstIDSuite) doRequest() {
	var reqType string
	var token string

	if suite.Type == "random" {
		reqType = suite.randomType()
	} else {
		reqType = suite.Type
	}

	if suite.Iface == "verify" {
		token = datagen.RandomConstID().(string)
	} else {
		switch reqType {
		case "web":
			token = suite.doC1()
		case "android":
			token = suite.doM1("android")
		case "ios":
			token = suite.doM1("ios")
		case "miniProgram":
			token = suite.doW1()
		}
	}

	if suite.Iface != "token" && token != "" {
		suite.doGetDeviceInfo(token)
	}
}

func (suite *ConstIDSuite) randomType() string {
	return types[rand.Intn(len(types))]
}

// web client
func (suite *ConstIDSuite) doC1() string {
	param, err := datagen.RandomWebRequest(suite.AppKey)
	if err != nil {
		boomer.RecordFailure("/udid/c1(web)", "err", 0, err.Error())
		return ""
	}
	start := boomer.Now()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(suite.C1URL)
	req.Header.Set("Param", base64Encoder.EncodeToString(param))
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)
	err = client.HTTPClient.Do(req, resp)
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure("/udid/c1(web)", "err", elapsed, err.Error())
		return ""
	}

	statusCode := resp.StatusCode()
	if statusCode == fasthttp.StatusOK {
		var webResp webResponse
		if err := json.Unmarshal(resp.Body(), &webResp); err == nil && webResp.Status == 2 {
			boomer.RecordSuccess("/udid/c1(web)", "200", elapsed, int64(0))
			return webResp.Data
		}
		boomer.RecordFailure("/udid/c1(web)", strconv.Itoa(webResp.Status), elapsed, webResp.Msg)
	} else {
		boomer.RecordFailure("/udid/c1(web)", strconv.Itoa(statusCode), elapsed, "")
	}

	return ""
}

//miniProgram client
func (suite *ConstIDSuite ) doW1() string{
	param, err := datagen.RandomminiProgramRequest(suite.AppKey)
	if err != nil {
		boomer.RecordFailure("/udid/w1(miniProgram)", "err", 0, err.Error())
		return ""
	}
	start := boomer.Now()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(suite.W1URL)
	req.Header.Set("Param", base64Encoder.EncodeToString(param))
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)
	err = client.HTTPClient.Do(req, resp)
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure("/udid/w1(miniProgram)", "err", elapsed, err.Error())
		return ""
	}

	statusCode := resp.StatusCode()
	if statusCode == fasthttp.StatusOK {
		var miniProgramResp miniProgramResponse
		if err := json.Unmarshal(resp.Body(), &miniProgramResp); err == nil && miniProgramResp.Status == 2 {
			boomer.RecordSuccess("/udid/w1(miniProgram)", "200", elapsed, int64(0))
			return miniProgramResp.Data
		}
		boomer.RecordFailure("/udid/w1(miniProgram)", strconv.Itoa(miniProgramResp.Status), elapsed, miniProgramResp.Msg)
	} else {
		boomer.RecordFailure("/udid/w1(miniProgram)", strconv.Itoa(statusCode), elapsed, "")
	}

	return ""

}

// mobile client
func (suite *ConstIDSuite) doM1(platform string) string {
	var (
		data  []byte
		err   error
		topic string
	)

	if platform == "android" {
		data, err = datagen.RandomAndroidRequest(suite.AppKey)
		topic = "/udid/m1(android)"
	} else {
		data, err = datagen.RandomIOSRequest(suite.AppKey)
		topic = "/udid/m1(ios)"
	}

	if err != nil {
		boomer.RecordFailure(topic, "err", 0, err.Error())
		return ""
	}

	sign := suite.getM1Sign(data)
	compressData := suite.compress(data)
	encryptedData, err := xxtea.Encrypt(compressData, suite.SecretKey)
	if err != nil {
		boomer.RecordFailure(topic, "err", 0, err.Error())
		return ""
	}

	start := boomer.Now()
	token, err := suite.callM1(encryptedData, sign)
	elapsed := boomer.Now() - start

	if err != nil {
		boomer.RecordFailure(topic, "err", elapsed, err.Error())
		return ""
	}
	if token == "" {
		boomer.RecordFailure(topic, "err", elapsed, "empty token")
		return ""
	}

	boomer.RecordSuccess(topic, "200", elapsed, int64(0))
	return token
}

func (suite *ConstIDSuite) doGetDeviceInfo(token string) {
	start := boomer.Now()
	statusCode, body, err := suite.callGetDeviceInfo(token)
	elapsed := boomer.Now() - start

	if err != nil {
		boomer.RecordFailure("/udid/api/getDeviceInfo", "err", elapsed, err.Error())
		return
	}

	devInfoResp := deviceInfoResponse{}
	if err := json.Unmarshal(body, &devInfoResp); err != nil {
		boomer.RecordFailure("/udid/api/getDeviceInfo", "err", elapsed, err.Error())
		return
	}

	if statusCode == fasthttp.StatusOK && devInfoResp.StateCode == fasthttp.StatusOK{
		boomer.RecordSuccess("/udid/api/getDeviceInfo", "200", elapsed, int64(0))
	}else if devInfoResp.StateCode != fasthttp.StatusOK{
		boomer.RecordFailure("/udid/api/getDeviceInfo", strconv.Itoa(devInfoResp.StateCode), elapsed, devInfoResp.Message)
		return
	}else if statusCode != fasthttp.StatusOK{
		boomer.RecordFailure("/udid/api/getDeviceInfo", strconv.Itoa(statusCode), elapsed, devInfoResp.Message)
		return
	}

}

func (suite *ConstIDSuite) callGetDeviceInfo(token string) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(suite.DeviceInfoURL+string(suite.getDeviceInfoPostData(token)))
	req.Header.SetMethod("POST")
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if err := client.HTTPClient.Do(req, resp); err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}

func (suite *ConstIDSuite) getDeviceInfoPostData(token string) []byte {
	var buf bytes.Buffer

	buf.WriteString("appId=")
	buf.WriteString(suite.AppKey)
	buf.WriteString("&sign=")
	buf.WriteString(suite.getDeviceInfoSign(token))
	buf.WriteString("&token=")
	buf.WriteString(token)

	return buf.Bytes()
}

func (suite *ConstIDSuite) getDeviceInfoSign(token string) string {
	buf := constidSignPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer constidSignPool.Put(buf)

	buf.WriteString(suite.AppSecret)
	buf.WriteString(token)
	buf.WriteString(suite.AppSecret)

	sum := md5.Sum(buf.Bytes())
	return hex.EncodeToString(sum[:])
}

func (suite *ConstIDSuite) callM1(data []byte, sign string) (string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(suite.M1URL + sign)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/octet-stream")
	req.SetBody(data)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	if err := client.HTTPClient.Do(req, resp); err != nil {
		return "", err
	}

	if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
		return "", errors.New(strconv.Itoa(statusCode))
	}

	return suite.parseResponse(resp.Body())
}

func (suite *ConstIDSuite) parseResponse(body []byte) (string, error) {
	var data []byte
	// 1. decrypt
	decryptedData, err := xxtea.Decrypt(body, suite.SecretKey)
	if err != nil {
		return "", err
	}

	// 2. uncompress
	if binary.LittleEndian.Uint32(decryptedData) == zipfileHeader {
		if data, err = suite.uncompress(decryptedData); err != nil {
			return "", err
		}
	} else {
		data = decryptedData
	}

	// 3. protobuf parsing
	response := &mobile.STEEResponse{}
	if err = proto.Unmarshal(data, response); err != nil {
		return "", err
	}
	if response.GetErrCode() != mobile.STEEErrorCode_ERR_NOERROR {
		return "", errors.New(strconv.Itoa(int(response.GetErrCode())))
	}
	// 4. constid read
	dataResp := &mobile.STEERiskMgrReportDataResponse{}
	if err = proto.Unmarshal(response.GetData(), dataResp); err != nil {
		return "", err
	}

	return dataResp.GetConstid(), nil
}

func (suite *ConstIDSuite) getM1Sign(reqData []byte) string {
	dataLen := len(reqData)
	if dataLen == 0 {
		return ""
	}
	keyLen := len(suite.AppKey)
	data := make([]byte, keyLen*2+dataLen)
	copy(data[:keyLen], suite.AppKey)
	copy(data[keyLen:keyLen+dataLen], reqData)
	copy(data[keyLen+dataLen:], suite.AppKey)

	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

func (suite *ConstIDSuite) compress(src []byte) []byte {
	buf := new(bytes.Buffer)
	// Create a new zip archive.
	w := zip.NewWriter(buf)
	f, _ := w.Create("data")
	f.Write(src)
	w.Close()

	return buf.Bytes()
}

func (suite *ConstIDSuite) uncompress(src []byte) ([]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(src), int64(len(src)))
	if err != nil {
		return nil, err
	}
	if len(reader.File) != 1 {
		return nil, errors.New("incorrect files in zip stream")
	}
	rc, err := reader.File[0].Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return ioutil.ReadAll(rc)
}

func generateSecretKey(appKey string) []byte {
	keyBytes := make([]byte, len(appKey)*3)
	copy(keyBytes[:len(appKey)], appKey)
	copy(keyBytes[len(appKey):2*len(appKey)], appKey)
	copy(keyBytes[2*len(appKey):], appKey)

	sum := sha1.Sum(keyBytes)
	secretKey := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(secretKey, sum[:])

	return secretKey
}

// HTTP /udid/c1 response json, {"data":"5d0c7594g2Pc7i9pfOWGKnWFWQLNfmIUDAoexPo1","msg":"success","status":2}
type webResponse struct {
	Data   string `json:"data"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
// HTTP /udid/w1 response json, {"data":"5d0c7594g2Pc7i9pfOWGKnWFWQLNfmIUDAoexPo1","msg":"success","status":2}
type miniProgramResponse struct {
	Data   string `json:"data"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

// Dubbo findByToken result
type constIdApiResult struct {
	Success bool
	Data    deviceInfo
	Msg     string
}

func (r constIdApiResult) JavaClassName() string {
	return "cn.securitystack.ctucommon.constid.api.util.ConstIdApiResult"
}

type deviceInfo struct {
	HardId     string
	Token      string
	AppKey     string
	UserId     string
	Scene      string
	CreateTime int
	CreateIp   string
	RiskCode   string
	Inet       string
	UserAgent  string
	Data       string
}

func (d deviceInfo) JavaClassName() string {
	return "cn.securitystack.ctucommon.constid.model.DeviceInfo"
}

// NewConstIDDubboSuite return a ConstIDDubboSuite
func NewConstIDDubboSuite() *ConstIDDubboSuite {
	return &ConstIDDubboSuite{}
}

// ConstIDDubboSuite object of constid pet tester
type ConstIDDubboSuite struct {
	Host         string
	FindTokenURL common.URL

	stopCh chan struct{} // handle start/stop event
}

// Init arguments parsing and initialization
func (suite *ConstIDDubboSuite) Init(flagSet *flag.FlagSet, args []string) error {
	flagSet.StringVar(&suite.Host, "host", "", "constid dubbo host and port, eg. 127.0.0.1:36502")

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if suite.Host == "" {
		return errors.New("host should be set")
	}

	durl, err := common.NewURL(context.Background(), fmt.Sprintf(dubboURLTpl, suite.Host))
	if err != nil {
		return err
	}
	suite.FindTokenURL = durl

	return nil
}

// GetTask return a boomer Task instance
func (suite *ConstIDDubboSuite) GetTask() *boomer.Task {
	return &boomer.Task{
		Name: "constid-dubbo",
		OnStart: func() {
			suite.stopCh = make(chan struct{})
		},
		OnStop: func() {
			close(suite.stopCh)
		},
		Fn: suite.doRequest,
	}
}

func (suite *ConstIDDubboSuite) doRequest() {
	dubboClient := client.NewDubboClient()
	for {
		select {
		case <-suite.stopCh:
			return
		default:
			token := datagen.RandomConstID().(string)
			suite.doFindByToken(dubboClient, token)
		}
	}
}

// dubbo verify
func (suite *ConstIDDubboSuite) doFindByToken(dubboClient *dubbo.Client, token string) {
	resp := &constIdApiResult{}
	start := boomer.Now()
	err := dubboClient.Call(suite.Host, suite.FindTokenURL, "findByToken", []interface{}{token}, resp)
	elapsed := boomer.Now() - start
	if err != nil {
		boomer.RecordFailure("findByToken", "err", elapsed, err.Error())
		return
	}
	if !resp.Success {
		boomer.RecordFailure("findByToken", "err", elapsed, resp.Msg)
		return
	}
	boomer.RecordSuccess("findByToken", "200", elapsed, int64(0))
}
