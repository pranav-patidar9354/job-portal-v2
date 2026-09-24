package repositories

import (
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/models"
	"gorm.io/gorm"
)

type ApplicationRepository struct{ DB *gorm.DB }

func (r *ApplicationRepository) Create(app *models.Application) error {
	return r.DB.Create(app).Error
}

func (r *ApplicationRepository) FindByJobAndCandidate(jobID, candidateID uint) (*models.Application, error) {
	var app models.Application
	err := r.DB.Where("job_id = ? AND candidate_id = ?", jobID, candidateID).First(&app).Error
	return &app, err
}

func (r *ApplicationRepository) FindByCandidate(id uint) ([]models.Application, error) {
	var apps []models.Application
	err := r.DB.Preload("Job").Where("candidate_id = ?", id).Order("created_at DESC").Find(&apps).Error
	return apps, err
}

func (r *ApplicationRepository) FindByID(id uint) (*models.Application, error) {
	var app models.Application
	err := r.DB.Preload("Job").Preload("Candidate").First(&app, id).Error
	return &app, err
}

func (r *ApplicationRepository) FindByJob(id uint) ([]models.Application, error) {
	var apps []models.Application
	err := r.DB.Preload("Candidate").Preload("Job").Where("job_id = ?", id).Order("created_at DESC").Find(&apps).Error
	return apps, err
}

func (r *ApplicationRepository) Update(app *models.Application) error {
	return r.DB.Save(app).Error
}

func (r *ApplicationRepository) Delete(app *models.Application) error {
	return r.DB.Delete(app).Error
}
