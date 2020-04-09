package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileName := flag.String("filename", "problems.csv", "The relative path to the csv containing quiz q&a")
	limit := flag.Int("limit", 30, "The time limit in seconds")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading CSV file", err)
	}

	c := make(chan bool)
	problems := parseLines(lines)
	var correct int
	go quiz(problems, &correct, c)
	go startTimer(*limit, c)
	_ = <-c
	fmt.Printf("Given %d correct answers out of %d\n", correct, len(problems))
}

func quiz(problems []problem, correct *int, c chan bool) {
	userScanner := bufio.NewScanner(os.Stdin)

	for _, p := range problems {
		fmt.Println(p.q)
		userScanner.Scan()
		if userScanner.Text() == p.a {
			*correct++
		}
	}

	c <- true
}

func startTimer(limit int, c chan bool) {
	time.Sleep(time.Duration(limit) * time.Second)
	fmt.Println("Run out of time.")
	c <- true
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return problems
}
