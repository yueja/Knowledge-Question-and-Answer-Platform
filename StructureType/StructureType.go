package StructureType

type Things struct {
	Result string
}

type Questioninfo struct {
	Question []string
}

type DetailedlistReply struct {
	Detailedlist []DetailedList
	Result       string
}
type DetailedList struct {
	Question   string
	Questioner string
	Answer     string
	Answerer   string
}

type AllMyQuestionReply struct {
	Question []Questionlist
	Result   string
}
type Questionlist struct {
	Id          string
	Question    string
	Questioner  string
	AnswerCount string
}

type AllMyAnswerReply struct {
	Answer []Answerlist
	Result string
}

type Answerlist struct {
	Num      string
	Answer   string
	Answerer string
}
