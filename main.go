/*下载工具*/
package main

import (
	"fmt"
	"encoding/base64"
)


func main() {

	var s, _ = base64.StdEncoding.DecodeString("a2V5b25lOTUyNw==")
	fmt.Println(string(s))
}
