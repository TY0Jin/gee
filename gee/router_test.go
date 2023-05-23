package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/user/jinxw", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hi/*", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/test/:name"), []string{"test", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/test/*"), []string{"test", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/test/*name/*"), []string{"test", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/jinxw")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "jinxw" {
		t.Fatal("name should be equal to 'jinxw'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}
