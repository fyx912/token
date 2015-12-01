package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"token/utils"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	utils.CheckError(err)
	t.Execute(w, nil)
	fmt.Println(" login.gtpl")
}
