package service

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"github.com/PreethiS-10/fiber-crud/internal/repository"
)

type ScheduleService interface {
	GetAll() ([]models.Schedule, error)
	GetByID(id uint) (*models.Schedule, error)
	Create(schedule *models.Schedule) error
	Update(schedule *models.Schedule) error
	Delete(id uint) error
}

type scheduleService struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(r repository.ScheduleRepository) ScheduleService {
	return &scheduleService{repo: r}
}

func (s *scheduleService) GetAll() ([]models.Schedule, error) {
	return s.repo.FindAll()
}

func (s *scheduleService) GetByID(id uint) (*models.Schedule, error) {
	return s.repo.FindByID(id)
}

func (s *scheduleService) Create(schedule *models.Schedule) error {
	return s.repo.Create(schedule)
}

func (s *scheduleService) Update(schedule *models.Schedule) error {
	return s.repo.Update(schedule)
}

func (s *scheduleService) Delete(id uint) error {
	return s.repo.Delete(id)
}
