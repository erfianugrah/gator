package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	err = cfg.SetUser("Erfi")
	if err != nil {
		return
	}
	finalCfg, err := config.Read()
	if err != nil {
		return
	}
	fmt.Println(finalCfg)
}
