package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

//Config will used to save all info needed to initiate a connection
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Dbname   string `json:"dbname"`
}

var configObj []Config

func validateConfig(objArr []Config) error {
	if len(objArr) == 0 {
		return errors.New("Config cannot be empty")
	}
	for _, element := range objArr {
		if element.Username == "" {
			return errors.New("Username name cannot be empty")
		}

		if element.Password == "" {
			return errors.New("Password name cannot be empty")
		}

		if element.Hostname == "" {
			return errors.New("Hostname name cannot be empty")
		}

		if element.Dbname == "" {
			return errors.New("Database name cannot be empty")
		}
	}

	return nil
}

//LoadConfig Will load config from json file
func LoadConfig(loc string) {
	log.Printf("Loading db config from file " + loc)
	dat, err := ioutil.ReadFile(loc)
	if err != nil {
		panic(err)
	}
	var obj []Config
	err = json.Unmarshal([]byte(dat), &obj)
	if err != nil {
		panic("Error unMarshaling config json object. " + err.Error())
	}
	err = validateConfig(obj)
	if err != nil {
		panic("Error validating config object. " + err.Error())
	}
	configObj = obj
	log.Printf("config successfully loaded")
}

//GetConfig should return the config object loaded from json file
func GetConfig() []Config {
	err := validateConfig(configObj)
	if err != nil {
		panic("config not loaded")
	}
	return configObj
}
