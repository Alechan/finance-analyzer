package main

import (
	"fmt"
	"os"

	"github.com/Alechan/finance-analyzer/pkg/internal/pdfcardsummaryio"
	"github.com/Alechan/finance-analyzer/pkg/internal/validation"
)

const (
	exitSuccess = 0
	exitFailure = 1
)

func main() {
	os.Exit(run())
}

// run executes the main program logic and returns the appropriate exit code.
func run() int {
	rawArgs := os.Args[1:]
	args, err := parseArgs(rawArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing arguments:", err)
		return exitFailure
	}

	reader, err := extractorFactory(args.Bank)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating reader:", err)
		return exitFailure
	}

	validator := validation.NewValidator()
	etl := pdfcardsummaryio.NewPDFCardSummaryETL(reader, validator)

	if args.JoinCSV != nil {
		// Join CSV mode: creates both individual CSVs and combined CSV
		err = etl.ETLFilesWithJoinedCSV(args.PDFs, *args.JoinCSV)
	} else {
		// Individual CSV mode (existing behavior)
		err = etl.ETLFilesIndependently(args.PDFs)
	}

	if err != nil {
		errMsg := fmt.Sprintf("Error processing files: %v", err)
		fmt.Fprintln(os.Stderr, errMsg)
		return exitFailure
	}

	return exitSuccess
}
