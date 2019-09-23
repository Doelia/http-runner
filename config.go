package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ConfigStruct struct {
	Port string
	Host string
	Security struct {
		Auth_type string
		Basic_auth struct {
			Login string
			Password string
		}
		Ip_authorised []string
	}
}

func copyTo(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func createConfig() error {
	fmt.Println("createConfig...")
	relativePath, _ := filepath.Abs("./")
	_ = os.MkdirAll(HOME_FOLDER, os.ModePerm)
	return copyTo(relativePath + "/config.example.yaml", HOME_FOLDER + "/config.yml")
}

func getConfigYaml() []byte {
	dat, err := ioutil.ReadFile(HOME_FOLDER + "/config.yml")
	if err != nil {
		log.Println("Create config.yaml...")
		err2 := createConfig()
		if err2 != nil {
			fmt.Printf("Error during create config.yml : %s", err2)
		} else {
			return getConfigYaml()
		}
	}

	return dat
}


func Config() (*ConfigStruct, error) {
	t := ConfigStruct{}

	data := getConfigYaml()
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

