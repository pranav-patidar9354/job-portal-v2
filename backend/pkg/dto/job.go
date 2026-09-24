package dto

type CreateJobRequest struct {
	Title       string  `json:"title" binding:"required,min=2,max=150"`
	Description string  `json:"description" binding:"required,min=10"`
	Company     string  `json:"company" binding:"required,min=2,max=150"`
	Location    string  `json:"location" binding:"required,min=2,max=150"`
	Salary      float64 `json:"salary" binding:"required,gt=0"`
	JobType     string  `json:"job_type" binding:"required,oneof=full-time part-time internship contract"`
}

type UpdateJobRequest struct {
	Title       string  `json:"title" binding:"omitempty,min=2,max=150"`
	Description string  `json:"description" binding:"omitempty,min=10"`
	Company     string  `json:"company" binding:"omitempty,min=2,max=150"`
	Location    string  `json:"location" binding:"omitempty,min=2,max=150"`
	Salary      float64 `json:"salary" binding:"omitempty,gt=0"`
	JobType     string  `json:"job_type" binding:"omitempty,oneof=full-time part-time internship contract"`
}

type JobQuery struct {
	Search    string
	Location  string
	Company   string
	JobType   string
	MinSalary float64
	MaxSalary float64
	Sort      string
	Page      int
	Limit     int
}
