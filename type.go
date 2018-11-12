package cn

import (
	"encoding/json"
	"io/ioutil"
)

type configuration struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (cfg *configuration) Load(configFilePath string) {
	//Read local Json-file
	var fcfg []byte
	var err error
	if configFilePath == "" {
		fcfg, err = ioutil.ReadFile("conf.json")
	} else {
		fcfg, err = ioutil.ReadFile(configFilePath)
	}
	CheckError(err)

	err = json.Unmarshal(fcfg, cfg)
	CheckError(err)
}
