package main

import (
	"fmt"
	"testing"

	"github.com/alexflint/go-arg"
	"github.com/stretchr/testify/require"
)

func TestReadArgs_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantBank BankType
		wantPDFs []string
	}{
		{
			name:     "Santander with multiple PDFs",
			args:     []string{"--bank", "santander", "file1.pdf", "file2.pdf"},
			wantBank: Santander,
			wantPDFs: []string{"file1.pdf", "file2.pdf"},
		},
		{
			name:     "Visa-Prisma with one PDF",
			args:     []string{"-b", "visa-prisma", "file3.pdf"},
			wantBank: VisaPrisma,
			wantPDFs: []string{"file3.pdf"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			gotArgs, err := parseArgs(tt.args)

			// Then
			require.NoError(t, err)
			require.Equal(t, tt.wantBank, gotArgs.Bank)
			require.Equal(t, tt.wantPDFs, gotArgs.PDFs)
		})
	}
}

func TestReadArgs_InvalidCases(t *testing.T) {
	helpMessage := "--bank BANK [--join-csvs JOIN-CSVS] PDFS [PDFS ...]\n\nPositional arguments:\n  PDFS                   List of PDF files to parse\n\nOptions:\n  --bank BANK, -b BANK   Bank type (santander or visa-prisma)\n  --join-csvs JOIN-CSVS\n                         Path to combined CSV output file. When specified, all PDFs are processed and their CSV outputs are combined into a single file with headers appearing only once. Individual CSV files are also created.\n  --help, -h             display this help and exit\n"
	tests := []struct {
		name      string
		args      []string
		baseError string
	}{
		{
			name:      "Missing --bank",
			args:      []string{"file1.pdf"},
			baseError: "error parsing arguments: --bank is required",
		},
		{
			name:      "Invalid bank type",
			args:      []string{"--bank", "invalid-bank", "file1.pdf"},
			baseError: "error validating arguments: error processing --bank: bank type invalid-bank is not one of the supported banks: [santander visa-prisma]",
		},
		{
			name:      "No PDFs provided",
			args:      []string{"-b", "santander"},
			baseError: "error parsing arguments: pdfs is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			_, err := parseArgs(tt.args)

			// Then
			require.Contains(t, err.Error(), tt.baseError)
			require.Contains(t, err.Error(), helpMessage)
		})
	}
}

func TestArgs_Validate(t *testing.T) {
	tests := []struct {
		name    string
		args    Args
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid Santander bank",
			args: Args{
				Bank: Santander,
				PDFs: []string{"file1.pdf"},
			},
			wantErr: false,
		},
		{
			name: "Valid Visa-Prisma bank",
			args: Args{
				Bank: VisaPrisma,
				PDFs: []string{"file1.pdf"},
			},
			wantErr: false,
		},
		{
			name: "Invalid bank type",
			args: Args{
				Bank: "invalid-bank",
				PDFs: []string{"file1.pdf"},
			},
			wantErr: true,
			errMsg:  "bank type invalid-bank is not one of the supported banks",
		},
		{
			name: "Empty bank type",
			args: Args{
				Bank: "",
				PDFs: []string{"file1.pdf"},
			},
			wantErr: true,
			errMsg:  "bank type  is not one of the supported banks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.Validate()
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetHelpString(t *testing.T) {
	// Create a parser with some args
	var args Args
	p, err := arg.NewParser(arg.Config{}, &args)
	require.NoError(t, err)

	// Test that help string contains expected elements
	helpStr := getHelpString(p)
	require.Contains(t, helpStr, "--bank")
	require.Contains(t, helpStr, "PDFS")
	require.Contains(t, helpStr, "Bank type")
	require.Contains(t, helpStr, "List of PDF files to parse")
}

func TestAddHelpToError(t *testing.T) {
	// Create a parser
	var args Args
	p, err := arg.NewParser(arg.Config{}, &args)
	require.NoError(t, err)

	// Test that error is properly wrapped with help
	baseErr := fmt.Errorf("test error")
	errWithHelp := addHelpToError(p, baseErr)

	require.NotNil(t, errWithHelp)

	require.Contains(t, errWithHelp.Error(), "test error")
	require.Contains(t, errWithHelp.Error(), "--bank")
	require.Contains(t, errWithHelp.Error(), "PDFS")
}

func TestReadArgs_AdditionalEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		baseError string
	}{
		{
			name:      "Empty PDF list",
			args:      []string{"--bank", "santander", ""},
			baseError: "empty PDF file name provided",
		},
		{
			name:      "Invalid PDF extension",
			args:      []string{"--bank", "santander", "file.txt"},
			baseError: "file file.txt does not have a .pdf extension",
		},
		{
			name:      "Multiple bank flags",
			args:      []string{"--bank", "santander", "-b", "visa-prisma", "file.pdf"},
			baseError: "error parsing arguments: bank flag (--bank or -b) can only be specified once",
		},
		{
			name:      "Bank type with spaces",
			args:      []string{"--bank", "santander bank", "file.pdf"},
			baseError: "bank type santander bank is not one of the supported banks",
		},
		{
			name:      "Bank type with special characters",
			args:      []string{"--bank", "santander@bank", "file.pdf"},
			baseError: "bank type santander@bank is not one of the supported banks",
		},
		{
			name:      "Empty bank type",
			args:      []string{"--bank", "", "file.pdf"},
			baseError: "bank type  is not one of the supported banks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			_, err := parseArgs(tt.args)

			// Then
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.baseError)
		})
	}
}

func TestReadArgs_ValidCases_Additional(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantBank BankType
		wantPDFs []string
	}{
		{
			name:     "Santander with PDF in different directory",
			args:     []string{"--bank", "santander", "/path/to/file.pdf"},
			wantBank: Santander,
			wantPDFs: []string{"/path/to/file.pdf"},
		},
		{
			name:     "Visa-Prisma with multiple PDFs in different directories",
			args:     []string{"-b", "visa-prisma", "file1.pdf", "/path/to/file2.pdf", "./relative/file3.pdf"},
			wantBank: VisaPrisma,
			wantPDFs: []string{"file1.pdf", "/path/to/file2.pdf", "./relative/file3.pdf"},
		},
		{
			name:     "Santander with PDF containing spaces",
			args:     []string{"--bank", "santander", "my file with spaces.pdf"},
			wantBank: Santander,
			wantPDFs: []string{"my file with spaces.pdf"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			gotArgs, err := parseArgs(tt.args)

			// Then
			require.NoError(t, err)
			require.Equal(t, tt.wantBank, gotArgs.Bank)
			require.Equal(t, tt.wantPDFs, gotArgs.PDFs)
		})
	}
}

func TestReadArgs_JoinCSVFlag(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantBank    BankType
		wantPDFs    []string
		wantJoinCSV *string
	}{
		{
			name:        "With join-csvs flag",
			args:        []string{"--bank", "santander", "--join-csvs", "combined.csv", "file1.pdf", "file2.pdf"},
			wantBank:    Santander,
			wantPDFs:    []string{"file1.pdf", "file2.pdf"},
			wantJoinCSV: stringPtr("combined.csv"),
		},
		{
			name:        "Without join-csvs flag",
			args:        []string{"--bank", "santander", "file1.pdf"},
			wantBank:    Santander,
			wantPDFs:    []string{"file1.pdf"},
			wantJoinCSV: nil,
		},
		{
			name:        "With join-csvs flag and visa-prisma",
			args:        []string{"-b", "visa-prisma", "--join-csvs", "/path/to/combined.csv", "file1.pdf"},
			wantBank:    VisaPrisma,
			wantPDFs:    []string{"file1.pdf"},
			wantJoinCSV: stringPtr("/path/to/combined.csv"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			gotArgs, err := parseArgs(tt.args)

			// Then
			require.NoError(t, err)
			require.Equal(t, tt.wantBank, gotArgs.Bank)
			require.Equal(t, tt.wantPDFs, gotArgs.PDFs)
			if tt.wantJoinCSV == nil {
				require.Nil(t, gotArgs.JoinCSV)
			} else {
				require.NotNil(t, gotArgs.JoinCSV)
				require.Equal(t, *tt.wantJoinCSV, *gotArgs.JoinCSV)
			}
		})
	}
}

func TestArgs_Validate_JoinCSV(t *testing.T) {
	tests := []struct {
		name    string
		args    Args
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid JoinCSV path",
			args: Args{
				Bank:    Santander,
				PDFs:    []string{"file1.pdf"},
				JoinCSV: stringPtr("combined.csv"),
			},
			wantErr: false,
		},
		{
			name: "Empty JoinCSV path",
			args: Args{
				Bank:    Santander,
				PDFs:    []string{"file1.pdf"},
				JoinCSV: stringPtr(""),
			},
			wantErr: true,
			errMsg:  "empty path provided for --join-csvs flag",
		},
		{
			name: "JoinCSV path with whitespace only",
			args: Args{
				Bank:    Santander,
				PDFs:    []string{"file1.pdf"},
				JoinCSV: stringPtr("   "),
			},
			wantErr: true,
			errMsg:  "empty path provided for --join-csvs flag",
		},
		{
			name: "Nil JoinCSV (flag not provided)",
			args: Args{
				Bank:    Santander,
				PDFs:    []string{"file1.pdf"},
				JoinCSV: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.Validate()
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetHelpString_JoinCSV(t *testing.T) {
	// Create a parser with some args
	var args Args
	p, err := arg.NewParser(arg.Config{}, &args)
	require.NoError(t, err)

	// Test that help string contains expected elements
	helpStr := getHelpString(p)
	require.Contains(t, helpStr, "--join-csvs")
	require.Contains(t, helpStr, "combined CSV output file")
}

// stringPtr is a helper function to create a pointer to a string
func stringPtr(s string) *string {
	return &s
}
