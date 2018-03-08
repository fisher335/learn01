package main

import (
	"os/exec"

	"io/ioutil"
)

func main() {

	var cmd= exec.Command("dir")
		var a,_=cmd.StdoutPipe()
		defer a.Close()
		var v,_ = ioutil.ReadAll(a)
		println(string(v))

	}
