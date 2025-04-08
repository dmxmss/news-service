package entities

type Tag struct {
	Name string `gorm:"primaryKey" json:"name"`
}
