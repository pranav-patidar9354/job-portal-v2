package repositories

import (
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/models"
	"gorm.io/gorm"
)

type SavedJobRepository struct{ DB *gorm.DB }

func (r *SavedJobRepository) Create(s *models.SavedJob) error {
	return r.DB.Create(s).Error
}

func (r *SavedJobRepository) Find(jobID, candidateID uint) (*models.SavedJob, error) {
	var s models.SavedJob
	err := r.DB.Where("job_id = ? AND candidate_id = ?", jobID, candidateID).First(&s).Error
	return &s, err
}

func (r *SavedJobRepository) List(candidateID uint) ([]models.SavedJob, error) {
	var s []models.SavedJob
	err := r.DB.Preload("Job").Where("candidate_id = ?", candidateID).Order("created_at DESC").Find(&s).Error
	return s, err
}

func (r *SavedJobRepository) Delete(s *models.SavedJob) error {
	return r.DB.Delete(s).Error
}
