package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello Andrey Ferubko")
	file, err := os.Open("Station.xlsx")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
