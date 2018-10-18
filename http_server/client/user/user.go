package user

//client.go

import (
	"github.com/go-chi/render"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"study0/proto/user"
	"study0/structure_type"
	"time"
)

type UserClientHandle struct {
	c user.GreeterClient
}

func NewUserClientHandle(c user.GreeterClient) *UserClientHandle {
	return &UserClientHandle{c}
}

func HttpServer(address string) user.GreeterClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	c := user.NewGreeterClient(conn)
	return c
}

func (u *UserClientHandle) RegisterServer(w http.ResponseWriter, r *http.Request) {
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	r.ParseForm()
	num := r.Form["num"][0]
	password := r.Form["password"][0]

	re, err := u.c.RegisteredUser(ctx, &user.RegisteredUserRequest{Num: num, Password: password})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}

func (u *UserClientHandle) LoginServer(w http.ResponseWriter, r *http.Request) {
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	r.ParseForm()
	num := r.Form["num"][0]
	password := r.Form["password"][0]

	re, err := u.c.LoginUser(ctx, &user.LoginUserRequest{Num: num, Password: password})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	s := structure_type.Things{Result: re.Result, Message: re.Message}
	render.JSON(w, r, s)
}
