package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func getConfigYaml() []byte {
	dat, err := ioutil.ReadFile(HOME_FOLDER + "/config.yml")
	if err != nil {
		fmt.Printf("Error during getting config.yml : %s", err)
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

