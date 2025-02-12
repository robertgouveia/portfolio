package usecase

import (
	"embed"
	"encoding/json"
	"github.com/robertgouveia/portfolio/internal/blog"
	"github.com/robertgouveia/portfolio/internal/models"
)

//go:embed data/blogs.json
var dataFile embed.FS

type BlogUseCase struct {
}

func NewBlogUseCase() blog.UseCase {
	return &BlogUseCase{}
}

func (u *BlogUseCase) GetBlogs() ([]models.Blog, error) {
	file, err := dataFile.ReadFile("data/blogs.json")
	if err != nil {
		return nil, err
	}

	var blogs []models.Blog
	err = json.Unmarshal(file, &blogs)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
