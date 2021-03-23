package suite

import (
	"flag"
	"github.com/feiyuw/xxtea"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConstIDSuiteInit(t *testing.T) {
	suite := NewConstIDSuite()

	Convey("init constid suite with default type", t, func() {
		err := suite.Init(flag.NewFlagSet("constid", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-app-key", "90b3a12f68187787bc8d52cf7e7b6366",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "token")
		So(suite.AppKey, ShouldEqual, "90b3a12f68187787bc8d52cf7e7b6366")
		So(suite.Type, ShouldEqual, "web")
		So(suite.C1URL, ShouldEqual, "http://127.0.0.1:7776/udid/c1")
		So(suite.M1URL, ShouldEqual, "http://127.0.0.1:7776/udid/m1?appKey=90b3a12f68187787bc8d52cf7e7b6366&version=5.5.0&sign=")
		So(suite.SecretKey, ShouldResemble, []byte("84c72b020844588f4e89043d150ef5ba0f808606"))
	})

	Convey("init constid suite with specified type", t, func() {
		err := suite.Init(flag.NewFlagSet("constid", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "all",
			"-app-key", "90b3a12f68187787bc8d52cf7e7b6366",
			"-app-secret", "b2a8a90190fff591bd93bfd99e268438",
			"-type", "android",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "all")
		So(suite.AppKey, ShouldEqual, "90b3a12f68187787bc8d52cf7e7b6366")
		So(suite.Type, ShouldEqual, "android")
		So(suite.C1URL, ShouldEqual, "http://127.0.0.1:7776/udid/c1")
		So(suite.M1URL, ShouldEqual, "http://127.0.0.1:7776/udid/m1?appKey=90b3a12f68187787bc8d52cf7e7b6366&version=5.5.0&sign=")
		So(suite.SecretKey, ShouldResemble, []byte("84c72b020844588f4e89043d150ef5ba0f808606"))

		data, err := xxtea.Encrypt([]byte("abc"), []byte("84c72b020844588f4e89043d150ef5ba0f808606"))
		So(err, ShouldBeNil)
		So(data, ShouldResemble, []byte{256 - 5, 93, 256 - 34, 256 - 111, 256 - 9, 256 - 70, 256 - 76, 256 - 14})
	})
}

func TestBase64Encode(t *testing.T) {
	Convey("custom base64 encode", t, func() {
		So(base64Encoder.EncodeToString([]byte("abcde")), ShouldEqual, "GvJCl9o=")
	})
}

func TestConstIDGetM1Sign(t *testing.T) {
	suite := &ConstIDSuite{AppKey: "abcde"}

	Convey("sign for empty data is empty", t, func() {
		So(suite.getM1Sign([]byte{}), ShouldEqual, "")
	})

	Convey("sign for normal data is 16bytes", t, func() {
		So(suite.getM1Sign([]byte("defghijklmn")), ShouldEqual, "a6d171140d24558b2ccf0b07912917cd")

		So(suite.getM1Sign([]byte("def")), ShouldEqual, "fe533a2bab895236a5aaf6f8f0e7eaf0")
	})
}

func TestConstIDGetDeviceInfoSign(t *testing.T) {
	Convey("generate sign for token", t, func() {
		suite := &ConstIDSuite{AppSecret: "b2a8a90190fff591bd93bfd99e268438"}
		So(suite.getDeviceInfoSign("5d10736cb2zjBzFjNY6WSTYRueOT13nvtroq8aw1"), ShouldEqual, "2a8ff3b4b66908ea75c30e2ffe3f9103")
	})
}

func TestConstIDCompress(t *testing.T) {
	suite := &ConstIDSuite{AppKey: "abcde"}

	Convey("compress request data", t, func() {
		So(suite.compress([]byte("abcdefghijklmn")), ShouldResemble, []byte{80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 100, 97, 116, 97, 74, 76, 74, 78, 73, 77, 75, 207, 200, 204, 202, 206, 201, 205, 3, 4, 0, 0, 255, 255, 80, 75, 7, 8, 120, 149, 13, 64, 20, 0, 0, 0, 14, 0, 0, 0, 80, 75, 1, 2, 20, 0, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 120, 149, 13, 64, 20, 0, 0, 0, 14, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 97, 116, 97, 80, 75, 5, 6, 0, 0, 0, 0, 1, 0, 1, 0, 50, 0, 0, 0, 70, 0, 0, 0, 0, 0})
	})
}

func TestConstIDUncompress(t *testing.T) {
	suite := &ConstIDSuite{AppKey: "abcde"}

	Convey("uncompress request data", t, func() {
		data, err := suite.uncompress([]byte{80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 100, 97, 116, 97, 74, 76, 74, 78, 73, 77, 75, 207, 200, 204, 202, 206, 201, 205, 3, 4, 0, 0, 255, 255, 80, 75, 7, 8, 120, 149, 13, 64, 20, 0, 0, 0, 14, 0, 0, 0, 80, 75, 1, 2, 20, 0, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 120, 149, 13, 64, 20, 0, 0, 0, 14, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 97, 116, 97, 80, 75, 5, 6, 0, 0, 0, 0, 1, 0, 1, 0, 50, 0, 0, 0, 70, 0, 0, 0, 0, 0})
		So(err, ShouldBeNil)
		So(data, ShouldResemble, []byte("abcdefghijklmn"))
	})
}

func TestParseResponse(t *testing.T) {
	suite := &ConstIDSuite{AppKey: "90b3a12f68187787bc8d52cf7e7b6366"}
	suite.SecretKey = generateSecretKey("90b3a12f68187787bc8d52cf7e7b6366")

	Convey("parse 200 response", t, func() {
		data := []byte{
			0x68, 0x37, 0xef, 0xbf, 0xd5, 0xe3, 0x70, 0x0d,
			0xf1, 0x03, 0x96, 0xde, 0x67, 0x3b, 0xb5, 0x57,
			0xcd, 0xcf, 0x7a, 0x47, 0xb5, 0xe5, 0x41, 0x2f,
			0xb0, 0xac, 0x2a, 0xf6, 0x11, 0x82, 0x72, 0x5d,
			0x72, 0x6b, 0xb4, 0x8b, 0x62, 0xda, 0x47, 0x2a,
			0x2b, 0x7c, 0x2e, 0xf8, 0x3a, 0x22, 0xa1, 0x65,
			0x80, 0x5f, 0x4c, 0xbb, 0xd8, 0x43, 0x88, 0x50,
			0xbb, 0x44, 0xbe, 0x9a, 0x6d, 0xf1, 0x0c, 0x4c,
			0x5c, 0x4b, 0x2e, 0xc2, 0xda, 0x72, 0xe7, 0x86,
			0x1a, 0x3f, 0xc7, 0x48, 0x81, 0x2c, 0x73, 0x60,
			0xda, 0x78, 0x40, 0x9b, 0x9f, 0x83, 0xf9, 0xa8,
			0x2d, 0xca, 0x67, 0x7c, 0x52, 0x12, 0xa1, 0x24,
			0xbc, 0x1f, 0x93, 0x2b, 0xc2, 0x3c, 0xfe, 0xd6,
			0x91, 0xd8, 0xb1, 0x23, 0x77, 0x4b, 0x7e, 0x25,
			0x00, 0xb1, 0x25, 0x1c, 0x08, 0x3f, 0x0b, 0xda,
			0xbd, 0x1b, 0xbf, 0x31, 0xa0, 0x51, 0x2d, 0x82,
			0x2d, 0xc6, 0x8f, 0xd1, 0x4e, 0x7d, 0x54, 0x1d,
			0x75, 0xc6, 0xfa, 0xf0, 0x6f, 0x01, 0x0c, 0x60,
			0x82, 0xde, 0xa4, 0xab, 0xe7, 0x72, 0xf6, 0x60,
			0x05, 0x75, 0x2c, 0xa6, 0xb7, 0xe8, 0x89, 0x1e,
			0x92, 0xf9, 0x2f, 0x15, 0x13, 0x88, 0xe5, 0x95,
			0xef, 0x4e, 0x75, 0xc6, 0x06, 0x31, 0x3b, 0xe6,
			0x99, 0xb6, 0x2b, 0x71}
		constID, err := suite.parseResponse(data)
		So(err, ShouldBeNil)
		So(constID, ShouldEqual, "5c32b143lBxSxqTSQ2p1WRxyTXUYR4khFtFEGGm3")
	})

	Convey("parse check sign error response", t, func() {
		data := []byte{
			0x7b, 0x7d, 0x55, 0x66, 0xe1, 0xaa, 0x47, 0x20,
			0x63, 0xb6, 0xdc, 0x70, 0x60, 0x7a, 0x9e, 0x73,
			0xbe, 0x35, 0xe7, 0xa1, 0xdb, 0xee, 0x10, 0x49,
			0xe5, 0xaa, 0x97, 0x32, 0xef, 0x6a, 0x73, 0x09,
			0xc7, 0xe6, 0x20, 0x28, 0xf5, 0x02, 0xd2, 0x37,
			0x40, 0x2e, 0x5f, 0xa4, 0xe8, 0x2f, 0x61, 0x0f,
			0x51, 0x48, 0x97, 0x1e, 0xf5, 0xfd, 0xc6, 0x74,
			0xda, 0x1c, 0x5c, 0x1c, 0xa1, 0xeb, 0x82, 0xce,
			0xeb, 0x8f, 0x10, 0x9a, 0x6f, 0x06, 0xa2, 0x8f,
			0x3a, 0xec, 0xe8, 0x74, 0x75, 0x6c, 0x55, 0xe4,
			0xf4, 0x83, 0x7e, 0x00, 0xa5, 0xff, 0x93, 0x42,
			0x5a, 0x55, 0x34, 0x03, 0x97, 0x4f, 0x3e, 0x0a,
			0x6d, 0xbb, 0xfa, 0x98, 0x72, 0x20, 0x8c, 0x38,
			0x31, 0xcb, 0x8c, 0x33, 0x0d, 0x49, 0x68, 0x7e,
			0x9b, 0xa4, 0x94, 0x5d, 0xe1, 0x03, 0x0e, 0xfa,
			0x46, 0xd2, 0x4c, 0x80, 0x64, 0x3f, 0x56, 0x56,
			0x9e, 0x66, 0x3c, 0x6d, 0xc7, 0x7f, 0x9b, 0x82,
			0x1a, 0x17, 0x06, 0xc8}
		constID, err := suite.parseResponse(data)
		So(err.Error(), ShouldEqual, "-10002")
		So(constID, ShouldEqual, "")
	})
}

func TestRandomRequestType(t *testing.T) {
	suite := &ConstIDSuite{AppKey: "abcde"}
	Convey("randomType should be one of web|android|ios|miniProgram", t, func() {
		reqType := suite.randomType()
		So(reqType == "web" || reqType == "ios" || reqType == "android" || reqType == "miniProgram", ShouldBeTrue)
	})
}

func BenchmarkConstIDGetM1Sign(b *testing.B) {
	suite := &ConstIDSuite{AppKey: "abcde"}
	data := []byte("defghijklmn")

	for i := 0; i < b.N; i++ {
		suite.getM1Sign(data)
	}
}

func BenchmarkConstIDGetSecretKey(b *testing.B) {
	appKey := "90b3a12f68187787bc8d52cf7e7b6366"
	for i := 0; i < b.N; i++ {
		generateSecretKey(appKey)
	}
}

func BenchmarkCompressZip(b *testing.B) {
	suite := &ConstIDSuite{}
	data := []byte("abcdefghijk")

	for i := 0; i < b.N; i++ {
		suite.compress(data)
	}
}

func BenchmarkUncompressZip(b *testing.B) {
	suite := &ConstIDSuite{}
	data := []byte{80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 100, 97, 116, 97, 74, 76, 74, 78, 73, 77, 75, 207, 200, 204, 202, 206, 201, 205, 3, 4, 0, 0, 255, 255, 80, 75, 7, 8, 120, 149, 13, 64, 20, 0, 0, 0, 14, 0, 0, 0, 80, 75, 1, 2, 20, 0, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 120, 149, 13, 64, 20, 0, 0, 0, 14, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 97, 116, 97, 80, 75, 5, 6, 0, 0, 0, 0, 1, 0, 1, 0, 50, 0, 0, 0, 70, 0, 0, 0, 0, 0}

	for i := 0; i < b.N; i++ {
		suite.uncompress(data)
	}
}

func BenchmarkGetDeviceInfoPostData(b *testing.B) {
	suite := &ConstIDSuite{AppKey: "05622f1ab6be69567d65a6e377edfef0", AppSecret: "b2a8a90190fff591bd93bfd99e268438"}
	token := "5d10736cb2zjBzFjNY6WSTYRueOT13nvtroq8aw1"

	for i := 0; i < b.N; i++ {
		suite.getDeviceInfoPostData(token)
	}
}

func BenchmarkConstIDParseResponse(b *testing.B) {
	suite := &ConstIDSuite{AppKey: "90b3a12f68187787bc8d52cf7e7b6366"}
	suite.SecretKey = generateSecretKey("90b3a12f68187787bc8d52cf7e7b6366")

	data := []byte{
		0x68, 0x37, 0xef, 0xbf, 0xd5, 0xe3, 0x70, 0x0d,
		0xf1, 0x03, 0x96, 0xde, 0x67, 0x3b, 0xb5, 0x57,
		0xcd, 0xcf, 0x7a, 0x47, 0xb5, 0xe5, 0x41, 0x2f,
		0xb0, 0xac, 0x2a, 0xf6, 0x11, 0x82, 0x72, 0x5d,
		0x72, 0x6b, 0xb4, 0x8b, 0x62, 0xda, 0x47, 0x2a,
		0x2b, 0x7c, 0x2e, 0xf8, 0x3a, 0x22, 0xa1, 0x65,
		0x80, 0x5f, 0x4c, 0xbb, 0xd8, 0x43, 0x88, 0x50,
		0xbb, 0x44, 0xbe, 0x9a, 0x6d, 0xf1, 0x0c, 0x4c,
		0x5c, 0x4b, 0x2e, 0xc2, 0xda, 0x72, 0xe7, 0x86,
		0x1a, 0x3f, 0xc7, 0x48, 0x81, 0x2c, 0x73, 0x60,
		0xda, 0x78, 0x40, 0x9b, 0x9f, 0x83, 0xf9, 0xa8,
		0x2d, 0xca, 0x67, 0x7c, 0x52, 0x12, 0xa1, 0x24,
		0xbc, 0x1f, 0x93, 0x2b, 0xc2, 0x3c, 0xfe, 0xd6,
		0x91, 0xd8, 0xb1, 0x23, 0x77, 0x4b, 0x7e, 0x25,
		0x00, 0xb1, 0x25, 0x1c, 0x08, 0x3f, 0x0b, 0xda,
		0xbd, 0x1b, 0xbf, 0x31, 0xa0, 0x51, 0x2d, 0x82,
		0x2d, 0xc6, 0x8f, 0xd1, 0x4e, 0x7d, 0x54, 0x1d,
		0x75, 0xc6, 0xfa, 0xf0, 0x6f, 0x01, 0x0c, 0x60,
		0x82, 0xde, 0xa4, 0xab, 0xe7, 0x72, 0xf6, 0x60,
		0x05, 0x75, 0x2c, 0xa6, 0xb7, 0xe8, 0x89, 0x1e,
		0x92, 0xf9, 0x2f, 0x15, 0x13, 0x88, 0xe5, 0x95,
		0xef, 0x4e, 0x75, 0xc6, 0x06, 0x31, 0x3b, 0xe6,
		0x99, 0xb6, 0x2b, 0x71}

	for i := 0; i < b.N; i++ {
		suite.parseResponse(data)
	}
}

func TestConstIDDubboSuite(t *testing.T) {
	Convey("init constid dubbo suite", t, func() {
		suite := NewConstIDDubboSuite()
		err := suite.Init(flag.NewFlagSet("constid-dubbo", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
		})
		So(err, ShouldBeNil)
		err = suite.Init(flag.NewFlagSet("constid-dubbo", flag.ContinueOnError), []string{})
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "host should be set")
	})
}
