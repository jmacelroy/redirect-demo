
.PHONY: build
build:
	go build -o load-data cmd/loot-data/main.go

.PHONY: start
start:
	./load-data

.PHONY: push
push:
	okteto build -t okteto.dev/loot-data -f Dockerfile.loot-data
