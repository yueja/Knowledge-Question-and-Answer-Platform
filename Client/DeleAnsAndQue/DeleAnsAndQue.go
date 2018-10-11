package DeleAnsAndQue

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/StructureType"
	"study0/proto/DeleteAnswerAndQuestion"
)

type DeleAnsAndQueClientHandle struct {
	c DeleteAnswerAndQuestion.GreeterClient
}

func NewDeleAnsAndQueClientHandle(c DeleteAnswerAndQuestion.GreeterClient) *DeleAnsAndQueClientHandle {
	return &DeleAnsAndQueClientHandle{c}
}
func Client(address string) DeleteAnswerAndQuestion.GreeterClient {
	conn, err := grpc.Dial("localhost"+address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := DeleteAnswerAndQuestion.NewGreeterClient(conn)
	return c
}

//写一个客户端
func (de *DeleAnsAndQueClientHandle) DeleteAnswerClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	question := r.Form["Question"][0]
	answerer := r.Form["Answerer"][0]

	r_1, err := de.c.DeleteAnswer(context.Background(), &DeleteAnswerAndQuestion.DeleteAnswerRequest{Question: question, Answerer: answerer})
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

func (de *DeleAnsAndQueClientHandle) DeleteQuestionClient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	question := r.Form["Question"][0]
	questioner := r.Form["Questioner"][0]

	r_1, err := de.c.DeleteQuestion(context.Background(), &DeleteAnswerAndQuestion.DeleteQuestionRequest{Question: question, Questioner: questioner})
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
