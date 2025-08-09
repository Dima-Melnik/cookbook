package services

import (
	r "cook_book/backend/internal/controllers/repositories"
	"cook_book/backend/internal/model"
)

type CookBookServices interface {
	GetAll() ([]*model.CookBook, error)
	GetByID(id uint) (*model.CookBook, error)
	Create(recipe *model.CookBook) error
	Update(recipe *model.CookBook, id uint) error
	Delete(id uint) error
}

type cookBookServices struct {
	repo r.CookBookRepositories
}

type CookBookServiceConfig struct {
	Repo r.CookBookRepositories
}

func NewCookBookServices(s *CookBookServiceConfig) CookBookServices {
	return &cookBookServices{
		repo: s.Repo,
	}
}

func (s *cookBookServices) GetAll() ([]*model.CookBook, error) {
	result, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cookBookServices) GetByID(id uint) (*model.CookBook, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cookBookServices) Create(recipe *model.CookBook) error {
	if err := s.repo.Create(recipe); err != nil {
		return err
	}

	return nil
}

func (s *cookBookServices) Update(recipe *model.CookBook, id uint) error {
	if err := s.repo.Update(recipe, id); err != nil {
		return err
	}

	return nil
}

func (s *cookBookServices) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
