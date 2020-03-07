package question

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Question is an individual q&a combo
type Question struct {
	question string
	answer   string
	response string
	correct  bool
	asked    bool
	ChScore  chan int
}

// NewQuestion constructs a new Question struct
func NewQuestion(question, answer string) *Question {
	q := new(Question)
	q.question = strings.TrimSpace(question)
	q.answer = strings.TrimSpace(answer)
	q.ChScore = make(chan int)
	return q
}

// Ask handles asking (and scoring the answer of) each question
// It should be called as a goroutine, and return its score on the channel
func (q *Question) Ask() {
	score := 0

	fmt.Printf("%s = ? ", q.question)
	q.asked = true
	c := make(chan string)

	go q.UserInput(c)
	q.response = <-c

	if strings.ToLower(q.response) == strings.ToLower(q.answer) {
		score = 1
		q.correct = true
	}

	q.ChScore <- score
}

// UserInput reads the user's response
func (q *Question) UserInput(c chan string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1) // for all platforms
	text = strings.Replace(text, "\r", "", -1) // for Windows users
	c <- strings.TrimSpace(text)
}

// ShowCorrect displays the correct answer for any question
// the user got wrong
func (q *Question) ShowCorrect() {
	if q.asked && !q.correct {
		fmt.Printf(" '%s' is '%s'; you said '%s'\n", q.question, q.answer, q.response)
	}
}
