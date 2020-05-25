package main

import (
"fmt"
"log"
"os"
"io"
"strings"
)

const DUPEFILE_VERSION = "v1.0"

var (
	sourceFile string
	destinationLocations []string
)

func main () {
	displayBanner ()

	count := len (os.Args)

	if (count == 1) { // just the appname no args
		showSyntax ()
		os.Exit (0)
	}

	if (count == 2) {// only appname and file name no target locations
		showError ("No destinations specified", true, -1)
    }

	sourceFile = os.Args[1] // first argument is the source
	if !fileExists (sourceFile) {
		showError ("Source file not found", true, -2)
	}

    // now copy all the destination locations into an array checking they exist first
	for i:= 2; i < count; i++ {

	    if folderExists (os.Args[i]) {
	  		destinationLocations = append(destinationLocations, os.Args[i])
	  	} else {
	  		tmp := fmt.Sprintf ("Folder '%s' does not exist. Skipping ...", strings.ToUpper (os.Args[i]))
	  		showError (tmp, false, 0)
	  	}

		
	}
}

func folderExists (folder string) bool {
	return false
}

func showError (errorText string, exitApp bool, exitCode int) {
	
	fmt.Printf ("*** Error : %s\n", errorText)
	if exitApp {
		os.Exit (exitCode)
	}
}

func displayBanner () {
	fmt.Printf ("DupeFile, file multicopier %s\n\n", DUPEFILE_VERSION)
}

func showSyntax () {
	fmt.Println ("DupeFile [<source files> <destination path>{<destination path>}*]\n")
}

func fileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
    	return true
  	} 
  	return false
}

func copyFile () {
	sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()
 
	// Create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
 
	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}
