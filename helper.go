package main

import (
	"encoding/base64"
	"fmt"
)

//Ni ni
func Ni() {
	println("中国你最牛")
}

//Why why
func Why() {
	var s, _ = base64.StdEncoding.DecodeString("a2V5b25lOTUyNw==")

	var a = base64.StdEncoding.EncodeToString([]byte("15110202919"))
	fmt.Println(a)
	fmt.Println(string(s))
}
