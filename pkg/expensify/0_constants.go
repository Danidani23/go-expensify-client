package expensify

const (
	baseURL    = "https://integrations.expensify.com/Integration-Server/ExpensifyIntegrations"
	timeFormat = "2006-01-02"
)

var validReportStates = []string{"OPEN", "SUBMITTED", "APPROVED", "REIMBURSED", "ARCHIVED"}

var validFileExtensions = []string{"pdf", "json"}
