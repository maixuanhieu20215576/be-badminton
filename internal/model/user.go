// internal/model/user.go

package model

// User là mô hình dữ liệu cho người dùng trong cơ sở dữ liệu
type User struct {
	ID   int    `json:"id"`   // ID của người dùng
	Name string `json:"name"` // Tên người dùng
}
