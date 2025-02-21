package handlers
import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Password string
}

func main() {
    db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect database")
    }

    db.AutoMigrate(&User{})
}
