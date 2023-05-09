package main

import "time"

type unsplashPhoto struct {
	Id                     string           `json:"id"`
	CreatedAt              string           `json:"created_at"`
	UpdatedAt              string           `json:"updated_at"`
	Width                  int64            `json:"width"`
	Height                 int64            `json:"height"`
	Color                  string           `json:"color"`
	BlurHash               string           `json:"blur_hash"`
	Downloads              int64            `json:"downloads"`
	Likes                  int64            `json:"likes"`
	LikedByUser            bool             `json:"liked_by_user"`
	Description            string           `json:"description"`
	Exif                   exif             `json:"exif"`
	Location               locationUnsplash `json:"location"`
	CurrentUserCollections []collection     `json:"current_user_collections"`
	Urls                   urls             `json:"urls"`
	Links                  links            `json:"links"`
	User                   userUnsplash     `json:"user"`
}

type exif struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	ExposureTime string `json:"exposure_time"`
	Aperture     string `json:"aperture"`
	FocalLength  string `json:"focal_length"`
	Iso          int    `json:"iso"`
}

type locationUnsplash struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Position struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"position"`
}

type urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type collection struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	PublishedAt     int64  `json:"published_at"`
	LastCollectedAt int64  `json:"last_collected_at"`
	UpdatedAt       int64  `json:"updated_at"`
	CoverPhoto      string `json:"cover_photo"`
	User            string `json:"user"`
}

type links struct {
	Self             string `json:"self"`
	Html             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}
type userUnsplash struct {
	Id                string    `json:"id"`
	UpdatedAt         time.Time `json:"updated_at"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	PortfolioUrl      string    `json:"portfolio_url"`
	Bio               string    `json:"bio"`
	Location          string    `json:"location"`
	TotalLikes        int       `json:"total_likes"`
	TotalPhotos       int       `json:"total_photos"`
	TotalCollections  int       `json:"total_collections"`
	InstagramUsername string    `json:"instagram_username"`
	TwitterUsername   string    `json:"twitter_username"`
	Links             links     `json:"links"`
}
