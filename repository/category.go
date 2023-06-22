package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	err := c.db.Create(Category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	return c.db.Where(id).Updates(category).Error // TODO: replace this

}

func (c *categoryRepository) Delete(id int) error {
	return c.db.Where(id).Delete(&model.Category{}).Error // TODO: replace this
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	var Category model.Category
	err := c.db.Where("id = ?", id).First(&Category).Error
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	var list []model.Category
	res, err := c.db.Table("categories").Rows()
	if err != nil {
		return []model.Category{}, err
	}
	defer res.Close()

	for res.Next() {
		c.db.ScanRows(res, &list)
	}
	return list, nil // TODO: replace this
}
