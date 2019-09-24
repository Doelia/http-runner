package main

import (
	"os"
)

var HOME_FOLDER = os.Getenv("HOME") + "/.http-runner"

func main() {
	Server()
	//RunScript("toto.sh")
}
