package utils

import (
	"encoding/json"
	"fmt"
	//"io"
	"log"
	"token/utils/session"
)

//公共方法中的session管理器
var GlobalSessions *session.Manager

//然后在init函数中初始化
func init() {
	//使用内存保存session,默认值c存活是 3600 秒
	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go GlobalSessions.GC()
}

//主要go外部文件引用函数,函数名称头字母要大写func C
func CheckError(err error) {
	if err != nil {
		fmt.Println("err==", err)
		log.Println(err)
	}
}

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