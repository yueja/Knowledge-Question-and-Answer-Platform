package pull_question

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/pull"
	"study0/structure_type"
	"time"
)

type PullClientHandle struct {
	c pull.GreeterClient
}

func NewPullClientHandle(c pull.GreeterClient) *PullClientHandle {
	return &PullClientHandle{c}
}
func HttpServer(address string) pull.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := pull.NewGreeterClient(conn)
	return c
}

func (p *PullClientHandle) AllMyQuestionServer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.ParseForm()
	questioner := r.Form["questioner"][0]

	re, err := p.c.AllMyQuestion(ctx, &pull.AllMyQuestionRequest{Questioner: questioner})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.AllMyQuestionReply{Result: re.Result, Message: re.Message}
	for i := 0; i < len(re.Question); i++ {
		s.Question = append(s.Question, structure_type.Questionlist{Id: re.Question[i].Id, Question: re.Question[i].Question,
			Questioner: re.Question[i].Questioner, AnswerCount: re.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}

func (p *PullClientHandle) AllMyAnswererServer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.ParseForm()
	answerer := r.Form["answerer"][0]

	re, err := p.c.AllMyAnswer(ctx, &pull.AllMyAnswerRequest{Answerer: answerer})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.AllMyAnswerReply{Result: re.Result, Message: re.Message}
	for i := 0; i < len(re.Answer); i++ {
		s.Answer = append(s.Answer, structure_type.Answerlist{Num: re.Answer[i].Num, Answer: re.Answer[i].Answer, Answerer: re.Answer[i].Answerer})
	}
	render.JSON(w, r, s)
}

func (p *PullClientHandle) HighestRankingServer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r_1, err := p.c.HighestRanking(ctx, &pull.HighestRankingRequest{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.AllMyQuestionReply{Result: r_1.Result, Message: r_1.Message}
	for i := 0; i < len(r_1.Question); i++ {
		s.Question = append(s.Question, structure_type.Questionlist{Id: r_1.Question[i].Id, Question: r_1.Question[i].Question,
			Questioner: r_1.Question[i].Questioner, AnswerCount: r_1.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}

func (p *PullClientHandle) RedisSortServer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	re, err := p.c.RedisSort(ctx, &pull.RedisSortRequest{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}

	s := structure_type.AllMyQuestionReply{Result: re.Result, Message: re.Message}
	for i := 0; i < len(re.Question); i++ {
		s.Question = append(s.Question, structure_type.Questionlist{Id: re.Question[i].Id, Question: re.Question[i].Question,
			Questioner: re.Question[i].Questioner, AnswerCount: re.Question[i].AnswerCount})
	}
	render.JSON(w, r, s)
}
