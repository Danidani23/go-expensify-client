package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinate(t *testing.T) {
	words := []string{"OPEN", "SUBMITTED", "APPROVED", "REIMBURSED", "ARCHIVED"}
	var result []string
	Combinate(words, []string{}, &result)

	assert.Equal(t, []string{"OPEN", "OPEN,SUBMITTED", "OPEN,SUBMITTED,APPROVED",
		"OPEN,SUBMITTED,APPROVED,REIMBURSED", "OPEN,SUBMITTED,APPROVED,REIMBURSED,ARCHIVED",
		"OPEN,SUBMITTED,APPROVED,ARCHIVED", "OPEN,SUBMITTED,REIMBURSED", "OPEN,SUBMITTED,REIMBURSED,ARCHIVED",
		"OPEN,SUBMITTED,ARCHIVED", "OPEN,APPROVED", "OPEN,APPROVED,REIMBURSED", "OPEN,APPROVED,REIMBURSED,ARCHIVED",
		"OPEN,APPROVED,ARCHIVED", "OPEN,REIMBURSED", "OPEN,REIMBURSED,ARCHIVED", "OPEN,ARCHIVED", "SUBMITTED",
		"SUBMITTED,APPROVED", "SUBMITTED,APPROVED,REIMBURSED", "SUBMITTED,APPROVED,REIMBURSED,ARCHIVED",
		"SUBMITTED,APPROVED,ARCHIVED", "SUBMITTED,REIMBURSED", "SUBMITTED,REIMBURSED,ARCHIVED",
		"SUBMITTED,ARCHIVED", "APPROVED", "APPROVED,REIMBURSED", "APPROVED,REIMBURSED,ARCHIVED",
		"APPROVED,ARCHIVED", "REIMBURSED", "REIMBURSED,ARCHIVED", "ARCHIVED"}, result)
}
