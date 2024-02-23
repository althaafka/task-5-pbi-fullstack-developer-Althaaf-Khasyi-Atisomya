package models

type Photo struct {
	ID       uint   `gorm:"primaryKey; autoIncrement" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	Url	  string `json:"url"`
	UserID   uint   `json:"user_id"`
}
