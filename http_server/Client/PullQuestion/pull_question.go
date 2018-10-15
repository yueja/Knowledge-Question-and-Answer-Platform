package pull_question

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/PullQuestion"
	"study0/structure_type"
)

type PullQuestionClientHandle struct {
	c pull_question.GreeterClient
}

func NewPullQuestionClientHandle(c pull_question.GreeterClient) *PullQuestionClientHandle {
	return &PullQuestionClientHandle{c}
}
func HttpServer(address string) pull_question.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := pull_question.NewGreeterClient(conn)
	return c
}

func (pu *PullQuestionClientHandle) AllMyQuestionServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	questioner := r.Form["questioner"][0]

	r_1, err := pu.c.AllMyQuestion(context.Background(), &pull_question.AllMyQuestionRequest{Questioner: questioner})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.AllMyQuestionReply{Result: r_1.Result, IsSuccess: r_1.Message}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, structure_type.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question,
			Questioner: r_1.Question[i].Questioner, AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}

func (pu *PullQuestionClientHandle) AllMyAnswererServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	answerer := r.Form["answerer"][0]

	r_1, err := pu.c.AllMyAnswer(context.Background(), &pull_question.AllMyAnswerRequest{Answerer: answerer})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.AllMyAnswerReply{Result: r_1.Result, IsSuccess: r_1.Message}
	for i := 0; i < len(r_1.Answer); i++ {
		s.Answer = append(s.Answer, structure_type.Answerlist{Num: r_1.Answer[i].Num, Answer: r_1.Answer[i].Answer, Answerer: r_1.Answer[i].Answerer})
	}
	render.JSON(w, r, s)
}

func (pu *PullQuestionClientHandle) HighestRankingServer(w http.ResponseWriter, r *http.Request) {
	r_1, err := pu.c.HighestRanking(context.Background(), &pull_question.HighestRankingRequest{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.AllMyQuestionReply{Result: r_1.Result, IsSuccess: r_1.Message}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, structure_type.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question,
			Questioner: r_1.Question[i].Questioner, AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}

func (pu *PullQuestionClientHandle) RedisSortServer(w http.ResponseWriter, r *http.Request) {

	r_1, err := pu.c.RedisSort(context.Background(), &pull_question.RedisSortRequest{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}

	s := structure_type.AllMyQuestionReply{Result: r_1.Result, IsSuccess: r_1.Message}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, structure_type.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question,
			Questioner: r_1.Question[i].Questioner, AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}
