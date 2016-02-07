package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-transpiler"
	app.Author = "Theodore Butler"
	app.Usage = "compile Go to Java, C++, or JS using tardisgo or gopherjs"
	app.Version = "0.3.0"
	app.Commands = []cli.Command{
		{
			Name:  "java",
			Usage: "Compile Go source to Java target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					tardis("java", "")
					return
				}
				if len(ctx.Args()) > 1 {
					fmt.Println("Too many arguments!")
					return
				}
				tardis("java", ctx.Args()[0])
			},
		}, {
			Name:  "cpp",
			Usage: "Compile Go source to C++ target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					tardis("cpp", "")
					return
				}
				if len(ctx.Args()) > 1 {
					fmt.Println("Too many arguments!")
					return
				}
				tardis("cpp", ctx.Args()[0])
			},
		}, {
			Name:  "js",
			Usage: "Compile Go source to JS target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					gopherjs("")
					return
				}
				if len(ctx.Args()) > 1 {
					fmt.Println("Too many arguments!")
					return
				}
				gopherjs(ctx.Args()[0])
			},
		},
	}
	app.Run(os.Args)
}
