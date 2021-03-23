package datagen

import (
	"pef/mobile"

	"github.com/golang/protobuf/proto"
)

var (
	androidKeyMap = map[string]string{
		"app_package":                  "K2",
		"app_name":                     "K3",
		"app_version":                  "K4",
		"app_process_name":             "K6",
		"app_sign_md5":                 "K7",
		"risk_version":                 "K8",
		"kernelVersion":                "K10",
		"build_version_release":        "K11",
		"build_version_security_patch": "K12",
		"build_fingerprint":            "K13",
		"build_hardware":               "K14",
		"build_host":                   "K15",
		"build_time":                   "K16",
		"build_device":                 "K17",
		"build_model":                  "K18",
		"build_brand":                  "K19",
		"build_product":                "K20",
		"build_cpu_abis":               "K21",
		"build_display":                "K22",
		"build_id":                     "K23",
		"build_manufacturer":           "K24",
		"build_board":                  "K25",
		"cores":                        "K30",
		"features":                     "K31",
		"flags":                        "K32",
		"hardware":                     "K33",
		"max_freq":                     "K34",
		"min_freq":                     "K35",
		"module_name":                  "K36",
		"processor":                    "K37",
		"vendor_id":                    "K38",
		"phone_number":                 "K44",
		"voice_number":                 "K45",
		"netType1":                     "K46",
		"netType2":                     "K47",
		"operator1":                    "K48",
		"operator2":                    "K49",
		"device_id1":                   "K50",
		"subscriberId1":                "K51",
		"simSerialNumber1":             "K52",
		"device_id2":                   "K53",
		"subscriberId2":                "K54",
		"simSerialNumber2":             "K55",
		"aid_android_id":               "K56",
		"sid_build_serial":             "K57",
		"uuid":                         "K58",
		"font_hash":                    "K60",
		"gles":                         "K61",
		"ram_rom_sdcard":               "K65",
		"ip":                           "K70",
		"port":                         "K71",
		"vpn":                          "K72",
		"wifiMac":                      "K75",
		"maclist_name_mac_from_native": "K76",
		"maclist_name_mac_from_java":   "K77",
		"api1_mac":                     "K80",
		"api1_name":                    "K81",
		"api2_mac":                     "K82",
		"api2_name":                    "K83",
		"status":                       "K90",
		"health":                       "K91",
		"present":                      "K92",
		"level":                        "K93",
		"scale":                        "K94",
		"plugged":                      "K95",
		"voltage":                      "K96",
		"temperature":                  "K97",
		"technology":                   "K98",
		"http_agent":                   "K100",
		"name":                         "K110",
		"size":                         "K111",
		"dpi":                          "K112",
		"density":                      "K113",
		"gps_info":                     "K125",
		"isemulator":                   "K130",
		"inject":                       "K131",
		"memdump":                      "K132",
		"debug":                        "K133",
		"multirun":                     "K134",
		"flaw_janus":                   "K135",
		"device_root":                  "K136",
		"createTime":                   "K200",
	}
	constAndroidInfoMap = map[string]string{
		"http_agent":                   "Dalvik/2.1.0 (Linux; U; Android 8.0.0; SM-G9550 Build/R16NW)",
		"build_version_release":        "8.0.0",
		"name":                         "nil",
		"size":                         "1080x2220",
		"dpi":                          "420",
		"density":                      "2.625",
		"build_version_security_patch": "2018-04-01",
		"gps_info":                     "23.143755,113.316933",
		"build_fingerprint":            "samsung/dream2qltezc/dream2qltechn:8.0.0/R16NW/G9550ZCU2CRD8:user/release-keys",
		"build_hardware":               "qcom",
		"build_host":                   "SWDG9716",
		"build_time":                   "1524835612000",
		"build_device":                 "dream2qltechn",
		"build_model":                  "SM-G9550",
		"build_brand":                  "samsung",
		"app_package":                  "com.dingxiang.app",
		"build_product":                "dream2qltezc",
		"createTime":                   "1528184572170",
		"build_cpu_abis":               "arm64-v8a,armeabi-v7a,armeabi",
		"build_display":                "R16NW.G9550ZCU2CRD8",
		"build_id":                     "R16NW",
		"build_manufacturer":           "samsung",
		"build_board":                  "msm8998",
		"app_name":                     "DeviceFingerprint",
		"cores":                        "8",
		"features":                     "fp asimd evtstrm aes pmull sha1 sha2 crc32\n",
		"hardware":                     "Qualcomm Technologies, Inc MSM8998\n",
		"max_freq":                     "1900800",
		"min_freq":                     "300000",
		"processor":                    "AArch64 Processor rev 1 (aarch64)\n",
		"app_version":                  "1.0",
		"netType1":                     "unavailable",
		"netType2":                     "unavailable",
		"font_hash":                    "781103542",
		"gles":                         "3.2",
		"maclist_name_mac_from_java":   `["dummy0-76:f7:57:24:f2:36","wlan0-a0:cc:2b:bf:90:85","p2p0-a2:cc:2b:bf:90:85","bond0-b6:2d:99:ac:49:be"]`,
		"status":                       "5",
		"health":                       "2",
		"present":                      "true",
		"level":                        "100",
		"scale":                        "100",
		"plugged":                      "2",
		"voltage":                      "4306",
		"temperature":                  "0",
		"technology":                   "Li-ion",
	}
	randomItems = []string{
		"device_id1",
		"device_id2",
		"aid_android_id",
		"sid_build_serial",
		"uuid",
		"ram_rom_sdcard",
		"wifiMac",
	}

	randomBoolItems = []string{
		"isemulator",
		"device_root",
	}

	constAndroidInfoItems = make([]*mobile.STEERiskMgrReportDataRequest_STEEInfoItem, len(constAndroidInfoMap))

	sdkVer        = "5.5.0"
	appVerCode    = int32(1)
	protoVersion  = int32(3)
	appCode       = "test"
	androidOsType = mobile.STEERequestHeader_Android
	androidOsVer  = "11.0"
)

func init() {
	// init constAndroidInfoItems
	idx := 0
	for k, v := range constAndroidInfoMap {
		mappedKey := androidKeyMap[k]
		mappedValue := v
		constAndroidInfoItems[idx] = &mobile.STEERiskMgrReportDataRequest_STEEInfoItem{
			Name:  &mappedKey,
			Value: &mappedValue,
		}
		idx++
	}
}

// RandomAndroidRequest is used to generate a android request data
func RandomAndroidRequest(appKey string) ([]byte, error) {
	reqData := randomAndroidRequestData()
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
		OsType:       &androidOsType,
		OsVer:        &androidOsVer,
		ProtoVersion: &protoVersion,
	}
	request := &mobile.STEERequest{
		Header: reqHeader,
		Type:   &reqType,
		Data:   reqBytes,
	}

	return proto.Marshal(request)
}

func randomAndroidRequestData() *mobile.STEERiskMgrReportDataRequest {
	return &mobile.STEERiskMgrReportDataRequest{
		Infos: randomAndroidDeviceInfo(),
	}
}

func randomAndroidDeviceInfo() []*mobile.STEERiskMgrReportDataRequest_STEEInfoItem {
	randomInfoItems := make([]*mobile.STEERiskMgrReportDataRequest_STEEInfoItem, len(randomItems))
	randomBoolInfoItems := make([]*mobile.STEERiskMgrReportDataRequest_STEEInfoItem, len(randomBoolItems))

	for idx, k := range randomItems {
		mappedKey := androidKeyMap[k]
		value := randomUUID()
		randomInfoItems[idx] = &mobile.STEERiskMgrReportDataRequest_STEEInfoItem{
			Name:  &mappedKey,
			Value: &value,
		}
	}

	for idx, k := range randomBoolItems {
		mappedKey := androidKeyMap[k]
		value := randomStrBoolean()
		randomBoolInfoItems[idx] = &mobile.STEERiskMgrReportDataRequest_STEEInfoItem{
			Name:  &mappedKey,
			Value: &value,
		}
	}
	newConstAndroidInfoItems:= append(constAndroidInfoItems, randomInfoItems...)
	return append(newConstAndroidInfoItems, randomBoolInfoItems...)
}
