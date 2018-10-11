package main

import (
	"flag"
	"fmt"
	"net/http"
	"study0/Http_Server/Client/AskAndAnswerQuestion"
	"study0/Http_Server/Client/DeleAnsAndQue"
	"study0/Http_Server/Client/PullQuestion"
	"study0/Http_Server/Client/RegiAndLog"
)

func main() {
	mux := http.NewServeMux()
	address := flag.String("address", "1994", "Input your username")
	flag.Parse()

	RegiAndLogClientHandle := RegiAndLog.NewRegiAndLogClientHandle(RegiAndLog.HttpServer(*address))
	//注册用户
	mux.Handle("/Register", http.HandlerFunc(RegiAndLogClientHandle.RegisterServer))
	//登录用户
	mux.Handle("/Login", http.HandlerFunc(RegiAndLogClientHandle.LoginServer))

	AskAndAnswerQuestionClientHandle := AskAndAnswerQuestion.NewAskAndAnswerQuestionClientHandle(AskAndAnswerQuestion.HttpServer(*address))
	//写一个提出问题服务
	mux.Handle("/AskQuestion", http.HandlerFunc(AskAndAnswerQuestionClientHandle.AskQuestionServer))
	//浏览问题列表
	mux.Handle("/BrowseQuestion", http.HandlerFunc(AskAndAnswerQuestionClientHandle.BrowseQuestionServer))
	//回答某问题
	mux.Handle("/AnswerQuestion", http.HandlerFunc(AskAndAnswerQuestionClientHandle.AnswerQuestionServer))
	//浏览单个问题详细内容(包括问题提问者、所有回答-回答者)
	mux.Handle("/DetailedList", http.HandlerFunc(AskAndAnswerQuestionClientHandle.DetailedListServer))

	DeleAnsAndQueClientHandle := DeleAnsAndQue.NewDeleAnsAndQueClientHandle(DeleAnsAndQue.HttpServer(*address))
	//用户删除自己在某问题下的回答.
	mux.Handle("/DeleteAnswer", http.HandlerFunc(DeleAnsAndQueClientHandle.DeleteAnswerServer))
	//用户删除自己所题的某问题
	mux.Handle("/DeleteQuestion", http.HandlerFunc(DeleAnsAndQueClientHandle.DeleteQuestionServer))

	PullQuestionClientHandle := PullQuestion.NewPullQuestionClientHandle(PullQuestion.HttpServer(*address))
	//拉取用户自己提出的全部问题.
	mux.Handle("/AllMyQuestion", http.HandlerFunc(PullQuestionClientHandle.AllMyQuestionServer))
	// 拉取用户自己所回答的全部回答
	mux.Handle("/AllMyAnswerer", http.HandlerFunc(PullQuestionClientHandle.AllMyAnswererServer))
	// 拉取用户自己所回答的全部回答
	mux.Handle("/HighestRanking", http.HandlerFunc(PullQuestionClientHandle.HighestRankingServer))
	//拉取最新的10个问题（借助Redis排名）
	mux.Handle("/RedisSort", http.HandlerFunc(PullQuestionClientHandle.RedisSortServer))

	if err := http.ListenAndServe(":2018", mux); err != nil {
		fmt.Println(err)
	}
}
