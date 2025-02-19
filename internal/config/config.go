package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	config_fp, err := getFilePath()
	if err != nil {
		fmt.Print("error getting config filepath")
	}
	dat, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		fmt.Print("error marshalling struct data to json bytes")
	}
	//fmt.Printf("config json bytes: %v", dat)
	err = os.WriteFile(config_fp, dat, 0644)
	if err != nil {
		fmt.Print("error writing json to config file")
	}
	return nil
}
