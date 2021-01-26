package copy

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestCopy(t *testing.T) {
	type student struct {
		Name string
	}
	stus := make([]student, 0, 1)
	stus = append(stus, student{Name: "ab"})
	stus1 := make([]student, 0, 1)
	Copy(&stus1, stus)
	spew.Dump(stus1)
}
