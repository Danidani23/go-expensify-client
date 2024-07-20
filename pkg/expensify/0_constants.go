package expensify

const (
	baseURL = "https://integrations.expensify.com/Integration-Server/ExpensifyIntegrations"
)

var validReportStates = []string{"OPEN", "SUBMITTED", "APPROVED", "REIMBURSED", "ARCHIVED"}

var timeFormat = "2006-01-02"

var validFileExtensions = []string{"csv", "xls", "xlsx", "txt", "pdf", "json", "xml"}
