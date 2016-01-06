// +build js

package router

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/gopherjs/gopherjs/js"
)

// This is an ugly hack to force jQuery initialization even earlier than init() would do it
// so that init() in other files will be run later
var _ = initjQuery()
func initjQuery() bool {
	cwd := js.Global.Get("process").Call("cwd").String()
	jsdom := js.Global.Call("require", cwd+"/node_modules/jsdom")
	doc := jsdom.Call("jsdom", "<html></html>")
	js.Global.Set("window", doc.Get("defaultView"))
	jQuery := js.Global.Call("require", cwd+"/node_modules/jquery")
	js.Global.Set("jQuery", jQuery)
	return true
}

func TestFoo(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		So( js.Global.Get("window"), ShouldNotBeNil )
	})
}
