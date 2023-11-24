package main

import (
	"fmt"
)

// Người dùng là một struct đại diện cho thông tin của mỗi người dùng trong danh bạ
type User struct {
	Name  string
	Email string
	Phone string
}

// Danh bạ là một map lưu trữ người dùng dựa trên email
var addressBook = make(map[string]User)

// Hàm chính để chạy ứng dụng
func main() {
	for {
		fmt.Println("1. Thêm người dùng")
		fmt.Println("2. Sửa người dùng")
		fmt.Println("3. Xóa người dùng")
		fmt.Println("4. Hiển thị danh bạ")
		fmt.Println("5. Thoát")

		var choice int
		fmt.Print("Nhập lựa chọn: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addUser()
		case 2:
			editUser()
		case 3:
			deleteUser()
		case 4:
			displayAddressBook()
		case 5:
			fmt.Println("Chương trình kết thúc.")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng chọn lại.")
		}
	}
}

// Hàm thêm người dùng vào danh bạ
func addUser() {
	var newUser User
	fmt.Print("Nhập tên: ")
	fmt.Scan(&newUser.Name)
	fmt.Print("Nhập email: ")
	fmt.Scan(&newUser.Email)
	fmt.Print("Nhập số điện thoại: ")
	fmt.Scan(&newUser.Phone)

	// Kiểm tra xem người dùng đã tồn tại hay chưa
	if _, exist := addressBook[newUser.Email]; exist {
		fmt.Println("Người dùng đã tồn tại trong danh bạ.")
		return
	}

	addressBook[newUser.Email] = newUser
	fmt.Println("Người dùng đã được thêm vào danh bạ.")
}

// Hàm sửa thông tin người dùng trong danh bạ
func editUser() {
	var email string
	fmt.Print("Nhập email của người dùng cần sửa: ")
	fmt.Scan(&email)

	// Kiểm tra xem người dùng có tồn tại hay không
	user, exist := addressBook[email]
	if !exist {
		fmt.Println("Không tìm thấy người dùng trong danh bạ.")
		return
	}

	fmt.Println("Thông tin người dùng:")
	fmt.Printf("Tên: %s\nEmail: %s\nSố điện thoại: %s\n", user.Name, user.Email, user.Phone)

	// Nhập thông tin mới
	fmt.Print("Nhập tên mới: ")
	fmt.Scan(&user.Name)
	fmt.Print("Nhập số điện thoại mới: ")
	fmt.Scan(&user.Phone)

	// Cập nhật thông tin người dùng
	// addressBook[email] = user
	fmt.Println("Thông tin người dùng đã được cập nhật.")
}

// Hàm xóa người dùng khỏi danh bạ
func deleteUser() {
	var email string
	fmt.Print("Nhập email của người dùng cần xóa: ")
	fmt.Scan(&email)

	// Kiểm tra xem người dùng có tồn tại hay không
	if _, exist := addressBook[email]; !exist {
		fmt.Println("Không tìm thấy người dùng trong danh bạ.")
		return
	}

	// Xóa người dùng
	delete(addressBook, email)
	fmt.Println("Người dùng đã được xóa khỏi danh bạ.")
}

// Hàm hiển thị toàn bộ danh bạ
func displayAddressBook() {
	fmt.Println("Danh bạ:")
	for email, user := range addressBook {
		fmt.Printf("Email: %s\nTên: %s\nSố điện thoại: %s\n\n", email, user.Name, user.Phone)
	}
}
