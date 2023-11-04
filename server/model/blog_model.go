package model

// blog will contain
// - unique id
// - title
// - post
// - image

type Blog struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Title string `json:"title" gorm:"not null;column:title;size:225"`
	Post  string `json:"post" gorm:"not null;column:post;size:225"`
	// Image string 
}