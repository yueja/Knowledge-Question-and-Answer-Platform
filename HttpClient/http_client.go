package main

import (
	"flag"
	"fmt"
	"net/http"
	"study0/Client/AskAndAnswerQuestion"
	"study0/Client/DeleAnsAndQue"
	"study0/Client/PullQuestion"
	"study0/Client/RegiAndLog"
)

func main() {
	mux := http.NewServeMux()
	address := flag.String("address", "1994", "Input your username")
	flag.Parse()

	RegiAndLogClientHandle := RegiAndLog.NewRegiAndLogClientHandle(RegiAndLog.Client(*address))
	//注册用户
	mux.Handle("/Register", http.HandlerFunc(RegiAndLogClientHandle.RegisterClient))
	//登录用户
	mux.Handle("/Login", http.HandlerFunc(RegiAndLogClientHandle.LoginClient))

	AskAndAnswerQuestionClientHandle := AskAndAnswerQuestion.NewAskAndAnswerQuestionClientHandle(AskAndAnswerQuestion.Client(*address))
	//写一个提出问题服务
	mux.Handle("/AskQuestion", http.HandlerFunc(AskAndAnswerQuestionClientHandle.AskQuestionClient))
	//浏览问题列表
	mux.Handle("/BrowseQuestion", http.HandlerFunc(AskAndAnswerQuestionClientHandle.BrowseQuestionClient))
	//回答某问题
	mux.Handle("/AnswerQuestion", http.HandlerFunc(AskAndAnswerQuestionClientHandle.AnswerQuestionClient))
	//浏览单个问题详细内容(包括问题提问者、所有回答-回答者)
	mux.Handle("/DetailedList", http.HandlerFunc(AskAndAnswerQuestionClientHandle.DetailedListClient))

	DeleAnsAndQueClientHandle := DeleAnsAndQue.NewDeleAnsAndQueClientHandle(DeleAnsAndQue.Client(*address))
	//用户删除自己在某问题下的回答.
	mux.Handle("/DeleteAnswer", http.HandlerFunc(DeleAnsAndQueClientHandle.DeleteAnswerClient))
	//用户删除自己所题的某问题
	mux.Handle("/DeleteQuestion", http.HandlerFunc(DeleAnsAndQueClientHandle.DeleteQuestionClient))

	PullQuestionClientHandle := PullQuestion.NewPullQuestionClientHandle(PullQuestion.Client(*address))
	//拉取用户自己提出的全部问题.
	mux.Handle("/AllMyQuestion", http.HandlerFunc(PullQuestionClientHandle.AllMyQuestionClient))
	// 拉取用户自己所回答的全部回答
	mux.Handle("/AllMyAnswerer", http.HandlerFunc(PullQuestionClientHandle.AllMyAnswererClient))
	// 拉取用户自己所回答的全部回答
	mux.Handle("/HighestRanking", http.HandlerFunc(PullQuestionClientHandle.HighestRankingClient))
	//拉取最新的10个问题（借助Redis排名）
	mux.Handle("/RedisSort", http.HandlerFunc(PullQuestionClientHandle.RedisSortClient))

	if err := http.ListenAndServe(":2018", mux); err != nil {
		fmt.Println(err)
	}
}
