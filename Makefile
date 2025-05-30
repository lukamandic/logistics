ifeq ($(OS),Windows_NT)
	COPY_CMD = xcopy frontend\dist backend\static /E
	RM_CMD = if exist backend\static rmdir /S /Q backend\static && if exist frontend\dist rmdir /S /Q frontend\dist
	START_BACKEND_CMD = start cmd /c "go mod tidy && go run cmd/main.go"
else
	COPY_CMD = cp -r frontend/dist/* backend/static/
	RM_CMD = rm -rf backend/static/* frontend/dist/*
	START_BACKEND_CMD = go mod tidy && go run cmd/main.go &
endif

start-frontend:
	@echo "Starting frontend dev server"
	cd frontend && pnpm run dev

start-backend:
	@echo "Starting backend dev server"
	cd backend && cmd /c go mod tidy && go run cmd/main.go

dev: start-backend-async start-frontend

start-backend-async:
	@echo "Starting backend dev server (async) ..."
	cd backend && $(START_BACKEND_CMD)

build-frontend:
	@echo "Building frontend..."
	cd frontend && pnpm run build
	@echo "Copying build to backend static directory..."
	$(COPY_CMD)

serve: build-frontend start-backend

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t logistics-app:latest .

docker-run:
	@echo "Running Docker container..."
	docker run -p 80:8080 --env-file backend/.env logistics-app:latest

docker-build-run: build-frontend docker-build docker-run

# Clean up commands
clean:
	@echo "Cleaning up build artifacts..."
	$(RM_CMD)

.PHONY: start-frontend start-backend dev start-backend-async build-frontend serve docker-build docker-run docker-build-run clean
