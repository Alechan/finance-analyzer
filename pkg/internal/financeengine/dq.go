package financeengine

import (
	"fmt"
	"slices"
	"time"

	"github.com/Alechan/finance-analyzer/pkg/internal/pdfcardsummary"
)

const (
	DQRuleMissingCategoryMappingForCardMovement = "DQ003"
	DQRuleMissingOwnerMappingForCardMovement    = "DQ004"
)

type Mappings struct {
	OwnersByCardOwner  map[string]string `json:"ownersByCardOwner"`
	OwnersByCardNumber map[string]string `json:"ownersByCardNumber"`
	CategoryByDetail   map[string]string `json:"categoryByDetail"`
}

type DQIssue struct {
	RuleID       string
	Message      string
	MovementType pdfcardsummary.MovementType
	CloseDate    time.Time
	Detail       string
	CardOwner    string
	CardNumber   *string
}

type DQSummaryByRuleRow struct {
	RuleID string
	Count  int
}

func (e *Engine) DataQuality(rows []pdfcardsummary.MovementWithCardContext, mappings Mappings) ([]DQIssue, []DQSummaryByRuleRow) {
	issues := make([]DQIssue, 0)
	countByRule := make(map[string]int)

	for _, row := range rows {
		if row.MovementType != pdfcardsummary.MovementTypeCard || row.CardContext == nil {
			continue
		}

		if !isMappedOwner(row, mappings) {
			issue := DQIssue{
				RuleID:       DQRuleMissingOwnerMappingForCardMovement,
				Message:      fmt.Sprintf("missing owner mapping for CardMovement (owner=%q, card=%q)", row.CardOwner, derefString(row.CardNumber)),
				MovementType: row.MovementType,
				CloseDate:    row.CloseDate,
				Detail:       row.Movement.Detail,
				CardOwner:    row.CardOwner,
				CardNumber:   row.CardNumber,
			}
			issues = append(issues, issue)
			countByRule[issue.RuleID]++
		}

		if !isMappedCategory(row, mappings) {
			issue := DQIssue{
				RuleID:       DQRuleMissingCategoryMappingForCardMovement,
				Message:      fmt.Sprintf("missing category mapping for CardMovement (detail=%q)", row.Movement.Detail),
				MovementType: row.MovementType,
				CloseDate:    row.CloseDate,
				Detail:       row.Movement.Detail,
				CardOwner:    row.CardOwner,
				CardNumber:   row.CardNumber,
			}
			issues = append(issues, issue)
			countByRule[issue.RuleID]++
		}
	}

	slices.SortFunc(issues, func(a, b DQIssue) int {
		if a.RuleID < b.RuleID {
			return -1
		}
		if a.RuleID > b.RuleID {
			return 1
		}
		if a.CloseDate.Before(b.CloseDate) {
			return -1
		}
		if a.CloseDate.After(b.CloseDate) {
			return 1
		}
		if a.Detail < b.Detail {
			return -1
		}
		if a.Detail > b.Detail {
			return 1
		}
		if a.CardOwner < b.CardOwner {
			return -1
		}
		if a.CardOwner > b.CardOwner {
			return 1
		}
		return 0
	})

	summary := make([]DQSummaryByRuleRow, 0, len(countByRule))
	for ruleID, count := range countByRule {
		summary = append(summary, DQSummaryByRuleRow{
			RuleID: ruleID,
			Count:  count,
		})
	}
	slices.SortFunc(summary, func(a, b DQSummaryByRuleRow) int {
		if a.RuleID < b.RuleID {
			return -1
		}
		if a.RuleID > b.RuleID {
			return 1
		}
		return 0
	})

	return issues, summary
}

func isMappedOwner(row pdfcardsummary.MovementWithCardContext, mappings Mappings) bool {
	if mappings.OwnersByCardOwner != nil {
		if _, ok := mappings.OwnersByCardOwner[row.CardOwner]; ok {
			return true
		}
	}

	if mappings.OwnersByCardNumber != nil && row.CardNumber != nil {
		if _, ok := mappings.OwnersByCardNumber[*row.CardNumber]; ok {
			return true
		}
	}

	return false
}

func isMappedCategory(row pdfcardsummary.MovementWithCardContext, mappings Mappings) bool {
	if mappings.CategoryByDetail == nil {
		return false
	}
	_, ok := mappings.CategoryByDetail[row.Movement.Detail]
	return ok
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
