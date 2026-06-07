# Advanced Terraform Setup / MuchToDo

Full-stack task management application with a Go backend API, a React TypeScript frontend, Docker-based local services, and GitHub Actions workflows for AWS deployment.

The application is organized as a frontend/backend project. The backend provides authentication, user profile management, todo/task CRUD, Swagger documentation, MongoDB persistence, and optional Redis caching. The frontend is a Vite React app that talks to the API with cookie-based authentication.

## Project Structure

```text
.
|-- backend/
|   |-- README.md
|   `-- MuchToDo/
|       |-- cmd/api/              # API entry point
|       |-- docs/                 # Generated Swagger/OpenAPI files
|       |-- internal/
|       |   |-- auth/             # Authentication tests and helpers
|       |   |-- cache/            # Redis cache layer
|       |   |-- config/           # Environment configuration
|       |   |-- database/         # MongoDB connection logic
|       |   |-- handlers/         # HTTP handlers
|       |   |-- logger/           # Structured logging
|       |   |-- middleware/       # CORS, auth, and request logging
|       |   |-- models/           # User and todo models
|       |   |-- routes/           # Gin route registration
|       |   `-- utils/            # Shared utilities
|       |-- .env.example          # Backend environment template
|       |-- docker-compose.yaml   # Backend, MongoDB, Redis, and admin UIs
|       |-- Dockerfile
|       |-- go.mod
|       `-- Makefile
|-- frontend/
|   |-- src/
|   |   |-- components/           # React and UI components
|   |   |-- context/              # Auth context
|   |   |-- hooks/                # Custom React hooks
|   |   |-- lib/                  # API client and utilities
|   |   |-- routes/               # TanStack Router pages
|   |   `-- types/                # TypeScript types
|   |-- package.json
|   |-- vite.config.ts
|   `-- tsconfig.json
|-- scripts/                     # Deployment helper scripts
`-- .github/workflows/           # Backend and frontend CI/CD workflows
```

## Features

- Go REST API using Gin
- JWT authentication with httpOnly cookie support
- User registration, login, logout, profile update, password change, and account deletion
- Authenticated task CRUD endpoints
- MongoDB persistence
- Optional Redis caching
- Swagger UI at `/swagger/index.html`
- React 19, TypeScript, Vite, TanStack Router, TanStack Query, Tailwind CSS, and Radix UI
- Docker Compose stack for backend, MongoDB, Redis, Mongo Express, and Redis Commander
- GitHub Actions workflows for build, test, image publishing, S3 deployment, CloudFront invalidation, and ASG refresh

## Prerequisites

- Go 1.25.1 or compatible with the version in `backend/MuchToDo/go.mod`
- Node.js 20 or compatible with the frontend workflow
- npm
- Docker and Docker Compose
- Make, if you want to use the backend Makefile targets
- Swag CLI for Swagger generation:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Backend Setup

```bash
cd backend/MuchToDo
cp .env.example .env
go mod download
make generate-docs
make run
```

The API runs on `http://localhost:8080` by default.

Important backend environment variables are defined in `backend/MuchToDo/.env.example`:

- `PORT`
- `MONGO_URI`
- `DB_NAME`
- `JWT_SECRET_KEY`
- `JWT_EXPIRATION_HOURS`
- `ENABLE_CACHE`
- `REDIS_ADDR`
- `REDIS_PASSWORD`
- `ALLOWED_ORIGINS`
- `COOKIE_DOMAINS`
- `SECURE_COOKIE`

For local development, set a strong `JWT_SECRET_KEY` and make sure `MONGO_URI` points to your MongoDB instance.

## Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

The frontend runs on `http://localhost:5173` by default.

The frontend API base URL defaults to `http://localhost:8080`. Override it with:

```bash
VITE_API_BASE_URL=http://localhost:8080
```

## Docker Development

Start the backend stack from the backend directory:

```bash
cd backend/MuchToDo
make dc-up
```

This starts:

- Backend API: `http://localhost:8080`
- MongoDB: `localhost:27017`
- Mongo Express: `http://localhost:8081`
- Redis: `localhost:6379`
- Redis Commander: `http://localhost:8082`

Stop the stack with:

```bash
make dc-down
```

## API Endpoints

### Public

- `GET /health` - Health check
- `GET /swagger/index.html` - Swagger UI
- `POST /auth/register` - Register user
- `POST /auth/login` - Login user
- `POST /auth/logout` - Logout user
- `GET /auth/username-check/:username` - Check username availability

### Protected

- `GET /users/me` - Get current user
- `PUT /users/me` - Update current user
- `PUT /users/me/password` - Change current user's password
- `DELETE /users/me` - Delete current user
- `GET /tasks` - List tasks
- `POST /tasks` - Create task
- `GET /tasks/:id` - Get task
- `PUT /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task

Protected endpoints require the authentication cookie set by login, or a valid bearer token where supported by the backend middleware.

## Common Commands

### Backend

```bash
cd backend/MuchToDo
make run              # Generate Swagger docs and run the API
make build            # Build the backend binary into ./bin
make unit-test        # Run Go unit tests
make integration-test # Run integration tests
make tidy             # Tidy Go modules
make dc-up            # Start Docker Compose stack
make dc-down          # Stop Docker Compose stack
```

### Frontend

```bash
cd frontend
npm run dev       # Start Vite dev server
npm run build     # Type-check and build production assets
npm run lint      # Run ESLint
npm run preview   # Preview production build locally
```

There is no frontend test script in `frontend/package.json` yet.

## CI/CD

The repository includes two GitHub Actions workflows:

- `.github/workflows/backend-ci-cd.yml`
- `.github/workflows/frontend-ci-cd.yml`

The backend workflow runs on changes to `backend/MuchToDo/**`, builds the Go API, builds and scans a Docker image, pushes it to ECR, refreshes an Auto Scaling Group, and smoke-tests `/health` through an ALB.

The frontend workflow runs on changes to `frontend/**`, installs dependencies, builds the Vite app with `VITE_API_BASE_URL`, syncs `dist/` to S3, and invalidates CloudFront.

Required repository secrets include:

- `AWS_ROLE_TO_ASSUME`
- `AWS_REGION`
- `ALB_DNS_NAME`
- `ECR_REPOSITORY_URL`
- `ASG_NAME`
- `S3_BUCKET_NAME`
- `CLOUDFRONT_DISTRIBUTION_ID`

## Deployment Scripts

The `scripts/` directory contains helper scripts. Currently, `deploy-backend.sh` builds and pushes the backend image using the `ECR_REPOSITORY` environment variable:

```bash
ECR_REPOSITORY=<account-id>.dkr.ecr.<region>.amazonaws.com/<repo> ./scripts/deploy-backend.sh
```

The other script files are present as placeholders and can be extended as the deployment process evolves.

## License

No license file is included yet. Add one before publishing or distributing the project.

