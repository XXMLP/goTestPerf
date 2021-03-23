package suite

import (
	"flag"
	"github.com/feiyuw/boomer"
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

type newSuite struct{}

func (s *newSuite) Init(flagSet *flag.FlagSet, args []string) error {
	return nil
}

func (s *newSuite) GetTask() *boomer.Task {
	return nil
}

func TestSuiteMap(t *testing.T) {
	Convey("get exist suite", t, func() {
		suite, err := SuiteMap.Get("engine")
		So(err, ShouldEqual, nil)
		So(suite, ShouldHaveSameTypeAs, &EngineSuite{})
	})

	Convey("get unknown suite", t, func() {
		suite, err := SuiteMap.Get("does-not-exist")
		So(err, ShouldNotEqual, nil)
		So(suite, ShouldEqual, nil)
	})

	Convey("add new suite", t, func() {
		suite, err := SuiteMap.Get("new-not-exist")
		So(err, ShouldNotEqual, nil)
		So(suite, ShouldEqual, nil)
		SuiteMap.Add("new-not-exist", &newSuite{})
		suite, err = SuiteMap.Get("new-not-exist")
		So(err, ShouldBeNil)
		So(suite, ShouldNotBeNil)
	})

	Convey("list suite names", t, func() {
		m := suiteMap{}
		So(len(m.List()), ShouldEqual, 0)
		m.Add("new-not-exist", &newSuite{})
		So(m.List(), ShouldResemble, []string{"new-not-exist"})
		m.Add("new-not-exist2", &newSuite{})
		mlist := m.List()
		sort.SliceStable(mlist, func(i, j int) bool {
			return mlist[i] < mlist[j]
		})
		So(mlist, ShouldResemble, []string{"new-not-exist", "new-not-exist2"})
	})
}
