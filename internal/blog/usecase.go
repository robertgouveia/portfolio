package blog

import "github.com/robertgouveia/portfolio/internal/models"

type UseCase interface {
	GetBlogs(search string) ([]models.Blog, error)
	GetBlog(id int) (*models.Blog, error)
}
