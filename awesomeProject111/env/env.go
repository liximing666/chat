package env

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var env string

type Enviroment struct {
	ENV string `json:"ENV"`
}

var (
	DevEnv    = "dev"
	TestEnv   = "test"
	OnlineEnv = "online"
)

func GetEnv() string {
	if env == "" {
		load()
	}

	return strings.ToLower(env)
}

func IsDevEnv() bool {
	return GetEnv() == string(DevEnv)

}

func IsTestEnv() bool {
	return GetEnv() == string(TestEnv)

}

func IsProductionEnv() bool {
	return GetEnv() == string(OnlineEnv)
}

func load() {
	path := "./settings/config.json"

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var envStruct Enviroment

	err = json.Unmarshal(bs, &envStruct)
	if err != nil {
		panic(err)
	}

	env = envStruct.ENV
}
