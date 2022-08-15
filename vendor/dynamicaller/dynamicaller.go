package dynamicaller

import (
	"fmt"
	"log"
	"reflect"
)

type Dynamicaller struct {}

func (dynamicaller Dynamicaller) Lembrete(arg0 string) {
	fmt.Print("Lembretando! hehe")
}

func (dynamicaller Dynamicaller) DynamicCall(obj interface{}, fn string, args string) (res []reflect.Value) {
	method := reflect.ValueOf(obj).MethodByName(fn)
	var inputs []reflect.Value
	inputs = append(inputs, reflect.ValueOf(args))
	return method.Call(inputs)
}

func (dynamicaller Dynamicaller) LembreteHandlerFunc(arg0 string) (res0 string, res1 string, err error) {
	log.Print("Handling command /lembrete with received args:", arg0)
	return "some value", "another value", nil
}
