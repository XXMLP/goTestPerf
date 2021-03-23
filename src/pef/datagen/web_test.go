package datagen

import (
	"encoding/hex"
	"github.com/feiyuw/xxtea"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestRandomServerLid(t *testing.T) {
	Convey("random server lid length should be 72", t, func() {
		lid := randomServerLid()
		So(len(lid), ShouldEqual, 72)
		encryptedBytes, err := hex.DecodeString(lid)
		So(err, ShouldBeNil)
		src, err := xxtea.Decrypt(encryptedBytes, _Key)
		So(err, ShouldBeNil)
		So(len(src), ShouldEqual, 31)
		_, err = strconv.ParseInt(string(src[:8]), 16, 32)
		So(err, ShouldBeNil)
	})
}

func BenchmarkRandomServerLid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomServerLid()
	}
}
