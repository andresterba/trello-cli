package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

type WorkConfig struct {
	BoardID string `json:"board_id,omitempty"`
}

type PersonalConfig struct {
	BoardID        string           `json:"board_id,omitempty"`
	RecurringTasks []RecurringTasks `json:"recurring_tasks,omitempty"`
}

type ShoppingConfig struct {
	BoardID      string `json:"board_id"`
	ListCardName string `json:"list_card_name"`
}

type RecurringTasks struct {
	Name   string   `json:"name,omitempty"`
	ListID string   `json:"listid,omitempty"`
	Labels []string `json:"labels,omitempty"`
}

type Config struct {
	DefaultContext string         `json:"default_context,omitempty"`
	AppKey         string         `json:"app_key,omitempty"`
	Token          string         `json:"token,omitempty"`
	WorkConfig     WorkConfig     `json:"work_config,omitempty"`
	PersonalConfig PersonalConfig `json:"personal_config,omitempty"`
	ShoppingConfig ShoppingConfig `json:"shopping_config,omitempty"`
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

func (c *Config) WriteConfig() error {
	file, _ := json.MarshalIndent(c, "", " ")
	err := ioutil.WriteFile(GetConfigPath(), file, 0644)
	if err != nil {
		return err
	}

	return nil
}
