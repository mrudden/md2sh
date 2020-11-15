package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"regexp"
)


func main() {
	fmt.Println("~~~ Here are your lists! ~~~\n")

	inputFile := os.Args[1]
		
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Variable for formatting output
	headerFound := false
	re := regexp.MustCompile(`\x60(.*)\x60`)

	// Scanner will read the file line by line
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {

			line := scanner.Text()
			//fmt.Println(line)

			//skip any lines with length of 0
			if len(line) > 0 {

				// If there's any backticks, transpose directly to output
				// Try to match
				stringMatch := re.MatchString(line)

				// Test regex match
				//fmt.Println(stringMatch)
				
				// if there's a match, print the part of the string from the backticks
				if stringMatch {
					codePrint := re.FindString(line)
					fmt.Println(codePrint[1:len(codePrint)-1])
				}

				// Remove any parenthesis and spaces to not break list
				line = strings.Replace(line, "(", "", -1)
				line = strings.Replace(line, ")", "", -1)
				

				if headerFound == false{

					// If line starts with "###", create array from following unordered list
					if strings.HasPrefix(line, "### ") {
						//fmt.Println("!!!!!H3 Detected!!!!!")

						// Replace spaces in header with underscores
						line = strings.Replace(line, " ", "_", -1)
						
						header := strings.ToUpper(line[4:])
						fmt.Println(header + "=(")
						headerFound = true
					}

				} else {
					if strings.Contains(line[0:2], "*") || strings.Contains(line[0:2], "-") || strings.Contains(line[0:2], "+") {
						//fmt.Println("!!!!!List Item Detected!!!!!")

						// Replace spaces in value with hyphens
						line = strings.Replace(line, " ", "-", -1)

						fmt.Println("  " + line[2:])
					} 
				}
			} else {
				if headerFound {
					fmt.Println(")")
					fmt.Println("# Add logic to process this array\n")
				}
				
				headerFound = false
			}
		}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}