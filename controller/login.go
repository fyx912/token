package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"token/conf"
	"token/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	w.Header().Add("Content-Type", "text/html;charset=UTF-8")
	method := r.Method
	if method == "GET" {
		t, _ := template.ParseFiles("views/login.html")
		t.Execute(w, nil)
	} else {
		contentType := r.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			jsonData, err := ioutil.ReadAll(r.Body)
			utils.CheckError(err)
			defer r.Body.Close()
			if string(jsonData) != "" {
				user := make(map[string]string)
				err = json.Unmarshal(jsonData, &user)
				utils.CheckError(err)
				fmt.Println("body=", r.Body)
				fmt.Println("body..jsonData=", string(jsonData))
				requseUsername := user["username"]
				requestPassword := user["password"]
				fmt.Println("body=", requseUsername, user["password"])
				username, password := conf.GetUsers()
				if requseUsername == username {
					if requestPassword == password {
						// t, _ := template.ParseFiles("views/login.gtpl")
						// t.Execute(w, nil)
						data := utils.OutputJsonData(0, "登陆成功!")
						fmt.Fprintln(w, data)
						fmt.Println("登陆成功!")
					} else {
						data := utils.OutputJsonData(1, "输入的密码有误!")
						fmt.Fprintln(w, data)
					}
				} else {
					data := utils.OutputJsonData(1, "账号不存在!")
					fmt.Fprintln(w, data)
				}
			} else {
				data := utils.OutputJsonData(1, "请求的内容不能为空!")
				fmt.Fprintln(w, data)
			}
		} else {
			data := utils.OutputJsonData(1, "请求的格式必须为json!")
			fmt.Fprintln(w, data)
		}
	}
}
