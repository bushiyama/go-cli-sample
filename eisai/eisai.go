package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

const version = "1.1.0"

var (
	flag1 bool
)

func main() {
	app := &cli.App{
		Usage: "If you type eisai, haramascoi will be returned. option \"-p\"",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "p",
				Usage:       "hayaranai",
				Destination: &flag1,
				Required:    false,
			},
		},
		Action: func(c *cli.Context) error {
			ret := "haramasuko-i"
			if flag1 {
				ret = "e-nn"
			}
			fmt.Println(ret)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
