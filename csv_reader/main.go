package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	file, err := os.Open("items.csv")

	if (err != nil) {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if (err != nil) {
		log.Fatal(err)
	}

	for _, line := range data {
		fmt.Println(line)
	}

	new_file, err := os.Open("items.csv")
	if (err != nil) {
		log.Fatal(err)
	}

	dataframe := dataframe.ReadCSV(new_file)
	fmt.Println("Table --->", dataframe)
}