package main

import (
	"fmt"
	"os"
	"os/exec"
)

var scriptsFolder = HOME_FOLDER + "/scripts"
var logsFolder = HOME_FOLDER + "/logs"

func RunScript(id string) {
	path := scriptsFolder + "/" + id
	cmd := exec.Command("/bin/sh", path)

	// open the out file for writing
	_ = os.MkdirAll(logsFolder + "/" + id, os.ModePerm)
	outfile, err := os.Create(logsFolder + "/" + id + "/" + "log.log")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	if err != nil {
		fmt.Printf("%s", err)
	}

	err = cmd.Start(); if err != nil {
		panic(err)
	}
	cmd.Wait()

	fmt.Println("Command Successfully Executed")

}

func ScriptExists(id string) bool {
	return fileExists(scriptsFolder + "/" + id)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}