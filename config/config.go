package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Conf configStruct

type configStruct struct {
	DiscordToken string `json : "discordToken"`
	ClashToken   string `json : "clashToken"`
	Prefix       string `json : "prefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &Conf)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
