package main

import (
	"go-iris/web/models"
	"reflect"
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func Test1(t *testing.T) {
	user := models.User{}
	t.Log(user.Username == "")
	t.Log(reflect.TypeOf(user.Username))
	t.Log(reflect.TypeOf(user).String())
	t.Log(reflect.TypeOf(user))
}

func TestCasbinMiddleware(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.Debug(false))

	type ttcasbin struct {
		username string
		path     string
		method   string
		status   int
	}

	tt := []ttcasbin{
		{"alice", "/dataset1/resource1", "GET", 200},
		{"alice", "/dataset1/resource1", "POST", 200},
		{"alice", "/dataset1/resource2", "GET", 200},
		{"alice", "/dataset1/resource2", "POST", 404},

		{"bob", "/dataset2/resource1", "GET", 200},
		{"bob", "/dataset2/resource1", "POST", 200},
		{"bob", "/dataset2/resource1", "DELETE", 200},
		{"bob", "/dataset2/resource2", "GET", 200},
		{"bob", "/dataset2/resource2", "POST", 404},
		{"bob", "/dataset2/resource2", "DELETE", 404},

		{"bob", "/dataset2/folder1/item1", "GET", 404},
		{"bob", "/dataset2/folder1/item1", "POST", 200},
		{"bob", "/dataset2/folder1/item1", "DELETE", 404},
		{"bob", "/dataset2/folder1/item2", "GET", 404},
		{"bob", "/dataset2/folder1/item2", "POST", 200},
		{"bob", "/dataset2/folder1/item2", "DELETE", 404},
	}

	for _, tt := range tt {
		check(e, tt.method, tt.path, tt.username, tt.status)
	}
}

func check(e *httptest.Expect, method, path, username string, status int) {
	e.Request(method, path).WithBasicAuth(username, "password").Expect().Status(status)
}
