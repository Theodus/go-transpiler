# go-transpiler

Command-Line utility for compiling Go to Java or C++ using tardisgo or to JavaScript using GopherJS

## Usage

Be sure that your GOPATH is set and the bin directory is added to PATH:
```
export GOPATH="path/to/gocode"
export PATH=$PATH:$GOPATH/bin
```
Be sure that Haxe version 3.2.0 or greater is installed for tardisgo

Go get and install this package:
```
go get -u github.com/theodus/go-transpiler
```
Compile your code to Java:
```
go-transpiler java myGoPackage
```
Or C++:
```
go-transpiler cpp myGoPackage
```
Or JS:
```
go-transpiler js myGoPackage
```

## More information on TARDIS Go

Check out the GitHub repository at https://github.com/tardisgo/tardisgo
or the website at https://tardisgo.github.io

## More information on GopherJS

Check out the GitHub repository at https://github.com/gopherjs/gopherjs

## License

MIT license, found in LICENSE file

### Any information on bugs or reccomendations on how to make this tool better would be greatly appreciated.
