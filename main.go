package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Parse the command-line arguments
	force := flag.Bool("force", false, "Overwrite the file if it already exists")
	flag.BoolVar(force, "f", false, "Overwrite the file if it already exists (shorthand)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] FILENAME\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Get the filename from the command-line arguments
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	filename := flag.Arg(0)

	// Check if the file already exists
	_, err := os.Stat(filename)
	if err == nil && !*force {
		log.Fatalf("File '%s' already exists. Use --force or -f option to overwrite.", filename)
	}

	// Create or open the file for writing
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Capture Ctrl+C signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Create a goroutine to copy input from stdin to the file
	go func() {
		_, err := io.Copy(file, os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Enter input (Ctrl+C to exit):")

	// Wait for Ctrl+C signal
	<-sigCh

	fmt.Println("Output has been written to the file.")
}
