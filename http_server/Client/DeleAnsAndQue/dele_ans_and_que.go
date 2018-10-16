package dele_ans_and_que

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/DeleteAnswerAndQuestion"
	"study0/structure_type"
	"time"
)

type DeleAnsAndQueClientHandle struct {
	c delete_answer_and_question.GreeterClient
}

func NewDeleAnsAndQueClientHandle(c delete_answer_and_question.GreeterClient) *DeleAnsAndQueClientHandle {
	return &DeleAnsAndQueClientHandle{c}
}
func HttpServer(address string) delete_answer_and_question.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := delete_answer_and_question.NewGreeterClient(conn)
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

//写一个客户端
func (de *DeleAnsAndQueClientHandle) DeleteAnswerServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]
	answerer := r.Form["answerer"][0]

	r_1, err := de.c.DeleteAnswer(context.Background(), &delete_answer_and_question.DeleteAnswerRequest{Question: question, Answerer: answerer})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}

func (de *DeleAnsAndQueClientHandle) DeleteQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]
	questioner := r.Form["questioner"][0]

	r_1, err := de.c.DeleteQuestion(context.Background(), &delete_answer_and_question.DeleteQuestionRequest{Question: question, Questioner: questioner})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}
