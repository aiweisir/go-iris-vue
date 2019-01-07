package utils

import (
	"fmt"
	"strconv"
)

const (
	rolePrefix string = "role_"
)

//
func FmtRolePrefix(sub interface{}) string {
	var s string
	switch sub.(type) {
	case int64:
		uid := sub.(int64)
		s = strconv.FormatInt(uid, 10)
	case string:
		s = sub.(string)
	}
	return fmt.Sprintf("%s%s", rolePrefix, s)
}
