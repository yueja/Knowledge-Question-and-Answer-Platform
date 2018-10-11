package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/DataConn"
	pb "study0/proto/AskAndAnswerQuestion"
	"study0/server/AskAndAnswerQuestion/api"
	"flag"
	)

func main() {
	db := DataConn.DB_Mysql()
	AskAndAnswerQuestion := api.Make_db(db)

	//注册
	port:= flag.String("port", ":1995", "Input your username")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, AskAndAnswerQuestion)
	s.Serve(lis)
}
