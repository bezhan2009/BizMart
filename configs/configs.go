package configs

import (
	"BizMart/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func ReadSettings() (models.Configs, error) {
	var AppSettings models.Configs

	configFile, err := os.Open("configs/configs.json")
	if err != nil {
		configFile, err = os.Open("configs/example.json")
		if err != nil {
			return models.Configs{}, errors.New(fmt.Sprintf("Couldn't open config file. Error is: %s", err.Error()))
		}
	}

	defer func(configFile *os.File) {
		err = configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}
	}(configFile)

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return models.Configs{}, errors.New(fmt.Sprintf("Couldn't decode settings json file. Error is: %s", err.Error()))
	}
	fmt.Println("Read: ", AppSettings)
	return AppSettings, nil
}
