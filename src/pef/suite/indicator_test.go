package suite

import (
	"flag"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	peflag "pef/flag"
)

func TestIndicatorSuiteInit(t *testing.T) {
	Convey("init indicator suite with multi indicator code and fields", t, func() {
		suite := NewIndicatorSuite()
		err := suite.Init(flag.NewFlagSet("indicator_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-dc", "ds-1",
			"-ic", "ic_1",
			"-ic", "ic_2",
			"-field", "ip",
			"-field", "email",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:7776")
		So(suite.Iface, ShouldEqual, "all")
		So(suite.DataSourceCode, ShouldEqual, "ds-1")
		So(suite.IndicatorCodeSet, ShouldResemble, peflag.ListFlags{"ic_1", "ic_2"})
		So(len(suite.FieldsFn), ShouldEqual, 2)
	})

	Convey("init indicator processMessage suite with no codeset", t, func() {
		suite := NewIndicatorSuite()
		err := suite.Init(flag.NewFlagSet("indicator_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "processMessage",
			"-dc", "ds-1",
			"-field", "email",
		})
		So(err, ShouldBeNil)
		So(suite.Iface, ShouldEqual, "processMessage")
		So(len(suite.IndicatorCodeSet), ShouldEqual, 0)
	})

	Convey("init indicator mget suite with fail on no codeset", t, func() {
		suite := NewIndicatorSuite()
		err := suite.Init(flag.NewFlagSet("indicator_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "mget",
			"-dc", "ds-1",
			"-field", "email",
		})
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "indicator codeset is mandentary when iface is mget or all")
	})

	Convey("init indicator suite with invalid field", t, func() {
		suite := NewIndicatorSuite()
		err := suite.Init(flag.NewFlagSet("indicator_test", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7776",
			"-iface", "processMessage",
			"-dc", "ds-1",
			"-field", "invalid_ip",
			"-field", "email",
		})
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "invalid_ip is not supported!")
	})
}

func BenchmarkCurrentTimestampMills(b *testing.B) {
	for i := 0; i < b.N; i++ {
		currentTimestampMills()
	}
}
