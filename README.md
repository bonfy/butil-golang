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

### Log


* func InitLogFileWithPrefix(f *os.File, prefix string)
* func InitLogFile(f *os.File)


### Type

* func StringToUint(s string) (uint, error)


### Web

#### Return Json

* func WriteJsonOK(w http.ResponseWriter) 
* func WriteJsonErr(w http.ResponseWriter, msg string)
* func WriteJsonObject(w http.ResponseWriter, obj interface{}) 