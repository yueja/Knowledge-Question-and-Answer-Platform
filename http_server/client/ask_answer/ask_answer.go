package ask_answer

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/ask_answer"
	"study0/structure_type"
	"time"
)

type AskAnswerClientHandle struct {
	c ask_answer.GreeterClient
}

func NewAskAnswerClientHandle(c ask_answer.GreeterClient) *AskAnswerClientHandle {
	return &AskAnswerClientHandle{c}
}

func HttpServer(address string) ask_answer.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := ask_answer.NewGreeterClient(conn)
	return c
}

func WithTime() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		log.Printf("Time out")
	default:
		break
	}
}

//提出问题
func (As *AskAnswerClientHandle) AskQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	num := r.Form["num"][0]
	question := r.Form["question"][0]

	re, err := As.c.AskQuestion(context.Background(), &ask_answer.AskQuestionRequest{Num: num, Question: question})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}

//浏览问题列表
func (As *AskAnswerClientHandle) BrowseQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	re, err := As.c.BrowseQuestion(context.Background(), &ask_answer.BrowseQuestionRequest{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.QuestionInfo{Question: re.Question, Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}

//回答某问题
func (As *AskAnswerClientHandle) AnswerQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]
	answer := r.Form["answer"][0]
	answerer := r.Form["answerer"][0]

	a := &ask_answer.AnswerQuestionRequest{Question: question, Answer: answer, Answerer: answerer}
	re, err := As.c.AnswerQuestion(context.Background(), a)
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}

//浏览问题详情
func (As *AskAnswerClientHandle) DetailedListServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]

	re, err := As.c.DetailedList(context.Background(), &ask_answer.DetailedListRequest{Question: question})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.DetailedListReply{Result: re.Result, Message: re.Message}
	for i := 0; i < len(re.Detailedlist); i++ {
		s.DetailedList = append(s.DetailedList, structure_type.DetailedList{Question: re.Detailedlist[i].Question,
			Questioner: re.Detailedlist[i].Questioner, Answer: re.Detailedlist[i].Answer, Answerer: re.Detailedlist[i].Answerer})
	}
	render.JSON(w, r, s)
}
