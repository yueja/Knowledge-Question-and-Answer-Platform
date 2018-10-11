package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/DataConn"
	pb "study0/proto/DeleteAnswerAndQuestion"
	"study0/server/DeleteAnswerAndQuestion/api"
	"flag"
)

func main() {
	db := DataConn.DB_Mysql()
	DeleteAnswerAndQuestion := api.Make_db(db)

	//注册
	port:= flag.String("port", ":1996", "Input your username")
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, DeleteAnswerAndQuestion)
	s.Serve(lis)
}
