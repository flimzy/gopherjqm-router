// +build js

package router

import (
	"sync"
	"log"
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/gopherjs/gopherjs/js"
)

// This is an ugly hack to force jQuery initialization even earlier than init() would do it
// so that init() in other files will be run later
var _ = initjQuery()
func initjQuery() bool {
	var wg sync.WaitGroup
	wg.Add(1)
	fs := js.Global.Call("require", "fs")
	fs.Call("readFile", "test/index.html", func(_ *js.Object, data *js.Object) {
		defer wg.Done()
		jsdom := js.Global.Call("require", "jsdom")
		doc := jsdom.Call("jsdom", data)
		js.Global.Set("window", doc.Get("defaultView"))
		jQuery := js.Global.Call("require", "jquery")
		js.Global.Set("jQuery", jQuery)
	})
	log.Print("Waiting")
	wg.Wait()
	log.Print("Done waiting")
	return true
}

func TestFoo(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		So( js.Global.Get("window"), ShouldNotBeNil )
	})
	// Close the window and terminate the program
	js.Global.Get("window").Call("close")
}
