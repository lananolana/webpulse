run:
	cd backend && \
	go run ./cmd/main.go

load-test:
	k6 run backend/tests/load/load_get_status.js