package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/delete"
	"study0/server/delete/api"
)

func main() {
	db := data_conn.MakeMySqlDB()
	deleteAQ := delete_api.MakeObject(db)

	//注册
	addr := flag.String("addr", ":1996", "DeleAnsAndQue")
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, deleteAQ)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("err: %v", err)
	}
}
