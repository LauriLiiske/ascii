package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//opening and reading standart.txt
	filename := "standard.txt"
	ascii_symbols := ReadFile(filename) // slice of all the contents of the banner file

	//making users input into slice splitting it on newlines
	input := strings.Split(os.Args[1], "\\n")

	//printing out the answer
	for _, line := range input { // we are going by each line
		if line != "" {
			//if line is not empty print each line of the letters in ascii-art
			PrintFile(line, ascii_symbols)
		} else {
			fmt.Print("\n")
		}
	}
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lineArray []string //string array where we store the result
	for scanner.Scan() {
		if scanner.Text() != "" { //if line is not empty then add to the result
			lineArray = append(lineArray, scanner.Text())
		}
	}
	return lineArray
}

func PrintFile(line string, ascii []string) {
	//convert line input string into a rune array
	runes := []rune(line)
	for i := 0; i < 8; i++ {
		//for each element of the rune array
		for _, element := range runes {
			// artLineNumber is the start line of the letters banner file lines, so for O letter the first line would be  ____ and...
			artLineNumber := (int(element)-32)*8 + i //...the next  / __ \  etc and the + i brings the next line till 8 lines have been printed for the element
			fmt.Print(ascii[artLineNumber])          // print the ascii-art line of the character on artLineNumber
		}
		fmt.Print("\n")
	}
}
