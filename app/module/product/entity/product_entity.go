package entity

import "time"

var post Post

type Post struct {
	ID         string    `bson:"_id"`
	CreatedAt  time.Time `bson:"created_at"`
	IsDeleted  bool      `bson:"is_deleted"`
	PostDetail struct {
		Category      string `bson:"category"`
		Content       string `bson:"content"`
		Country       string `bson:"country"`
		Location      string `bson:"location"`
		Title         string `bson:"title"`
		UploadedFiles []struct {
			ID       string `bson:"_id"`
			ImageURL string `bson:"imageUrl"`
		} `bson:"uploaded_files"`
	} `bson:"post_detail"`
	UpdatedAt  time.Time `bson:"updated_at"`
	UserDetail struct {
		DisplayName string `bson:"displayName"`
		Thumbnail   string `bson:"thumbnail"`
	} `bson:"user_detail"`
	UserID    string `bson:"user_id"`
	Version   int    `bson:"version"`
	ViewCount int    `bson:"view_count"`
}
