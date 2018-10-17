package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/user"
	"study0/server/user/api"
)

func main() {
	db := data_conn.MakeMySqlDB()
	re := data_conn.MakeRedis()
	user := user_api.MakeDb(db, re)

	//注册
	addr := flag.String("addr", ":1994", "RegiAndLog")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, user)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("err: %v", err)
	}
}
