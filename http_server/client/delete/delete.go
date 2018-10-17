package delete

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/delete"
	"study0/structure_type"
	"time"
)

type DeleteClientHandle struct {
	c delete.GreeterClient
}

func NewDeleteClientHandle(c delete.GreeterClient) *DeleteClientHandle {
	return &DeleteClientHandle{c}
}
func HttpServer(address string) delete.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := delete.NewGreeterClient(conn)
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

//写一个客户端
func (de *DeleteClientHandle) DeleteAnswerServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]
	answerer := r.Form["answerer"][0]

	re, err := de.c.DeleteAnswer(context.Background(), &delete.DeleteAnswerRequest{Question: question, Answerer: answerer})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}

func (de *DeleteClientHandle) DeleteQuestionServer(w http.ResponseWriter, r *http.Request) {
	WithTime()
	r.ParseForm()
	question := r.Form["question"][0]
	questioner := r.Form["questioner"][0]

	re, err := de.c.DeleteQuestion(context.Background(), &delete.DeleteQuestionRequest{Question: question, Questioner: questioner})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}
