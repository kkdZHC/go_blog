package service

import (
	"errors"

	"github.com/kkdZHC/go_blog/dao"
	"github.com/kkdZHC/go_blog/models"
	"github.com/kkdZHC/go_blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "aijdl") //密码使用md5加密
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//使用jwt生成 token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	userInfo := models.UserInfo{
		Uid:      user.Uid,
		UserName: user.UserName,
		Avatar:   user.Avatar,
	}

	lr := &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}

	return lr, nil
}
