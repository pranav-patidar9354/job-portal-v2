# Job Portal V2

Full-stack Job Portal using Go, Gin, GORM, MySQL, JWT, bcrypt and React.

## Backend

```bash
cd backend
go mod tidy
go run ./cmd
```

Backend: http://localhost:8080

## Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend: http://localhost:5173

## MySQL

Create a database:

```sql
CREATE DATABASE job_portal_v2;
```

Copy `backend/.env.example` to `backend/.env` and configure MySQL credentials.

The backend uses GORM AutoMigrate for tables.
