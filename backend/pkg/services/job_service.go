package services

import (
	"errors"

	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/dto"
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/models"
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/repositories"
)

type JobService struct{ Jobs *repositories.JobRepository }

func (s *JobService) Create(req dto.CreateJobRequest, recruiterID uint) (*models.Job, error) {
	job := &models.Job{
		Title: req.Title, Description: req.Description,
		Company: req.Company, Location: req.Location,
		Salary: req.Salary, JobType: req.JobType,
		PostedBy: recruiterID,
	}
	if err := s.Jobs.Create(job); err != nil {
		return nil, err
	}
	return s.Jobs.FindByID(job.ID)
}

func (s *JobService) List(q dto.JobQuery) ([]models.Job, int64, error) {
	return s.Jobs.List(q)
}

func (s *JobService) Get(id uint) (*models.Job, error) {
	return s.Jobs.FindByID(id)
}

func (s *JobService) OwnJobs(id uint) ([]models.Job, error) {
	return s.Jobs.FindByRecruiter(id)
}

func (s *JobService) Update(id, recruiterID uint, req dto.UpdateJobRequest) (*models.Job, error) {
	job, err := s.Jobs.FindByID(id)
	if err != nil {
		return nil, err
	}
	if job.PostedBy != recruiterID {
		return nil, errors.New("you can only update your own jobs")
	}

	if req.Title != "" { job.Title = req.Title }
	if req.Description != "" { job.Description = req.Description }
	if req.Company != "" { job.Company = req.Company }
	if req.Location != "" { job.Location = req.Location }
	if req.Salary > 0 { job.Salary = req.Salary }
	if req.JobType != "" { job.JobType = req.JobType }

	if err := s.Jobs.Update(job); err != nil {
		return nil, err
	}
	return s.Jobs.FindByID(id)
}

func (s *JobService) Delete(id, recruiterID uint) error {
	job, err := s.Jobs.FindByID(id)
	if err != nil {
		return err
	}
	if job.PostedBy != recruiterID {
		return errors.New("you can only delete your own jobs")
	}
	return s.Jobs.Delete(job)
}
