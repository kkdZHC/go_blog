package views

import (
	"net/http"

	"github.com/kkdZHC/go_blog/common"
	"github.com/kkdZHC/go_blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)

}
