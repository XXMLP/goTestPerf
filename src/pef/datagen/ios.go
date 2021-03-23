package datagen

import (
	"pef/mobile"

	"github.com/golang/protobuf/proto"
)

var (
	iosKeyMap = map[string]string{
		"macAddress":        "K1",
		"proxyType":         "K2",
		"proxyUrl":          "K3",
		"idfa":              "K4",
		"idfv":              "K5",
		"uuid":              "K6",
		"cpuType":           "K7",
		"cpuSubType":        "K8",
		"maxCpus":           "K9",
		"availCpus":         "K10",
		"totalSpace":        "K11",
		"freeSpace":         "K12",
		"memory":            "K13",
		"resolution":        "K14",
		"networkType":       "K15",
		"rssi":              "K16",
		"battery":           "K17",
		"os":                "K18",
		"netWorkNode":       "K19",
		"release":           "K20",
		"sysVersion":        "K21",
		"hardWareType":      "K22",
		"jailBreak":         "K23",
		"aslr":              "K24",
		"bootTime":          "K25",
		"cellIp":            "K26",
		"wifiIp":            "K27",
		"wifiNetMask":       "K28",
		"vpnIp":             "K29",
		"vpnNetMask":        "K30",
		"language":          "K31",
		"deviceName":        "K32",
		"carrier":           "K33",
		"mnc":               "K34",
		"mcc":               "K35",
		"radioAccessType":   "K36",
		"countryIso":        "K37",
		"bssId":             "K38",
		"ssId":              "K39",
		"dns":               "K40",
		"brightNess":        "K41",
		"jbPrint":           "K42",
		"gpsLocation":       "K43",
		"gpsStatus":         "K44",
		"timeZone":          "K45",
		"bundle":            "K46",
		"osVersion":         "K47",
		"platform":          "K48",
		"sdkVersion":        "K49",
		"hookInline":        "K50",
		"hookObjective":     "K51",
		"debug":             "K52",
		"inject":            "K53",
		"hook":              "K54",
		"collect_cost_time": "K300",
	}
	constIOSInfoMap = map[string]string{
		"wifiNetMask":  "255.255.252.0",
		"cpuType":      "16777228",
		"memory":       "2097152000",
		"freeSpace":    "4100239360",
		"release":      "16.7.0",
		"maxCpus":      "2",
		"cpuSubType":   "1",
		"language":     "zh-Hans-CN",
		"battery":      "unplugged/81%",
		"mcc":          "460",
		"resolution":   "750X1334",
		"deviceName":   "Q",
		"platform":     "iPhone9,1",
		"wifiIp":       "10.0.1.45",
		"aslr":         "170491904",
		"countryIso":   "CN",
		"gpsStatus":    "16",
		"cellIp":       "10.156.153.174",
		"networkType":  "Wifi",
		"netWorkNode":  "Q",
		"bundle":       "com.dingxiang.app1.new_1.0",
		"sysVersion":   "Darwin Kernel Version 16.7.0: Thu Jun 15 18:33:36 PDT 2017; root:xnu-3789.70.16~4/RELEASE_ARM64_T8010",
		"rssi":         "-70",
		"brightNess":   "0.424114",
		"mnc":          "02",
		"osVersion":    "Darwin",
		"bootTime":     "1505399451",
		"hardWareType": "iPhone9,1",
		"proxyType":    "none",
		"dns":          "10.0.0.254",
		"jailBreak":    "false",
	}
	constIOSInfoItems = make([]*mobile.STEERiskMgrReportDataRequest_STEEInfoItem, len(constIOSInfoMap))
	randomIOSItems    = []string{"uuid", "idfv", "bssId", "macAddress"}
	randomBoolIOSItems    = []string{"hook", "jailBreak", "inject"}

	iosOsType = mobile.STEERequestHeader_iOS
	iosOsVer  = "11.0"
)

func init() {
	// init constInfoItems
	idx := 0
	for k, v := range constIOSInfoMap {
		mappedKey := iosKeyMap[k]
		mappedValue := v
		constIOSInfoItems[idx] = &mobile.STEERiskMgrReportDataRequest_STEEInfoItem{
			Name:  &mappedKey,
			Value: &mappedValue,
		}
		idx++
	}
}

// RandomIOSRequest is used to generate a ios request data
func RandomIOSRequest(appKey string) ([]byte, error) {
	reqData := randomIOSRequestData()
	reqBytes, err := proto.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	reqType := mobile.STEEDataType_DATATYPE_DO_REPORT_DATA
	reqHeader := &mobile.STEERequestHeader{
		SdkVer:       &sdkVer,
		AppCode:      &appCode,
		AppVerCode:   &appVerCode,
		AppKey:       &appKey,
		OsType:       &iosOsType,
		OsVer:        &iosOsVer,
		ProtoVersion: &protoVersion,
	}
	request := &mobile.STEERequest{
		Header: reqHeader,
		Type:   &reqType,
		Data:   reqBytes,
	}

	return proto.Marshal(request)
}

func randomIOSRequestData() *mobile.STEERiskMgrReportDataRequest {
	infos := randomIOSDeviceInfo()
	return &mobile.STEERiskMgrReportDataRequest{
		Infos: infos,
	}
}

func randomIOSDeviceInfo() []*mobile.STEERiskMgrReportDataRequest_STEEInfoItem {
	randomInfoItems := make([]*mobile.STEERiskMgrReportDataRequest_STEEInfoItem, len(randomIOSItems))
	randomBoolInfoItems := make([]*mobile.STEERiskMgrReportDataRequest_STEEInfoItem, len(randomBoolIOSItems))

	for idx, k := range randomIOSItems {
		mappedKey := iosKeyMap[k]
		value := randomUUID()
		randomInfoItems[idx] = &mobile.STEERiskMgrReportDataRequest_STEEInfoItem{
			Name:  &mappedKey,
			Value: &value,
		}
	}

	for idx, k := range randomBoolIOSItems {
		mappedKey := iosKeyMap[k]
		value := randomStrBoolean()
		randomBoolInfoItems[idx] = &mobile.STEERiskMgrReportDataRequest_STEEInfoItem{
			Name:  &mappedKey,
			Value: &value,
		}
	}
	newConstIOSInfoItems := append(constIOSInfoItems, randomInfoItems...)
	return append(newConstIOSInfoItems, randomBoolInfoItems...)
}
