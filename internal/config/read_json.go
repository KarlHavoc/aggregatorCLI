package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

const config_filepath = "/.gatorconfig.json"

func getFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
	}
	config_path := home_dir + config_filepath
	return config_path, nil
}

func ReadConfig() (Config, error) {
	new_config := Config{}
	fp, err := getFilePath()
	if err != nil {
		return Config{}, err
	}
	dat, err := os.ReadFile(fp)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(dat, &new_config)
	if err != nil {
		return Config{}, err
	}
	return new_config, nil
}
