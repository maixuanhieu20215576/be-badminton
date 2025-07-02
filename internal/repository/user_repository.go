package repository

import (
	"go-sql-project/internal/model"
	"database/sql"
	"log"
)

// Giả sử db là kết nối cơ sở dữ liệu đã được thiết lập
var db *sql.DB

func FetchUsers() ([]model.User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
