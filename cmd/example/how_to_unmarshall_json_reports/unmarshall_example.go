package main

import (
	"encoding/json"
	"log"
	"os"
)

// Report represents a single report in the JSON structure
type Report struct {
	AccountEmail         string        `json:"accountEmail"`
	AccountID            string        `json:"accountID"`
	ActionList           []Action      `json:"actionList"`
	Approved             string        `json:"approved"`
	Approvers            []Approver    `json:"approvers"`
	Created              string        `json:"created"`
	Currency             string        `json:"currency"`
	CustomField          string        `json:"customField.Name_of_Report_Field"`
	EntryID              string        `json:"entryID"`
	IsACHReimbursed      bool          `json:"isACHReimbursed"`
	ManagerEmail         string        `json:"managerEmail"`
	ManagerUserID        string        `json:"managerUserID"`
	ManagerPayrollID     string        `json:"managerPayrollID"`
	ManagerFirstName     string        `json:"manager.firstName"`
	ManagerLastName      string        `json:"manager.lastName"`
	ManagerFullName      string        `json:"manager.fullName"`
	PolicyName           string        `json:"policyName"`
	PolicyID             string        `json:"policyID"`
	Reimbursed           string        `json:"reimbursed"`
	ReportID             string        `json:"reportID"`
	OldReportID          string        `json:"oldReportID"`
	ReportName           string        `json:"reportName"`
	Status               string        `json:"status"`
	Submitted            string        `json:"submitted"`
	EmployeeCustomField1 string        `json:"employeeCustomField1"`
	EmployeeCustomField2 string        `json:"employeeCustomField2"`
	SubmitterFirstName   string        `json:"submitter.firstName"`
	SubmitterLastName    string        `json:"submitter.lastName"`
	SubmitterFullName    string        `json:"submitter.fullName"`
	Total                float64       `json:"total"`
	TransactionList      []Transaction `json:"transactionList"`
}

// Action represents an action in the actionList
type Action struct {
	Action string `json:"action"`
}

// Approver represents an approver in the approvers list
type Approver struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

// Transaction represents each transaction in the transactionList
type Transaction struct {
	Amount                 string `json:"amount"`
	Attendees              string `json:"attendees"`
	Bank                   string `json:"bank"`
	Billable               string `json:"billable"`
	Category               string `json:"category"`
	CategoryGlCode         string `json:"categoryGlCode"`
	CategoryPayrollCode    string `json:"categoryPayrollCode"`
	Comment                string `json:"comment"`
	ConvertedAmount        string `json:"convertedAmount"`
	Created                string `json:"created"`
	Currency               string `json:"currency"`
	CurrencyConversionRate string `json:"currencyConversionRate"`
	HasTax                 string `json:"hasTax"`
	Inserted               string `json:"inserted"`
	MCC                    string `json:"mcc"`
	Merchant               string `json:"merchant"`
	ModifiedAmount         string `json:"modifiedAmount"`
	ModifiedCreated        string `json:"modifiedCreated"`
	ModifiedMCC            string `json:"modifiedMCC"`
	ModifiedMerchant       string `json:"modifiedMerchant"`
	NtagX                  string `json:"ntagX"`
	NtagXGlCode            string `json:"ntagXGlCode"`
	ReceiptFilename        string `json:"receiptFilename"`
	ReceiptID              string `json:"receiptID"`
	ReceiptSmallThumbnail  string `json:"receiptObject.smallThumbnail"`
	ReceiptThumbnail       string `json:"receiptObject.thumbnail"`
	ReceiptTransactionID   string `json:"receiptObject.transactionID"`
	ReceiptType            string `json:"receiptObject.type"`
	ReceiptURL             string `json:"receiptObject.url"`
	Reimbursable           string `json:"reimbursable"`
	ReportID               string `json:"reportID"`
	Tag                    string `json:"tag"`
	TagGlCode              string `json:"tagGlCode"`
	TaxAmount              string `json:"taxAmount"`
	ModifiedTaxAmount      string `json:"modifiedTaxAmount"`
	TaxName                string `json:"taxName"`
	TaxRate                string `json:"taxRate"`
	TaxRateName            string `json:"taxRateName"`
	TaxCode                string `json:"taxCode"`
	TransactionID          string `json:"transactionID"`
	Type                   string `json:"type"`
	UnitsCount             string `json:"units.count"`
	UnitsRate              string `json:"units.rate"`
	UnitsUnit              string `json:"units.unit"`
}

func main() {

	// I am reading a previously exported file this file contains all standard Report and Expense fields
	inBytes, err := os.ReadFile("temp/export4d21c24c-6257-48d0-8545-d3262756c0ab.json")
	if err != nil {
		log.Fatalln("error while opening the file: ", err)
	}

	myReport := make([]Report, 0)
	err = json.Unmarshal(inBytes, &myReport)
	if err != nil {
		log.Fatalln("error while unmarshalling the incoming data: ", err)
	}

	log.Printf("%v", myReport)

	log.Println("process finished successfully")

}
