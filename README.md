# tardisgo-compileutil

Command-Line utility for compiling Go to Java, C++, or C# using tardisgo
or to JavaScript using GopherJS

## Usage

Be sure that your GOPATH is set and the bin directory is added to PATH:
```
export GOPATH="path/to/gocode"
export PATH=$PATH:$GOPATH/bin
```
Be sure that Haxe version 3.2.0 is installed

Go get and install this package:
```
go get -u github.com/theodus/tardisgo-compileutil
```
Compile your code to Java:
```
tardisgo-compileutil java myGoPackage
```
Or C++:
```
tardisgo-compileutil cpp myGoPackage
```
Or C#:
```
tardisgo-compileutil cs myGoPackage
```
Or JS:
```
tardisgo-compileutil js myGoPackage
```

## More information on TARDIS Go

Check out the Github repository at https://github.com/tardisgo/tardisgo
or the website at https://tardisgo.github.io

## More information on GopherJS

Check out the Github repository at https://github.com/gopherjs/gopherjs

## License

MIT license, found in LICENSE file

### Any information on bugs or reccomendations on how to make this tool better would be greatly appreciated.
