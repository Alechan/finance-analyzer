package demodataset

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Alechan/finance-analyzer/pkg/internal/pdfcardsummary"
	"github.com/stretchr/testify/require"
)

func TestPublicWebDemoDataset_CanBeParsedIntoRepoCompatibleStructs(t *testing.T) {
	csvPath := filepath.Join("..", "..", "web", "mockups_lab", "tmp_public_data", "current", "demo_extracted.csv")
	csvBytes, err := os.ReadFile(csvPath)
	require.NoError(t, err)

	rows, err := pdfcardsummary.ParseMovementsWithCardContextCSV(csvBytes)
	require.NoError(t, err)
	require.NotEmpty(t, rows)
}
