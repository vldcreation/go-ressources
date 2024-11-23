package util

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

func LoadEnvFromPath(path string, cfg any) error {
	v := viper.New()
	sourcePath := RootPath() + "/" + path
	v.SetConfigFile(sourcePath)
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		log.Printf("Error unmarshalling config file, %s", err)
		return err
	}

	return nil
}

func LoadEnvFromBytes(data []byte, cfg any) error {
	v := viper.New()
	stream := bytes.NewBuffer(data)
	v.SetConfigType("json")

	if err := v.ReadConfig(stream); err != nil {
		log.Printf("Error reading config file, %+v", err)
		return err
	}

	if err := v.Unmarshal(cfg); err != nil {
		log.Printf("Error unmarshalling config file, %+v", err)
		return err
	}

	return nil
}
