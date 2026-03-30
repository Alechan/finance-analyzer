package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/alexflint/go-arg"
)

// parseArgs parses the command line arguments. It is usually called passing os.Args[1:].
func parseArgs(rawArgs []string) (Args, error) {
	// Check for duplicate bank flags
	a, err2 := preValidation(rawArgs)
	if err2 != nil {
		return a, err2
	}

	var args Args
	p, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		return Args{}, fmt.Errorf("error creating parser: %w", err)
	}
	err = p.Parse(rawArgs)
	if err != nil {
		baseError := fmt.Errorf("error parsing arguments: %w", err)
		return Args{}, addHelpToError(p, baseError)
	}

	err = args.Validate()
	if err != nil {
		baseError := fmt.Errorf("error validating arguments: %w", err)
		return Args{}, addHelpToError(p, baseError)
	}
	return args, nil
}

// preValidation performs preliminary checks on the raw arguments to ensure they meet basic requirements.
func preValidation(rawArgs []string) (Args, error) {
	bankFlagCount := 0
	for i := 0; i < len(rawArgs); i++ {
		if rawArgs[i] == "--bank" || rawArgs[i] == "-b" {
			bankFlagCount++
			if bankFlagCount > 1 {
				return Args{}, fmt.Errorf("error parsing arguments: bank flag (--bank or -b) can only be specified once")
			}
		}
	}
	return Args{}, nil
}

// Args defines the command line arguments.
type Args struct {
	Bank    BankType `arg:"-b,--bank,required" help:"Bank type (santander or visa-prisma)"`
	PDFs    []string `arg:"positional,required" help:"List of PDF files to parse"`
	JoinCSV *string  `arg:"--join-csvs" help:"Path to combined CSV output file. When specified, all PDFs are processed and their CSV outputs are combined into a single file with headers appearing only once. Individual CSV files are also created."`
}

func (a Args) Validate() error {
	// Validate bank is one of the valid banks
	if !slices.Contains(validBanks, a.Bank) {
		return fmt.Errorf(
			"error processing --bank: bank type %s is not one of the supported banks: %v",
			a.Bank,
			validBanks,
		)
	}

	// Validate PDFs
	if len(a.PDFs) == 0 {
		return fmt.Errorf("no PDF files provided")
	}

	for _, pdf := range a.PDFs {
		// Check for empty PDF name
		if strings.TrimSpace(pdf) == "" {
			return fmt.Errorf("empty PDF file name provided")
		}

		// Check file extension
		ext := strings.ToLower(filepath.Ext(pdf))
		if ext != ".pdf" {
			return fmt.Errorf("file %s does not have a .pdf extension", pdf)
		}
	}

	// Validate JoinCSV path if provided
	if a.JoinCSV != nil {
		if strings.TrimSpace(*a.JoinCSV) == "" {
			return fmt.Errorf("empty path provided for --join-csvs flag")
		}
	}

	return nil
}

func getHelpString(p *arg.Parser) string {
	var buf bytes.Buffer
	p.WriteHelp(&buf)
	result := buf.String() // Output is in result
	return result
}

func addHelpToError(p *arg.Parser, baseError error) error {
	helpStr := getHelpString(p)
	errorWithHelp := fmt.Errorf("%w\n\n%s", baseError, helpStr)
	return errorWithHelp
}
