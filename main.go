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

	  	// run through those that exist copying the file to each folder
	  	for _,fldr := range destinationLocations {
	  		copyFile (sourceFile, fldr)
	  	}

		
	}
}

func folderExists (folder string) bool {
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		return true
	}
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

// need to return rather than exit
func copyFile (filepath string, folder string) bool {
	sourceFile, err := os.Open(filepath)
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

	return true
}
