package models

type Blog struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
}
