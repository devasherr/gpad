package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/devasherr/gpad/internal/files"
	"github.com/devasherr/gpad/internal/parser"
)

func printVerbose(start time.Time, initialAllocation, currentAllocation uintptr) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "──────────────────────────────────────────────────────────")
	fmt.Fprintf(w, "Time Taken:\t%s\n", time.Since(start))
	fmt.Fprintf(w, "Initial Allocation:\t%d bytes\n", initialAllocation)
	fmt.Fprintf(w, "Current Allocation:\t%d bytes\n", currentAllocation)

	diff := int(initialAllocation) - int(currentAllocation)
	percent := float64(diff) / float64(initialAllocation) * 100
	fmt.Fprintf(w, "Saved:\t%d bytes (%.2f%%)\n", diff, percent)
	fmt.Fprintln(w, "──────────────────────────────────────────────────────────")
	w.Flush()
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	input := flag.String("path", "", "directory entry")
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	filePaths := files.CollectFiles(path + *input)

	start := time.Now()
	initialAllocation, currentAllocation := parser.ParseFiles(filePaths, *verbose)

	if *verbose {
		printVerbose(start, uintptr(initialAllocation), uintptr(currentAllocation))
	}
}
