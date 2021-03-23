package suite

import (
	"flag"
	"testing"

	"github.com/miekg/dns"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDNSSuite(t *testing.T) {
	Convey("init dns suite with correct data", t, func() {
		suite := NewDNSSuite()

		err := suite.Init(flag.NewFlagSet("dns", flag.ContinueOnError), []string{
			"-protocol", "udp",
			"-host", "127.0.0.1:53",
			"-type", "A",
			"-record", "yum.dx.corp.",
			"-record", "pypi.dx.corp",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:53")
		So(suite.Protocol, ShouldEqual, "udp")
		So(suite.Type, ShouldEqual, dns.TypeA)
		So([]string(suite.Records), ShouldResemble, []string{"yum.dx.corp.", "pypi.dx.corp."})
	})

	Convey("init dns suite with default field", t, func() {
		suite := NewDNSSuite()

		err := suite.Init(flag.NewFlagSet("dns", flag.ContinueOnError), []string{
			"-host", "127.0.0.1",
			"-record", "yum.dx.corp",
			"-record", "pypi.dx.corp.",
		})
		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:53")
		So(suite.Protocol, ShouldEqual, "udp")
		So(suite.Type, ShouldEqual, dns.TypeA)
		So([]string(suite.Records), ShouldResemble, []string{"yum.dx.corp.", "pypi.dx.corp."})
	})

	Convey("init dns suite with missing record", t, func() {
		So(NewDNSSuite().Init(flag.NewFlagSet("dns", flag.ContinueOnError), []string{}), ShouldNotBeNil)
	})

	Convey("incorrect type", t, func() {
		So(NewDNSSuite().Init(flag.NewFlagSet("dns", flag.ContinueOnError), []string{
			"-type", "AA",
			"-record", "yum.dx.corp",
		}), ShouldNotBeNil)
	})

	Convey("incorrect protocol", t, func() {
		So(NewDNSSuite().Init(flag.NewFlagSet("dns", flag.ContinueOnError), []string{
			"-protocol", "sip",
			"-record", "yum.dx.corp",
		}), ShouldNotBeNil)
	})
}
