//go:build !go1.18
// +build !go1.18

package sqlconv

import "reflect"

type any = interface{}

const reflectPointer = reflect.Ptr
