package datagen

import (
	json "github.com/json-iterator/go"
)


type miniProgramRequest struct {
	Sv     string `json:"sv"`
	Ac     string `json:"ac"`
	Att    string `json:"att"`
	Al     string `json:"al"`
	Bl     string `json:"bl"`
	Bml    string `json:"bml"`
	Bd     string `json:"bd"`
	Bs     string `json:"bs"`
	Ct     string `json:"ct"`
	Dc     string `json:"dc"`
	Fss    string `json:"fss"`
	Ha     string `json:"ha"`
	Lang   string `json:"lang"`
	Lt     string `json:"lt"`
	Lgt    string `json:"lgt"`
	Md     string `json:"md"`
	Nt     string `json:"nt"`
	Pr     string `json:"pr"`
	Pf     string `json:"pf"`
	Sh     string `json:"sh"`
	Sw     string `json:"sw"`
	Se     string `json:"se"`
	Sp     string `json:"sp"`
	Ss     string `json:"ss"`
	Sm     string `json:"sm"`
	Sy     string `json:"sy"`
	Si     string `json:"si"`
	Vs     string `json:"vs"`
	Va     string `json:"va"`
	Wh     string `json:"wh"`
	Ww     string `json:"ww"`
	Od     string `json:"od"`
	SdkVer string `json:"sdkVer"`
	Gps    string `json:"gps"`
	Lid    string `json:"lid"`
	AppKey    string `json:"appKey"`
}

// RandomminiProgramRequest generate a random constid request of web client
func RandomminiProgramRequest(appKey string) ([]byte, error) {
	request := miniProgramRequest{
		Sv:     "5.6.0",
		Ac:     "unknown",
		Att:    "unknown",
		Al:     "false",
		Bl:     "100",
		Bml:    "50",
		Bd:     "iphone",
		Bs:     "unknown",
		Ct:     "unknown",
		Dc:     "false",
		Fss:    "16",
		Ha:     "unknown",
		Lang:   "CN",
		Lt:     "23.143755",
		Lgt:    "113.316933",
		Md:     "unknown",
		Nt:     "wifi",
		Pr:     "2",
		Pf:     "wechat",
		Sh:     "640",
		Sw:     "360",
		Se:     "true",
		Sp:     "unknown",
		Ss:     "unknown",
		Sm:     "unknown",
		Sy:     "10.15",
		Si:     "unknown",
		Vs:     "7.0.15",
		Va:     "0",
		Wh:     "640",
		Ww:     "360",
		Od:     "a12345",
		SdkVer: "1.0",
		Gps:    "23.143755,113.316933",
		Lid:    randomServerLid(),
		AppKey: appKey,

	}

	return json.Marshal(request)
}
