package services

import (
	"errors"

	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/dto"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/models"
	"github.com/pranav-patidar9354/job-portal-v2/backend/internal/repositories"
)

type ApplicationService struct {
	Apps *repositories.ApplicationRepository
	Jobs *repositories.JobRepository
}

func (s *ApplicationService) Apply(jobID, candidateID uint, req dto.ApplyRequest) (*models.Application, error) {
	if _, err := s.Jobs.FindByID(jobID); err != nil {
		return nil, errors.New("job not found")
	}

	if _, err := s.Apps.FindByJobAndCandidate(jobID, candidateID); err == nil {
		return nil, errors.New("you have already applied to this job")
	}

	app := &models.Application{
		JobID: jobID, CandidateID: candidateID,
		Status: "applied", CoverLetter: req.CoverLetter,
	}

	if err := s.Apps.Create(app); err != nil {
		return nil, err
	}
	return s.Apps.FindByID(app.ID)
}

func (s *ApplicationService) Mine(id uint) ([]models.Application, error) {
	return s.Apps.FindByCandidate(id)
}

func (s *ApplicationService) Withdraw(id, candidateID uint) error {
	app, err := s.Apps.FindByID(id)
	if err != nil {
		return err
	}
	if app.CandidateID != candidateID {
		return errors.New("you can only withdraw your own application")
	}
	return s.Apps.Delete(app)
}

func (s *ApplicationService) Applicants(jobID, recruiterID uint) ([]models.Application, error) {
	job, err := s.Jobs.FindByID(jobID)
	if err != nil {
		return nil, err
	}
	if job.PostedBy != recruiterID {
		return nil, errors.New("you can only view applicants for your own jobs")
	}
	return s.Apps.FindByJob(jobID)
}

func (s *ApplicationService) UpdateStatus(id, recruiterID uint, req dto.UpdateApplicationStatusRequest) (*models.Application, error) {
	app, err := s.Apps.FindByID(id)
	if err != nil {
		return nil, err
	}
	if app.Job.PostedBy != recruiterID {
		return nil, errors.New("you can only update applications for your own jobs")
	}

	app.Status = req.Status
	if err := s.Apps.Update(app); err != nil {
		return nil, err
	}
	return s.Apps.FindByID(id)
}
