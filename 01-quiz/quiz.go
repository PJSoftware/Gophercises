package main

import (
	"flag"

	"./quiz"
)

var q quiz.Quiz

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file containing 'question,answer' pairs")
	flag.Parse()

	q.Import(*csvFilename)
	q.Play()
	q.Score()
}
