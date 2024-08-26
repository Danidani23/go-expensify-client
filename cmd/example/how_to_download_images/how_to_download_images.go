package main

import (
	"github.com/Danidani23/go-expensify-client/v2/pkg/expensify"
	"log"
)

func main() {
	// Step 1: Log in to your expensify account and export the session cookies (you can use a tool like 'EditThisCookie'
	//as a browser extension. It offers JSON export of your cookies.

	// Step 2: Load cookies from JSON file
	cookies, err := expensify.LoadCookiesFromJSON("cmd/example/how_to_download_images/cookies.json") // Replace with your cookies JSON file path
	if err != nil {
		log.Fatalf("error while loading the cookies: %s", err)
	}

	myImage, err := expensify.GetImage("https://www.expensify.com/receipts/w_0413aeadd2d0cf3df6280d3f05955798509daf19.pdf", cookies)
	if err != nil {
		log.Fatalf("error while fetching the image: %s", err)
	}

	err = myImage.WriteToDisk("temp") // we write it out into the temp folder
	if err != nil {
		log.Fatalf("error while writing out the file: %s", err)
	}

}
