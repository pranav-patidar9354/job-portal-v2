Job Portal V2

A full-stack job portal built with Go, Gin, GORM, MySQL, React, Vite, JWT, and Axios.

The application provides separate workflows for Candidates and Recruiters, including authentication, job management, job discovery, applications, application-status tracking, saved jobs, search, filtering, sorting, pagination, and role-based authorization.

🚀 Features

Authentication & Authorization

Candidate and Recruiter registration

Login with email and password

Password hashing using bcrypt

JWT-based authentication

Protected routes

Role-based authorization

Candidate and Recruiter have separate API access

Public job browsing without authentication

👨‍💻 Candidate Features

Browse all available jobs

Search jobs by keyword

Filter by location, company, job type, and salary range

Sort jobs

Pagination

View complete job details

Apply for jobs with an optional cover letter

View submitted applications

Track application status

Withdraw applications

Save jobs

Unsave jobs

View saved jobs

🏢 Recruiter Features

Recruiter dashboard

Create jobs

View own jobs

Edit own jobs

Delete own jobs

View applicants for own jobs

Update application status

Ownership protection so recruiters cannot modify another recruiter's jobs

🔐 Security

JWT authentication

bcrypt password hashing

Role-based middleware

Protected backend APIs

Recruiter job ownership checks

Candidate/recruiter route protection

CORS configuration

Environment-based configuration

🛠️ Tech Stack

Backend

Go

Gin — HTTP web framework

GORM — ORM

MySQL — relational database

JWT — authentication

bcrypt — password hashing

godotenv — environment configuration

Frontend

React

Vite

React Router

Axios

CSS

Development Tools

Git

GitHub

VS Code

MySQL Workbench

🏗️ Architecture

The application follows a layered backend architecture:

                    React Frontend
                         │
                         │ Axios / HTTP
                         ▼
                  Gin REST API
                         │
                ┌────────┴────────┐
                │                 │
          JWT Middleware     Role Middleware
                │                 │
                └────────┬────────┘
                         ▼
                      Handler
                         │
                         ▼
                      Service
                         │
                         ▼
                    Repository
                         │
                         ▼
                       GORM
                         │
                         ▼
                       MySQL

Request Flow

Frontend
   ↓
Axios
   ↓
Gin Router
   ↓
JWT / Role Middleware
   ↓
Handler
   ↓
Service
   ↓
Repository
   ↓
GORM
   ↓
MySQL

📁 Project Structure

job-portal-v2/
│
├── backend/
│   ├── cmd/
│   │   └── main.go
│   │
│   ├── config/
│   │   └── database.go
│   │
│   ├── internal/
│   │   ├── dto/
│   │   │   ├── application.go
│   │   │   ├── auth.go
│   │   │   └── job.go
│   │   │
│   │   ├── handlers/
│   │   │   ├── application_handler.go
│   │   │   ├── auth_handler.go
│   │   │   ├── job_handler.go
│   │   │   └── saved_job_handler.go
│   │   │
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   └── cors.go
│   │   │
│   │   ├── models/
│   │   │   ├── application.go
│   │   │   ├── job.go
│   │   │   ├── saved_job.go
│   │   │   └── user.go
│   │   │
│   │   ├── repositories/
│   │   │   ├── application_repository.go
│   │   │   ├── job_repository.go
│   │   │   ├── saved_job_repository.go
│   │   │   └── user_repository.go
│   │   │
│   │   ├── routes/
│   │   │   └── routes.go
│   │   │
│   │   └── services/
│   │       ├── application_service.go
│   │       ├── auth_service.go
│   │       ├── job_service.go
│   │       └── saved_job_service.go
│   │
│   ├── .env.example
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── context/
│   │   ├── pages/
│   │   ├── services/
│   │   ├── App.jsx
│   │   ├── main.jsx
│   │   └── styles.css
│   │
│   ├── .env.example
│   ├── index.html
│   ├── package.json
│   └── package-lock.json
│
├── .gitignore
└── README.md

🗄️ Database Design

The backend uses four main entities.

User
 │
 ├─────────────── creates ───────────────► Job
 │                                          │
 │                                          │
 │                                  receives applications
 │                                          │
 ▼                                          ▼
Candidate ───────── applies ──────────► Application
 │
 │
 └────────────── saves ────────────────► SavedJob

User

Stores:

ID

Name

Email

Password hash

Role

Timestamps

Roles:

candidate
recruiter

Job

Stores:

Title

Description

Company

Location

Salary

Job type

Recruiter ID

Timestamps

Application

Stores:

Job ID

Candidate ID

Application status

Cover letter

Timestamps

Application statuses:

applied
reviewing
shortlisted
accepted
rejected

A composite uniqueness constraint prevents the same candidate from applying to the same job multiple times.

SavedJob

Stores:

Job ID

Candidate ID

Created timestamp

A composite uniqueness constraint prevents duplicate saved-job records.

🔑 Authentication Flow

Candidate / Recruiter
        │
        ▼
     Login
        │
        ▼
  Email + Password
        │
        ▼
 bcrypt verification
        │
        ▼
      JWT
        │
        ▼
Frontend stores token
        │
        ▼
Authorization: Bearer <token>
        │
        ▼
JWT Middleware
        │
        ├── verifies signature
        ├── validates token
        ├── extracts user_id
        └── extracts role
        │
        ▼
Role-based authorization
        │
        ▼
Protected Handler

🌐 API Endpoints

Base URL:

http://localhost:8080/api/v1

Authentication

Method

Endpoint

Access

POST

/auth/register

Public

POST

/auth/login

Public

GET

/auth/me

Authenticated

Jobs

Method

Endpoint

Access

GET

/jobs

Public

GET

/jobs/:id

Public

POST

/recruiter/jobs

Recruiter

GET

/recruiter/jobs

Recruiter

PUT

/recruiter/jobs/:id

Recruiter

DELETE

/recruiter/jobs/:id

Recruiter

Applications

Method

Endpoint

Access

POST

/candidate/jobs/:jobID/apply

Candidate

GET

/candidate/applications

Candidate

DELETE

/candidate/applications/:id

Candidate

GET

/recruiter/jobs/:jobID/applicants

Recruiter

PATCH

/recruiter/applications/:id/status

Recruiter

Saved Jobs

Method

Endpoint

Access

POST

/candidate/jobs/:jobID/save

Candidate

GET

/candidate/saved-jobs

Candidate

🔎 Job Search

The jobs API supports:

Keyword search

Location filtering

Company filtering

Job type filtering

Minimum salary

Maximum salary

Sorting

Pagination

Example:

GET /api/v1/jobs?search=backend&location=Delhi&min_salary=500000&page=1&limit=6

⚙️ Local Setup

Prerequisites

Install:

Go

Node.js

npm

MySQL

Git

1. Clone the repository

git clone https://github.com/pranav-patidar9354/job-portal-v2.git
cd job-portal-v2

2. Create the MySQL database

Open MySQL and run:

CREATE DATABASE job_portal_v2;

3. Configure backend environment

Go to:

backend/

Create .env from .env.example.

Example:

PORT=8080
JWT_SECRET=your-secret-key

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=job_portal_v2

Do not commit .env to GitHub.

4. Start the backend

cd backend
go mod tidy
go run ./cmd

Backend:

http://localhost:8080

5. Start the frontend

Open another terminal:

cd frontend
npm install
npm run dev

Frontend:

http://localhost:5173

🔗 Frontend ↔ Backend

The frontend communicates with the backend using Axios.

React
  │
  ▼
Axios
  │
  ▼
http://localhost:8080/api/v1
  │
  ▼
Gin REST API

For authenticated requests, Axios automatically sends:

Authorization: Bearer <JWT>

The backend then verifies the token before allowing access to protected APIs.

🧪 Tested Application Flow

The following end-to-end flows have been tested:

Candidate Registration
        ↓
Candidate Login
        ↓
Browse Jobs
        ↓
Search / Filter / Sort
        ↓
View Job
        ↓
Apply
        ↓
Track Application
        ↓
Withdraw Application

Recruiter flow:

Recruiter Registration
        ↓
Recruiter Login
        ↓
Create Job
        ↓
Edit Job
        ↓
Delete Job
        ↓
View Applicants
        ↓
Update Application Status

Saved jobs:

Save Job
   ↓
Saved Jobs
   ↓
Unsave Job

Authorization:

Candidate
   ↓
Candidate APIs

Recruiter
   ↓
Recruiter APIs

Not Logged In
   ↓
Public APIs

📌 Key Backend Design Decisions

Layered Architecture

The backend separates responsibilities into:

Handler
  ↓
Service
  ↓
Repository
  ↓
Database

This keeps HTTP handling, business logic, database operations, and data models separated.

Repository Pattern

Repositories contain database-specific operations, while services contain business rules.

JWT Authentication

JWT is used so authenticated requests can carry the user's identity and role without sending the password again.

Role-Based Authorization

Middleware verifies whether the authenticated user has permission to access a specific API.

Job Ownership

Recruiters can modify and delete only jobs that belong to them.

🔮 Future Improvements

Possible extensions include:

Docker and Docker Compose

Swagger / OpenAPI documentation

Resume upload

Recruiter company profiles

Email notifications

Advanced recruiter analytics

Redis caching

Full-text job search

Refresh tokens

Password reset

Cloud deployment

Automated tests and CI/CD

🚀 Deployment Guide

This project is a decoupled full-stack application consisting of:
- **Frontend**: React + Vite SPA
- **Backend**: Go (Gin) + GORM REST API
- **Database**: MySQL

### 1. Deploy Frontend on Vercel

The frontend is ready for Vercel with SPA routing rewrite configuration in `frontend/vercel.json`.

#### Option A: Via Vercel Dashboard (Recommended)
1. Go to [vercel.com](https://vercel.com) and log in with your GitHub account.
2. Click **Add New...** > **Project**.
3. Import the repository: `pranav-patidar9354/job-portal-v2`.
4. In **Project Settings**:
   - **Root Directory**: Click *Edit* and select `frontend`.
   - **Framework Preset**: Vite (detected automatically).
   - **Build Command**: `npm run build`
   - **Output Directory**: `dist`
5. Under **Environment Variables**, add:
   - `VITE_API_URL`: Your deployed backend API URL (e.g., `https://your-backend.onrender.com/api/v1` or `http://localhost:8080/api/v1` for local testing).
6. Click **Deploy**.

#### Option B: Via Vercel CLI
```bash
cd frontend
npx vercel
```
Follow the interactive prompts to link and deploy the project.

---

### 2. Deploy Backend & Database

Because the backend is a persistent Go HTTP service using MySQL, deploy it to a service like **Render**, **Railway**, **Fly.io**, or a VPS:

1. **Database**: Create a managed MySQL database (e.g. on [Aiven](https://aiven.io), [Railway](https://railway.app), or [TiDB Cloud](https://tidbcloud.com)).
2. **Backend Environment Variables**:
   - `PORT`: `8080` (or platform default)
   - `JWT_SECRET`: A long random secret key
   - `CLIENT_URL`: `https://your-frontend.vercel.app` (enables CORS for your Vercel site)
   - `DB_HOST`: Your cloud database host
   - `DB_PORT`: `3306`
   - `DB_USER`: Your cloud database username
   - `DB_PASSWORD`: Your cloud database password
   - `DB_NAME`: Your database name
3. Deploy the backend using the root directory `backend/` and start command:
   ```bash
   go run ./cmd
   ```
   (or build binary with `go build -o server ./cmd` and run `./server`).

📄 License

This project is intended for learning, portfolio, and demonstration purposes.