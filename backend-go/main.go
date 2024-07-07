package main

import (
    "fmt"
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

func main() {
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        "localhost", 5422, "myuser", "mypassword", "mydatabasebook")

    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // db.Create(&books)

    // Migrate schema
    db.AutoMigrate(&Book{})
    fmt.Println("Database migration completed!")

    // Setup routes
    r := gin.Default()

    // CORS middleware
    r.Use(cors.Default())

    r.GET("/", getBooks)

    // Run server
    r.Run(":8080")
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
