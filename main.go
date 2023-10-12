package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// 1. user option i.e wc -l
	// 2. file open / read
	// 3. line by line (stream / buffer)
	line := flag.Bool("l", true, "line count")

	*line = true

	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	file, err := os.Open(files[0])
	if err != nil {
		if os.IsPermission(err) {
			fmt.Fprintf(os.Stderr, "%s: %s: open: Permission denied\n", os.Args[0], files[0])
			os.Exit(-1)
		} else if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s: %s: open: No such file or directories\n", os.Args[0], files[0])
			os.Exit(-1)
		} else {
			fmt.Fprintf(os.Stderr, "%s: %s: open: %s\n", os.Args[0], files[0], err.Error())
			os.Exit(-1)
		}

	}
	defer file.Close() // Ensure the file is closed when we're done with it

	fileInfo, err := file.Stat()
	if err != nil {
		os.Exit(-1)
	}
	if fileInfo.IsDir() {
		fmt.Fprintf(os.Stderr, "%s: %s: read: is a directory\n", os.Args[0], files[0])
		os.Exit(-1)
	}

	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}

	fmt.Printf("%8d %s\n", lines, files[0])
}

func init() {
	// Customize the Usage function to include information about required arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  <file>   The file to process\n")
	}
}
