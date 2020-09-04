package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func main() {
	pathOnLoadGoDotEndConcise()

	fmt.Println(os.Getenv("FOUND"))
}

//         >> godotenv <<
// https://github.com/joho/godotenv

// import as always and call a second function
// that imports envs if a ADDITIONAL_ENV_PATH env exists
// the function tries to import envs from full path
// if it does not succced then tries relative path
//
// so both are accceptable
// a. ADDITIONAL_ENV_PATH=additional/.env
// b. ADDITIONAL_ENV_PATH=/home/blabla/blabla/.env

func pathOnLoadGoDotEndConcise() {
	err := godotenv.Load()
	if err != nil || loadExtraEnv() != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

func loadExtraEnv() error {
	envPath := os.Getenv("ADDITIONAL_ENV_PATH")
	if envPath == "" {
		return nil
	}

	err := godotenv.Load(envPath)
	if err == nil {
		return nil
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envPath = path.Join(dir, envPath)
	fmt.Println(envPath)
	return godotenv.Load(envPath)
}
