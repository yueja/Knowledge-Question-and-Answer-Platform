package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/AskAndAnswerQuestion"
	"study0/server/ask_answer/api"
)

func main() {
	db := data_conn.MakeMySqlDB()
	askAndAnswerQuestion := ask_answer_api.MakeDb(db)

	addr := flag.String("addr", ":1995", "AskAndAnswerQuestion")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, askAndAnswerQuestion)
	err=s.Serve(lis)
	if err != nil {
		log.Fatal("err: %v", err)
	}
}
