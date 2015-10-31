package main

import (
	"testing"
)

func TestCompileutil(t *testing.T) {
	out, err := cmd("go", "install", "github.com/theodus/tardisgo-compileutil")
	if err != nil {
		t.Fatal(out, err)
	}
	out, err = cmd("tardisgo-compileutil", "java", "github.com/theodus/tardisgo-compileutil/test")
	if err != nil {
		t.Fatal(out, err)
	}
	out, err = cmd("tardisgo-compileutil", "cpp", "github.com/theodus/tardisgo-compileutil/test")
	if err != nil {
		t.Fatal(out, err)
	}
	out, err = cmd("tardisgo-compileutil", "cs", "github.com/theodus/tardisgo-compileutil/test")
	if err != nil {
		t.Fatal(out, err)
	}
}
