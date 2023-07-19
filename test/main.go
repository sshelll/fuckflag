package main

import (
	"fmt"
	"os"

	"github.com/sshelll/fuckflag"
)

var (
	id      = 0
	version = fuckflag.Bool("v", false, "print version")
	echo    = fuckflag.String("echo", "", "echo string")
)

func main() {
	os.Args = []string{"./test"}
	runTest()
	os.Args = []string{"./test", "-v", "-echo=hello world", "-hack", "--bench", "-test.run=x", "-gcflags=\"all=-l -N\"", "-race"}
	runTest()
	os.Args = []string{"./test", "-echo=hello world", "-x", "-gcflags=\"all=-l -N\""}
	runTest()
}

func runTest() {
	id++
	fmt.Printf("id: %v\n", id)
	fuckflag.Parse()
	if *version {
		println("v0.0.1")
	}
	if *echo != "" {
		println("echo: " + *echo)
	}
	exts := fuckflag.Extends()
	fmt.Printf("fuckflag.Extends() size=%d, value=%v\n", len(exts), exts)
	fmt.Printf("version.set=%v, echo.set=%v\n\n", fuckflag.IsSet("v"), fuckflag.IsSet("echo"))
	// reset
	fuckflag.CommandLine = fuckflag.NewFlagSet(os.Args[0], fuckflag.ExitOnError)
	version = fuckflag.Bool("v", false, "print version")
	echo = fuckflag.String("echo", "", "echo string")
}
