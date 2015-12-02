package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"token/utils"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	getSessions, err := utils.GlobalSessions.SessionStart(w, r) //公共方法中的session管理器,判断session是否访问
	utils.CheckError(err)
	defer getSessions.SessionRelease(w) //释放session
	username := getSessions.Get("username")
	if str, ok := username.(string); ok {
		fmt.Println(" userList  is session :", str)
		session, _ := USER_SESSION[str]
		fmt.Println(" userList  is sessionId  :", session.Get("sessionId"))
	}
	if username == nil { //判断session中是否存在用户
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Println(" login.gtpl")
		return
	} else {
		t, err := template.ParseFiles("views/index.html")
		utils.CheckError(err)
		t.Execute(w, nil)
		fmt.Println(" index.html")
	}

}
