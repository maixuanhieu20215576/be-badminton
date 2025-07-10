package model

// User model đại diện cho bảng users trong cơ sở dữ liệu
type User struct {
	ID                int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username          string `json:"username" gorm:"not null"`
	Avatar            string `json:"avatar"`
	ReferenceUserID   int    `json:"reference_user_id"`
	Status            string `json:"status" gorm:"default:'Playing'"`
}

