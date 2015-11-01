package main

import (
	"fmt"
	"os"
	"testing"
)

func TestJava(t *testing.T) {
	out, err := cmd("go", "install", "github.com/theodus/tardisgo-compileutil")
	if err != nil {
		t.Fatal(out, err)
	}
	out, err = cmd("tardisgo-compileutil", "java", "github.com/theodus/tardisgo-compileutil/test")
	if err != nil {
		t.Fatal(out, err)
	}
	binDir := fmt.Sprintf("%s/bin/java/Go.jar", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("jar file not in ", binDir)
	}
}

func TestCPP(t *testing.T) {
	out, err := cmd("go", "install", "github.com/theodus/tardisgo-compileutil")
	if err != nil {
		t.Fatal(out, err)
	}
	out, err = cmd("tardisgo-compileutil", "cpp", "github.com/theodus/tardisgo-compileutil/test")
	if err != nil {
		t.Fatal(out, err)
	}
	binDir := fmt.Sprintf("%s/bin/cpp/Go", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("C++ binary not in ", binDir)
	}
}

func TestCS(t *testing.T) {
	out, err := cmd("go", "install", "github.com/theodus/tardisgo-compileutil")
	if err != nil {
		t.Fatal(out, err)
	}
	out, err = cmd("tardisgo-compileutil", "cs", "github.com/theodus/tardisgo-compileutil/test")
	if err != nil {
		t.Fatal(out, err)
	}
	binDir := fmt.Sprintf("%s/bin/cs/Go.csproj", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		t.Fatal("C# file not in ", binDir)
	}
}
