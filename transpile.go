package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func tardis(lang, pkg string) {
	_, err := os.Stat(fmt.Sprintf("%s/bin/tardisgo", os.Getenv("GOPATH")))
	if os.IsNotExist(err) {
		verbose("go", "get", "-u", "-v", "github.com/tardisgo/tardisgo")
	}
	if pkg == "" {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		pref := fmt.Sprintf("%s%s", os.Getenv("GOPATH"), "/src/")
		pkg = strings.TrimPrefix(wd, pref)
	}
	verbose("tardisgo", pkg)
	verbose("haxe", "-main", "tardis.Go", "-cp", "tardis", "-dce", "full", fmt.Sprintf("-%s", lang), fmt.Sprintf("tardis/%s", lang))
	out := pkg[strings.LastIndex(pkg, "/")+1:]
	switch lang {
	case "java":
		out += ".jar"
		f, err := os.Create(out)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := copyFile("tardis/java/Go.jar", f.Name()); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "cpp":
		f, err := os.Create(out)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := copyFile("tardis/java/Go", f.Name()); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func gopherjs(pkg string) {
	_, err := os.Stat(fmt.Sprintf("%s/bin/gopherjs", os.Getenv("GOPATH")))
	if os.IsNotExist(err) {
		verbose("go", "get", "-u", "-v", "github.com/gopherjs/gopherjs")
	}
	if pkg == "" {
		verbose("gopherjs", "build")
		return
	}
	verbose("gopherjs", "build", pkg)
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

func verbose(cmdName string, cmdArgs ...string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdStr := cmdName
	for _, s := range cmdArgs {
		cmdStr += fmt.Sprintf(" %s", s)
	}
	fmt.Println(cmdStr)
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	outScanner := bufio.NewScanner(outReader)
	go func() {
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()
	errReader, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()
	if err = cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err = cmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
