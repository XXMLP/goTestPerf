package suite

import (
	"encoding/base64"
	"flag"
	"pef/datagen"
	"sort"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	engineSuite *EngineSuite
	getGen      = func(field string) func() interface{} {
		_, handler, _ := datagen.GetGenerator(field)
		return handler
	}
)

func init() {
	engineSuite = &EngineSuite{
		Host:      "127.0.0.1:7090",
		AppKey:    "testappid",
		AppSecret: "testappkey",
		EventCodeList: "test_event",
		EventCode: strings.Fields("test_event")[0],
		FieldsFn: map[string]func() interface{}{
			"ip":           getGen("ip"),
			"email":        getGen("email"),
			"phone_number": getGen("phone_number"),
			"const_id":     getGen("const_id"),
		},

		URLWithoutSign: "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=",
	}
}

func TestEngineSuiteForCredit(t *testing.T) {
	suite := &EngineSuite{
		Host:      "127.0.0.1:7090",
		AppCode:   "testappcode",
		AppKey:    "testappid",
		AppSecret: "testappkey",
		EventCodeList: "test_event",
		EventCode: strings.Fields("test_event")[0],
		FieldsFn: map[string]func() interface{}{
			"ip":           getGen("ip"),
			"email":        getGen("email"),
			"phone_number": getGen("phone_number"),
			"const_id":     getGen("const_id"),
			"ext_name":     getGen("ext_name:string_64"),
		},

		URLWithoutSign: "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=",
	}

	Convey("get sign of multi fields for credit", t, func() {
		sign := suite.getSign(map[string]interface{}{
			"ip":      "1.2.3.4",
			"email":   "a@b.c",
			"user_id": "hello_world",
		})
		So(sign, ShouldEqual, "0306fff69f1af0f005e0d297b8ffc4df")
	})

	Convey("get data of multi fields for credit", t, func() {
		data, err := suite.getData(map[string]interface{}{
			"ip":          "1.2.3.4",
			"androidRoot": true,
		})
		So(err, ShouldEqual, nil)
		dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
		size, err := base64.StdEncoding.Decode(dst, data)
		So(err, ShouldEqual, nil)
		So(size, ShouldEqual, 113)
	})
}

func TestEngineSuiteRandomFields(t *testing.T) {
	suite := &EngineSuite{
		Host:      "127.0.0.1:7090",
		AppKey:    "testappid",
		AppSecret: "testappkey",
		EventCodeList: "test_event",
		EventCode: strings.Fields("test_event")[0],
		FieldsFn: map[string]func() interface{}{
			"ip":           getGen("ip"),
			"email":        getGen("email"),
			"phone_number": getGen("phone_number"),
			"const_id":     getGen("const_id"),
			"ext_name":     getGen("ext_name:string_64"),
		},

		URLWithoutSign: "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=",
	}

	Convey("generate random Fields", t, func() {
		fields := suite.randomFields()
		So(len(fields["ext_name"].(string)), ShouldBeLessThanOrEqualTo, 64)
		So(len(fields["ext_name"].(string)), ShouldBeGreaterThan, 0)
		So(fields["ip"].(string), ShouldNotBeBlank)
		So(fields["email"].(string), ShouldNotBeBlank)
		So(fields["phone_number"].(string), ShouldNotBeBlank)
		So(fields["const_id"].(string), ShouldNotBeBlank)
	})
}

func TestEngineSuiteInit(t *testing.T) {
	suite := NewEngineSuite()

	Convey("init engine suite with correct random data", t, func() {
		err := suite.Init(flag.NewFlagSet("engine", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7090",
			"-app-key", "testappid",
			"-app-secret", "testappkey",
			"-event-codes", "test_event",
			"-field", "ip",
			"-field", "email",
			"-field", "ext_name:string",
			"-field", "ext_amount:float_5000",
		})

		So(err, ShouldBeNil)
		So(suite.Host, ShouldEqual, "127.0.0.1:7090")
		So(suite.AppKey, ShouldEqual, "testappid")
		So(suite.AppSecret, ShouldEqual, "testappkey")
		So(suite.EventCodeList, ShouldEqual, "test_event")
		keys := make([]string, 4)
		idx := 0
		for name := range suite.FieldsFn {
			keys[idx] = name
			idx++
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		So(keys, ShouldResemble, []string{"email", "ext_amount", "ext_name", "ip"})
		So(suite.URLWithoutSign, ShouldEqual, "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=")
	})

	Convey("init engine suite missing field", t, func() {
		So(NewEngineSuite().Init(flag.NewFlagSet("engine", flag.ContinueOnError), []string{
			"-app-key", "testappid",
			"-app-secret", "testappkey",
			"-event-codes", "test_event",
			"-field", "ip",
			"-field", "email",
		}), ShouldNotBeNil)
		So(NewEngineSuite().Init(flag.NewFlagSet("engine", flag.ContinueOnError), []string{
			"-host", "testappid",
			"-app-secret", "testappkey",
			"-event-codes", "test_event",
			"-field", "ip",
			"-field", "email",
		}), ShouldNotBeNil)
		So(NewEngineSuite().Init(flag.NewFlagSet("engine", flag.ContinueOnError), []string{
			"-host", "testappid",
			"-app-key", "testappid",
			"-app-secret", "testappkey",
			"-event-codes", "test_event",
		}), ShouldNotBeNil)
	})

	Convey("init engine suite with random data source", t, func() {
		err := suite.Init(flag.NewFlagSet("engine", flag.ContinueOnError), []string{
			"-host", "127.0.0.1:7090",
			"-app-key", "testappid",
			"-app-secret", "testappkey",
			"-event-codes", "test_event",
			"-data-source", "random",
			"-field", "ip",
		})

		So(err, ShouldBeNil)
	})
}

func TestGetSign(t *testing.T) {
	Convey("get sign of one field", t, func() {
		sign := engineSuite.getSign(map[string]interface{}{
			"ip": "1.2.3.4",
		})
		So(sign, ShouldEqual, "5db1434d41ba1d8029de5abdde40060b")
	})

	Convey("get sign of multi fields", t, func() {
		sign := engineSuite.getSign(map[string]interface{}{
			"ip":      "1.2.3.4",
			"email":   "a@b.c",
			"user_id": "hello_world",
		})
		So(sign, ShouldEqual, "2d4d5c7245ece671c359f454e565e223")
	})

	Convey("get sign of non-string fields", t, func() {
		sign := engineSuite.getSign(map[string]interface{}{
			"ip":          "1.2.3.4",
			"androidRoot": false,
			"ratio":       1,
		})
		So(sign, ShouldEqual, "996aa5fdffcbb8e5e0e832cbcbd19c44")
	})

	Convey("get sign of null fields", t, func() {
		sign := engineSuite.getSign(map[string]interface{}{
			"ip":          nil,
			"androidRoot": false,
			"ratio":       1,
		})
		So(sign, ShouldEqual, "bc11ed1d9dff21f03642ba9ee1684e52")
	})
}

func TestGetData(t *testing.T) {
	Convey("get data of single field", t, func() {
		data, err := engineSuite.getData(map[string]interface{}{
			"ip": "1.2.3.4",
		})
		So(err, ShouldEqual, nil)
		dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
		size, err := base64.StdEncoding.Decode(dst, data)
		So(err, ShouldEqual, nil)
		So(size, ShouldEqual, 70)
	})

	Convey("get data of multi field", t, func() {
		data, err := engineSuite.getData(map[string]interface{}{
			"ip":          "1.2.3.4",
			"androidRoot": true,
		})
		So(err, ShouldEqual, nil)
		dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
		size, err := base64.StdEncoding.Decode(dst, data)
		So(err, ShouldEqual, nil)
		So(size, ShouldEqual, 89)
	})
}

func TestGetURL(t *testing.T) {
	Convey("get url of fields", t, func() {
		url := engineSuite.getURL(map[string]interface{}{
			"ip": "1.2.3.4",
		})
		So(url, ShouldEqual, "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=5db1434d41ba1d8029de5abdde40060b")
	})
}

func TestRunOnceStartStopHandling(t *testing.T) {
	suite := &EngineSuite{
		Host:       "127.0.0.1:7090",
		AppCode:    "testappcode",
		AppKey:     "testappid",
		AppSecret:  "testappkey",
		EventCodeList:  "test_event",
		EventCode: strings.Fields("test_event")[0],
		DataSource: "test.data",
		RunOnce:    true,
		FieldsFn:   map[string]func() interface{}{},

		URLWithoutSign: "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=",
	}
	clearFieldsBuffer()

	Convey("rehatching run once task, the data count should be correct", t, func() {
		go suite.readFromDataSource()

		for idx := 0; idx < 5; idx++ {
			hatchEvt <- struct{}{}
			time.Sleep(100 * time.Millisecond)
			So(len(fieldsBuffer), ShouldEqual, 2)
			stopEvt <- struct{}{}
			clearFieldsBuffer()
			So(len(fieldsBuffer), ShouldEqual, 0)
		}
	})
}

// NOTE: this case should be after TestRunOnceStartStopHandling
func TestReadFromDataSourceLooped(t *testing.T) {
	suite := &EngineSuite{
		Host:       "127.0.0.1:7090",
		AppCode:    "testappcode",
		AppKey:     "testappid",
		AppSecret:  "testappkey",
		EventCodeList:  "test_event",
		EventCode: strings.Fields("test_event")[0],
		DataSource: "test.data",
		FieldsFn:   map[string]func() interface{}{},

		URLWithoutSign: "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=",
	}
	go suite.readFromDataSource()

	Convey("read data from test.data rotately", t, func() {
		So(suite.getFields(), ShouldResemble, map[string]interface{}{
			"ip": "1.1.1.1",
		})
		So(suite.getFields(), ShouldResemble, map[string]interface{}{
			"ip":    "2.2.2.2",
			"email": "a@b.c",
		})
		So(suite.getFields(), ShouldResemble, map[string]interface{}{
			"ip": "1.1.1.1",
		})
	})
}

func BenchmarkGetURL(b *testing.B) {
	fields := engineSuite.randomFields()
	for i := 0; i < b.N; i++ {
		engineSuite.getURL(fields)
	}
}

func BenchmarkGetData(b *testing.B) {
	fields := engineSuite.randomFields()
	for i := 0; i < b.N; i++ {
		engineSuite.getData(fields)
	}
}

func BenchmarkGetSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		engineSuite.getSign(map[string]interface{}{
			"ip":      "1.2.3.4",
			"email":   "a@b.c",
			"user_id": "hello_world",
		})
	}
}

func BenchmarkGenerateRandomFields(b *testing.B) {
	suite := &EngineSuite{
		Host:      "127.0.0.1:7090",
		AppKey:    "testappid",
		AppSecret: "testappkey",
		EventCodeList: "test_event",
		EventCode: strings.Fields("test_event")[0],
		FieldsFn: map[string]func() interface{}{
			"ip":           getGen("ip"),
			"email":        getGen("email"),
			"phone_number": getGen("phone_number"),
			"const_id":     getGen("const_id"),
			"ext_name":     getGen("ext_name:string_64"),
			"ext_amount":   getGen("ext_amount:float_5000"),
		},

		URLWithoutSign: "http://127.0.0.1:7090/ctu/event.do?appKey=testappid&version=1&sign=",
	}

	for i := 0; i < b.N; i++ {
		suite.randomFields()
	}
}
