package main

import (
	"fmt"
	"os"
	"testing"
)

func TestJava(t *testing.T) {
	err := cmd("go", "install", "github.com/theodus/go-transpiler")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd("go-transpiler", "java", "github.com/theodus/go-transpiler/test")
	if err != nil {
		t.Fatal(err)
	}
	binDir := fmt.Sprintf("%s/bin/java/test.jar", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("jar file not in ", binDir)
	}
}

func TestCPP(t *testing.T) {
	err := cmd("go", "install", "github.com/theodus/go-transpiler")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd("go-transpiler", "cpp", "github.com/theodus/go-transpiler/test")
	if err != nil {
		t.Fatal(err)
	}
	binDir := fmt.Sprintf("%s/bin/cpp/test", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("C++ binary not in ", binDir)
	}
}

func TestCS(t *testing.T) {
	err := cmd("go", "install", "github.com/theodus/go-transpiler")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd("go-transpiler", "cs", "github.com/theodus/go-transpiler/test")
	if err != nil {
		t.Fatal(err)
	}
	binDir := fmt.Sprintf("%s/bin/cs/test.csproj", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("C# file not in ", binDir)
	}
}

func TestJS(t *testing.T) {
	err := cmd("go", "install", "github.com/theodus/go-transpiler")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd("go-transpiler", "js", "github.com/theodus/go-transpiler/test")
	if err != nil {
		t.Fatal(err)
	}
	binDir := fmt.Sprintf("%s/bin/js/test.js", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("JS file not in ", binDir)
	}
}