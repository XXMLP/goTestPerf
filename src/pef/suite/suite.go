package suite

import (
	"errors"
	"flag"

	"github.com/feiyuw/boomer"
)

// SuiteMap is global suites map
var SuiteMap = suiteMap{}

// Suite is interface used by all test suite
type Suite interface {
	Init(*flag.FlagSet, []string) error
	GetTask() *boomer.Task
}

type suiteMap map[string]Suite

func (m suiteMap) Add(name string, suite Suite) {
	m[name] = suite
}

func (m suiteMap) Get(name string) (Suite, error) {
	entry, exists := m[name]
	if !exists {
		return nil, errors.New("unknown suite " + name)
	}

	return entry, nil
}

func (m suiteMap) List() []string {
	names := make([]string, len(m))

	idx := 0
	for name := range m {
		names[idx] = name
		idx++
	}

	return names
}
