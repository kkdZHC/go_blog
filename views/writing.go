package views

import (
	"log"
	"net/http"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr, err := service.Writing()
	if err != nil {
		log.Println(err)
		return
	}
	writing.WriteData(w, wr)

}
