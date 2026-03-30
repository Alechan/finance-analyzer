package main

import (
	"fmt"

	"github.com/Alechan/finance-analyzer/pkg/internal/extractor/santander"
	"github.com/Alechan/finance-analyzer/pkg/internal/extractor/visaprisma"
	"github.com/Alechan/finance-analyzer/pkg/internal/pdfcardsummaryio"
)

// extractorFactory creates a new PDF card summary extractor based on the bank type.
// This function will be extended as more bank types are supported.
func extractorFactory(bankType BankType) (pdfcardsummaryio.Extractor, error) {
	switch bankType {
	case Santander:
		return santander.NewSantanderExtractorFromDefaultCfg(), nil
	case VisaPrisma:
		return visaprisma.NewVisaprismaExtractor(), nil
	default:
		return nil, fmt.Errorf("unsupported bank type: %s", bankType)
	}
}
