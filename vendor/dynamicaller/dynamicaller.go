package dynamicaller

import (
	"reflect"
)

type Dynamicaller struct {}

func (dynamicaller Dynamicaller) DynamicCall(obj interface{}, fn string, args string) (res []reflect.Value) {
	method := reflect.ValueOf(obj).MethodByName(fn)
	var inputs []reflect.Value
	inputs = append(inputs, reflect.ValueOf(args))
	return method.Call(inputs)
}
