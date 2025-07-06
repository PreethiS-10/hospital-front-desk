package repository

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindAll() ([]models.Department, error)
	FindByID(id uint) (*models.Department, error)
	Create(dept *models.Department) error
	Update(dept *models.Department) error
	Delete(id uint) error
}

type departmentRepo struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepo{db}
}

func (r *departmentRepo) FindAll() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.Find(&departments).Error
	return departments, err
}

func (r *departmentRepo) FindByID(id uint) (*models.Department, error) {
	var department models.Department
	err := r.db.First(&department, id).Error
	return &department, err
}

func (r *departmentRepo) Create(dept *models.Department) error {
	return r.db.Create(dept).Error
}

func (r *departmentRepo) Update(dept *models.Department) error {
	return r.db.Save(dept).Error
}

func (r *departmentRepo) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}
