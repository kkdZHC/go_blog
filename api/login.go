package api

import (
	"net/http"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//接受用户名密码, 返回json数据
	params := common.GetPequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(username, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
