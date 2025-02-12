package blog

import "github.com/robertgouveia/portfolio/internal/models"

type UseCase interface {
	GetBlogs() ([]models.Blog, error)
}
