## For running application
.PHONY: run
run:
	go run cmd/app/main.go

## For Cleaning unnecessary packages
.PHONY: clean
clean:
	go mod tidy

### Migrations
# ----------------------------------------------------------------
.PHONY: migration-create
migration-create:
	migrate create -ext sql -dir internal/migrations/ -seq init

.PHONY: migration-up
migration-up:
	migrate -path internal/migrations/ -database "mysql://root:NewPassword!123@tcp(localhost:3306)/todo_app" -verbose up
		
.PHONY: migration-down
migration-down:
	migrate -path internal/migrations/ -database "mysql://root:NewPassword!123@tcp(localhost:3306)/todo_app" -verbose down

.PHONY: force-version
force-version:
		@if [ -z "$(VERSION)" ]; then \
		echo "VERSION is not set. Usage: make force-version VERSION=<version>"; \
	else \
		migrate -path internal/migrations/ -database "mysql://root:NewPassword!123@tcp(localhost:3306)/todo_app" -verbose force $(VERSION); \
	fi				