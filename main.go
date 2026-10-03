package main

import (
	"fmt"
	"gatorcli/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.SetUser("Jeremy")
	if err != nil {
		log.Fatal(err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("test for push")
	fmt.Println("test for user push")
	fmt.Println(updatedCfg.DBUrl)
	fmt.Println(updatedCfg.CurrentUserName)
}
