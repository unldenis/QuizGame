package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := "30"

	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	timer, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal("Time argument must be an integer")
	}

	t1 := time.NewTimer(time.Duration(timer) * time.Second)

	records := readCsvFile("./problems.csv")

	var input string
	correctAnswers := 0
	for _, value := range records {
		question := value[0]
		fmt.Print("what is " + question + " sir? ")

		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal("Error handling answer")
		}

		select {
		case <-t1.C:
			log.Fatal("Time expired")
		default:
		}

		answer := value[1]
		if answer == input {
			correctAnswers++
		}

	}
	fmt.Println("Total number of questions correct:", correctAnswers)
	fmt.Println("Total number of questions:", len(records))
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}
