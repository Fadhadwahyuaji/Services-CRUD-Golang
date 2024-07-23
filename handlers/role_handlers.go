package handlers

import (
	"log"
	"net/http"

	"crud-golang/db"

	"crud-golang/models"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.CreatedAt = ""
	role.UpdatedAt = ""
	_, err := db.DB.NamedExec(`INSERT INTO role (role_name)
                            VALUES (:role_name)`, &role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func GetRole(c *gin.Context) {
	var role []models.Role
	err := db.DB.Select(&role, "SELECT * FROM role WHERE is_deleted = false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

func GetRoleByID(c *gin.Context) {
	id := c.Param("id")

	var role models.Role
	err := db.DB.Get(&role, "SELECT * FROM role WHERE id = $1 AND is_deleted = false", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the role struct to ensure values are correct
	log.Printf("Updating role: %+v\n", role)

	// Updated query string with correct column names
	query := `UPDATE role SET role_name=:role_name WHERE id=:id`

	// Log the query and parameters
	log.Printf("Executing query: %s\nWith parameters: %+v\n", query, role)

	_, err := db.DB.NamedExec(query, &role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("UPDATE role SET is_deleted = true WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role deleted successfully"})
}
