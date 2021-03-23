package datagen

import (
	"pef/mobile"
	"testing"

	"github.com/golang/protobuf/proto"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomAndroidDeviceInfo(t *testing.T) {
	Convey("random android device info", t, func() {
		infos := randomAndroidDeviceInfo()
		foundK11 := false
		foundK77 := false
		for _, info := range infos {
			if *info.Name == "K11" {
				foundK11 = true
				So(*info.Value, ShouldEqual, "8.0.0")
			} else if *info.Name == "K77" {
				foundK77 = true
				So(*info.Value, ShouldEqual, `["dummy0-76:f7:57:24:f2:36","wlan0-a0:cc:2b:bf:90:85","p2p0-a2:cc:2b:bf:90:85","bond0-b6:2d:99:ac:49:be"]`)
			}
		}
		So(foundK11, ShouldBeTrue)
		So(foundK77, ShouldBeTrue)
		So(len(infos), ShouldEqual, 54)
		infos2 := randomAndroidDeviceInfo()
		So(len(infos2), ShouldEqual, 54)
	})
}

func TestRandomAndroidRequestData(t *testing.T) {
	Convey("generate random request data", t, func() {
		reqData := randomAndroidRequestData()
		So(len(reqData.Infos), ShouldEqual, 54)
		foundK11 := false
		foundK77 := false
		for _, info := range reqData.Infos {
			if *info.Name == "K11" {
				foundK11 = true
				So(*info.Value, ShouldEqual, "8.0.0")
			} else if *info.Name == "K77" {
				foundK77 = true
				So(*info.Value, ShouldEqual, `["dummy0-76:f7:57:24:f2:36","wlan0-a0:cc:2b:bf:90:85","p2p0-a2:cc:2b:bf:90:85","bond0-b6:2d:99:ac:49:be"]`)
			}
		}
		So(foundK11, ShouldBeTrue)
		So(foundK77, ShouldBeTrue)
	})

	Convey("random request data Marshal", t, func() {
		reqData := randomAndroidRequestData()
		reqBytes, err := proto.Marshal(reqData)
		So(err, ShouldBeNil)
		So(len(reqBytes), ShouldBeGreaterThan, 0)
	})
}

func TestRandomAndroidRequest(t *testing.T) {
	Convey("generate random android request", t, func() {
		reqBytes, err := RandomAndroidRequest("abcdefg")
		So(err, ShouldBeNil)
		So(len(reqBytes), ShouldBeGreaterThan, 0)

		request := &mobile.STEERequest{}
		err = proto.Unmarshal(reqBytes, request)
		So(err, ShouldBeNil)
		So(request.GetType(), ShouldEqual, mobile.STEEDataType_DATATYPE_DO_REPORT_DATA)
		So(request.Header.GetOsType(), ShouldEqual, mobile.STEERequestHeader_Android)
	})
}

func BenchmarkRandomAndroidDeviceInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomAndroidDeviceInfo()
	}
}
