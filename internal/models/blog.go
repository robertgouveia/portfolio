package models

type Blog struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
	Paragraphs  []string `json:"paragraphs"`
	Image       string   `json:"image"`
	Caption     string   `json:"caption"`
}
