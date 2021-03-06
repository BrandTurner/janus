package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jijeshmohan/janus/config"
	"github.com/jijeshmohan/janus/server"
)

func main() {
	fmt.Printf("Janus - fake rest api server (%s) \n", VERSION)
	c := getConfig()

	server.StartServer(c)
}

// getConfig get the configuration from the config file.
func getConfig() *config.Config {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to get the current working directory")
		os.Exit(1)
	}

	c, err := config.ParseFile(filepath.Join(path, "config.json"))

	if err != nil {
		fmt.Printf("Config file error: %s\n", err.Error())
		os.Exit(1)
	}

	// Add current path to the config
	c.Path = path

	return c
}
