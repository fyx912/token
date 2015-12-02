package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"token/conf"
	"token/utils"
)

//生成uuid,根据当前系统时间的MD5值
func GetUUID() string {
	cruntime := time.Now().Unix()
	fmt.Println(cruntime)
	data := []byte(strconv.FormatInt(cruntime, 10))
	//%x    表示为十六进制，使用a-f
	md5Value := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	return md5Value
}

//获取token,支持GET或POST方法,如:http://192.168.40.116:9999/getToken
func GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	maps := make(map[string]interface{})
	method := r.Method
	if method == "GET" || method == "POST" {
		conf.GetUsers()
		getSessions, err := utils.GlobalSessions.SessionStart(w, r) //公共方法中的session管理器,判断session是否访问

		// fmt.Println(getSessions.SessionID(r.Cookies()))
		utils.CheckError(err)
		defer getSessions.SessionRelease(w)           //释放session
		accessToken := getSessions.Get("accessToken") //获取session中的accessToken
		maps["code"] = 0
		if accessToken != nil && len(accessToken.(string)) > 0 { //判断当前session是否存在accessToken
			maps["token"] = accessToken
		} else {
			token := GetUUID()
			getSessions.Set("accessToken", token)
			maps["token"] = token
		}
	} else {
		maps["code"] = 1
		maps["message"] = "只支持GET或POST方法!"
	}
	jsonData, _ := json.Marshal(maps)
	fmt.Fprintln(w, string(jsonData))
}

//验证token,如:http://192.168.40.116:9999/token?token=743856b55b9d51b23b085cf12f7e2583
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	maps := make(map[string]interface{})
	method := r.Method
	fmt.Println(r.Method)
	if method == "GET" || method == "POST" {
		getSessions, err := utils.GlobalSessions.SessionStart(w, r) //判断session是否访问
		utils.CheckError(err)
		defer getSessions.SessionRelease(w) //释放session
		token := r.FormValue("token")
		fmt.Println(" get  token=", token)
		accessToken := getSessions.Get("accessToken")
		flag := false
		if token == accessToken {
			flag = true
			maps["code"] = 0
		} else {
			maps["code"] = 1
		}
		maps["message"] = flag
	} else {
		maps["code"] = 1
		maps["message"] = "只支持GET或POST方法!"
	}
	jsonData, _ := json.Marshal(maps)
	fmt.Fprintln(w, string(jsonData))
	// defer w.Header().Set("Connection", "close")
}

//销毁token,如:http://192.168.40.116:9999/destroyToken?token=e016679a2b3606a2b6a26b759958e891
func DestroyToken(w http.ResponseWriter, r *http.Request) {
	var data = ""
	method := r.Method
	if method == "GET" || method == "POST" {
		getSessions, err := utils.GlobalSessions.SessionStart(w, r) //判断session是否访问
		utils.CheckError(err)
		defer getSessions.SessionRelease(w) //释放session
		token := r.FormValue("token")
		getSessions.Delete(token)
		fmt.Println(r.Method)
		data = utils.OutputJsonData(1, "token销毁成功!")
	} else {
		data = utils.OutputJsonData(1, "只支持GET或POST方法!")
	}
	fmt.Fprintln(w, data)
}
