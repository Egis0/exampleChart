package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Correction struct {
	Date  string `json:"date"`
	Title string `json:"title"`
	Value int    `json:"value"`
}

type Corrections []Correction

func returnAllCorrections(w http.ResponseWriter, r *http.Request) {
	corrections := Corrections{
		Correction{Date: "2019-03-08", Title: "Corrected", Value: 50},
		Correction{Date: "2019-03-08", Title: "Uncorrected", Value: 50},
		Correction{Date: "2019-03-09", Title: "Corrected", Value: 37},
		Correction{Date: "2019-03-09", Title: "Uncorrected", Value: 63},
		Correction{Date: "2019-03-10", Title: "Corrected", Value: 31},
		Correction{Date: "2019-03-10", Title: "Uncorrected", Value: 69},
		Correction{Date: "2019-03-11", Title: "Corrected", Value: 68},
		Correction{Date: "2019-03-11", Title: "Uncorrected", Value: 32},
		Correction{Date: "2019-03-12", Title: "Corrected", Value: 75},
		Correction{Date: "2019-03-12", Title: "Uncorrected", Value: 25},
		Correction{Date: "2019-03-13", Title: "Corrected", Value: 77},
		Correction{Date: "2019-03-13", Title: "Uncorrected", Value: 23},
		Correction{Date: "2019-03-14", Title: "Corrected", Value: 58},
		Correction{Date: "2019-03-14", Title: "Uncorrected", Value: 42},
	}

	json.NewEncoder(w).Encode(corrections)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, re *http.Request) {
		http.ServeFile(w, re, "../frontend/build/index.html")
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/build/static/"))))
	http.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("../frontend/build/"))))
	http.HandleFunc("/corrections", returnAllCorrections)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Couldnt start server: %q\n", err)
	}
}
