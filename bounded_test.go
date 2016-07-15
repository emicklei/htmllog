package htmllog

import (
	"testing"

	"github.com/emicklei/assert"
)

func TestWriteLimited(t *testing.T) {
	s := LimitedSprintf(4, "%s", "123456789")
	assert.That(t, "first 4", s).Equals("1234")
}

type node struct {
	parent *node
}

func TestRecursive(t *testing.T) {
	root := new(node)
	root.parent = root
	s := LimitedSprintf(40, "%#v", root)
	assert.That(t, "capped recursive print", s).Len(40)
}
