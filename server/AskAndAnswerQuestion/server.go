package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/AskAndAnswerQuestion"
	"study0/server/AskAndAnswerQuestion/api"
)

func main() {
	db := data_conn.DB_Mysql()
	askAndAnswerQuestion := api.MakeDb(db)

	port := flag.String("port", ":1995", "AskAndAnswerQuestion")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, askAndAnswerQuestion)
	s.Serve(lis)
}
