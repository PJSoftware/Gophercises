package main

import (
	"flag"

	"./quiz"
)

var q quiz.Quiz

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file containing 'question,answer' pairs")
	timeLimit := flag.Int("limit", 30, "maximum time (seconds) to allow for quiz [0 = no limit]")
	flag.Parse()

	q.Import(*csvFilename)
	q.Play(*timeLimit)
	q.Score()
}
