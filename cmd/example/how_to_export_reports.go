package main

import (
	"context"
	"fmt"
	"github.com/Danidani23/go-expensify-client/pkg/expensify"
	"log"
	"time"
)

func main() {

	// Set up your client
	c, err := expensify.NewClient("userID", "userSecret")
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
		OutputFileExtension:           "pdf",
		OutputFileBaseName:            nil,
		OutputIncludeFullPageReceipts: true,
		IsThisAtestCall:               true,
	}

	emailConf := expensify.OnFinishSendEmail{
		Message:    "hello from Go",
		Recipients: []string{"email@domain.com"},
	}
	// Commit your configuration
	err = c.ConfigureFileExport(config, &emailConf, nil, nil)

	if err != nil {
		log.Fatalln("error while configuring file export: ", err)
	}
	log.Println("configuration successful")

	// Execute the export
	incData, err := c.ExecuteFileExport(context.Background())
	if err != nil {
		log.Fatalln("error while fetching the file: ", err)
	}
	log.Println("response received")
	fmt.Println(string(incData))
}
