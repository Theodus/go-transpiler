package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/theodus/command"
)

var gopath = os.Getenv("GOPATH")

func TestJava(t *testing.T) {
	command.Verbose("go", "install", "github.com/theodus/go-transpiler")
	command.Verbose("go-transpiler", "java", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/java/test.jar", gopath)
	_, err := os.Stat(binDir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCPP(t *testing.T) {
	command.Verbose("go", "install", "github.com/theodus/go-transpiler")
	command.Verbose("go-transpiler", "cpp", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/cpp/test", gopath)
	_, err := os.Stat(binDir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJS(t *testing.T) {
	command.Verbose("go", "install", "github.com/theodus/go-transpiler")
	command.Verbose("go-transpiler", "js", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/js/test.js", gopath)
	_, err := os.Stat(binDir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCleanup(t *testing.T) {
	if err := os.RemoveAll(gopath + "/bin/java"); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(gopath + "/bin/cpp"); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(gopath + "/bin/js"); err != nil {
		t.Fatal(err)
	}
}
