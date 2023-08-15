ROOT := $(CURDIR)

build:
	@cd src && go build -o $(ROOT)/bin/rss-aggregator

run:
	@./bin/rss-aggregator

dev: build run

dev-http-only: build
	@export HTTP_ONLY=true && $(MAKE) run

prod: build
	@export PRODUCTION="true" && export GIN_MODE=release && $(MAKE) run

test:
	@go test -v ./...

generate-TLS:
	@openssl genpkey -algorithm RSA -out server.key
	@openssl req -new -key server.key -out server.csr
	@openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

delete-TLS:
	@rm server.crt
	@rm server.csr
	@rm server.key