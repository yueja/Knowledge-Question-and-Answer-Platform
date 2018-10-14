package api

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"regexp"
	"study0/data_conn"
	pb "study0/proto/RegiAndLog"
)

type server struct {
	db *gorm.DB
	re *redis.Client
}

func MakeDb(db *gorm.DB, re *redis.Client) *server {
	DB := &server{db, re}
	return DB
}

//注册
func (s *server) RegisteredUser(ctx context.Context, in *pb.RegisteredUserRequest) (*pb.RegisteredUserReply, error) {
	var num string
	if m, _ := regexp.MatchString("^[0-9]+$", in.Num); !m {
		return &pb.RegisteredUserReply{Result: "请输入数字号码", Message: false}, nil
	}

	if len(in.Num) != 8 {
		return &pb.RegisteredUserReply{Result: "请输入八位有效数字号码", Message: false}, nil
	}
	rows, err := s.db.Model(&data_conn.User{}).Where(" Num=?", in.Num).Select("Num").Rows()
	if err != nil {
		return &pb.RegisteredUserReply{Result: "出错", Message: false}, nil
	}

	for rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			return &pb.RegisteredUserReply{Result: "出错", Message: false}, nil
		}
	}
	if num != "" {
		return &pb.RegisteredUserReply{Result: "该账户已注册", Message: false}, nil
	}

	if m, _ := regexp.MatchString("^[a-zA-Z]+$", in.Password); !m {
		return &pb.RegisteredUserReply{Result: "请输入英文密码", Message: false}, nil
	}

	err = s.db.Create(&data_conn.User{Num: in.Num, Password: in.Password}).Error
	if err != nil {
		return &pb.RegisteredUserReply{Result: "注册账号失败", Message: false}, nil
	}
	return &pb.RegisteredUserReply{Result: "注册成功", Message: true}, nil
}

//登陆
func (s *server) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserReply, error) {
	var num, password string
	if in.Num == "" {
		return &pb.LoginUserReply{Result: "账号为空", Message: false}, nil
	}
	//取缓存账户
	result, err := s.re.HMGet("NumPassword", "Num", "Password").Result()
	if err != nil {
		return &pb.LoginUserReply{Result: "出错", Message: false}, nil
	}

	if in.Num == result[0] && in.Password == result[1] {
		return &pb.LoginUserReply{Result: "登录成功", Message: false}, nil
	}

	rows, err := s.db.Model(&data_conn.User{}).Where("Num=?", in.Num).Select("Num,Password").Rows()
	if err != nil {
		return &pb.LoginUserReply{Result: "出错", Message: false}, nil
	}

	for rows.Next() {
		err = rows.Scan(&num, &password)
		if err != nil {
			return &pb.LoginUserReply{Result: "出错", Message: false}, nil
		}
	}
	if num == "" {
		return &pb.LoginUserReply{Result: "用户名不存在", Message: false}, nil
	}
	if password != in.Password {
		return &pb.LoginUserReply{Result: "密码错误", Message: false}, nil
	}
	//缓存账号
	s.re.HMSet("NumPassword", map[string]interface{}{"Num": in.Num, "Password": in.Password})
	return &pb.LoginUserReply{Result: "登录成功", Message: true}, nil
}
