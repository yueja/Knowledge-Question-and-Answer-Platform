package pull_api

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"log"
	"study0/data_conn"
	pb3 "study0/proto/pull"
)

type server struct {
	db *gorm.DB
	re *redis.Client
}

func MakeObject(db *gorm.DB, re *redis.Client) *server {
	DB := &server{db, re}
	return DB
}

//拉取用户自己提出的全部问题.
func (s *server) AllMyQuestion(ctx context.Context, in *pb3.AllMyQuestionRequest) (*pb3.AllMyQuestionReply, error) {
	reply := &pb3.AllMyQuestionReply{}
	tmp := &pb3.QuestionList{}

	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Where("Questioner=?", in.Questioner).Select("Id," +
		"Question,Questioner,AnswerCount").Rows()
	if err != nil {
		log.Printf("err: %s", err)
	}
	for rows.Next() {
		err = rows.Scan(&tmp.Id, &tmp.Question, &tmp.Questioner, &tmp.AnswerCount)
		if err != nil {
			log.Printf("err: %s", err)
		}
		reply.Question = append(reply.Question, tmp)
	}
	reply.Result = true
	reply.Message = "拉取问题成功"
	return reply, nil
}

// 拉取用户自己所回答的全部回答
func (s *server) AllMyAnswer(ctx context.Context, in *pb3.AllMyAnswerRequest) (*pb3.AllMyAnswerReply, error) {
	reply := &pb3.AllMyAnswerReply{}
	tmp := &pb3.AnswerList{}

	rows, err := s.db.Model(&data_conn.AnswerInfo{}).Where("Answerer=?", in.Answerer).Select("Id,Answer,Answerer").Rows()

	if err != nil {
		log.Printf("err: %s", err)
	}
	for rows.Next() {
		err = rows.Scan(&tmp.Num, &tmp.Answer, &tmp.Answerer)
		if err != nil {
			log.Printf("err: %s", err)
		}
		reply.Answer = append(reply.Answer, tmp)
	}
	reply.Result = true
	reply.Message = "拉取回答成功"
	return reply, nil
}

//按照回答数排名最高的10个问题。
func (s *server) HighestRanking(ctx context.Context, in *pb3.HighestRankingRequest) (*pb3.HighestRankingReply, error) {
	reply := &pb3.HighestRankingReply{}
	tmp := &pb3.QuestionList{}

	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Order("AnswerCount desc ").Limit(10).Select("Id," +
		"Question,Questioner,AnswerCount").Rows()
	if err != nil {
		log.Printf("err: %s", err)
	}
	for rows.Next() {
		err = rows.Scan(&tmp.Id, &tmp.Question, &tmp.Questioner, &tmp.AnswerCount)
		if err != nil {
			log.Printf("err: %s", err)
		}
		reply.Question = append(reply.Question, tmp)
	}
	reply.Result = true
	reply.Message = "拉取排名最高10个问题成功"
	return reply, nil
}

//拉取最新的10个问题（借助Redis排名）
func (s *server) RedisSort(ctx context.Context, in *pb3.RedisSortRequest) (*pb3.RedisSortReply, error) {
	reply := &pb3.RedisSortReply{}
	tmp := &pb3.QuestionList{}

	type RedisSortRequest struct {
		Id int `json:"id"`
	}
	type RedisSortResponse struct {
		Response []RedisSortRequest
	}
	var s_1 RedisSortResponse

	rows, err := s.db.Model(&data_conn.QuestionInfo{}).Select("Id").Rows()
	if err != nil {
		log.Printf("err: %s", err)
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("err: %s", err)
		}
		s_1.Response = append(s_1.Response, RedisSortRequest{Id: id})
	}
	maxId := s_1.Response[len(s_1.Response)-1].Id
	for i := 0; i < 10; {
		rows, err = s.db.Model(&data_conn.QuestionInfo{}).Where("Id=?", maxId).Select("Id,Question,Questioner,AnswerCount").Rows()
		if err != nil {
			log.Printf("err: %s", err)
		}
		for rows.Next() {
			err = rows.Scan(&tmp.Id, &tmp.Question, &tmp.Questioner, &tmp.AnswerCount)
			if err != nil {
				log.Printf("err: %s", err)
			}
			if tmp.Id == "" {
				maxId = maxId - 1
				continue
			}
			reply.Question = append(reply.Question, tmp)
			i++
		}
	}
	reply.Result = true
	reply.Message = "拉取最新10个问题成功"
	return reply, nil
}
