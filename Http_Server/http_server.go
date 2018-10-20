package main

import (
	"flag"
	"fmt"
	"net/http"
	"study0/http_server/client/ask_answer"
	"study0/http_server/client/delete"
	"study0/http_server/client/pull"
	"study0/http_server/client/user"
)

func main(){
	mux := http.NewServeMux()
	address := flag.String("address", ":2018", "http_server")
	userAddress := flag.String("userAddress", "localhost:1994", "RegiAndLog")
	askAnswerAddress := flag.String("askAnswerAddress", "localhost:1995", "AskAndAnswerQuestion")
	deleteAddress := flag.String("deleteAddress", "localhost:1996", "DeleAnsAndQue")
	pullAddress := flag.String("pullAddress", "localhost:1997", "PullQuestion")
	flag.Parse()

	userClient := user.MakeStub(*userAddress)
	askAnswerClient := ask_answer.MakeStub(*askAnswerAddress)
	deleteClient := delete.MakeStub(*deleteAddress)
	pullClient := pull_question.MakeStub(*pullAddress)

	userClientHandle := user.NewUserClientHandle(userClient)
	//注册用户
	mux.Handle("/register", http.HandlerFunc(userClientHandle.RegisterServer))
	//登录用户
	mux.Handle("/login", http.HandlerFunc(userClientHandle.LoginServer))

	askAnswerClientHandle := ask_answer.NewAskAnswerClientHandle(askAnswerClient)
	//写一个提出问题服务
	mux.Handle("/ask_question", http.HandlerFunc(askAnswerClientHandle.AskQuestionServer))
	//浏览问题列表
	mux.Handle("/browse_question", http.HandlerFunc(askAnswerClientHandle.BrowseQuestionServer))
	//回答某问题
	mux.Handle("/answer_question", http.HandlerFunc(askAnswerClientHandle.AnswerQuestionServer))
	//浏览单个问题详细内容(包括问题提问者、所有回答-回答者)
	mux.Handle("/detailed_list", http.HandlerFunc(askAnswerClientHandle.DetailedListServer))

	deleteClientHandle := delete.NewDeleteClientHandle(deleteClient)
	//用户删除自己在某问题下的回答.
	mux.Handle("/delete_answer", http.HandlerFunc(deleteClientHandle.DeleteAnswerServer))
	//用户删除自己所题的某问题
	mux.Handle("/delete_question", http.HandlerFunc(deleteClientHandle.DeleteQuestionServer))

	pullClientHandle := pull_question.NewPullClientHandle(pullClient)
	//拉取用户自己提出的全部问题.
	mux.Handle("/all_my_question", http.HandlerFunc(pullClientHandle.AllMyQuestionServer))
	// 拉取用户自己所回答的全部回答.
	mux.Handle("/all_my_answerer", http.HandlerFunc(pullClientHandle.AllMyAnswererServer))
	// 拉取用户自己所回答的全部回答.
	mux.Handle("/highest_ranking", http.HandlerFunc(pullClientHandle.HighestRankingServer))
	//拉取最新的10个问题（借助Redis排名）.
	mux.Handle("/redis_sort", http.HandlerFunc(pullClientHandle.RedisSortServer))

	if err := http.ListenAndServe(*address, mux); err != nil {
		fmt.Println(err)
	}
}
