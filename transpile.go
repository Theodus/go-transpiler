package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-transpiler"
	app.Usage = "compile Go to Java, C++, or JS using tardisgo/gopherjs"
	app.Version = "0.2.3"
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
				tardis("java", ctx.Args()[0])
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
				tardis("cpp", ctx.Args()[0])
			},
		}, {
			Name:  "js",
			Usage: "Compile Go source to JS target",
			Action: func(ctx *cli.Context) {
				if len(ctx.Args()) < 1 {
					fmt.Println("Please specify a package.")
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

func tardis(lang, pkg string) {
	_, err := os.Stat(fmt.Sprintf("%s/bin/tardisgo", os.Getenv("GOPATH")))
	if os.IsNotExist(err) {
		cmd("go", "get", "-u", "-v", "github.com/tardisgo/tardisgo")
	}
	cmd("tardisgo", pkg)
	cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", fmt.Sprintf("-%s", lang), fmt.Sprintf("tardis/%s", lang))
	binDir := fmt.Sprintf("%s/bin/%s", os.Getenv("GOPATH"), lang)
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		cmd("mkdir", binDir)
	}
	if strings.LastIndex(pkg, "/") == len(pkg)-1 {
		pkg = pkg[:len(pkg)-1]
	}
	end := pkg[strings.LastIndex(pkg, "/")+1:]
	switch lang {
	case "java":
		cmd("cp", fmt.Sprintf("tardis/%s/Go.jar", lang), fmt.Sprintf("%s/%s.jar", binDir, end))
	case "cpp":
		cmd("cp", fmt.Sprintf("tardis/%s/Go", lang), fmt.Sprintf("%s/%s", binDir, end))
	}
	fmt.Println("binary placed in", binDir)
}

func gopherjs(pkg string) {
	_, err := os.Stat(fmt.Sprintf("%s/bin/gopherjs", os.Getenv("GOPATH")))
	if os.IsNotExist(err) {
		cmd("go", "get", "-u", "-v", "github.com/gopherjs/gopherjs")
	}
	cmd("gopherjs", "install", pkg)
	binDir := fmt.Sprintf("%s/bin/js", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		cmd("mkdir", binDir)
	}
	if strings.LastIndex(pkg, "/") == len(pkg)-1 {
		pkg = pkg[:len(pkg)-1]
	}
	end := pkg[strings.LastIndex(pkg, "/")+1:]
	cmd("mv", fmt.Sprintf("%s/bin/%s.js", os.Getenv("GOPATH"), end), fmt.Sprintf("%s/%s.js", binDir, end))
	cmd("mv", fmt.Sprintf("%s/bin/%s.js.map", os.Getenv("GOPATH"), end), fmt.Sprintf("%s/%s.js.map", binDir, end))
	fmt.Println("JS placed in", binDir)
}

func cmd(cmdName string, cmdArgs ...string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()
	errReader, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()
	if err = cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = cmd.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
