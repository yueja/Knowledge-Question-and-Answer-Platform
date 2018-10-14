package main

import (
	"flag"
	"fmt"
	"net/http"
	"study0/Http_Server/Client/AskAndAnswerQuestion"
	"study0/Http_Server/Client/DeleAnsAndQue"
	"study0/http_server/Client/PullQuestion"
	"study0/http_server/Client/RegiAndLog"
)

func main() {
	mux := http.NewServeMux()
	address := flag.String("address", ":2018", "http_server")
	address1 := flag.String("address1", "localhost:1994", "RegiAndLog")
	address2 := flag.String("address2", "localhost:1995", "AskAndAnswerQuestion")
	address3 := flag.String("address3", "localhost:1996", "DeleAnsAndQue")
	address4 := flag.String("address4", "localhost:1997", "PullQuestion")
	flag.Parse()

	client1 := regi_and_log.HttpServer(*address1)
	client2 := ask_and_answer_question.HttpServer(*address2)
	client3 := dele_ans_and_que.HttpServer(*address3)
	client4 := pull_question.HttpServer(*address4)

	regiAndLogClientHandle := regi_and_log.NewRegiAndLogClientHandle(client1)
	//注册用户
	mux.Handle("/register", http.HandlerFunc(regiAndLogClientHandle.RegisterServer))
	//登录用户
	mux.Handle("/login", http.HandlerFunc(regiAndLogClientHandle.LoginServer))

	askAndAnswerQuestionClientHandle := ask_and_answer_question.NewAskAndAnswerQuestionClientHandle(client2)
	//写一个提出问题服务
	mux.Handle("/ask_question", http.HandlerFunc(askAndAnswerQuestionClientHandle.AskQuestionServer))
	//浏览问题列表
	mux.Handle("/browse_question", http.HandlerFunc(askAndAnswerQuestionClientHandle.BrowseQuestionServer))
	//回答某问题
	mux.Handle("/answer_question", http.HandlerFunc(askAndAnswerQuestionClientHandle.AnswerQuestionServer))
	//浏览单个问题详细内容(包括问题提问者、所有回答-回答者)
	mux.Handle("/detailed_list", http.HandlerFunc(askAndAnswerQuestionClientHandle.DetailedListServer))

	deleAnsAndQueClientHandle := dele_ans_and_que.NewDeleAnsAndQueClientHandle(client3)
	//用户删除自己在某问题下的回答.
	mux.Handle("/delete_answer", http.HandlerFunc(deleAnsAndQueClientHandle.DeleteAnswerServer))
	//用户删除自己所题的某问题
	mux.Handle("/delete_question", http.HandlerFunc(deleAnsAndQueClientHandle.DeleteQuestionServer))

	pullQuestionClientHandle := pull_question.NewPullQuestionClientHandle(client4)
	//拉取用户自己提出的全部问题.
	mux.Handle("/all_my_question", http.HandlerFunc(pullQuestionClientHandle.AllMyQuestionServer))
	// 拉取用户自己所回答的全部回答
	mux.Handle("/all_my_answerer", http.HandlerFunc(pullQuestionClientHandle.AllMyAnswererServer))
	// 拉取用户自己所回答的全部回答
	mux.Handle("/highest_ranking", http.HandlerFunc(pullQuestionClientHandle.HighestRankingServer))
	//拉取最新的10个问题（借助Redis排名）
	mux.Handle("/redis_sort", http.HandlerFunc(pullQuestionClientHandle.RedisSortServer))

	if err := http.ListenAndServe(*address, mux); err != nil {
		fmt.Println(err)
	}
}
