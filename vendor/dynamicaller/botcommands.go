package dynamicaller

import (
	"log"
)

func (dynamicaller Dynamicaller) LembreteHandlerFunc(arg0 string) (res0 string, res1 string, err error) {
	log.Print("Handling command /lembrete with received args:", arg0)
	return "some value", "another value", nil
}