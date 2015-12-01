package conf

import (
	"fmt"
	"token/utils/config"
)

/**
 * [GetUsers description]
 *获取配置文件
 * @param {[type]} ) (username, password string [description]
 */
func GetUsers() (username, password string) {
	iniconf, err := config.NewConfig("ini", "./conf/ini.conf")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	username = iniconf.String("user::username")
	password = iniconf.String("user::password")
	return username, password
}

// func main() {
// 	username, passsword := GetUsers()
// 	fmt.Println("username=", username)
// 	fmt.Println("passsword=", passsword)
// }
