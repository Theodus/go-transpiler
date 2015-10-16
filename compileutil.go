package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Please specify the package and target language. (java, cpp, cs)")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments!")
		return
	}
	_, err := os.Stat("src/github.com/tardisgo/tardisgo")
	if os.IsNotExist(err) {
		cmd("go", []string{"get", "-u", "github.com/tardisgo/tardisgo"})
	}

	compArgs := []string{"-main", "tardis.Go", "-cp", "tardis", "-dce", "full"}
	lang := args[1]
	switch lang {
	case "java":
		compArgs = append(compArgs, "-java", "tardis/java")
		cmd("./bin/tardisgo", []string{args[0]})
		cmd("haxe", compArgs)
		_, err = os.Stat("bin/java")
		if os.IsNotExist(err) {
			cmd("mkdir", []string{"bin/java"})
		}
		cmd("cp", []string{"tardis/java/Go.jar", "bin/java/"})
	case "cpp":
		compArgs = append(compArgs, "-cpp", "tardis/cpp")
		cmd("./bin/tardisgo", []string{args[0]})
		cmd("haxe", compArgs)
		cmd("mkdir", []string{"bin/cpp"})
		_, err = os.Stat("bin/cpp")
		if os.IsNotExist(err) {
			cmd("mkdir", []string{"bin/cpp"})
		}
		cmd("cp", []string{"tardis/cpp/Go", "bin/cpp/"})
	case "cs":
		compArgs = append(compArgs, "-cs", "tardis/cs")
		cmd("./bin/tardisgo", []string{args[0]})
		cmd("haxe", compArgs)
	default:
		fmt.Println(lang, "is not a target language. (java, cpp, cs)")
		return
	}
}

func cmd(cmdName string, cmdArgs []string) {
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
