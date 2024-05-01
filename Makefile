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
# # BUILD
# ==================================================================================== #
## build/app: build the cmd/web application
.PHONY: build/app 
build/app:
	@echo 'Building cmd/web...'
	go build -ldflags='-s' -o=./bin/web ./cmd/web
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/web ./cmd/web

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
	ssh -i ~/.ssh/id_rsa_${production_non_root_user} ${production_non_root_user}@${production_host_ip}

## production/deploy/app: deploy the app to production using the server specific binary
.PHONY: production/deploy/app
production/deploy/app: confirm
	@echo 'Minifying CSS...'
	npx tailwindcss -i ./ui/static/input.css -o ./ui/static/style.css --minify
	@echo 'Minifying JS...'
	terser ./ui/static/input.js -o ./ui/static/script.js --compress --mangle
	rsync -P ./bin/linux_amd64/web ${production_non_root_user}@${production_host_ip}:~
	rsync -P ./ui -r --delete --exclude "./ui/static/input.*" ${production_non_root_user}@${production_host_ip}:~
	rsync -P ./remote/production/web.service ${production_non_root_user}@${production_host_ip}:~
	ssh -t ${production_non_root_user}@${production_host_ip} '\
	sudo mv ~/web.service /etc/systemd/system/ \
		&& sudo setcap 'cap_net_bind_service=+ep' ~/web \
		&& sudo systemctl daemon-reload \
		&& sudo systemctl enable web \
		&& sudo systemctl restart web \
		'


# ==================================================================================== #
# # QUALITY CONTROL
# ==================================================================================== #
## vendor: tidy and vendor dependencies
.PHONY: vendor 
vendor:
	@echo 'Tidying and verifying module dependencies...' 
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit 
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
