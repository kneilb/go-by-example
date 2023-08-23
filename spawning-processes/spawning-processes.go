// https://gobyexample.com/spawning-processes

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Simple command, takes no arguments & just prints to stdout.
	// The exec.Command helper creates an object to represent this external process.
	dateCmd := exec.Command("date")

	// The Output method runs the command, waits for it to finish and collects its Stdout.
	// If there were no errors, dateOut will hold bytes with the data info.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Output and other Command methods return *exec.Error if there was a problem executing the command (wrong path),
	// and *exec.ExitError if the command ran but exited with a non-zero return code.
	//
	// Exits with rc = 1, because date doesn't have a -x option.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// A more involved case, where we pipe data to the process on stdin, and collect output from stdout.
	grepCmd := exec.Command("grep", "hello")

	// Explicitly grab the In/Out pipes, start the process, write some input, read the output, and wait for it to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// We omitted any checks, but should use the usual err != nil thing.
	// We also only collected StdoutPipe, but could also collect StderrPipe.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// When spawning commands, we need to provide a command & argument array, rather than a string.
	// If you want to run a full command with a string, use bash's -c option:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
