package main

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "livereload",
		Usage:       "Start a `livereload` server",
		Description: "Reload webpages on changes, without hitting refresh in your browser",
		Authors: []*cli.Author{
			{
				Name:  "Tacey Wong",
				Email: "xinyong.wang@qq.com",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "127.0.0.1",
				Usage: "`Hostname` to run livereload server on",
			},
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   35729,
				Usage:   "`Port` to run livereload server on",
			},
			&cli.PathFlag{
				Name:  "directory",
				Value: ".",
				Usage: "`Directory` to serve files from",
			},
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Usage:   "`File` or `directory` to watch for changes",
			},
			&cli.Float64Flag{
				Name:    "wait",
				Aliases: []string{"w"},
				Value:   0.0,
				Usage:   "Time delay in `seconds` before reloading",
			},
			&cli.Float64Flag{
				Name:    "open-url-delay",
				Aliases: []string{"o"},
				Usage:   `If set, triggers browser opening <D> seconds after starting`,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Enable  pretty logging",
			},
		},
	}
	license := "BSD-3-Clause License"
	app.Copyright = "2020 - Now " + app.Authors[0].Name + " , " + license
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
