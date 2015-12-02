package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	//"reflect"
	"strings"
	"token/conf"
	"token/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	getSessions, sessionId := utils.GetSessionAndSessionId(w, r)
	defer getSessions.SessionRelease(w) //释放session

	fmt.Println(" session ID=", sessionId)
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

				//去除string的前后空白,通过map获取json数据
				requseUsername := strings.TrimSpace(user["username"])
				requestPassword := strings.TrimSpace(user["password"])

				//处理用户重复登录
				userHandle(r, requseUsername)
				// getSessions = USER_SESSION[requseUsername]
				//保存username到sessionId
				//SESSIONID_USER = new map[string]session.SessionStore
				utils.SESSIONID_USER[sessionId] = requseUsername
				//保存session到uesrname
				getSessions.Set("sessionId", sessionId)
				utils.USER_SESSION[requseUsername] = getSessions

				fmt.Println("body=", requseUsername, user["password"])

				username, password := conf.GetUsers()
				if requseUsername == username {
					if requestPassword == password {
						// t, _ := template.ParseFiles("views/login.gtpl")
						// t.Execute(w, nil)
						getSessions.Set("username", requseUsername)
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

/**
 * [userHandle description] 重复登陆处理
 * @param  {[type]} r *http.Request [description]
 * @return {[type]}   [description]
 */
func userHandle(r *http.Request, username string) {
	gosessionid, err := r.Cookie("gosessionid")
	if err != nil {
		fmt.Println(" Cookie  err=", err)
	}
	//当前的sessionId
	sessionId := gosessionid.Value
	//删除当前sessionId绑定的用户
	delete(utils.SESSIONID_USER, sessionId)
	getMapSession, ok := utils.USER_SESSION[username]
	if ok {
		fmt.Println(" 删除之前保存的sessionId: ", getMapSession.Get("sessionId"))
		str, _ := getMapSession.Get("sessionId").(string)
		fmt.Println(" 断言 之前保存的sessionId: ", str)
		//删除之前sessionId保存的uesrname
		delete(utils.SESSIONID_USER, str)
		getMapSession.Delete(username)
		fmt.Println("您得账号在另一处登录,您被迫下线!")
	}
	//删除当前用户绑定的session
	delete(utils.USER_SESSION, username)
}
