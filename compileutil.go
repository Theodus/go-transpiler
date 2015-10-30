package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func main() {
	cmd("go", "get", "-u", "-v", "github.com/tardisgo/tardisgo")
	app := cli.NewApp()
	app.Name = "tardisgo-compileutil"
	app.Usage = "compile Go to Java, C++, and C# using tardisgo"
	app.Commands = []cli.Command{
		{
			Name:  "java",
			Usage: "Compile Go source to Java target",
			Action: func(ctx *cli.Context) {
				cmd("./bin/tardisgo", ctx.Args()[0])
				cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", "-java", "tardis/java")
				_, err := os.Stat("bin/java")
				if os.IsNotExist(err) {
					cmd("mkdir", "bin/java")
				}
				cmd("cp", "tardis/java/Go.jar", "bin/java/")
			},
		}, {
			Name:  "cpp",
			Usage: "Compile Go source to C++ target",
			Action: func(ctx *cli.Context) {
				cmd("./bin/tardisgo", ctx.Args()[0])
				cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", "-cpp", "tardis/cpp")
				cmd("mkdir", "bin/cpp")
				_, err := os.Stat("bin/cpp")
				if os.IsNotExist(err) {
					cmd("mkdir", "bin/cpp")
				}
				cmd("cp", "tardis/cpp/Go", "bin/cpp/")
			},
		}, {
			Name:  "cs",
			Usage: "Compile Go source to C# target",
			Action: func(ctx *cli.Context) {
				cmd("./bin/tardisgo", ctx.Args()[0])
				cmd("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", "-cs", "tardis/cs")
				cmd("mkdir", "bin/cs")
				_, err := os.Stat("bin/cs")
				if os.IsNotExist(err) {
					cmd("mkdir", "bin/cs")
				}
				cmd("cp", "tardis/cs/Go", "bin/cs/")
			},
		},
	}
	app.Run(os.Args)
}

func cmd(cmdName string, cmdArgs ...string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", cmdName)
		os.Exit(1)
	}
}
