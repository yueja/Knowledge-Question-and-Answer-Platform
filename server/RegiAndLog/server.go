package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"study0/data_conn"
	pb "study0/proto/RegiAndLog"
	"study0/server/RegiAndLog/api"
)

func main() {
	db := data_conn.DB_Mysql()
	re := data_conn.RED()
	regiAndLog := api.MakeDb(db, re)

	//注册
	port := flag.String("port", ":1994", "RegiAndLog")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, regiAndLog)
	s.Serve(lis)
}
