package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/service"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/:id
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不能识别此路径"))
		return
	}
	//处理分页
	err = r.ParseForm()
	if err != nil {
		log.Println("获取表单数据出错", err)
		categoryTemplate.WriteError(w, err)
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr) // 页数
	//每页显示的数量
	pageSize := 10
	//调用service返回数据
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
