# Job Portal V2

A modern, full-stack recruitment platform connecting job seekers and employers. Built with a high-performance **Go (Gin + GORM)** backend and a responsive **React 19 (Vite)** frontend, backed by **PostgreSQL (Neon)**. Designed for speed, clean layered architecture, and seamless cloud deployment on Vercel.

---

## 📌 Table of Contents

- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Key Features](#key-features)
  - [For Candidates](#for-candidates)
  - [For Recruiters](#for-recruiters)
  - [Core Platform & Security](#core-platform--security)
- [Architecture & Folder Structure](#architecture--folder-structure)
- [Database Schema](#database-schema)
- [API Reference](#api-reference)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Environment Variables](#environment-variables)
- [Deployment](#deployment)
- [License](#license)

---

## 🚀 Overview

Job Portal V2 provides an end-to-end recruitment workflow. Recruiters can post openings, manage listings, and review candidate applications with a real-time status update pipeline. Candidates can discover opportunities through search and multi-parameter filtering, save favorite listings, apply with personalized cover letters, and track their application progress.

---

## 🛠 Tech Stack

### Backend
- **Language:** Go (1.25+)
- **Web Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io/) (PostgreSQL Driver)
- **Database:** PostgreSQL / [Neon Serverless Postgres](https://neon.tech/)
- **Authentication:** JSON Web Tokens (JWT) via [golang-jwt/jwt](https://github.com/golang-jwt/jwt) (HMAC-SHA256, 24-hour expiration)
- **Password Hashing:** [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- **CORS:** [gin-contrib/cors](https://github.com/gin-contrib/cors) with dynamic origin matching (localhost, Vercel deployments, custom origins)
- **Configuration:** [godotenv](https://github.com/joho/godotenv)
- **Deployment:** Vercel Serverless Functions (`api/index.go`) or Standalone Go HTTP server (`cmd/main.go`)

### Frontend
- **Framework:** [React 19](https://react.dev/)
- **Build Tool:** [Vite 7](https://vitejs.dev/)
- **Routing:** [React Router v7](https://reactrouter.com/)
- **HTTP Client:** [Axios](https://axios-http.com/) (configured with auth token request interceptors and error handlers)
- **Styling:** Custom responsive CSS design system with card layouts, tags, and status badges

---

## ✨ Key Features

### For Candidates
- **Browse & Search Jobs:** Full-text search across job titles and descriptions.
- **Advanced Filtering & Sorting:**
  - Filter by location, company, and job type (`full-time`, `part-time`, `internship`, `contract`).
  - Filter by minimum and maximum salary thresholds.
  - Sort by newest, oldest, highest salary, or lowest salary.
- **Server-Side Pagination:** Fast and scalable job listing pagination.
- **One-Click Application:** Apply to any listing with an optional cover letter (duplicate applications are automatically prevented).
- **Application Tracking:** Real-time visibility into application status (`applied`, `reviewing`, `shortlisted`, `accepted`, `rejected`).
- **Application Withdrawal:** Cancel or withdraw submitted applications at any time.
- **Saved Jobs (Bookmarks):** Save listings to a personalized dashboard for easy access later.

### For Recruiters
- **Recruiter Dashboard:** Centralized view of all jobs posted by the logged-in recruiter.
- **Job Management (CRUD):**
  - Create new job postings with detailed metadata (title, company, location, salary, job type, description).
  - Edit existing job posts directly from the dashboard.
  - Delete listings with automatic cleanup.
- **Applicant Review:** View all candidates who applied to a specific opening along with their contact details and cover letters.
- **Application Status Pipeline:** Update candidate status across stages (`applied` ➔ `reviewing` ➔ `shortlisted` ➔ `accepted` / `rejected`).

### Core Platform & Security
- **Role-Based Access Control (RBAC):** Distinct roles (`candidate` and `recruiter`) enforced at both backend middleware and frontend route levels.
- **Secure Authentication:** Passwords hashed with bcrypt; stateless authorization using JWT bearer tokens.
- **Database Connection Pooling:** Configured with max open/idle connection limits and connection lifetimes optimized for serverless and long-running servers.
- **Automatic Migrations:** GORM automatically creates and synchronizes database tables and indexes on startup.

---

## 📁 Architecture & Folder Structure

```text
job-portal-v2/
├── backend/
│   ├── api/
│   │   └── index.go                 # Vercel serverless entry point
│   ├── cmd/
│   │   └── main.go                  # Local standalone HTTP server entry point
│   ├── config/
│   │   └── database.go              # Database connection pool & auto-migrations
│   ├── pkg/
│   │   ├── dto/                     # Request validation & transfer objects
│   │   │   ├── application.go
│   │   │   ├── auth.go
│   │   │   └── job.go
│   │   ├── handlers/                # HTTP route controllers
│   │   │   ├── application_handler.go
│   │   │   ├── auth_handler.go
│   │   │   ├── job_handler.go
│   │   │   └── saved_job_handler.go
│   │   ├── middleware/              # JWT auth, RBAC & CORS policies
│   │   │   ├── auth.go
│   │   │   └── cors.go
│   │   ├── models/                  # Database entity definitions (GORM)
│   │   │   ├── application.go
│   │   │   ├── job.go
│   │   │   ├── saved_job.go
│   │   │   └── user.go
│   │   ├── repositories/            # Data persistence & database queries
│   │   │   ├── application_repository.go
│   │   │   ├── job_repository.go
│   │   │   ├── saved_job_repository.go
│   │   │   └── user_repository.go
│   │   ├── routes/                  # Route registry and Gin router setup
│   │   │   └── routes.go
│   │   └── services/                # Business logic layer
│   │       ├── application_service.go
│   │       ├── auth_service.go
│   │       ├── job_service.go
│   │       └── saved_job_service.go
│   ├── .env.example
│   ├── go.mod
│   ├── go.sum
│   └── vercel.json                  # Backend Vercel serverless configuration
│
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── components/              # Shared UI components
│   │   │   ├── Navbar.jsx
│   │   │   └── ProtectedRoute.jsx
│   │   ├── context/
│   │   │   └── AuthContext.jsx      # Global authentication state
│   │   ├── pages/                   # Application views
│   │   │   ├── Applicants.jsx       # Recruiter applicant review
│   │   │   ├── Applications.jsx     # Candidate application list
│   │   │   ├── Home.jsx             # Landing page
│   │   │   ├── JobDetails.jsx       # Job details & application submission
│   │   │   ├── Jobs.jsx             # Public job search & filtering
│   │   │   ├── Login.jsx            # Sign in
│   │   │   ├── RecruiterDashboard.jsx # Recruiter job management
│   │   │   ├── Register.jsx         # Sign up
│   │   │   └── SavedJobs.jsx        # Bookmarked jobs
│   │   ├── services/
│   │   │   └── api.js               # Axios instance with interceptors
│   │   ├── App.jsx                  # Client routes definition
│   │   ├── main.jsx                 # React root mount
│   │   └── styles.css               # Global application styling
│   ├── .env.example
│   ├── index.html
│   ├── package.json
│   └── vercel.json                  # Frontend SPA routing configuration
│
└── README.md
```

---

## 🗄 Database Schema

The database consists of 4 main relational models:

```mermaid
erDiagram
    USERS ||--o{ JOBS : "posts (recruiter)"
    USERS ||--o{ APPLICATIONS : "submits (candidate)"
    USERS ||--o{ SAVED_JOBS : "bookmarks (candidate)"
    JOBS ||--o{ APPLICATIONS : "receives"
    JOBS ||--o{ SAVED_JOBS : "referenced in"

    USERS {
        uint id PK
        string name
        string email UK
        string password
        string role "candidate | recruiter"
        datetime created_at
        datetime updated_at
    }

    JOBS {
        uint id PK
        string title
        text description
        string company
        string location
        float salary
        string job_type "full-time | part-time | internship | contract"
        uint posted_by FK
        datetime created_at
        datetime updated_at
    }

    APPLICATIONS {
        uint id PK
        uint job_id FK
        uint candidate_id FK
        string status "applied | reviewing | shortlisted | accepted | rejected"
        text cover_letter
        datetime created_at
        datetime updated_at
    }

    SAVED_JOBS {
        uint id PK
        uint job_id FK
        uint candidate_id FK
        datetime created_at
    }
```

---

## 🔌 API Reference

**Base URL:** `/api/v1`

### 1. Authentication
| Method | Endpoint | Access | Description |
|---|---|---|---|
| `POST` | `/auth/register` | Public | Register a new user (`name`, `email`, `password`, `role`) |
| `POST` | `/auth/login` | Public | Authenticate user & return JWT token |
| `GET` | `/auth/me` | Authenticated | Fetch current user profile |

### 2. Jobs (Public)
| Method | Endpoint | Access | Description |
|---|---|---|---|
| `GET` | `/jobs` | Public | List jobs with pagination, filtering & sorting |
| `GET` | `/jobs/:id` | Public | Retrieve detailed information for a single job |

**Query Parameters for `GET /jobs`:**
- `search` *(string)*: Search keywords across title and description
- `location` *(string)*: Filter by location (case-insensitive substring)
- `company` *(string)*: Filter by company name (case-insensitive substring)
- `job_type` *(string)*: `full-time`, `part-time`, `internship`, or `contract`
- `min_salary` *(number)*: Minimum salary
- `max_salary` *(number)*: Maximum salary
- `sort` *(string)*: `newest` (default), `oldest`, `salary_desc`, `salary_asc`
- `page` *(number)*: Page number (default: 1)
- `limit` *(number)*: Items per page (default: 9, max: 50)

### 3. Recruiter Endpoints (Auth + Role: `recruiter`)
| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/recruiter/jobs` | Retrieve all jobs posted by the authenticated recruiter |
| `POST` | `/recruiter/jobs` | Create a new job listing |
| `PUT` | `/recruiter/jobs/:id` | Update an existing job listing |
| `DELETE` | `/recruiter/jobs/:id` | Delete a job listing |
| `GET` | `/recruiter/jobs/:jobID/applicants` | View all candidate applications for a specific job |
| `PATCH` | `/recruiter/applications/:id/status` | Update applicant status (`applied`, `reviewing`, `shortlisted`, `accepted`, `rejected`) |

### 4. Candidate Endpoints (Auth + Role: `candidate`)
| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/candidate/jobs/:jobID/apply` | Apply to a job with optional `cover_letter` |
| `GET` | `/candidate/applications` | List all applications submitted by candidate |
| `DELETE` | `/candidate/applications/:id` | Withdraw an existing application |
| `POST` | `/candidate/jobs/:jobID/save` | Toggle save/bookmark status for a job |
| `GET` | `/candidate/saved-jobs` | List all bookmarked jobs for candidate |

---

## 💻 Getting Started

### Prerequisites
- **Go:** 1.25 or newer installed ([Download Go](https://go.dev/dl/))
- **Node.js:** 18.x or newer & npm ([Download Node.js](https://nodejs.org/))
- **PostgreSQL:** Local PostgreSQL instance or a free cloud database (e.g., [Neon](https://neon.tech/))

---

### Backend Setup

1. **Navigate to the backend directory:**
   ```bash
   cd backend
   ```

2. **Install Go dependencies:**
   ```bash
   go mod download
   ```

3. **Configure Environment Variables:**
   Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
   Open `.env` and configure your database and JWT secret:
   ```env
   PORT=8080
   JWT_SECRET=your-super-secret-jwt-key
   CLIENT_URL=http://localhost:5173

   # Neon or PostgreSQL connection string
   DATABASE_URL=postgresql://user:password@ep-sample.us-east-2.aws.neon.tech/neondb?sslmode=require
   ```

4. **Run the Backend Server:**
   ```bash
   go run cmd/main.go
   ```
   The backend server will start on `http://localhost:8080`.

---

### Frontend Setup

1. **Navigate to the frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Configure Environment Variables:**
   Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
   Ensure the API endpoint matches your backend:
   ```env
   VITE_API_URL=http://localhost:8080/api/v1
   ```

4. **Start the Development Server:**
   ```bash
   npm run dev
   ```
   Open your browser and navigate to `http://localhost:5173`.

---

## ⚙️ Environment Variables

### Backend (`backend/.env`)

| Variable | Required | Default | Description |
|---|---|---|---|
| `PORT` | No | `8080` | Port for the local HTTP server |
| `JWT_SECRET` | **Yes** | — | Secret key used for signing HMAC-SHA256 JWT tokens |
| `CLIENT_URL` | No | `http://localhost:5173` | Allowed frontend origin for CORS |
| `ALLOWED_ORIGINS`| No | — | Comma-separated list of additional allowed origins |
| `DATABASE_URL` | Conditional | — | Direct PostgreSQL / Neon URI (recommended) |
| `DB_HOST` | Conditional | `localhost` | Database host (used if `DATABASE_URL` is omitted) |
| `DB_PORT` | Conditional | `5432` | Database port |
| `DB_USER` | Conditional | `postgres` | Database username |
| `DB_PASSWORD` | Conditional | — | Database password |
| `DB_NAME` | Conditional | `job_portal_v2` | Database name |
| `DB_SSLMODE` | Conditional | `disable` | SSL mode (`disable`, `require`, etc.) |

### Frontend (`frontend/.env`)

| Variable | Required | Default | Description |
|---|---|---|---|
| `VITE_API_URL` | **Yes** | `http://localhost:8080/api/v1` | Base URL pointing to the backend API (`/api/v1`) |

---

## 🚀 Deployment

### Deploy Backend on Vercel
The backend includes `backend/vercel.json` and a serverless entry point at `backend/api/index.go`:
1. Import the `backend` folder as a project in Vercel.
2. In Project Settings ➔ Environment Variables, add:
   - `DATABASE_URL`: Your Neon PostgreSQL connection string.
   - `JWT_SECRET`: A secure random secret string.
   - `CLIENT_URL`: Your deployed frontend URL (optional; Vercel preview/production domains are automatically allowed).
3. Deploy!

### Deploy Frontend on Vercel
The frontend includes `frontend/vercel.json` configured for SPA client-side routing:
1. Import the `frontend` folder as a project in Vercel.
2. In Project Settings ➔ Environment Variables, add:
   - `VITE_API_URL`: Your deployed backend API URL (e.g. `https://your-backend.vercel.app/api/v1`).
3. Set Framework Preset to **Vite**.
4. Deploy!

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
