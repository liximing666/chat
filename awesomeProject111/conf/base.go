package conf

import (
	"encoding/json"
	"io/ioutil"
)

var BaseConf *BaseConfig

type DBItem struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type BaseConfig struct {
	ImgDir           string `json:"img_dir"`
	SystemAdmin      []int  `json:"system_admin"`
	HttpServerConfig struct {
		Addr         string `json:"addr"`
		ReadTimeout  int    `json:"read_timeout"`
		WriteTimeout int    `json:"write_timeout"`
	} `json:"http_server_config"`
	Mysql struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"redis"`
}

func Init(env string) {
	path := "./settings/base_config_" + env + ".json"
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, &BaseConf)
	if err != nil {
		panic(err)
	}
}
