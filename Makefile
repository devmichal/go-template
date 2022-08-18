SWAGGER_OUT_DIR ?= api/swagger

shell:
	docker-compose exec notification /bin/sh

hot:
	nodemon --exec go run -mod=vendor cmd/main.go server $* --signal SIGTERM

local-setup:
	cat .env.dist > .env

vendor:
	GOPRIVATE="github.com/chemiq/notification" go mod vendor

clean:
	rm -rf bin/

run-%:
	go run -mod=vendor cmd/main.go $*

# swagger
swagger:
	swag init \
		--parseDependency \
		--output api/swagger \
		--dir cmd/server/ \
	&& rm api/swagger/swagger.yaml

test:
	grc go test -v -failfast -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
