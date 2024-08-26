package main

import (
	"context"
	"github.com/Danidani23/go-expensify-client/pkg/expensify"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

// Export level field names in Expensify
var reportFieldNames = []string{
	"accountEmail",
	"accountID",
	"actionList",
	"approved",
	"approvers",
	"created",
	"currency",
	"customField.Name_of_Report_Field",
	"entryID",
	"isACHReimbursed",
	"managerEmail",
	"managerUserID",
	"managerPayrollID",
	"manager.firstName",
	"manager.lastName",
	"manager.fullName",
	"policyName",
	"policyID",
	"reimbursed",
	"reportID",
	"oldReportID",
	"reportName",
	"status",
	"submitted",
	"employeeCustomField1",
	"employeeCustomField2",
	"submitter.firstName",
	"submitter.lastName",
	"submitter.fullName",
	"total",
	"transactionList",
}

var expenseFieldNames = []string{
	"amount",
	"attendees",
	"bank",
	"billable",
	"category",
	"categoryGlCode",
	"categoryPayrollCode",
	"comment",
	"convertedAmount",
	"created",
	"currency",
	"currencyConversionRate",
	"hasTax",
	"inserted",
	"mcc",
	"merchant",
	"modifiedAmount",
	"modifiedCreated",
	"modifiedMCC",
	"modifiedMerchant",
	"ntagX",
	"ntagXGlCode",
	"receiptFilename",
	"receiptID",
	"receiptObject.smallThumbnail",
	"receiptObject.thumbnail",
	"receiptObject.transactionID",
	"receiptObject.type",
	"receiptObject.url",
	"reimbursable",
	"reportID",
	"tag",
	"tagGlCode",
	"taxAmount",
	"modifiedTaxAmount",
	"taxName",
	"taxRate",
	"taxRateName",
	"taxCode",
	"transactionID",
	"type",
	"units.count",
	"units.rate",
	"units.unit",
}

func main() {
	// Load the .env file - you can ignore that part, I am using a .env file to store my secrets for testing
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// you can use context if you like
	ctx := context.Background()

	// Get your secrets from env variables (or whichever way you prefer)
	userID := os.Getenv("USERID")
	userSecret := os.Getenv("USERSECRET")
	myEmail := os.Getenv("MYEMAIL")

	if userID == "" || userSecret == "" {
		log.Fatalln("either userId or user userSecret is empty!")
	}

	// Set up your client
	c, err := expensify.NewClient(userID, userSecret)
	if err != nil {
		log.Fatalln(err)
	}

	// Configure your export
	filterDate, err := time.Parse(time.DateOnly, "2024-01-01")
	if err != nil {
		log.Fatalln("couldn't parse date: ", err)
	}

	config := expensify.FileExportConfig{
		FilterByReportId:              nil,
		FilterByPolicyId:              nil,
		FilterByStartDate:             &filterDate,
		FilterByEndDate:               nil,
		FilterByApprovedAfterDate:     nil,
		FilterByMarkedAsApprovedTag:   nil,
		FilterByEmployeeEmail:         nil,
		FilterByReportState:           nil,
		LimitNumberOfReportsExported:  nil,
		OutputFileExtension:           "json", // we only allow PDF or JSON
		OutputFileBaseName:            nil,
		OutputIncludeFullPageReceipts: false,
		IsThisAtestCall:               true,
	}

	emailConf := expensify.OnFinishSendEmail{
		Message:    "hello from Go",
		Recipients: []string{myEmail},
	}
	// Commit your configuration
	err = c.ConfigureFileExport(config, &emailConf, nil, nil)

	if err != nil {
		log.Fatalln("error while configuring file export: ", err)
	}
	log.Println("configuration successful")

	// Execute the export - you can configure which report fields and which expense fields you need
	createdReports, err := c.ExecuteFileExport(ctx, reportFieldNames, expenseFieldNames)
	if err != nil {
		log.Fatalln("error while fetching the file: ", err)
	}
	log.Println("response received")
	log.Printf("created reports are: %s", createdReports)

	// DownloadReport the report(s)
	for _, report := range createdReports {
		log.Printf("starting to download this report: %s", report)
		err = c.DownloadReport(ctx, report)
		if err != nil {
			log.Fatalf("error while downloading this report: '%v', msg: %s", report, err)
		}
	}
	log.Println("reports successfully downloaded")

	// Write them out to disk if you like
	for _, report := range createdReports {
		err = report.WriteToDisk(ctx, "temp")
		if err != nil {
			log.Fatalln(err)
		}
	}

}
