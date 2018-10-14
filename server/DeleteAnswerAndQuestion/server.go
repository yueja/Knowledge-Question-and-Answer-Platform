package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/DeleteAnswerAndQuestion"
	"study0/server/DeleteAnswerAndQuestion/api"
)

func main() {
	db := data_conn.DB_Mysql()
	deleteAnswerAndQuestion := api.MakeDb(db)

	//注册
	port := flag.String("port", ":1996", "DeleAnsAndQue")
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, deleteAnswerAndQuestion)
	s.Serve(lis)
}
