package usecase

import (
	"encoding/json"
	"github.com/robertgouveia/portfolio/internal/blog"
	"github.com/robertgouveia/portfolio/internal/models"
	"os"
)

type BlogUseCase struct {
}

func NewBlogUseCase() blog.UseCase {
	return &BlogUseCase{}
}

func (u *BlogUseCase) GetBlogs() ([]models.Blog, error) {
	file, err := os.ReadFile("./data/blogs.json")
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
