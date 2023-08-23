// https://gobyexample.com/logging
//
// Go supports logging!
// log package for freeform, slog for structured.

package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {

	// The standard logger is preconfigured for "reasonable logging output" to os.Stderr.
	// Additional methods like Fatal* & Panic* will exit the program after logging.
	log.Println("standard logger")

	// Configure loggers with flags to set their output format.
	// By default, the standard logger has log.Ldate & log.Ltime set, which are collected as log.LstdFlags.
	// This changes the default logger to emit microseconds, too.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// We can also emit the file name & line where the log function is called.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// We can also create custom loggers, and pass them around.
	// When creating one, we set a prefix to distinguish it.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// We can also change the prefix (even on the standard logger...)
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Loggers can have custom output targets; any io.Writer works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// This writes the log output into buf.
	buflog.Println("hello")

	// This shows it on stdout.
	fmt.Print("from buflog:", buf.String())

	// The slog package provides structured log output, for example JSON.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// In addition to the message, slog output can contain arbitrary key/value pairs.
	myslog.Info("hello again", "key", "val", "age", 25)
}
