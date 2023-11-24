package main

import "fmt"

type User struct {
	Name  string
	Email string
	Phone string
}

var addressBook = make(map[string]User)

func main() {
	var userTest User

	userTest.Email = "trhthang0401@gmail.com"
	userTest.Name = "Thang"
	userTest.Phone = "01056893798"

	addressBook[userTest.Email] = userTest

	user, exist := addressBook["trhthang0401@gmail.com"]
	if !exist {
		fmt.Println("Không tìm thấy người dùng trong danh bạ.")
		return
	}
	fmt.Println("Nhập thông tin mới: ")

	fmt.Print("Nhập tên mới: ")
	fmt.Scan(&user.Name)

	fmt.Print("Nhập sdt mới: ")

	fmt.Scan(&user.Phone)

	fmt.Println("Kết quả", userTest)
}
