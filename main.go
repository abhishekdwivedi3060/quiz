package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "takes a csv file with question and answer")
	timeLimit := flag.Int("limit", 30, "time limit for the test")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Could not open csv file")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	qna := parser(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	count := 0
	answerChan := make(chan string)
	for i, list := range qna {
		fmt.Printf("Question %d is %s ", i+1, list.q)

		go func() {
			var ans string
			fmt.Scanf("%s \n", &ans)
			answerChan <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTime out...You scored %d out of %d ", count, len(qna))
			return
		case ans := <-answerChan:
			if list.a == ans {
				fmt.Println("correct")
				count++
			}

		}
	}
	fmt.Printf("You scored %d out of %d \n", count, len(qna))
}

type csvStruct struct {
	q string
	a string
}

func parser(lines [][]string) []csvStruct {
	ret := make([]csvStruct, len(lines))
	for i, line := range lines {
		ret[i] = csvStruct{
			q: line[0],
			a: line[1],
		}

	}
	return ret
}
