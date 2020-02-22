package goweb

import (
	"fmt"
	"testing"
)

func fakeHandler(c *Context) {
	fmt.Println("fakeHandler")
}

func TestTreeInsert(t *testing.T) {

	root := initTree()

	root.insert("/hello", HandlerChain{fakeHandler})
	root.insert("/he/ha/hb", HandlerChain{fakeHandler})
	root.insert("/he/:name", HandlerChain{fakeHandler})

	// ans := root.search("/hello")
	ans := root.search("/he/aaa")
	t.Log(ans)
}