package views

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/service"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//获取路径参数
	//http://localhost:8080/p/7.html
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/p/")    //7.html
	cIdStr = strings.TrimSuffix(cIdStr, ".html") //7
	pId, err := strconv.Atoi(cIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不能识别此路径"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
