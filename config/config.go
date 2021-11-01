package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
)

var (
	errCouldNotOpenConfigFile  = errors.New("could not open configuration file")
	errCouldNotParseConfigFile = errors.New("could not parse configuration file")
)

const (
	configFileName = ".trello-cli.json"
)

type Config struct {
	AppKey               string `json:"app_key"`
	Token                string `json:"token"`
	BoardID              string `json:"board_id"`
	ShoppingListCardName string `json:"shopping_list_card_name"`
}

func GetConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/%s", usr.HomeDir, configFileName)
}

func LoadConfig(pathToConfig string) (*Config, error) {
	file, err := os.Open(pathToConfig)
	if err != nil {
		return nil, errCouldNotOpenConfigFile
	}

	config := &Config{}

	jsonDecoder := json.NewDecoder(file)

	err = jsonDecoder.Decode(&config)
	if err != nil {
		return nil, errCouldNotParseConfigFile
	}

	return config, nil
}
