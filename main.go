package main

import (
"fmt"
"os"
"io"
"strings"
)

const DUPEFILE_VERSION = "v1.0"

const (
 	KErrorNone = 0
 	KErrorProblemOpeningFile   = -1
 	KErrorProblemCopyingData   = -2
 	KErrorProblemClosingFile   = -3
 	KErrorUnableToWriteToFile  = -4
 	KErrorUnableToWriteData    = -5
)

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


	if len(destinationLocations) > 0 {// at least one existing folder
		for i:= 0; i < len(destinationLocations); i++{
			if _, status := copyFile (sourceFile, destinationLocations[i]); !status {
				fmt.Printf ("Problem copying file '%s' to '%s'\n", strings.ToUpper(sourceFile), strings.ToUpper(destinationLocations[i]))
			}
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
	fmt.Printf ("DupeFile, File Multicopier %s\n\n", DUPEFILE_VERSION)
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

func getFileSize (filepath string) (int64, error) {
	f, err := os.Open(filepath)
    if err != nil {
        return 0, err
    }
    defer f.Close()
    fi, err := f.Stat()
    if err != nil {
        return 0, err
    }

    return fi.Size(), nil
}


// need to return rather than exit
func copyFile (filename string, folder string) (int, bool) {
	fmt.Printf ("Copying %s to '%s'\n", strings.ToUpper(filename), strings.ToUpper(folder))

	sourceFile, err := os.Open(filename)
	if err != nil {
		// problem opening file
		return KErrorProblemOpeningFile, false
	}
	defer sourceFile.Close()
 
	// Create new file
	newFilePath := folder
	if folder[len(folder)-1] != '\\' {
		newFilePath += "\\" // append a trailing slash
	}
	newFilePath += filename

	newFile, err := os.Create(newFilePath)
	if err != nil {
		return KErrorUnableToWriteToFile, false
	}
	defer newFile.Close()
 
	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		return KErrorUnableToWriteData, false
	}
	
	var byteCount int64
	byteCount, _ = getFileSize (filename)

	if byteCount == bytesCopied {
		return KErrorNone, true
	} else {
		return KErrorProblemCopyingData, false
	}
}
