package models

type Role struct {
	ID        int    `db:"id" json:"id"`
	RoleName  string `db:"role_name" json:"role_name"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type Position struct {
	ID           int    `db:"id" json:"id"`
	PositionName string `db:"position_name" json:"position_name"`
	IsDeleted    bool   `db:"is_deleted" json:"is_deleted"`
	UpdatedAt    string `db:"updated_at" json:"updated_at"`
	CreatedAt    string `db:"created_at" json:"created_at"`
}

type User struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	RoleID    int    `db:"role_id" json:"role_id"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type Employee struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Email       string `db:"email" json:"email"`
	PositionID  int    `db:"position_id" json:"position_id"`
	IsDeleted   bool   `db:"is_deleted" json:"is_deleted"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}
