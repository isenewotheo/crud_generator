SERVICENAME=crud_generator
SERVICEURL=github.com/ManyakRus/$(SERVICENAME)

FILEMAIN=./cmd/$(SERVICENAME)/main.go
FILEAPP=./bin/$(SERVICENAME)

NEW_REPO=$(SERVICENAME)


run:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	#	cd ./bin && \
	./bin/$(SERVICENAME)
mod:
	clear
	go get -u ./internal/...
	go mod tidy -compat=1.18
	go mod vendor
	go fmt ./...
build:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	cd ./scripts && \
	./VersionToFile.py

lint:
	clear
	go fmt ./...
	golangci-lint run ./internal/...
	golangci-lint run ./pkg/...
	gocyclo -over 10 ./internal
	gocyclo -over 10 ./pkg
	gocritic check ./internal/...
	gocritic check ./pkg/...
	staticcheck ./internal/...
	staticcheck ./pkg/...
run.test:
	clear
	go fmt ./...
	go test -coverprofile cover.out ./internal/v0/app/...
	go tool cover -func=cover.out
newrepo:
	sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g' go.mod
	find -name *.go -not -path "*/vendor/*"|xargs sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g'
graph:
	clear
	image_packages ./ docs/packages.graphml
conn:
	clear
	image_connections ./internal docs/connections.graphml $(SERVICENAME)
lines:
	clear
	go_lines_count ./ ./docs/lines_count.txt 2
licenses:
	golicense -out-xlsx=./docs/licenses.xlsx $(FILEAPP)
