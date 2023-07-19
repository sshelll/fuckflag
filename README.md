# fuckflag

A golang `flag` lib which support undefined flags.



## Story

Here's the thing... When I tried to create a command line tool with golang, I found the `flag.Parse()` will throw an error like this:

```go
flag provided but not defined: -xxx
```

So I modified the src code, and that's how `fuckflag` was born.


## Install
`go get -u github.com/sshelll/fuckflag`



## Demo

```go
package main

var (
	version = fuckflag.Bool("v", false, "print version")
	echo    = fuckflag.String("echo", "", "echo string")
)

func main() {
	fuckflag.Parse()
	if *version {
		println("v0.0.1")
	}
	if *echo != "" {
		println("echo: " + *echo)
	}
	// call this func to get undefined flags
	exts := fuckflag.Extends()
	fmt.Printf("fuckflag.Extends() size=%d, value=%v\n", len(exts), exts)
	// call this the check if a flag was passed
	fmt.Printf("version.set=%v, echo.set=%v\n\n", fuckflag.IsSet("v"), fuckflag.IsSet("echo"))
}
```

**Exec Result**:

```sh
>> go build -o test main.go

>> ./test -v -echo=hello -undef --x -p=1

>> output:
v0.0.1
echo: hello
fuckflag.Extends() size=3, value=[-undef --x -p=1]
version.set=true, echo.set=true
```

