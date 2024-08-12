package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Envs map[string]string

func GetEnvs() map[string]string {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load it")
	}

	if Envs == nil {
		Envs = make(map[string]string)
	}

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			Envs[pair[0]] = pair[1]
		}
	}
	return Envs
}
