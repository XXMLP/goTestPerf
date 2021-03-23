package datagen

import (
	"pef/mobile"

	"github.com/golang/protobuf/proto"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRandomIOSDeviceInfo(t *testing.T) {
	Convey("random ios device info", t, func() {
		infos := randomIOSDeviceInfo()
		foundK26 := false
		foundK15 := false
		for _, info := range infos {
			if *info.Name == "K26" {
				foundK26 = true
				So(*info.Value, ShouldEqual, "10.156.153.174")
			} else if *info.Name == "K15" {
				foundK15 = true
				So(*info.Value, ShouldEqual, "Wifi")
			}
		}
		So(len(infos), ShouldEqual, 38)
		So(foundK26, ShouldBeTrue)
		So(foundK15, ShouldBeTrue)
		infos2 := randomIOSDeviceInfo()
		So(len(infos2), ShouldEqual, 38)
	})
}

func TestRandomIOSRequestData(t *testing.T) {
	Convey("generate random request data", t, func() {
		reqData := randomIOSRequestData()
		So(len(reqData.Infos), ShouldEqual, 38)
		foundK26 := false
		foundK15 := false
		for _, info := range reqData.Infos {
			if *info.Name == "K26" {
				foundK26 = true
				So(*info.Value, ShouldEqual, "10.156.153.174")
			} else if *info.Name == "K15" {
				foundK15 = true
				So(*info.Value, ShouldEqual, "Wifi")
			}
		}
		So(foundK26, ShouldBeTrue)
		So(foundK15, ShouldBeTrue)
	})

	Convey("random request data Marshal", t, func() {
		reqData := randomIOSRequestData()
		reqBytes, err := proto.Marshal(reqData)
		So(err, ShouldBeNil)
		So(len(reqBytes), ShouldBeGreaterThan, 0)
	})
}

func TestRandomIOSRequest(t *testing.T) {
	Convey("generate random android request", t, func() {
		reqBytes, err := RandomIOSRequest("abcdefg")
		So(err, ShouldBeNil)
		So(len(reqBytes), ShouldBeGreaterThan, 0)

		request := &mobile.STEERequest{}
		err = proto.Unmarshal(reqBytes, request)
		So(err, ShouldBeNil)
		So(request.GetType(), ShouldEqual, mobile.STEEDataType_DATATYPE_DO_REPORT_DATA)
		So(request.Header.GetOsType(), ShouldEqual, mobile.STEERequestHeader_iOS)
	})
}

func BenchmarkRandomIOSDeviceInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomIOSDeviceInfo()
	}
}
