package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Unable to Open CSV file:", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file %v as a CSV: %v", filePath, err)
	}
	return records
}

func main() {
	records := readCsvFile("./csvFile")
	log.Println(records)
	// creates a new instance of a mux router
	//myRouter := mux.NewRouter().StrictSlash(true)

	// calls to handlers functions
	//myRouter.HandleFunc("/", ServiceHandler).Methods("POST")
	http.HandleFunc("/", ServiceHandler)
	http.ListenAndServe(":8090", nil)
}

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	records := readCsvFile("./csvFile")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Index	Value \n")
	for _, val := range records {
		n, err := fmt.Fprintf(w, "%v	%v\n", val[0], val[1])
		if err != nil {
			log.Fatalln("Unable to write response: ", err)
		}
		log.Printf("%v byte written as a response on browser", n)
	}
}
