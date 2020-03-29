package util

import "github.com/golang/protobuf/ptypes/wrappers"

func WrapBool(value bool) *wrappers.BoolValue {
	return &wrappers.BoolValue{Value: value}
}
