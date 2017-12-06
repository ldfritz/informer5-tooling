package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	API   string `json:"api"`
	Token string `json:"token"`
}

func ParseConfig(contents []byte) (Config, error) {
	var config Config
	if err := json.Unmarshal(contents, &config); err != nil {
		return config, err
	}
	return config, nil
}

func LoadConfigFile(filename string) (Config, error) {
	contents, err := ioutil.ReadFile("config.json")
	if err != nil {
		return Config{}, err
	}
	return ParseConfig(contents)
}
