package PullQuestion

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/StructureType"
	"study0/proto/PullQuestion"
)

type PullQuestionClientHandle struct {
	c PullQuestion.GreeterClient
}

func NewPullQuestionClientHandle(c PullQuestion.GreeterClient) *PullQuestionClientHandle {
	return &PullQuestionClientHandle{c}
}
func Client(address string) PullQuestion.GreeterClient {
	conn, err := grpc.Dial("localhost"+address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := PullQuestion.NewGreeterClient(conn)
	return c
}

//写一个客户端
func (pu *PullQuestionClientHandle) AllMyQuestionClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	questioner := r.Form["Questioner"][0]

	r_1, err := pu.c.AllMyQuestion(context.Background(), &PullQuestion.AllMyQuestionRequest{Questioner: questioner})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r_1)

	s := StructureType.AllMyQuestionReply{Result: r_1.Result}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, StructureType.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question, Questioner: r_1.Question[i].Questioner,
			AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}

func (pu *PullQuestionClientHandle) AllMyAnswererClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	answerer := r.Form["Answerer"][0]

	r_1, err := pu.c.AllMyAnswer(context.Background(), &PullQuestion.AllMyAnswerRequest{Answerer: answerer})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.AllMyAnswerReply{Result: "could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1)
	s := StructureType.AllMyAnswerReply{Result: r_1.Result}
	for i := 0; i < len(r_1.Answer); i++ {
		s.Answer = append(s.Answer, StructureType.Answerlist{Num: r_1.Answer[i].Num, Answer: r_1.Answer[i].Answer, Answerer: r_1.Answer[i].Answerer})
	}
	render.JSON(w, r, s)
}

func (pu *PullQuestionClientHandle) HighestRankingClient(w http.ResponseWriter, r *http.Request) {
	r_1, err := pu.c.HighestRanking(context.Background(), &PullQuestion.HighestRankingRequest{})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.AllMyQuestionReply{Result: "could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1)
	s := StructureType.AllMyQuestionReply{Result: r_1.Result}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, StructureType.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question, Questioner: r_1.Question[i].Questioner,
			AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}

func (pu *PullQuestionClientHandle) RedisSortClient(w http.ResponseWriter, r *http.Request) {

	r_1, err := pu.c.RedisSort(context.Background(), &PullQuestion.RedisSortRequest{})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := StructureType.AllMyQuestionReply{Result: "could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1)

	s := StructureType.AllMyQuestionReply{Result: r_1.Result}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, StructureType.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question, Questioner: r_1.Question[i].Questioner,
			AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}
