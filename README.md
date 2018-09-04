# butil-golang
bonfy's own util for golang

## Usage

```golang
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
```

### Type

* func StringToUint(s string) (uint, error)


### Web

#### Return Json

* func WriteJsonOK(w http.ResponseWriter) 
* func WriteJsonErr(w http.ResponseWriter, msg string)
* func WriteSingleObjectStatus(w http.ResponseWriter, obj interface{})
* func WriteListObjectStatus(w http.ResponseWriter, obj []interface{})  