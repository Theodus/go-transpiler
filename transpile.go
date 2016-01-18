package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/theodus/command"
)

func tardis(lang, pkg string) {
	_, err := os.Stat(fmt.Sprintf("%s/bin/tardisgo", os.Getenv("GOPATH")))
	if os.IsNotExist(err) {
		command.Verbose("go", "get", "-u", "-v", "github.com/tardisgo/tardisgo")
	}
	command.Verbose("tardisgo", pkg)
	defer func() {
		if err := os.RemoveAll("tardis"); err != nil {
			fmt.Println(err)
			return
		}
	}()
	command.Verbose("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", fmt.Sprintf("-%s", lang), fmt.Sprintf("tardis/%s", lang))
	binDir := fmt.Sprintf("%s/bin/%s", os.Getenv("GOPATH"), lang)
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		if err = os.Mkdir(binDir, 0777); err != nil {
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
		if err = copyFile(fmt.Sprintf("tardis/%s/Go.jar", lang), fmt.Sprintf("%s/%s.jar", binDir, end)); err != nil {
			fmt.Println(err)
			return
		}
	case "cpp":
		if err = copyFile(fmt.Sprintf("tardis/%s/Go", lang), fmt.Sprintf("%s/%s", binDir, end)); err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("binary placed in", binDir)
}

func gopherjs(pkg string) {
	_, err := os.Stat(fmt.Sprintf("%s/bin/gopherjs", os.Getenv("GOPATH")))
	if os.IsNotExist(err) {
		command.Verbose("go", "get", "-u", "-v", "github.com/gopherjs/gopherjs")
	}
	command.Verbose("gopherjs", "install", pkg)
	binDir := fmt.Sprintf("%s/bin/js", os.Getenv("GOPATH"))
	_, err = os.Stat(binDir)
	if os.IsNotExist(err) {
		if err = os.Mkdir(binDir, 0777); err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.LastIndex(pkg, "/") == len(pkg)-1 {
		pkg = pkg[:len(pkg)-1]
	}
	end := pkg[strings.LastIndex(pkg, "/")+1:]
	if err = copyFile(fmt.Sprintf("%s/bin/%s.js", os.Getenv("GOPATH"), end), fmt.Sprintf("%s/%s.js", binDir, end)); err != nil {
		fmt.Println(err)
		return
	}
	if err = copyFile(fmt.Sprintf("%s/bin/%s.js.map", os.Getenv("GOPATH"), end), fmt.Sprintf("%s/%s.js.map", binDir, end)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("JS placed in", binDir)
}

func copyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return nil
		}
	}
	if err = os.Link(src, dst); err == nil {
		return err
	}
	return copyFileContents(src, dst)
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return err
}
