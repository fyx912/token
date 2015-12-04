package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"token/controller"
	"token/utils"
)

/**
 * 404
 * [NotFoundHandler description]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "login", http.StatusFound) //302
	}
	t, err := template.ParseFiles("views/404.html")
	utils.CheckError(err)
	t.Execute(w, nil)
	data := utils.OutputJsonData(404, "请求的方法不存在!")
	fmt.Fprintln(w, data)
}

//token认证服务
func main() {
	fmt.Println("服务开始启动.....")
	fmt.Println("......启动时间:", time.Now())
	//访问静态资源
	go http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//访问静态资源
	go http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./views/"))))
	go http.HandleFunc("/userList", controller.UserList)
	go http.HandleFunc("/getToken", controller.GetToken)
	go http.HandleFunc("/token", controller.ValidateToken)
	go http.HandleFunc("/destroyToken", controller.DestroyToken)
	go http.HandleFunc("/login", controller.Login)
	go http.HandleFunc("/", NotFoundHandler)

	services := &http.Server{
		Addr: ":8899",
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(services.ListenAndServe())
	fmt.Println("......启动完成:", time.Now())
	// err := http.ListenAndServe(":8899", nil)
	// fmt.Println("......启动完成:", time.Now())
	// if err != nil {
	// 	log.Println("ListenAndServe: ", err)
	// }
}
