package main

import (
	"fmt"
	"os"
	"testing"
)

func TestJava(t *testing.T) {
	cmd("go", "install", "github.com/theodus/go-transpiler")
	cmd("go-transpiler", "java", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/java/test.jar", os.Getenv("GOPATH"))
	_, err := os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("jar file not in ", binDir)
	}
}

func TestCPP(t *testing.T) {
	cmd("go", "install", "github.com/theodus/go-transpiler")
	cmd("go-transpiler", "cpp", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/cpp/test", os.Getenv("GOPATH"))
	_, err := os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("C++ binary not in ", binDir)
	}
}

func TestJS(t *testing.T) {
	cmd("go", "install", "github.com/theodus/go-transpiler")
	cmd("go-transpiler", "js", "github.com/theodus/go-transpiler/test")
	binDir := fmt.Sprintf("%s/bin/js/test.js", os.Getenv("GOPATH"))
	_, err := os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("JS file not in ", binDir)
	}
}
