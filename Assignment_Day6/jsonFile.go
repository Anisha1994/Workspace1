package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Employee struct {
	Name        string `json:"name"`
	Designation string `json:"designation"`
}

func (t Employee) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}


func getEmpDetails() []Employee {
	employee := make([]Employee, 3)
	raw, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &employee)
	return employee
}

func getCalculation() []Employee {
	calcul := make([]Employee, 4)
	raw, err := ioutil.ReadFile("Calculation.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &calcul)
	return calcul
}

func getPlainText() {
	file, err := os.Open("text.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}


func csvFile() {
	csv_file, _ := os.Open("csvFile.csv")
	r := csv.NewReader(csv_file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func main() {
	employee := getEmpDetails()
	//fmt.Println(employee)
	for _, te := range employee {
		fmt.Println(te.toString())
	}

	calcul := getCalculation()
	for _, te := range calcul {
		fmt.Println(te.toString())
	}
	getPlainText()
	getCalculation()
	csvFile()

}