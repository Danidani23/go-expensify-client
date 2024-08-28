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

	config := expensify.FileExportBaseConfig{
		FilterByReportId:              nil,
		FilterByPolicyId:              nil,
		FilterByStartDate:             &filterDate,
		FilterByEndDate:               nil,
		FilterByApprovedAfterDate:     nil,
		FilterByMarkedAsApprovedTag:   nil,
		FilterByEmployeeEmail:         nil,
		FilterByReportState:           nil,
		LimitNumberOfReportsExported:  nil,
		OutputFileBaseName:            nil,
		OutputIncludeFullPageReceipts: false,
		IsThisAtestCall:               true,
	}

	// you can configure the below parameters and pass it if you like, but they are optional
	/*
		client.OnFinishMarkAsExportedConfig{}
		client.OnFinishSftpUploadDataConfig{}
		client.OnFinishSendEmailConfig{}
	*/

	myReports, err := c.GetReportsInJson(ctx, config, nil, nil, nil, reportFieldNames, expenseFieldNames)
	if err != nil {
		log.Fatalf("error while fetching the reports: %s", err)
	}

	for _, report := range myReports {
		filename, err := report.WriteToDisk(ctx, "temp")
		if err != nil {
			log.Fatalln("error while writing the report to disk")
		}
		log.Println("the following file was written successfully: ", filename)
	}
}
