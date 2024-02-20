package main

import (
	"log"
	"net/http"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/router"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"describe"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	router.Router()

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
