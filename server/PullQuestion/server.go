package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/PullQuestion"
	"study0/server/PullQuestion/api"
)

func main() {
	db := data_conn.DB_Mysql()
	re := data_conn.RED()
	PullQuestion := api.MakeDb(db, re)

	//注册
	port := flag.String("port", ":1997", "")
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, PullQuestion)
	s.Serve(lis)
}
