package regi_and_log

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/RegiAndLog"
	"study0/structure_type"
)

type RegiAndLogClientHandle struct {
	c regi_and_log.GreeterClient
}

func NewRegiAndLogClientHandle(c regi_and_log.GreeterClient) *RegiAndLogClientHandle {
	return &RegiAndLogClientHandle{c}
}

func HttpServer(address string) regi_and_log.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := regi_and_log.NewGreeterClient(conn)
	return c
}

func (Re *RegiAndLogClientHandle) RegisterServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	num := r.Form["num"][0]
	password := r.Form["password"][0]

	r_1, err := Re.c.RegisteredUser(context.Background(), &regi_and_log.RegisteredUserRequest{Num: num, Password: password})
	if err != nil {
		//log.Fatal("could not greet: %v", err)
		s := structure_type.Things{Result: "could not greet", IsSuccess: false}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1.Result)
	s := structure_type.Things{Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}

func (Re *RegiAndLogClientHandle) LoginServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	num := r.Form["num"][0]
	password := r.Form["password"][0]

	r_1, err := Re.c.LoginUser(context.Background(), &regi_and_log.LoginUserRequest{Num: num, Password: password})
	if err != nil {
		// log.Fatal("could not greet: %v", err)
		s := structure_type.Things{Result: "could not greet", IsSuccess: false}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1.Result)
	s := structure_type.Things{Result: r_1.Result, IsSuccess: r_1.Message}
	render.JSON(w, r, s)
}
