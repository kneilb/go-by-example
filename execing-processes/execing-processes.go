// https://gobyexample.com/execing-processes
//
// Instead of spawning separate processes, we can also just exec them.
// This replaces the current Go process with the new one.
// We use the Go implementation of exec https://en.wikipedia.org/wiki/Exec_(operating_system) to do this.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// For our example, we'll exec ls.
	// Go requires an absolute path, so we use exec.LookPath to find it (probably /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires args in slice form (not one big string).
	// We give ls some common arguments - the first one is the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use.
	// Here we just provide our own environment.
	env := os.Environ()

	// This is the actual syscall.Exec call.
	// If it's successful, our process will end here and be replaced by the one defined above.
	// If there's an error, we'll get a return value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
