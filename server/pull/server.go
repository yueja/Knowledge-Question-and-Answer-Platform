package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/pull"
	"study0/server/pull/api"
)

func main() {
	db := data_conn.MakeMySqlDB()
	re := data_conn.MakeRedis()
	pull := pull_api.MakeDb(db, re)

	//注册
	addr := flag.String("addr", ":1997", "")
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, pull)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("err: %v", err)
	}
}
