package main

import (
	"casbin-demo/models"
	"reflect"
	"testing"
)

func Test1(t *testing.T)  {
	user := models.User{}
	t.Log(user.Username == "")
	t.Log(reflect.TypeOf(user.Username))
	t.Log(reflect.TypeOf(user).String())
	t.Log(reflect.TypeOf(user))
}
