package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var pkg string

func TestMain(m *testing.M) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := os.Chdir(wd + "/test/"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	pkg = strings.TrimPrefix(wd, os.Getenv("GOPATH")+"/src/")
	os.Exit(m.Run())
}

func TestJava(t *testing.T) {
	tardis("java", pkg)
	_, err := os.Stat("test.jar")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove("test.jar"); err != nil {
		t.Fatal(err)
	}
}

func TestCPP(t *testing.T) {
	tardis("cpp", pkg)
	_, err := os.Stat("test")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove("test"); err != nil {
		t.Fatal(err)
	}
}

func TestJS(t *testing.T) {
	gopherjs(pkg)
	_, err := os.Stat("test.js")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove("test.js"); err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat("test.js.map")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove("test.js.map"); err != nil {
		t.Fatal(err)
	}
}
