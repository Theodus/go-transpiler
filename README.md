# tardisgo-compileutil

Command-Line utility for compiling Go to Java, C++, and C# using tardisgo

## Usage:

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

## For more information on TARDIS Go:

Check out the github repository at https://github.com/tardisgo/tardisgo
or the website at https://tardisgo.github.io

### Any information on bugs or reccomendations on how to make this tool better would be greatly appreciated.
