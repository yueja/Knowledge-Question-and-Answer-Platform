package ask_and_answer_question

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/AskAndAnswerQuestion"
	"study0/structure_type"
	"time"
)

type AskAndAnswerQuestionClientHandle struct {
	c ask_and_answer_question.GreeterClient
}

func NewAskAndAnswerQuestionClientHandle(c ask_and_answer_question.GreeterClient) *AskAndAnswerQuestionClientHandle {
	return &AskAndAnswerQuestionClientHandle{c}
}

func HttpServer(address string) ask_and_answer_question.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := ask_and_answer_question.NewGreeterClient(conn)
	return c
}

func WithTime()  {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	select{
	case <-ctx.Done():
		log.Printf("Time out")
	default: break
	}
}
//提出问题
func (As *AskAndAnswerQuestionClientHandle) AskQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	num := r.Form["num"][0]
	question := r.Form["question"][0]

	r_1, err := As.c.AskQuestion(context.Background(), &ask_and_answer_question.AskQuestionRequest{Num: num, Question: question})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}

//浏览问题列表
func (As *AskAndAnswerQuestionClientHandle) BrowseQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r_1, err := As.c.BrowseQuestion(context.Background(), &ask_and_answer_question.BrowseQuestionRequest{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.QuestionInfo{Question: r_1.Question, Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}

//回答某问题
func (As *AskAndAnswerQuestionClientHandle) AnswerQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]
	answer := r.Form["answer"][0]
	answerer := r.Form["answerer"][0]

	a := &ask_and_answer_question.AnswerQuestionRequest{Question: question, Answer: answer, Answerer: answerer}
	r_1, err := As.c.AnswerQuestion(context.Background(), a)
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}

//浏览问题详情
func (As *AskAndAnswerQuestionClientHandle) DetailedListServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]

	r_1, err := As.c.DetailedList(context.Background(), &ask_and_answer_question.DetailedListRequest{Question: question})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.DetailedListReply{Result: r_1.Result, IsSuccess: r_1.Message}
	for i := 0; i < len(r_1.Detailedlist); i++ {
		s.DetailedList = append(s.DetailedList, structure_type.DetailedList{Question: r_1.Detailedlist[i].Question,
			Questioner: r_1.Detailedlist[i].Questioner, Answer: r_1.Detailedlist[i].Answer, Answerer: r_1.Detailedlist[i].Answerer})
	}
	render.JSON(w, r, s)
}
