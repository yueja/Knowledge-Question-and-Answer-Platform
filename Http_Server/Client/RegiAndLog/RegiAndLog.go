package RegiAndLog

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/StructureType"
	"study0/proto/RegiAndLog"
)

type RegiAndLogClientHandle struct {
	c RegiAndLog.GreeterClient
}

func NewRegiAndLogClientHandle(c RegiAndLog.GreeterClient) *RegiAndLogClientHandle {
	return &RegiAndLogClientHandle{c}
}

func HttpServer(address string) RegiAndLog.GreeterClient {
	conn, err := grpc.Dial("localhost"+address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := RegiAndLog.NewGreeterClient(conn)
	return c
}

//写一个客户端
func (Re *RegiAndLogClientHandle) RegisterServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	num := r.Form["Num"][0]
	password := r.Form["Password"][0]

	r_1, err := Re.c.RegisteredUser(context.Background(), &RegiAndLog.RegisteredUserRequest{Num: num, Password: password})
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

func (Re *RegiAndLogClientHandle) LoginServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	num := r.Form["Num"][0]
	password := r.Form["Password"][0]

	r_1, err := Re.c.LoginUser(context.Background(), &RegiAndLog.LoginUserRequest{Num: num, Password: password})
	if err != nil {
		// log.Fatal("could not greet: %v", err)
		s := StructureType.Things{"could not greet"}
		render.JSON(w, r, s)
		return
	}
	//log.Printf("Greeting: %s", r_1.Result)
	s := StructureType.Things{r_1.Result}
	render.JSON(w, r, s)
}
