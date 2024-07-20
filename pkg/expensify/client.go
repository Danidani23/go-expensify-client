package expensify

// ExpensifyClient struct to hold the client information
type ExpensifyClient struct {
	partnerUserID     string
	partnerUserSecret string
	fileExportConfig  *fileExportRequest
}

// NewClient creates a new ExpensifyClient
func NewClient(partnerUserID, partnerUserSecret string) (*ExpensifyClient, error) {
	return &ExpensifyClient{
		partnerUserID:     partnerUserID,
		partnerUserSecret: partnerUserSecret,
	}, nil
}
