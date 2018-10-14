package structure_type

type Things struct {
	Result string
	IsSuccess bool
}

type QuestionInfo struct {
	Question []string
	Result string
	IsSuccess bool
}

type DetailedListReply struct {
	DetailedList []DetailedList
	Result       string
	IsSuccess bool
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
	IsSuccess bool
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
	IsSuccess bool
}

type Answerlist struct {
	Num      string
	Answer   string
	Answerer string
}
