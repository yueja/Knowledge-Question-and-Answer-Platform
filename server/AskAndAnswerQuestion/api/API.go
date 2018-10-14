package api

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"study0/data_conn"
	pb1 "study0/proto/AskAndAnswerQuestion"
)

type server struct {
	db *gorm.DB
}

func MakeDb(db *gorm.DB) *server {
	DB := &server{db}
	return DB
}

//提出问题
func (s *server) AskQuestion(ctx context.Context, in *pb1.AskQuestionRequest) (*pb1.AskQuestionReply, error) {
	if in.Question == "" {
		return &pb1.AskQuestionReply{Result: "内容为空", Message: false}, nil
	}
	err := s.db.Create(&data_conn.QuestionInfo{Questioner: in.Num, Question: in.Question, AnswerCount: 0}).Error
	if err != nil {
		return &pb1.AskQuestionReply{Result: "提出问题失败", Message: false}, nil
	}
	return &pb1.AskQuestionReply{Result: "提出问题成功", Message: true}, nil
}

//浏览问题列表
func (s *server) BrowseQuestion(ctx context.Context, in *pb1.BrowseQuestionRequest) (*pb1.BrowseQuestionReply, error) {
	var question string
	var Questioninfo []string

	type BrowseQuestionRequest struct {
		Question string
	}
	type BrowseQuestionResponse struct {
		Response []BrowseQuestionRequest
	}
	var s_1 BrowseQuestionResponse

	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Select("Question").Rows()
	if err != nil {
		return &pb1.BrowseQuestionReply{Result: "查询错误，未查出相应结果", Message: false}, nil
	}
	for rows.Next() {
		err = rows.Scan(&question)
		if err != nil {
			return &pb1.BrowseQuestionReply{Result: "出错", Message: false}, nil
		}
		s_1.Response = append(s_1.Response, BrowseQuestionRequest{Question: question})
	}
	for i := 0; i < len(s_1.Response); i++ {
		Questioninfo = append(Questioninfo, s_1.Response[i].Question)
	}
	return &pb1.BrowseQuestionReply{Question: Questioninfo, Result: "浏览问题列表成功", Message: true}, nil
}

//回答某问题
func (s *server) AnswerQuestion(ctx context.Context, in *pb1.AnswerQuestionRequest) (*pb1.AnswerQuestionReply, error) {
	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Where("Question=?", in.Question).Select("Id,Answer_count").Rows()
	if err != nil {
		return &pb1.AnswerQuestionReply{Result: "提出问题失败", Message: false}, nil
	}

	var Id, Answer_count int
	for rows.Next() {
		err = rows.Scan(&Id, &Answer_count)
		if err != nil {
			return &pb1.AnswerQuestionReply{Result: "出错", Message: false}, nil
		}
	}
	//判断同一用户是否已回答同一问题
	rows, err = s.db.Model(&data_conn.AnswerInfo{}).Where(" Answerer=? and Id=?", in.Answerer, Id).Select("Id").Rows()
	if err != nil {
		return &pb1.AnswerQuestionReply{Result: "出错", Message: false}, nil
	}
	var NUM int
	for rows.Next() {
		err = rows.Scan(&NUM)
		if err != nil {
			return &pb1.AnswerQuestionReply{Result: "出错", Message: false}, nil
		}
	}
	if NUM != 0 {
		return &pb1.AnswerQuestionReply{Result: "该用户已经回答了该问题", Message: false}, nil
	}
	//根据编号增加答案列表信息
	err = s.db.Create(&data_conn.AnswerInfo{Id: Id, Answer: in.Answer, Answerer: in.Answerer}).Error
	if err != nil {
		return &pb1.AnswerQuestionReply{Result: "回答失败", Message: false}, nil
	}
	//答案个数计数
	Answer_count = Answer_count + 1
	err = s.db.Model(&data_conn.QuestionInfo{}).Where("Question=?", in.Question).Updates(data_conn.QuestionInfo{AnswerCount: Answer_count}).Error
	if err != nil {
		return &pb1.AnswerQuestionReply{Result: "更新失败", Message: false}, nil
	}
	return &pb1.AnswerQuestionReply{Result: "回答成功", Message: true}, nil
}

//浏览单个问题详细内容(包括问题提问者、所有回答-回答者)
func (s *server) DetailedList(ctx context.Context, in *pb1.DetailedListRequest) (*pb1.DetailedListReply, error) {
	reply := &pb1.DetailedListReply{}
	tmp := &pb1.DetailedList{}

	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Where("Question=?", in.Question).Select("Id,Questioner").Rows()
	if err != nil {
		return &pb1.DetailedListReply{Result: "查询出错", Message: false}, nil
	}
	for rows.Next() {
		var id int
		var questioner string
		err = rows.Scan(&id, &questioner)
		if err != nil {
			return &pb1.DetailedListReply{Result: "查询出错", Message: false}, nil
		}
		rows, err = s.db.Model(&data_conn.AnswerInfo{}).Where("Id=?", id).Select("Answer,Answerer").Rows()
		for rows.Next() {
			err = rows.Scan(&tmp.Answer, &tmp.Answerer)
			if err != nil {
				return &pb1.DetailedListReply{Result: "查询出错", Message: false}, nil
			}
			tmp.Question = in.Question
			tmp.Questioner = questioner
			reply.Detailedlist = append(reply.Detailedlist, tmp)
			reply.Result = "浏览成功"
			reply.Message = true
		}
	}
	return reply, nil
}
