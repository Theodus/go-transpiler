package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/theodus/command"
)

func TestJava(t *testing.T) {
	command.Verbose("go", "install", "github.com/theodus/go-transpiler")
	command.Verbose("go-transpiler", "java", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/java/test.jar", os.Getenv("GOPATH"))
	_, err := os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("jar file not in ", binDir)
	}
}

func TestCPP(t *testing.T) {
	command.Verbose("go", "install", "github.com/theodus/go-transpiler")
	command.Verbose("go-transpiler", "cpp", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/cpp/test", os.Getenv("GOPATH"))
	_, err := os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("C++ binary not in ", binDir)
	}
}

func TestJS(t *testing.T) {
	command.Verbose("go", "install", "github.com/theodus/go-transpiler")
	command.Verbose("go-transpiler", "js", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/js/test.js", os.Getenv("GOPATH"))
	_, err := os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("JS file not in ", binDir)
	}
}
