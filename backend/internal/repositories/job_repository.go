package repositories

import (
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/dto"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/models"
	"gorm.io/gorm"
)

type JobRepository struct{ DB *gorm.DB }

func (r *JobRepository) Create(job *models.Job) error {
	return r.DB.Create(job).Error
}

func (r *JobRepository) FindByID(id uint) (*models.Job, error) {
	var job models.Job
	err := r.DB.Preload("Recruiter").First(&job, id).Error
	return &job, err
}

func (r *JobRepository) Update(job *models.Job) error {
	return r.DB.Save(job).Error
}

func (r *JobRepository) Delete(job *models.Job) error {
	return r.DB.Delete(job).Error
}

func (r *JobRepository) FindByRecruiter(id uint) ([]models.Job, error) {
	var jobs []models.Job
	err := r.DB.Where("posted_by = ?", id).Order("created_at DESC").Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) List(q dto.JobQuery) ([]models.Job, int64, error) {
	db := r.DB.Model(&models.Job{}).Preload("Recruiter")

	if q.Search != "" {
		s := "%" + q.Search + "%"
		db = db.Where("title LIKE ? OR description LIKE ?", s, s)
	}
	if q.Location != "" {
		db = db.Where("location LIKE ?", "%"+q.Location+"%")
	}
	if q.Company != "" {
		db = db.Where("company LIKE ?", "%"+q.Company+"%")
	}
	if q.JobType != "" {
		db = db.Where("job_type = ?", q.JobType)
	}
	if q.MinSalary > 0 {
		db = db.Where("salary >= ?", q.MinSalary)
	}
	if q.MaxSalary > 0 {
		db = db.Where("salary <= ?", q.MaxSalary)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	switch q.Sort {
	case "salary_asc":
		db = db.Order("salary ASC")
	case "salary_desc":
		db = db.Order("salary DESC")
	case "oldest":
		db = db.Order("created_at ASC")
	default:
		db = db.Order("created_at DESC")
	}

	offset := (q.Page - 1) * q.Limit

	var jobs []models.Job
	err := db.Offset(offset).Limit(q.Limit).Find(&jobs).Error
	return jobs, total, err
}
