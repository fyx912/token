package utils

import (
	"encoding/json"
	"fmt"
	//"io"
	"log"
	"net/http"
	"token/utils/session"
)

//公共方法中的session管理器
var GlobalSessions *session.Manager

//保存用户的session,如username=session
var USER_SESSION = make(map[string]session.SessionStore)

//保存所有用户的sessionID,如username=sessionId
var SESSIONID_USER = make(map[string]string)

//然后在init函数中初始化
func init() {
	//使用内存保存session,默认值c存活是 3600 秒
	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go GlobalSessions.GC()
}

/**
 * [name description] 获取session和cookie中的sessionId
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func GetSessionAndSessionId(w http.ResponseWriter, r *http.Request) (session.SessionStore, string) {
	getSessions, err := GlobalSessions.SessionStart(w, r) //公共方法中的session管理器,判断session是否访问
	CheckError(err)
	sessionId, err := r.Cookie("gosessionid") //获取浏览器的sessionID
	CheckError(err)
	return getSessions, sessionId.Value
}

/**
 *  判断方法是否通过登录,false 则返回登录界面,true通过
 */
func VilidataMethodLoggedIn(w http.ResponseWriter, r *http.Request, getSessions session.SessionStore, sessionId string) bool {
	flag := false
	oldSessionId := ""
	username := getSessions.Get("username")
	if str, ok := username.(string); ok {
		fmt.Println(" userList  is session :", str)
		session, _ := USER_SESSION[str]
		oldSessionId = session.Get("sessionId").(string)
		fmt.Println(" userList  is sessionId  :", session.Get("sessionId"))
	}
	if username == nil { //判断session中是否存在用户
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Println(" login.gtpl")
	} else {
		if sessionId == oldSessionId {
			flag = true
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
			fmt.Println(" login.gtpl")
		}
	}
	return flag
}

//主要go外部文件引用函数,函数名称头字母要大写func C
func CheckError(err error) {
	if err != nil {
		fmt.Println("err==", err)
		log.Println(err)
	}
}

/**
 *输出json提示
 */
func OutputJsonData(code int, message string) string {
	mapData := make(map[string]interface{})
	mapData["code"] = code
	mapData["message"] = message
	jsonData, err := json.Marshal(mapData)
	CheckError(err)
	return string(jsonData)
	// str := "{\"code\":code,\"message\":\"+message+\"}"
	// return str
}
