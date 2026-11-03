package service

import (
	"chi-api/internal/models"
	"chi-api/internal/repository"
	"fmt"
)

type LinkServicer interface {
	ListAllLinks() ([]models.Link, error)
	Save(url string) (models.Link, error)
	GetURL(code string) (models.Link, error)
}

type LinkService struct {
	Repo repository.LinksRepo
}

func NewLinkService(repo repository.LinksRepo) LinkServicer {
	return &LinkService{
		Repo: repo,
	}
}

func (ls *LinkService) ListAllLinks() ([]models.Link, error) {
	links, err := ls.Repo.FindAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get all links: %w", err)
	}

	return links, nil
}

func (ls *LinkService) Save(url string) (models.Link, error) {
	link, err := ls.Repo.AddLink(url)

	if err != nil {
		return models.Link{}, err
	}

	return link, nil
}

func (ls *LinkService) GetURL(code string) (models.Link, error) {
	link, err := ls.Repo.FindLink(code)

	if err != nil {
		return models.Link{}, err
	}
	return link, nil
}
