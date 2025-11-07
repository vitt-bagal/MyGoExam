package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	age  int    `json:"Age"`
	name string `json:"Name"`
}

var p []Person

func main() {
	//count := 0
	p1 := Person{10, "abc"}
	p2 := Person{20, "mno"}
	p3 := Person{30, "pqr"}
	p = append(p, p1, p2, p3)
	//fmt.Println("All person data with json marshaling", p)
	//b, err := json.Marshal(p)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//mt.Println("All person data with json marshaling", string(b))
	r := mux.NewRouter()
	r.HandleFunc("/age", getPerson)
	http.ListenAndServe(":9090", r)

}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	age := vars["age"]
	a, err := strconv.Atoi(age)
	if err != nil {
		log.Fatalln("Age should be integer value and conversion failed")
	}
	for i := 0; i < len(p); i++ {
		if p[i].age == a {
			fmt.Fprintf(w, "Person is: %v\n", p[i])
		}
	}
}
