package main

import (
	"fmt"
	"log"

	"github.com/NikhilSharma03/Yeager/config"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
