package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func main() {
	out, err := cmd("go", "get", "-u", "-v", "github.com/tardisgo/tardisgo")
	if err != nil {
		fmt.Println(out, err)
	}
	app := cli.NewApp()
	app.Name = "tardisgo-compileutil"
	app.Usage = "compile Go to Java, C++, and C# using tardisgo"
	app.Commands = []cli.Command{
		{
			Name:  "java",
			Usage: "Compile Go source to Java target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					fmt.Println("Please specify a package.")
					return
				}
				if len(ctx.Args()) > 1 {
					fmt.Println("Too many arguments!")
					return
				}
				build("java", ctx.Args()[0], "jar")
			},
		}, {
			Name:  "cpp",
			Usage: "Compile Go source to C++ target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					fmt.Println("Please specify a package.")
					return
				}
				if len(ctx.Args()) > 1 {
					fmt.Println("Too many arguments!")
					return
				}
				build("cpp", ctx.Args()[0], "cpp")
			},
		}, {
			Name:  "cs",
			Usage: "Compile Go source to C# target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					fmt.Println("Please specify a package.")
					return
				}
				if len(ctx.Args()) > 1 {
					fmt.Println("Too many arguments!")
					return
				}
				build("cs", ctx.Args()[0], "cs")
			},
		},
	}
	app.Run(os.Args)
}

func build(lang, pkg, suf string) {
	out, err := cmd("tardisgo", pkg)
	if err != nil {
		fmt.Println(out, err)
		return
	}
	out, err = cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", fmt.Sprintf("-%s", lang), fmt.Sprintf("tardis/%s", lang))
	if err != nil {
		fmt.Println(out, err)
		return
	}
	binDir := fmt.Sprintf("%s/bin/%s", os.Getenv("GOPATH"), lang)
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		out, err = cmd("mkdir", binDir)
		if err != nil {
			fmt.Println(out, err)
			return
		}
	}
	if lang == "java" {
		out, err = cmd("cp", fmt.Sprintf("tardis/%s/Go.jar", lang), binDir)
		if err != nil {
			fmt.Println(out, err)
			return
		}
	} else if lang == "cpp" {
		out, err = cmd("cp", fmt.Sprintf("tardis/%s/Go", lang), binDir)
		if err != nil {
			fmt.Println(out, err)
			return
		}
	} else {
		out, err = cmd("cp", fmt.Sprintf("tardis/%s/Go.csproj", lang), binDir)
		if err != nil {
			fmt.Println(out, err)
			return
		}
	}
	fmt.Println("binary placed in", binDir)
}

func cmd(cmdName string, cmdArgs ...string) (out string, err error) {
	b, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
	out = string(b)
	return out, err
}
