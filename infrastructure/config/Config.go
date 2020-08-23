package config

import "github.com/maulibra/simple_movie_api/util"

type Env struct {
	DbUser string
	DbPassword string
	DbHost string
	DbPort string
	SchemaName string
	Driver string
}


func NewEnv() *Env{
	return &Env{
		DbUser:     util.GetEnv("dbUser", "defaultUser"),
		DbPassword: util.GetEnv("dbPassword", "defaultPass"),
		DbHost:     util.GetEnv("dbHost", "defaultHost"),
		DbPort:     util.GetEnv("dbPort", "12345"),
		SchemaName: util.GetEnv("dbSchema", "defaultScheme"),
		Driver: util.GetEnv("driver","mysql"),
	}
}
