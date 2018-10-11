package AskAndAnswerQuestion

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/StructureType"
	"study0/proto/AskAndAnswerQuestion"
)

type AskAndAnswerQuestionClientHandle struct {
	c AskAndAnswerQuestion.GreeterClient
}

func NewAskAndAnswerQuestionClientHandle(c AskAndAnswerQuestion.GreeterClient) *AskAndAnswerQuestionClientHandle {
	return &AskAndAnswerQuestionClientHandle{c}
}

func HttpServer(address string) AskAndAnswerQuestion.GreeterClient {
	conn, err := grpc.Dial("localhost"+address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := AskAndAnswerQuestion.NewGreeterClient(conn)
	return c
}

//提出问题
func (As *AskAndAnswerQuestionClientHandle) AskQuestionServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	num := r.Form["Num"][0]
	question := r.Form["Question"][0]

	r_1, err := As.c.AskQuestion(context.Background(), &AskAndAnswerQuestion.AskQuestionRequest{Num: num, Question: question})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.Things{"could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1.Result)
	s := StructureType.Things{r_1.Result}
	render.JSON(w, r, s)
}

//浏览问题列表
func (As *AskAndAnswerQuestionClientHandle) BrowseQuestionServer(w http.ResponseWriter, r *http.Request) {
	r_1, err := As.c.BrowseQuestion(context.Background(), &AskAndAnswerQuestion.BrowseQuestionRequest{})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.Things{"could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1.Question)
	s := StructureType.Questioninfo{r_1.Question}
	render.JSON(w, r, s)
}

//回答某问题
func (As *AskAndAnswerQuestionClientHandle) AnswerQuestionServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	question := r.Form["Question"][0]
	answer := r.Form["Answer"][0]
	answerer := r.Form["Answerer"][0]

	a := &AskAndAnswerQuestion.AnswerQuestionRequest{Question: question, Answer: answer, Answerer: answerer}
	r_1, err := As.c.AnswerQuestion(context.Background(), a)
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.Things{"could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1.Result)
	s := StructureType.Things{r_1.Result}
	render.JSON(w, r, s)
}

func (As *AskAndAnswerQuestionClientHandle) DetailedListServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	question := r.Form["Question"][0]

	r_1, err := As.c.DetailedList(context.Background(), &AskAndAnswerQuestion.DetailedListRequest{Question: question})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.DetailedlistReply{Result:"could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1)
	s := StructureType.DetailedlistReply{Result: r_1.Result}
	for i := 0; i < len(r_1.Detailedlist); i++ {
		s.Detailedlist = append(s.Detailedlist, StructureType.DetailedList{Question: r_1.Detailedlist[i].Question, Questioner: r_1.Detailedlist[i].Questioner,
			Answer: r_1.Detailedlist[i].Answer, Answerer: r_1.Detailedlist[i].Answerer})
	}
	render.JSON(w, r, s)
}
