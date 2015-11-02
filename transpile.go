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
	app.Version = "0.2.1"
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
				if err := cmd("go", "get", "-u", "github.com/tardisgo/tardisgo"); err != nil {
					fmt.Println(err)
					return
				}
				tardis("java", ctx.Args()[0], "jar")
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
				if err := cmd("go", "get", "-u", "github.com/tardisgo/tardisgo"); err != nil {
					fmt.Println(err)
					return
				}
				tardis("cpp", ctx.Args()[0], "cpp")
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
				if err := cmd("go", "get", "-u", "github.com/gopherjs/gopherjs"); err != nil {
					fmt.Println(err)
					return
				}
				gopherjs(ctx.Args()[0])
			},
		},
	}
	app.Run(os.Args)
}

func tardis(lang, pkg, suf string) {
	if err := cmd("tardisgo", pkg); err != nil {
		fmt.Println(err)
		return
	}
	err := cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", fmt.Sprintf("-%s", lang), fmt.Sprintf("tardis/%s", lang))
	if err != nil {
		fmt.Println(err)
		return
	}
	binDir := fmt.Sprintf("%sbin/%s", os.Getenv("GOPATH"), lang)
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		err = cmd("mkdir", binDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.LastIndex(pkg, "/") == len(pkg)-1 {
		pkg = pkg[:len(pkg)-1]
	}
	end := pkg[strings.LastIndex(pkg, "/")+1:]
	switch lang {
	case "java":
		err = cmd("cp", fmt.Sprintf("tardis/%s/Go.jar", lang), fmt.Sprintf("%s/%s.jar", binDir, end))
		if err != nil {
			fmt.Println(err)
			return
		}
	case "cpp":
		err = cmd("cp", fmt.Sprintf("tardis/%s/Go", lang), fmt.Sprintf("%s/%s", binDir, end))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("binary placed in", binDir)
}

func gopherjs(pkg string) {
	err := cmd("gopherjs", "install", pkg)
	if err != nil {
		fmt.Println(err)
		return
	}
	binDir := fmt.Sprintf("%sbin/js", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		err = cmd("mkdir", binDir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.LastIndex(pkg, "/") == len(pkg)-1 {
		pkg = pkg[:len(pkg)-1]
	}
	end := pkg[strings.LastIndex(pkg, "/")+1:]
	err = cmd("mv", fmt.Sprintf("%s/bin/%s.js", os.Getenv("GOPATH"), end), fmt.Sprintf("%s/%s.js", binDir, end))
	err = cmd("mv", fmt.Sprintf("%s/bin/%s.js.map", os.Getenv("GOPATH"), end), fmt.Sprintf("%s/%s.js.map", binDir, end))
	fmt.Println("JS placed in", binDir)
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
