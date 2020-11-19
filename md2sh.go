// md2sh.go
// Copyright 2020 Michael Rudden

package main

import (
	"fmt"
	"os"
	"flag"
	"bufio"
	"log"
	"strings"
	"regexp"
)

func main() {
	//fmt.Println("~~~ Here is your output! ~~~")
	//fmt.Println()

	inputFlag := flag.String("f", "README.md", "filepath to run md2sh against")
	silentFlag := flag.Bool("s", false, "silent mode; won't print output to terminal")
	headerForListFlag := flag.String("l", "###", "change which markdown header denotes the start of a list that should become an array")
	createOutputFileFlag := flag.Bool("o", false, "creates an output file in present directory called \"run-md2sh.sh\"")
	flag.Parse()

	// Dereference the pointers so that they can be used as normal types
	inputFile := *inputFlag
	silentMode := *silentFlag
	listHeaderPrefixInput := *headerForListFlag
	outputFile := *createOutputFileFlag

	// Testing assigning header from flag
	listHeaderPrefix := "### "
	if listHeaderPrefixInput == "###" || listHeaderPrefixInput == "h3" {
		listHeaderPrefix = "### "
	} else if listHeaderPrefixInput == "####" || listHeaderPrefixInput == "h4" {
		listHeaderPrefix = "#### "
	}

	// For debugging flags
	/*if silentMode {
		fmt.Println("silentMode = true")
	}
	if outputFile {
		fmt.Println("createOutputFile = true")
	}*/
	
	// Open input file
	inFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	// Declare output slice
	output := []string{}

	// Variable for detecting list headers
	headerFound := false

	// Regex for backticks with anything between them"``"
	re := regexp.MustCompile(`\x60(.*)\x60`)

	// Variable for detecting code blocks
	codeBlockHeaderFound := false

	// Scanner will read the file line by line
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {

		line := scanner.Text()
		//fmt.Println(line)

		// Boolean to see if line matches any rules below
		matchFound := false

		// Look for code blocks
		if !matchFound && !codeBlockHeaderFound {
			if line == "```" || line == "~~~" {
				//fmt.Println("Code block found!")
				line = "# BEGIN Code Block"
				codeBlockHeaderFound = true
				matchFound = true				
			}
		} else if !matchFound && codeBlockHeaderFound {
			// If we haven't matched this line, but we're in a code block, line should print literally. If it's the end of the code block, we need to set the variable to false to end it.
			if line == "```" || line == "~~~" {
				codeBlockHeaderFound = false
				line = "# END Code Block"
			}	
			matchFound = true
		}

		// If there's any backticks around content in a single line, transpose directly to output
		// Try to match
		stringMatch := false
		if !matchFound {
			stringMatch = re.MatchString(line)
		}
		
		// Test regex match
		//fmt.Println(stringMatch)
		
		// if there's a match, print the part of the string from the backticks
		if !matchFound && stringMatch {
			codePrint := re.FindString(line)
			codePrint = codePrint[1:len(codePrint)-1]
			line = "# " + line + "\n" + codePrint
			matchFound = true
		}
		
		if !matchFound && !headerFound {

			// If line starts with "###", create array from following unordered list
			if strings.HasPrefix(line, listHeaderPrefix) {
				//fmt.Println("!!!!!Header Detected!!!!!")

				// Remove any parenthesis to not break list
				line = strings.Replace(line, "(", "", -1)
				line = strings.Replace(line, ")", "", -1)

				// Replace spaces in header with underscores, make uppercase, and format for list
				line = strings.Replace(line, " ", "_", -1)
				line = strings.ToUpper(line[4:])
				line = line + "=("

				//fmt.Println(line)

				headerFound = true
				matchFound = true
			}

		} else if len(line) > 0 {
			if strings.Contains(line[0:2], "*") || strings.Contains(line[0:2], "-") || strings.Contains(line[0:2], "+") {
				// This if statement parses list items into the array
				//fmt.Println("!!!!!List Item Detected!!!!!")

				// Remove any parenthesis to not break list
				line = strings.Replace(line, "(", "", -1)
				line = strings.Replace(line, ")", "", -1)

				// Replace spaces in value with hyphens
				line = strings.Replace(line, " ", "-", -1)
				// Strips bulleted / unordered list character from front of line
				line = "  " + line[2:] 
				//fmt.Println(line)
				matchFound = true
			}
		} else {
			// End the array when the md list ends
			line = line + ")\n#TODO: Add logic to process this array\n"
			
			headerFound = false
			matchFound = true
		}

		// If line doesn't match any of the above criteria, print it as a comment
		if !matchFound && len(line) > 0 {
			line = "# " + line
			//fmt.Println(line)
		}
		
		// Append line to output slice
		//fmt.Println(line)
		output = append(output, line)
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	
	// Write to file from output slice if flag is used
	if outputFile {
		outFile, err := os.Create("./run-md2sh.sh")
		if err != nil {
			log.Fatal(err)
		}
		outFile.WriteString("#!/bin/sh\n\n")
		for _, outputLine := range output {
			outFile.WriteString(outputLine + "\n")
		}
		defer outFile.Close()
	}

	// Print lines out of output slice if not in silent mode
	if !silentMode {
		for _, outputLine := range output {
			fmt.Println(outputLine)
		}
	}
}