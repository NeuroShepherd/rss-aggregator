package main

import (
	"fmt"
	"os"

	"github.com/neuroshepherd/rss-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "read config:", err)
		os.Exit(1)
	}
	fmt.Println(cfg)

	err = cfg.SetUser("kenobi")
	if err != nil {
		fmt.Fprintln(os.Stderr, "set user:", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "read config:", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
