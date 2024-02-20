package router

import (
	"net/http"

	"github.com/kkdZHC/go_blog/api"
	"github.com/kkdZHC/go_blog/views"
)

func Router() {
	//1.返回页面  views 2.返回数据（json）api 3.返回静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/:id 文章分类
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/api/v1/login", api.API.Login)
	//http://localhost:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)

	http.HandleFunc("/writing", views.HTML.Writing)

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)

	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
