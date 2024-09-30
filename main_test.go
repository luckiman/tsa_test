package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateContactWithValidData(t *testing.T) {
	// Prepare test server
	router := gin.Default()

	// Initialize routes
	router.POST("/contacts", func(c *gin.Context) {
		var newContact Contact
		if err := c.ShouldBindJSON(&newContact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, phone := range newContact.PhoneNumbers {
			if !isValidAustralianPhoneNumber(phone) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Australian phone number"})
				return
			}
		}

		// Simulate successful response for test (skip DB interaction)
		c.JSON(http.StatusOK, gin.H{"status": "Contact saved"})
	})

	// Create valid contact
	contact := Contact{
		FullName:     "Alex Bell",
		Email:        "alex@bell-labs.com",
		PhoneNumbers: []string{"+61385786688", "+61412345678"},
	}

	// Marshal the contact into JSON
	jsonValue, _ := json.Marshal(contact)

	// Create a new HTTP request for testing
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Contact saved")
}

func TestCreateContactWithInvalidPhoneNumber(t *testing.T) {
	// Prepare test server
	router := gin.Default()

	// Initialize routes
	router.POST("/contacts", func(c *gin.Context) {
		var newContact Contact
		if err := c.ShouldBindJSON(&newContact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, phone := range newContact.PhoneNumbers {
			if !isValidAustralianPhoneNumber(phone) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Australian phone number"})
				return
			}
		}

		// Simulate successful response for test (skip DB interaction)
		c.JSON(http.StatusOK, gin.H{"status": "Contact saved"})
	})

	// Create a contact with invalid phone number
	contact := Contact{
		FullName:     "Alex Bell",
		Email:        "alex@bell-labs.com",
		PhoneNumbers: []string{"03 8578 6688"}, // Invalid format
	}

	// Marshal the contact into JSON
	jsonValue, _ := json.Marshal(contact)

	// Create a new HTTP request for testing
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid Australian phone number")
}

func TestCreateContactWithoutEmail(t *testing.T) {
	// Prepare test server
	router := gin.Default()

	// Initialize routes
	router.POST("/contacts", func(c *gin.Context) {
		var newContact Contact
		if err := c.ShouldBindJSON(&newContact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, phone := range newContact.PhoneNumbers {
			if !isValidAustralianPhoneNumber(phone) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Australian phone number"})
				return
			}
		}

		// Simulate successful response for test (skip DB interaction)
		c.JSON(http.StatusOK, gin.H{"status": "Contact saved"})
	})

	// Create a contact without an email and a valid phone number
	contact := Contact{
		FullName:     "Fredrik Idestam",
		PhoneNumbers: []string{"+61398889988"}, // Valid format with 9 digits
	}

	// Marshal the contact into JSON
	jsonValue, _ := json.Marshal(contact)

	// Create a new HTTP request for testing
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Contact saved")
}
