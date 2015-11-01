package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func main() {
	if err := cmd("go", "get", "-u", "-v", "github.com/tardisgo/tardisgo"); err != nil {
		fmt.Println(err)
		return
	}
	app := cli.NewApp()
	app.Name = "tardisgo-compileutil"
	app.Usage = "compile Go to Java, C++, and C# using tardisgo"
	app.Version = "0.1.0"
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
	if err := cmd("tardisgo", pkg); err != nil {
		fmt.Println(err)
		return
	}
	err := cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", fmt.Sprintf("-%s", lang), fmt.Sprintf("tardis/%s", lang))
	if err != nil {
		fmt.Println(err)
		return
	}
	binDir := fmt.Sprintf("%s/bin/%s", os.Getenv("GOPATH"), lang)
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		err = cmd("mkdir", binDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if lang == "java" {
		err = cmd("cp", fmt.Sprintf("tardis/%s/Go.jar", lang), binDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if lang == "cpp" {
		err = cmd("cp", fmt.Sprintf("tardis/%s/Go", lang), binDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = cmd("cp", fmt.Sprintf("tardis/%s/Go.csproj", lang), binDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("binary placed in", binDir)
}

func cmd(cmdName string, cmdArgs ...string) error {
	cmd := exec.Command(cmdName, cmdArgs...)
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()
	if err = cmd.Start(); err != nil {
		return err
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
