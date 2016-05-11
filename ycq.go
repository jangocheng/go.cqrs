package ycq

import (
	"github.com/jetbasrawi/yoono-uuid"
	"reflect"
)

func typeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

// A helper function to provide quick access to a uuid
// sorry about the naming. I could not resist the pun
func yooid() uuid.UUID {
	return uuid.NewV4()
}