// +build js

package router

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/gopherjs/gopherjs/js"
)

func TestFoo(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		So( js.Global.Get("window"), ShouldNotBeNil )
	})
}
