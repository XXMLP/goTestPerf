package flag

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListFlags(t *testing.T) {
	Convey("list flags string", t, func() {
		flags := ListFlags{}
		flags.Set("a")
		flags.Set("b")
		flags.Set("c")
		So(flags.String(), ShouldEqual, "a, b, c")
	})
}
