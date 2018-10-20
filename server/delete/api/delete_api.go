package delete_api

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"log"
	"study0/data_conn"
	pb2 "study0/proto/delete"
)

type server struct {
	db *gorm.DB
}

func MakeObject(db *gorm.DB) *server {
	DB := &server{db}
	return DB
}

//用户删除自己在某问题下的回答.
func (s *server) DeleteAnswer(ctx context.Context, in *pb2.DeleteAnswerRequest) (*pb2.DeleteAnswerReply, error) {
	//查询问题所在编号
	var Id, Answer_count int
	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Where("Question=?", in.Question).Select("Id,Answer_count").Rows()
	if err != nil {
		log.Printf("err: %s", err)
	}
	for rows.Next() {
		err = rows.Scan(&Id, &Answer_count)
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
	//根据编号查询答案并删除
	err = s.db.Model(&data_conn.AnswerInfo{}).Where("Answerer=? and Id=?", in.Answerer, Id).Delete(&data_conn.AnswerInfo{}).Error
	if err != nil {
		log.Printf("err: %s", err)
	}
	//更新答案个数
	Answer_count = Answer_count - 1
	err = s.db.Model(&data_conn.QuestionInfo{}).Where(" Id=?", Id).Updates(data_conn.QuestionInfo{AnswerCount: Answer_count}).Error
	if err != nil {
		log.Printf("err: %s", err)
	}
	return &pb2.DeleteAnswerReply{Result: true, Message: "删除回答成功"}, nil
}

//用户删除自己所题的某问题

func (s *server) DeleteQuestion(ctx context.Context, in *pb2.DeleteQuestionRequest) (*pb2.DeleteQuestionReply, error) {
	var Id int
	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Where("Question=? and Questioner=?", in.Question, in.Questioner).Select("Id").Rows()
	if err != nil {
		log.Printf("err: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&Id)
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
	//删除相应问题的答案
	err = s.db.Model(&data_conn.AnswerInfo{}).Where("Id=?", Id).Delete(&data_conn.AnswerInfo{}).Error
	if err != nil {
		log.Printf("err: %s", err)
	}
	//删除相应问题
	err = s.db.Model(&data_conn.QuestionInfo{}).Where("Id=?", Id).Delete(&data_conn.QuestionInfo{}).Error
	if err != nil {
		log.Printf("err: %s", err)
	}
	return &pb2.DeleteQuestionReply{Result: true, Message: "删除问题成功"}, nil
}
