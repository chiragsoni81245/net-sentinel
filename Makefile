build:
	@go build -o bin/net-sentinel ./main.go

run: build
	@./bin/net-sentinel start $(ARGS)

test:
	@go test ./... -v --race

# Database migrations
migrate-up: build
	@./bin/net-sentinel migrate up $(ARGS)

migrate-down: build
	@./bin/net-sentinel migrate down $(ARGS)
