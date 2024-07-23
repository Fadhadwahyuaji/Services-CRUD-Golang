package handlers

import (
	"net/http"

	"crud-golang/db"

	"crud-golang/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = ""
	user.UpdatedAt = ""
	_, err := db.DB.NamedExec(`INSERT INTO users (name, email, username, password, role_id, is_deleted)
                            VALUES (:name, :email, :username, :password, :role_id, :is_deleted)`, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	err := db.DB.Select(&users, "SELECT * FROM users WHERE is_deleted = false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id = $1 AND is_deleted = false", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil data pengguna yang ada dari database
	var existingUser models.User
	err := db.DB.Get(&existingUser, "SELECT * FROM users WHERE id = $1", input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	// Siapkan query untuk memperbarui data pengguna
	query := `UPDATE users SET name = :name, email = :email, username = :username, role_id = :role_id, is_deleted = :is_deleted`

	// Tambahkan password ke query jika disediakan
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		query += `, password = :password`
		input.Password = string(hashedPassword)
	}

	// Tambahkan kondisi WHERE untuk query
	query += ` WHERE id = :id`

	// Perbarui data pengguna di database
	_, err = db.DB.NamedExec(query, map[string]interface{}{
		"name":       input.Name,
		"email":      input.Email,
		"username":   input.Username,
		"password":   input.Password,
		"role_id":    input.RoleID,
		"is_deleted": input.IsDeleted,
		"id":         input.ID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("UPDATE users SET is_deleted = true WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
