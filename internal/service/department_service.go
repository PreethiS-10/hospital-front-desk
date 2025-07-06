package service

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/repository"
)

type DepartmentService interface {
	GetAll() ([]models.Department, error)
	GetByID(id uint) (*models.Department, error)
	Create(dept *models.Department) error
	Update(dept *models.Department) error
	Delete(id uint) error
}

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(r repository.DepartmentRepository) DepartmentService {
	return &departmentService{repo: r}
}

func (s *departmentService) GetAll() ([]models.Department, error) {
	return s.repo.FindAll()
}

func (s *departmentService) GetByID(id uint) (*models.Department, error) {
	return s.repo.FindByID(id)
}

func (s *departmentService) Create(dept *models.Department) error {
	return s.repo.Create(dept)
}

func (s *departmentService) Update(dept *models.Department) error {
	return s.repo.Update(dept)
}

func (s *departmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
