package main

import (
	"fmt"
	"reflect"

	butil "github.com/bonfy/butil-golang"
)

func main() {
	uid, _ := butil.StringToUint("8")
	fmt.Println(uid)
	fmt.Println(reflect.TypeOf(uid))
}
