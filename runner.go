package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var scriptsFolder = HOME_FOLDER + "/scripts"
var logsFolder = HOME_FOLDER + "/logs"

func CreateLog(id string) string {
	// open the out file for writing
	_ = os.MkdirAll(logsFolder + "/" + id, os.ModePerm)

	now := time.Now()
	return now.Format("2006-01-02_15-04-5") + strconv.Itoa(rand.Intn(10000)) + ".log"
}

func RunScript(id string, logname string, getsParam string, rawParam string) {
	log.Println("RunScript " + id + "...")

	path := scriptsFolder + "/" + id
	cmd := exec.Command("/bin/bash", path, getsParam, rawParam, logname)

	outfile, err := os.Create(logsFolder + "/" + id + "/" + logname)
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

func ListLogFiles(id string) []string {
	var list []string
	files, _ := ioutil.ReadDir(logsFolder + "/" + id)
	for _, f := range files {
		fmt.Println(f.Name())
		list = append(list, f.Name())
	}
	return list
}

func ContentLogFile(id string, file string) []byte {
	data, _ := ioutil.ReadFile(logsFolder + "/" + id + "/" + file)
	return data
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
