package main

import (
	"math/rand"
	"os"
	"time"
)

var HOME_FOLDER = os.Getenv("HOME") + "/.http-runner"

func main() {
	rand.Seed(time.Now().UnixNano())
	Server()
	//RunScript("toto.sh")
}
