# go-expensify-client

[![Build Status](https://github.com/Danidani23/go-expensify-client/actions/workflows/go.yml/badge.svg)](https://github.com/Danidani23/go-expensify-client/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Danidani23/go-expensify-client)](https://goreportcard.com/report/github.com/Danidani23/go-expensify-client)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Danidani23/go-expensify-client)](https://pkg.go.dev/github.com/Danidani23/go-expensify-client)

go-expensify is a Go package designed to provide a simple and efficient way to interact with Expensify’s API. This library enables developers to seamlessly integrate Expensify’s expense management capabilities into their Go applications, making it easier to manage expenses and reports programmatically.

## Features


### Report Exports

This wrapper is very opinionated, meaning it hides the complexity but also
some of the functions the API offers. I have done that because some of the
features does not work, and some others in my opinion not so useful.

Export extension:
In general the API offers 'csv' or 'pdf' export. In the documentation you
may read about other formats (like JSON) but you always get a 'csv' or 'pdf'
back, regardless of the configuration you set. 

I am working on the rest.


## Before you start

You can read the API docs of Expensify here:
https://integrations.expensify.com/Integration-Server/doc/#export

The client interacts with this API. (All features available at the day of release)

You need to create your credentials to interract with the API. You can do that here:
https://www.expensify.com/tools/integrations/

May the force be with you! :)

## How to install

```go get github.com/Danidani23/go-expensify-client```

## Examples

You can find examples on how to use the package in ./cmd/example

## Notes

### Timeout
The exporter API of Expensify may get a bit slow, so it may take a while before you get a response (mostly for PDF-s). 
You should consider that in case you are setting a context with timeout. 

#### Limitations

While the native API of expensify offers various export file formats, one hast to configure the export in Freemarker
templates. 
This wrapper hides this complexity, in exchange for limiting the output formats available. This package offers 2 formats:
- JSON
- PDF

Pleas see the attached examples how to get your reports.


## Scraping receipts and invoices

Unfortunately there is no official API endpoint to get the documents and images you have uploaded to your expenses.
I have created an automated scraper that fetches these documents for you.
It requires a set of session cookies to work. You can take a look at the provided example file.

## Known Issues

When you configure email sending at the end of an export, the email never arrives. We suspect this is an error  on 
Expensify's side, as we get no error messages back when we place the call.