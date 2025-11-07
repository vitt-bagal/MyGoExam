package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Store data to URLFile- short-url=long-url
func main() {
	//var shorturl string
	urlFile, err := os.OpenFile("URLFile.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	//urlFile, err := os.Open("./URLFile.txt")
	// check for error
	if err != nil {
		log.Fatalln("Could not open URLFile.txt")
	}

	//lurl := "https://www.guru3d.com/news-story/microsoft-latest-update-for-windows-11-includes-support-for-android-applications.html"
	scanner := bufio.NewScanner(urlFile)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
	urlFile.Close()
}
