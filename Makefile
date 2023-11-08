build:
	@go build -o bin/goco2

run: build
	@chmod +x bin/goco2
	./bin/goco2

test:
	go test -v ./...

docs:
	@docker run -p 8080:8080 -d --rm --name goco2-openapi-v3 -e URL=openapi.yaml -v ${PWD}/openapi.yaml:/usr/share/nginx/html/openapi.yaml swaggerapi/swagger-ui
