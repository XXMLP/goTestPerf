package suite

import (
	"flag"
	json "github.com/json-iterator/go"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCaptchaSuiteInit(t *testing.T) {
	suite := NewCaptchaSuite()

	Convey("init captcha suite without iface and jsv set", t, func() {
		err := suite.Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-ak", "testappid",
			"-c", "test_constid",
			"-ac", "xxxac",
			"-x", "33",
			"-y", "44",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "all")
		So(suite.Jsv, ShouldEqual, "1.3.9.91")
		So(suite.Ak, ShouldEqual, "testappid")
		So(suite.C, ShouldEqual, "test_constid")
		So(suite.Ac, ShouldEqual, "xxxac")
		So(suite.X, ShouldEqual, "33")
		So(suite.Y, ShouldEqual, "44")
		So(suite.AURLPrefix, ShouldEqual, "http://127.0.0.1:7776/api/a?jsv=1.3.9.91&c=test_constid&ak=testappid&s=50&h=150&w=300&aid=")
	})

	Convey("init captcha suite with iface as a and jsv set", t, func() {
		err := suite.Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "a",
			"-ak", "testappid",
			"-jsv", "1.2.26",
			"-c", "test_constid",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "a")
		So(suite.Jsv, ShouldEqual, "1.2.26")
		So(suite.Ak, ShouldEqual, "testappid")
		So(suite.C, ShouldEqual, "test_constid")
		So(suite.Ac, ShouldEqual, defaultAC)
		So(suite.X, ShouldEqual, defaultX)
		So(suite.Y, ShouldEqual, defaultY)
	})

	Convey("init captcha suite with iface as v1 and sid set", t, func() {
		err := suite.Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "v1",
			"-ak", "testappid",
			"-jsv", "1.2.26",
			"-c", "test_constid",
			"-sid", "test_sid",
		})

		So(err, ShouldBeNil)

		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "v1")
		So(suite.Jsv, ShouldEqual, "1.2.26")
		So(suite.Ak, ShouldEqual, "testappid")
		So(suite.C, ShouldEqual, "test_constid")
		So(suite.Sid, ShouldEqual, "test_sid")
		So(suite.Ac, ShouldEqual, defaultAC)
		So(suite.X, ShouldEqual, defaultX)
		So(suite.Y, ShouldEqual, defaultY)
	})

	Convey("init captcha suite with iface as tokenVerify", t, func() {
		err := suite.Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "tokenVerify",
			"-ak", "testappid",
			"-jsv", "1.2.26",
			"-c", "test_constid",
			"-sid", "test_sid",
		})

		So(err, ShouldBeNil)

		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "tokenVerify")
		So(suite.Jsv, ShouldEqual, "1.2.26")
		So(suite.Ak, ShouldEqual, "testappid")
		So(suite.C, ShouldEqual, "test_constid")
		So(suite.Sid, ShouldEqual, "test_sid")
		So(suite.Ac, ShouldEqual, defaultAC)
		So(suite.X, ShouldEqual, defaultX)
		So(suite.Y, ShouldEqual, defaultY)
	})

	Convey("init captcha suite missing field", t, func() {
		So(NewCaptchaSuite().Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-ak", "testappid",
			"-c", "test_constid",
		}), ShouldNotBeNil)
		So(NewCaptchaSuite().Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-c", "test_constid",
		}), ShouldNotBeNil)
		So(NewCaptchaSuite().Init(flag.NewFlagSet("captcha_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-ak", "testappid",
		}), ShouldNotBeNil)
	})
}

func TestCaptchaSuiteDoA(t *testing.T) {
	Convey("do a response decode", t, func() {
		data := `{"sid":"8fdb67cd4bf0c7742249db0459802567","y":37,"success":true,"p1":"/dx/vvXKj3SpFi/zib3/911f6790ee994969b419ac939dc749d4.webp","p2":"/dx/vvXKj3SpFi/zib3/7442e71dda6f4edca69ab6b7223d8d22.webp","p3":"/dx/vvXKj3SpFi/zib3/79684c5091b349069115ecad08467810.webp","msg":null,"t":null,"result":1,"type":0,"logo":null}`
		resp := aResponse{}
		err := json.Unmarshal([]byte(data), &resp)
		So(err, ShouldEqual, nil)
		So(resp.Sid, ShouldEqual, "8fdb67cd4bf0c7742249db0459802567")
	})
}

func BenchmarkRandomAID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomAID()
	}
}
