package main

import (
	"fmt"
	"os"
)

var HOME_FOLDER = os.Getenv("HOME") + "/.http-runner"

func main() {
	fmt.Printf("hello, world\n")
	//runCmd()
	Server()
	//RunScript("toto.sh")
}
