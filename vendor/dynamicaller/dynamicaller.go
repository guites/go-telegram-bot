package dynamicaller

import (
	"reflect"
)

type Dynamicaller struct {}

func (dynamicaller Dynamicaller) DynamicCall(obj interface{}, fn string, args0 string, chatId int) (res []reflect.Value) {
	method := reflect.ValueOf(obj).MethodByName(fn)
	var inputs []reflect.Value
	inputs = append(inputs, reflect.ValueOf(args0))
	inputs = append(inputs, reflect.ValueOf(chatId))
	return method.Call(inputs)
}
