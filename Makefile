## Load environment variables
include .env

# ==================================================================================== #
# # HELPERS
# ==================================================================================== #
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# ==================================================================================== #
# # DEVELOPMENT
# ==================================================================================== #
## run: run the application using development config
.PHONY: run
run:
	go run ./cmd/web -addr=":8080" -in-production="false"


# ==================================================================================== #
# # PRODUCTION
# ==================================================================================== #
production_host_ip = '209.38.16.135'
production_non_root_user = 'ljpurcell'

## production/connect: connect to the production server
.PHONY: production/connect 
production/connect:
	ssh ${production_non_root_user}@${production_host_ip}

## production/deploy/app: deploy the app to production using the server specific binary
.PHONY: production/deploy/app
production/deploy/app: confirm
	rsync -P ./bin/linux_amd64/api ${production_non_root_user}@${production_host_ip}


# ==================================================================================== #
# # QUALITY CONTROL
# ==================================================================================== #
## audit: tidy dependencies and format, vet and test all code
.PHONY: audit 
audit:
	@echo 'Tidying and verifying module dependencies...' go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
