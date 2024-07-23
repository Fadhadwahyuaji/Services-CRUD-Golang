package handlers

import (
	"log"
	"net/http"

	"crud-golang/db"

	"crud-golang/models"

	"github.com/gin-gonic/gin"
)

func CreatePosition(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position.CreatedAt = ""
	position.UpdatedAt = ""
	_, err := db.DB.NamedExec(`INSERT INTO position (position_name)
                            VALUES (:position_name)`, &position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, position)
}

func GetPosition(c *gin.Context) {
	var position []models.Position
	err := db.DB.Select(&position, "SELECT * FROM position WHERE is_deleted = false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, position)
}

func GetPositionByID(c *gin.Context) {
	id := c.Param("id")

	var position models.Position
	err := db.DB.Get(&position, "SELECT * FROM position WHERE id = $1 AND is_deleted = false", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "position not found"})
		return
	}
	c.JSON(http.StatusOK, position)
}

func UpdatePosition(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the position struct to ensure values are correct
	log.Printf("Updating position: %+v\n", position)

	// Updated query string with correct column names
	query := `UPDATE position SET position_name=:position_name, 
            is_deleted=:is_deleted, updated_at=CURRENT_TIMESTAMP WHERE id=:id`

	// Log the query and parameters
	log.Printf("Executing query: %s\nWith parameters: %+v\n", query, position)

	_, err := db.DB.NamedExec(query, &position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, position)
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("UPDATE position SET is_deleted = true WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "position deleted successfully"})
}
