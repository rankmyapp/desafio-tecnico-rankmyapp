run: ; go run ./cmd/api
worker: ; go run ./cmd/worker
test: ; go test ./... -v
docker-up: ; docker compose up -d --build
docker-down: ; docker compose down -v