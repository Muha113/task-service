package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	HttpAddr string `json:"http_host"`
	HttpPort string `json:"http_port"`
}

func InitConfig(cfgPath string) (*Config, error) {
	var config Config

	file, err := os.Open("./configs/config_debug.json")
	if err != nil {
		return nil, err
	}

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buff, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
