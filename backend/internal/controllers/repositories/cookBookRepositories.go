package repositories

import (
	"cook_book/backend/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CookBookRepositories interface {
	GetAll() ([]*model.CookBook, error)
	GetByID(id uint) (*model.CookBook, error)
	Create(recipe *model.CookBook) error
	Update(recipe *model.CookBook, id uint) error
	Delete(id uint) error
}

type cookBookRepositories struct {
	db *gorm.DB
}

type CookBookRepoConfig struct {
	DB *gorm.DB
}

func NewCookBookRepositories(r CookBookRepoConfig) CookBookRepositories {
	return &cookBookRepositories{
		db: r.DB,
	}
}

func (r *cookBookRepositories) GetAll() ([]*model.CookBook, error) {
	var recipes []*model.CookBook

	result := r.db.Find(&recipes)
	if result.Error != nil {
		return nil, result.Error
	}

	return recipes, nil
}

func (r *cookBookRepositories) GetByID(id uint) (*model.CookBook, error) {
	var recipe *model.CookBook

	result := r.db.First(&recipe, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return recipe, nil
}

func (r *cookBookRepositories) Create(recipe *model.CookBook) error {
	result := r.db.Create(&recipe)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *cookBookRepositories) Update(recipe *model.CookBook, id uint) error {
	result := r.db.Model(&model.CookBook{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&recipe)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *cookBookRepositories) Delete(id uint) error {
	result := r.db.Delete(&model.CookBook{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
