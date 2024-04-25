package model


type Service_model struct {
	Service_name string
	Service_id string
	Envs []Env_model
}

type Env_model struct {
	Env_name string
	Env_value string
}