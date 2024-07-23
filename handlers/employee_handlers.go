package handlers

import (
	"log"
	"net/http"

	"crud-golang/db"

	"crud-golang/models"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.CreatedAt = ""
	employee.UpdatedAt = ""
	_, err := db.DB.NamedExec(`INSERT INTO employee (name, phone_number, email, position_id)
                            VALUES (:name, :phone_number, :email, :position_id)`, &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func GetEmployee(c *gin.Context) {
	var employee []models.Employee
	err := db.DB.Select(&employee, "SELECT * FROM employee WHERE is_deleted = false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	err := db.DB.Get(&employee, "SELECT * FROM employee WHERE id = $1 AND is_deleted = false", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the employee struct to ensure values are correct
	log.Printf("Updating Employee: %+v\n", employee)

	// Updated query string with correct column names
	query := `UPDATE employee SET name=:name, phone_number=:phone_number, email=:email, position_id=:position_id, 
            is_deleted=:is_deleted, updated_at=CURRENT_TIMESTAMP WHERE id=:id`

	// Log the query and parameters
	log.Printf("Executing query: %s\nWith parameters: %+v\n", query, employee)

	_, err := db.DB.NamedExec(query, &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("UPDATE employee SET is_deleted = true WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
