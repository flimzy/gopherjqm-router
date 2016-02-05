// +build js

package router

import (
// 	"errors"

	"github.com/armon/go-radix"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

const dataKey = "gopherjqm-router"

type Handler func(*js.Object, *jquery.Event, *js.Object, Params) bool

type Router struct {
	element *jquery.JQuery
	trees   map[string]*radix.Tree
}

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// New creates a new Router object, and attaches it to the requested element
func New(elem *jquery.JQuery) *Router {
	r := &Router{
		element: elem,
	}
	elem.SetData(dataKey, r)
	return r
}

func (r *Router) Handle(path, event string, handler Handler) error {
	if r.trees == nil {
		r.trees = make(map[string]*radix.Tree)
	}

	root := r.trees[event]
	if root == nil {
		root = radix.New()
		r.trees[event] = root
		r.element.On("pagecontainer" + event, func(e *jquery.Event, ui *js.Object) bool {
			return r.handleRequest(event, e, ui)
		})
	}

	root.Insert(path, handler)
	return nil
}

func (r *Router) handleRequest(eventName string, eventObj *jquery.Event, ui *js.Object) bool {
	// Find and execute handler
	return true // To continue processing any other event handlers
}
