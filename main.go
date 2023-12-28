package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		age, _ := strconv.Atoi(r.FormValue("age"))
		weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
		height, _ := strconv.Atoi(r.FormValue("height"))
		gender := r.FormValue("gender")

		bmr := mifflin_st_jeor(age, weight, height, gender)
		fmt.Fprintf(w, "Your BMR is: %d", bmr)
	})

	log.Println("App starting...")
	log.Fatal(http.ListenAndServe(":8008", nil))
}

func mifflin_st_jeor(age int, weight float64, height int, gender string) int {
	if gender == "male" {
		return int((10 * weight) + (6.25 * float64(height)) - (5 * float64(age)) + 5)
	} else {
		return int((10 * weight) + (6.25 * float64(height)) - (5 * float64(age)) - 161)
	}
}
