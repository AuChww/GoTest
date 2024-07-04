package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "./API" 
)

const (
    host     = "localhost"  
    port     = 5422         
    user     = "myuser"     
    password = "mypassword" 
    dbname   = "mydatabasebook" 
)

var db *gorm.DB

func main() {
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // ทำการ Migrate schema
    db.AutoMigrate(&Book{})
    fmt.Println("Database migration completed!")

    // ตั้งค่า route และเปิดเซิร์ฟเวอร์
    r := gin.Default()
    r.GET("/books", API.getBooks) // เรียกใช้ getBooks จาก package API
    r.Run(":8080")
}
