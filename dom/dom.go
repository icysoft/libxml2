package dom

import (
	"sync"

	"github.com/icysoft/libxml2/xpath"
)

var docPool sync.Pool

func init() {
	SetupXPathCallback()
	docPool = sync.Pool{}
	docPool.New = func() interface{} {
		return Document{}
	}
}

func SetupXPathCallback() {
	xpath.WrapNodeFunc = WrapNode
}

func WrapDocument(n uintptr) *Document {
	doc := docPool.Get().(Document)
	doc.mortal = false
	doc.ptr = n
	return &doc
}
