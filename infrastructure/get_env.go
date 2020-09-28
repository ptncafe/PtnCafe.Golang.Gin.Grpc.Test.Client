package infrastructure

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func GetConfiguration() (map[string]string, error){
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file")
		return nil, err
	}
	env:=os.Getenv("ENV")
	err = godotenv.Load(fmt.Sprintf(".env.override.%s",env))
	if err != nil {
		logrus.Fatalf("Error loading .env external file %s %v", fmt.Sprintf(".env.override.%s",env), err)
		return nil, err
	}
	var configuration map[string]string
	configuration, err = godotenv.Read()
	if err != nil {
		logrus.Fatalf("Error loading .env godotenv.Read external file")
		return nil, err
	}
	return configuration, err
}