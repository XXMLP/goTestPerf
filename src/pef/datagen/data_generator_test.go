package datagen

import (
	"fmt"
	"github.com/icrowley/fake"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGetGenerator(t *testing.T) {
	Convey("get defined generator", t, func() {
		name, handler, err := GetGenerator("ip")
		So(name, ShouldEqual, "ip")
		So(err, ShouldBeNil)
		So(handler().(string), ShouldNotBeBlank)
	})

	Convey("get dynamic string generator without max", t, func() {
		name, handler, err := GetGenerator("ext_name:string")
		So(name, ShouldEqual, "ext_name")
		So(err, ShouldBeNil)
		So(handler().(string), ShouldNotBeBlank)
	})

	Convey("get dynamic string generator with max", t, func() {
		name, handler, err := GetGenerator("ext_name:string_10")
		So(name, ShouldEqual, "ext_name")
		So(err, ShouldBeNil)
		So(handler().(string), ShouldNotBeBlank)
		So(len(handler().(string)), ShouldBeLessThanOrEqualTo, 10)
	})

	Convey("get dynamic num generator with max", t, func() {
		name, handler, err := GetGenerator("ext_number:num_10")
		So(name, ShouldEqual, "ext_number")
		So(err, ShouldBeNil)
		So(handler().(string), ShouldNotBeBlank)
		So(len(handler().(string)), ShouldBeLessThanOrEqualTo, 10)
	})

	Convey("get dynamic int generator with max", t, func() {
		name, handler, err := GetGenerator("ext_amount:int_10")
		So(name, ShouldEqual, "ext_amount")
		So(err, ShouldBeNil)
		So(handler().(int), ShouldBeLessThanOrEqualTo, 10)
		So(handler().(int), ShouldBeGreaterThanOrEqualTo, 0)
	})

	Convey("get dynamic float generator with max", t, func() {
		name, handler, err := GetGenerator("ext_float:float_10")
		So(name, ShouldEqual, "ext_float")
		So(err, ShouldBeNil)
		So(handler().(float32), ShouldBeLessThanOrEqualTo, 10)
		So(handler().(float32), ShouldBeGreaterThanOrEqualTo, 0)
	})

	Convey("get dynamic double generator with max", t, func() {
		name, handler, err := GetGenerator("ext_double:double_10")
		So(name, ShouldEqual, "ext_double")
		So(err, ShouldBeNil)
		So(handler().(float64), ShouldBeLessThanOrEqualTo, 10)
		So(handler().(float64), ShouldBeGreaterThanOrEqualTo, 0)
	})

	Convey("get dynamic uuid generator", t, func() {
		name, handler, err := GetGenerator("ext_uid:uuid")
		So(name, ShouldEqual, "ext_uid")
		So(err, ShouldBeNil)
		So(len(handler().(string)), ShouldEqual, 36)
	})

	Convey("get dynamic uuid generator", t, func() {
		name, handler, err := GetGenerator("ext_bool:bool")
		So(name, ShouldEqual, "ext_bool")
		So(err, ShouldBeNil)
		So(handler().(bool), ShouldBeIn, []bool{true, false})
	})
}

func TestRandomBoolean(t *testing.T) {
	Convey("generate random boolean", t, func() {
		So(randomBoolean().(bool), ShouldBeIn, []bool{true, false})
	})
}

func TestRandomConstID(t *testing.T) {
	Convey("generate random const_id", t, func() {
		cID := RandomConstID()
		So(len(cID.(string)), ShouldEqual, 40)
	})
}

func TestRandomPhoneNumber(t *testing.T) {
	Convey("generate random phone_number", t, func() {
		phone := randomPhoneNumber()
		So(len(phone.(string)), ShouldEqual, 11)
	})
}

func TestRandomIPv4(t *testing.T) {
	Convey("generate random ipv4", t, func() {
		ip := randomIPv4()
		So(len(strings.Split(ip.(string), ".")), ShouldEqual, 4)
	})
}

func TestRandomBytes(t *testing.T) {
	Convey("generate random size of bytes", t, func() {
		So(randomBytes(0), ShouldResemble, []byte{})
		So(randomBytes(-1), ShouldResemble, []byte{})
		So(len(randomBytes(1)), ShouldEqual, 1)
		So(len(randomBytes(13)), ShouldEqual, 13)
	})
}

func BenchmarkRandomConstId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomConstID()
	}
}

func BenchmarkRandomPhoneNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomPhoneNumber()
	}
}

func BenchmarkFakeWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fake.Word()
	}
}

func BenchmarkFakeUserName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fake.UserName()
	}
}

func BenchmarkRandomIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomIPv4()
	}
}

func BenchmarkFakeIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fake.IPv4()
	}
}

func BenchmarkSprintf(b *testing.B) {
	num := time.Now().Unix()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%0x", num)
	}
}

func BenchmarkStrconv(b *testing.B) {
	num := time.Now().Unix()
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(num, 16)
	}
}
