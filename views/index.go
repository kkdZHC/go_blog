package views

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/service"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//处理分页
	err := r.ParseForm()
	if err != nil {
		log.Println("获取表单数据出错", err)
		index.WriteError(w, err)
	}
	pageStr := r.Form.Get("page")
	page := 1 // 页数
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	//获取路径参数
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	//数据库查询
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("index获取数据出错", err)
		index.WriteError(w, err)
	}
	index.WriteData(w, hr)

}
