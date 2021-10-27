package category

import (
	"go-hexagonal/business/category"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}
type CategoryTable struct {
	Id        int       `gorm:"id"`
	Name      string    `gorm:"name,index:unique"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func (col *CategoryTable) ToCategory() category.Category {
	var category category.Category

	category.CategoryID = col.Id
	category.CategoryName = col.Name

	return category
}

func newCategoryTable(category category.Category) *CategoryTable {

	return &CategoryTable{
		category.CategoryID,
		category.CategoryName,
		category.CreatedAt,
		category.UpdatedAt,
		category.DeletedAt,
	}

}

//NewGormDBRepository Generate Gorm DB pet repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) FindCategoryByID(id int) (*category.Category, error) {
	var categoryData CategoryTable

	err := repo.DB.Where("id = ?", id).First(&categoryData).Error

	if err != nil {
		return nil, err
	}

	category := categoryData.ToCategory()

	return &category, nil

}

func (repo *GormRepository) FindAllCategory(skip int, rowPerPage int) ([]category.Category, error) {

	var categories []CategoryTable

	err := repo.DB.Offset(skip).Limit(rowPerPage).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var result []category.Category
	for _, value := range categories {
		result = append(result, value.ToCategory())
	}

	return result, nil
}

func (repo *GormRepository) InsertCategory(category category.Category) error {

	categoryData := newCategoryTable(category)

	err := repo.DB.Create(categoryData).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *GormRepository) UpdateCategory(category category.Category) error {

	categoryData := newCategoryTable(category)

	err := repo.DB.Model(&categoryData).Where("id = ?", category.CategoryID).Updates(
		CategoryTable{
			Name:      category.CategoryName,
			UpdatedAt: category.UpdatedAt,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) DeleteCategory(id int) error {
	var category CategoryTable
	err := repo.DB.Where("id = ?", id).Delete(&category).Error

	if err != nil {
		return err
	}

	return nil
}
