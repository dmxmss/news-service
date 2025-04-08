package entities

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Email string
	Password string
	Role string `gorm:"type:json"`
	TelegramID *int
}
