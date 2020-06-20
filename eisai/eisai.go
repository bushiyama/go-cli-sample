package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

const version = "1.0.0"

func main() {
	app := &cli.App{
		Usage: "If you type eisai, haramascoi will be returned.",
		Action: func(c *cli.Context) error {
			fmt.Println("haramasuko-i")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
