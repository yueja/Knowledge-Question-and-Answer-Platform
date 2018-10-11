package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/DataConn"
	pb "study0/proto/RegiAndLog"
	"study0/server/RegiAndLog/api"
	"flag"
)

func main() {
	db := DataConn.DB_Mysql()
	re := DataConn.RED()
	RegiAndLog := api.Make_db(db, re)

	//注册
	port:= flag.String("port", ":1994", "Input your username")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,RegiAndLog)
	s.Serve(lis)
}
