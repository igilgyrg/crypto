APP_BIN = build/app

run-local: clean swagger local-env-up
	go build -o $(APP_BIN) github.com/igilgyrg/crypto/cmd
	./${APP_BIN}

local-env-up:
	docker compose -f ./deployment/docker-compose.yml up -d

swagger:
	swag init -g ./cmd/main.go

clean:
	rm -rf build || true
