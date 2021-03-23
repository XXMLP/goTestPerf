package suite

import (
	"errors"
	"flag"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneralSuiteInit(t *testing.T) {
	suite := NewGeneralSuite()

	Convey("init general suite with error data-source path", t, func() {
		err := suite.Init(flag.NewFlagSet("general", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:8080",
			"-api",   "/api/a",
			"-data-source", "error.data",
		})
		So(err, ShouldResemble, errors.New("datasource does not exist"))
	})

	Convey("init general suite missing field", t, func() {
		So(NewGeneralSuite().Init(flag.NewFlagSet("general", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:8080",
			"-api", "/api/a",
			}), ShouldResemble, errors.New("host, api, data-source should be set"))
		So(NewGeneralSuite().Init(flag.NewFlagSet("general", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:8080",
			"-data-source", "test.data",
		}), ShouldResemble, errors.New("host, api, data-source should be set"))
		So(NewGeneralSuite().Init(flag.NewFlagSet("general", flag.ContinueOnError), []string{
			"-api", "/api/a",
			"-data-source", "test.data",
		}), ShouldResemble, errors.New("host, api, data-source should be set"))
	})
}

func TestGeneralRunOnceStartStopHandling(t *testing.T) {
	suite := &GeneralSuite{
		Host:       "127.0.0.1:8080",
		DataSource: "test.data",
		RunOnce:    true,
		Api: "/api/a",
	}
	generalClearFieldsBuffer()

	Convey("rehatching run once task, the data count should be correct", t, func() {
		go suite.readFromDataSource()
		for idx := 0; idx < 5; idx++ {
			generalhatchEvt <- struct{}{}
			time.Sleep(100 * time.Millisecond)
			So(len(generalfieldsBuffer), ShouldEqual, 2)
			generalstopEvt <- struct{}{}
			generalClearFieldsBuffer()
			So(len(generalfieldsBuffer), ShouldEqual, 0)
		}
	})
}

// NOTE: this case should be after TestGeneralRunOnceStartStopHandling
func TestGeneralReadFromDataSourceLooped(t *testing.T) {
	suite := &GeneralSuite{
		Host:       "127.0.0.1:8080",
		DataSource: "test.data",
		Api: "/api/a",
	}
	go suite.readFromDataSource()

	Convey("read data from test.data rotately", t, func() {
		So(<-generalfieldsBuffer, ShouldResemble, map[string]interface{}{
			"ip": "1.1.1.1",
		})
		So(<-generalfieldsBuffer, ShouldResemble, map[string]interface{}{
			"ip":    "2.2.2.2",
			"email": "a@b.c",
		})
		So(<-generalfieldsBuffer, ShouldResemble, map[string]interface{}{
			"ip": "1.1.1.1",
		})
	})
}