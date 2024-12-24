package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := flag.String("file", "problems.csv", "set custom filename")
	timeDuration := flag.Int("time", 10, "set custom timer")

	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file read err: %v\n", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	problems := parseInput(records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "csv err: %v\n", err)
		return
	}
	score := 0
	input := ""
	duration := time.Duration(time.Duration(*timeDuration) * time.Second)

	Timer := time.NewTimer(duration)
	done := make(chan struct{})

	go func() {
		<-Timer.C
		done <- struct{}{}

	}()

	go func() {

		for _, problem := range problems {

			fmt.Printf("%v: ", problem.question)
			fmt.Scanln(&input)
			if input == problem.answer {
				score += 1
			}
			input = ""
		}
		done <- struct{}{}

		Timer.Stop()

	}()

	<-done
	fmt.Printf("%v/%v", score, len(records))

}

func parseInput(records [][]string) []Problem {
	var problems []Problem
	for _, record := range records {
		p := Problem{
			question: record[0],
			answer:   record[1],
		}
		problems = append(problems, p)
	}
	return problems
}

type Problem struct {
	question string
	answer   string
}
