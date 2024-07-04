package API

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Book struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Author      string `json:"author"`
    Publisher   string `json:"publisher"`
    Description string `json:"description"`
}

var db *gorm.DB

func SetDB(database *gorm.DB) {
    db = database
}

func getBooks(c *gin.Context) {
    var books []Book
    result := db.Find(&books)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}
