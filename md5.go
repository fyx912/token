package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
	"token/conf"
)

func main() {
	hash := md5.New()
	b := []byte("test")
	hash.Write(b)
	fmt.Printf("%x %x\n", hash.Sum(nil), md5.Sum(b))
	hash.Write(nil)
	fmt.Printf("%x %x\n", hash.Sum(b), hash.Sum(nil))

	fmt.Println(uuid())
	fmt.Println(name())

	username, passsword := conf.GetUsers()
	fmt.Println("username=", username)
	fmt.Println("passsword=", passsword)
}

func uuid() string {
	cruntime := time.Now().Unix()
	data := []byte(strconv.FormatInt(cruntime, 10))
	//%x    表示为十六进制，使用a-f
	return fmt.Sprintf("%x", md5.Sum(data))
}

func name() string {
	name := "dings"
	ding := name
	return ding
}
