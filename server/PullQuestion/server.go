package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/DataConn"
	pb "study0/proto/PullQuestion"
	"study0/server/PullQuestion/api"
	"flag"
)

func main() {
	db := DataConn.DB_Mysql()
	re := DataConn.RED()
	PullQuestion := api.Make_db(db, re)

	//注册
	port:= flag.String("port", ":1997", "Input your username")
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, PullQuestion)
	s.Serve(lis)
}
