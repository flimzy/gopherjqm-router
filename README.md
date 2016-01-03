# gopherjqm-router

Router for [jQuery Mobile](https://jquerymobile.com/) applications written for [GopherJS](http://www.gopherjs.org/) (Go compiled to JavaScript).

## Why?

Several routers exist for single-page web applications.  There's at least one for GopherJS, and there's at least one for jQuery Mobile, written in JavaScript.  But there are none (to my knowledge) with bridge both worlds.

## Design goals

The specific goals I intend this package to acheive are:

* hashchange-driven router for jQuery Mobile apps. [jquerymobile-router](https://github.com/azicchetti/jquerymobile-router) does this, but only for JavaScript code, and only for old versions of jQuery Mobile.
* The router should be written in Go. [humble-router](https://github.com/go-humble/router) does this, but does not support jQuery Mobile apps.
* All source (.go and .js files) should be compilable to a single .js bundle. This plays nicely with both jQuery Mobile and GopherJS.
* The router should support handlers written in either Go or JavaScript.

The latter point is specifically to bridge two realities: There are people like me, who hate writing JavaScript code, and there is the real world, where skilled web designers and front-end developers write code in JavaScript.

In "traditional" web apps, it's easy for back-end developers to user thier language of choice, and leave client-side JavaScript coding to the front-end guys.  In a [Single Page Application](https://en.wikipedia.org/wiki/Single-page_application) model, such as that provided by jQuery Mobile, no such division of coding technologies is traditionally available. This means that most SPAs are written entirely in a single language. This language is usually JavaScript, or sometimes another language, [compiled to JavaScript](https://github.com/jashkenas/coffeescript/wiki/List-of-languages-that-compile-to-JS).

One key goal of this project is to allow the traditional "client-side" development to use traditional tools, including JavaScript, while allowing the traditional "back end" work to be done in Go, while creating a single, unified output file (compiled JavaScript) which can run in the browser.

It might be enough to have a router written in Go, with only a JavaScript API, but once we've achieved this, supporting a Go API, and thus allowing request handlers written in Go, too, is a small step to take. This will also allow a choice, when it makes sense, between Go and JS for handler logic. And by making this a design goal from the beginning, I hope to ensure that both modes of operation (handlers in both Go and JS) are well supported.

# License

This software is released under the terms of the Apache 2.0 license. See LICENCE.md, or read the [full license](http://www.apache.org/licenses/LICENSE-2.0).
