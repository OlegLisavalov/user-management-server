package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Points     int    `gorm:"default:0"`
	ReferrerID *uint
}

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Points    int    `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	UserID    uint
}
