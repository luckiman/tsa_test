package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Contact struct {
	FullName     string   `json:"full_name" binding:"required"`
	Email        string   `json:"email"`
	PhoneNumbers []string `json:"phone_numbers" binding:"required"`
}

func main() {
	router := gin.Default()

	// Setup Database (replace with real connection details)
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=root")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	router.POST("/contacts", func(c *gin.Context) {
		var newContact Contact
		if err := c.ShouldBindJSON(&newContact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate phone numbers
		for _, phone := range newContact.PhoneNumbers {
			if !isValidAustralianPhoneNumber(phone) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Australian phone number"})
				return
			}
		}

		// Insert contact into the database (assumes a table "contacts" exists)
		if newContact.Email != "" {
			query := `INSERT INTO contacts (full_name, email, phone_numbers) VALUES ($1, $2, $3)`
			_, err = db.Exec(query, newContact.FullName, newContact.Email, strings.Join(newContact.PhoneNumbers, ","))
		} else {
			query := `INSERT INTO contacts (full_name, phone_numbers) VALUES ($1, $2)`
			_, err = db.Exec(query, newContact.FullName, strings.Join(newContact.PhoneNumbers, ","))
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Contact saved"})
	})

	router.Run(":8080")
}

// isValidAustralianPhoneNumber validates phone numbers in E.164 format and Australian numbers
func isValidAustralianPhoneNumber(phone string) bool {
	// E.164 format for Australian landline, mobile, and toll-free numbers
	fmt.Println(phone)
	re := regexp.MustCompile(`^\+61([2-478]\d{8}|1800\d{6})$`)
	return re.MatchString(phone)
}
