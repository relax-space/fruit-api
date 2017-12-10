package main

import "sync"

var ()

type EnvParamDto struct {
	AppEnv  string `json:"app_env"`
	ConnEnv string `json:"conn_env"`
}

var Env *EnvParamDto
var onceEnv sync.Once

func InitEnv(env *EnvParamDto) {
	onceEnv.Do(func() {
		Env = env
	})
}
