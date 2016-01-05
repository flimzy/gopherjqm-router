// +build js

package router

// This file contains jQuery wrappers for the standard API

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

// convenience
var jQuery = jquery.NewJQuery
var document *js.Object = js.Global.Get("document")

func init() {
	// Register as a jQuery plugin, for the benefit of JavaScript code
	js.Global.Get("jQuery").Get("fn").Set("router", map[string]interface{}{
		"handle": js.MakeFunc(jsHandle),
	})
}

func jsHandle(this *js.Object, arguments []*js.Object) interface{} {
	elem := jQuery(this)
	r := GetRouter(&elem)
	if r == nil {
		r = New(&elem)
	}
	path := arguments[0].String()
	eventName := arguments[1].String()
	handler := arguments[2]
	return r.Handle(path, eventName, func(element *js.Object, event *jquery.Event, ui *js.Object, p Params) bool {
		// Set 'this' to the element
		return handler.Call("call", element, event, ui, p).Bool()
	})
}

// GetRouter fetches a router object from the specified jQuery DOM
// element. If no router is set, nil is returned.
func GetRouter(elem *jquery.JQuery) *Router {
	return elem.Data(dataKey).(*Router)
}
