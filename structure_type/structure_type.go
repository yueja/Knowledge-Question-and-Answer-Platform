package structure_type

type Things struct {
	Result  bool
	Message string
}

type QuestionInfo struct {
	Question []string
	Result   bool
	Message  string
}

type DetailedListReply struct {
	DetailedList []DetailedList
	Result       bool
	Message      string
}

type DetailedList struct {
	Question   string
	Questioner string
	Answer     string
	Answerer   string
}

type AllMyQuestionReply struct {
	Question []Questionlist
	Result   bool
	Message  string
}

type Questionlist struct {
	Id          string
	Question    string
	Questioner  string
	AnswerCount string
}

type AllMyAnswerReply struct {
	Answer  []Answerlist
	Result  bool
	Message string
}

type Answerlist struct {
	Num      string
	Answer   string
	Answerer string
}
