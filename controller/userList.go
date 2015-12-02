package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"token/utils"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	getSessions, sessionId := utils.GetSessionAndSessionId(w, r)
	defer getSessions.SessionRelease(w) //释放session
	if utils.VilidataMethodLoggedIn(w, r, getSessions, sessionId) {
		t, err := template.ParseFiles("views/index.html")
		utils.CheckError(err)
		t.Execute(w, nil)
		fmt.Println(" index.html")
	}
}
