package repository

import (
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	FindAll() ([]models.Schedule, error)
	FindByID(id uint) (*models.Schedule, error)
	Create(schedule *models.Schedule) error
	Update(schedule *models.Schedule) error
	Delete(id uint) error
}

type scheduleRepo struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepo{db}
}

func (r *scheduleRepo) FindAll() ([]models.Schedule, error) {
	var schedules []models.Schedule
	err := r.db.Preload("Doctor").Find(&schedules).Error
	return schedules, err
}

func (r *scheduleRepo) FindByID(id uint) (*models.Schedule, error) {
	var schedule models.Schedule
	err := r.db.Preload("Doctor").First(&schedule, id).Error
	return &schedule, err
}

func (r *scheduleRepo) Create(schedule *models.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepo) Update(schedule *models.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepo) Delete(id uint) error {
	return r.db.Delete(&models.Schedule{}, id).Error
}
