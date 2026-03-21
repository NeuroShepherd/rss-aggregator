package main

import (
	"fmt"

	"github.com/neuroshepherd/rss-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}
