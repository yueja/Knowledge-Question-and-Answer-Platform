package api

import (
		"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"study0/DataConn"
	pb2 "study0/proto/DeleteAnswerAndQuestion"
)

type server struct {
	db *gorm.DB
}

func Make_db(db *gorm.DB) *server {
	DB := &server{db}
	return DB
}

//用户删除自己在某问题下的回答.
func (s *server) DeleteAnswer(ctx context.Context, in *pb2.DeleteAnswerRequest) (*pb2.DeleteAnswerReply, error) {
	//查询问题所在编号
	var Id ,Answer_count int
	rows, err := s.db.Model(&DataConn.Questioninfo{}).Where("Question=?", in.Question).Select("Id,Answer_count").Rows()
	if err != nil {
		return &pb2.DeleteAnswerReply{Result: "出错"}, nil
	}
	for rows.Next() {
		err = rows.Scan(&Id, &Answer_count)
		if err != nil {
			return &pb2.DeleteAnswerReply{Result: "出错"}, nil
		}
	}
	//根据编号查询答案并删除
	err = s.db.Model(&DataConn.Answerinfo{}).Where("Answerer=? and Id=?", in.Answerer, Id).Delete(&DataConn.Answerinfo{}).Error
	if err != nil {
		return &pb2.DeleteAnswerReply{Result: "出错"}, nil
	}
	//更新答案个数
	Answer_count = Answer_count - 1
	err = s.db.Model(&DataConn.Questioninfo{}).Where(" Id=?", Id).Updates(DataConn.Questioninfo{Answer_count: Answer_count}).Error
	if err != nil {
		return &pb2.DeleteAnswerReply{Result: "更新失败"}, nil
	}
	return &pb2.DeleteAnswerReply{Result: "删除回答成功"}, nil
}

//用户删除自己所题的某问题

func (s *server) DeleteQuestion(ctx context.Context, in *pb2.DeleteQuestionRequest) (*pb2.DeleteQuestionReply, error) {
	var Id int
	rows, err := s.db.Model(&DataConn.Questioninfo{}).Where("Question=? and Questioner=?", in.Question, in.Questioner).Select("Id").Rows()
	if err != nil {
		return &pb2.DeleteQuestionReply{Result: "出错"}, nil
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&Id)
		if err != nil {
			return &pb2.DeleteQuestionReply{Result: "出错"}, nil
		}
	}
	//删除相应问题的答案
	err = s.db.Model(&DataConn.Answerinfo{}).Where("Id=?", Id).Delete(&DataConn.Answerinfo{}).Error
	if err != nil {
		return &pb2.DeleteQuestionReply{Result: "删除问题失败"}, nil
	}
	//删除相应问题
	err = s.db.Model(&DataConn.Questioninfo{}).Where("Id=?", Id).Delete(&DataConn.Questioninfo{}).Error
	if err != nil {
		return &pb2.DeleteQuestionReply{Result: "删除问题失败"}, nil
	}
	return &pb2.DeleteQuestionReply{Result: "删除问题成功"}, nil
}

