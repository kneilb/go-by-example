// https://gobyexample.com/command-line-subcommands
//
// go build command-line-subcommands.go
//
// Try out foo:
//   ./command-line-subcommands foo -enable -name=joe a1 a2
//
// Try out bar:
//   ./command-line-subcommands bar -level 8 a1
//
// But bar won't accept foo's flags.
//   ./command-line-subcommands bar -enable a1

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Define a subcommand using NewFlagSet, and define flags specific to this subcommand.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Add another subcommand, with its own flags.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// The subcommand is the first argument to the program.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Check which is invoked, and deal with it.
	//
	// For each, we parse its flags, and can access trailing positional arguments.
	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())

	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())

	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
