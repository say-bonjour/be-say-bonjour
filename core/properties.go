package core

import (
	"encoding/json"
	"log"
	"os"
)

type Properties struct {
	Port     string             `json:"port"`
	Database DatabaseProperties `json:"database"`
}

type DatabaseProperties struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	PoolSize int    `json:"poolSize"`
}

func LoadConfiguration(filename string) (Properties, error) {
	var properties Properties

	jsonFile, err := os.Open(filename)
	if err != nil {
		return properties, err
	}

	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&properties)
	log.Println(properties)
	return properties, err
}
