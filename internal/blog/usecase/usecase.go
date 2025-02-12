package usecase

import (
	"embed"
	"encoding/json"
	"errors"
	"github.com/robertgouveia/portfolio/internal/blog"
	"github.com/robertgouveia/portfolio/internal/models"
	"log"
	"strings"
)

//go:embed data/blogs.json
var dataFile embed.FS

type BlogUseCase struct {
	log *log.Logger
}

func NewBlogUseCase(log *log.Logger) blog.UseCase {
	return &BlogUseCase{
		log: log,
	}
}

func (u *BlogUseCase) GetBlogs(search string) ([]models.Blog, error) {
	file, err := dataFile.ReadFile("data/blogs.json")
	if err != nil {
		return nil, err
	}

	var blogs []models.Blog
	err = json.Unmarshal(file, &blogs)
	if err != nil {
		return nil, err
	}

	if search == "" {
		return blogs, nil
	}

	var match []models.Blog
	for _, b := range blogs {
		if strings.Contains(strings.ToLower(b.Title), search) || strings.Contains(strings.ToLower(b.Description), search) {
			match = append(match, b)
		}
	}

	return match, nil
}

func (u *BlogUseCase) GetBlog(id int) (*models.Blog, error) {
	file, err := dataFile.ReadFile("data/blogs.json")
	if err != nil {
		return nil, err
	}

	var blogs []models.Blog
	err = json.Unmarshal(file, &blogs)
	if err != nil {
		return nil, err
	}

	for _, b := range blogs {
		if b.Id == id {
			return &b, nil
		}
	}

	return nil, errors.New("not found")
}
