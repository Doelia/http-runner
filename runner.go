package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

var scriptsFolder = HOME_FOLDER + "/scripts"
var logsFolder = HOME_FOLDER + "/logs"

func RunScript(id string) {
	log.Println("RunScript " + id + "...")

	path := scriptsFolder + "/" + id
	cmd := exec.Command("/bin/sh", path)

	// open the out file for writing
	_ = os.MkdirAll(logsFolder + "/" + id, os.ModePerm)

	now := time.Now()
	filename := now.Format("2006-01-02_15-04-5") + ".log"

	outfile, err := os.Create(logsFolder + "/" + id + "/" + filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile
	cmd.Stderr = outfile

	err = cmd.Start(); if err != nil {
		panic(err)
	}
	_ = cmd.Wait()

	log.Println("RunScript " + id + " finished.")
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