build: clean
	go build -o build/go-web -race ./cmd/main.go

.PHONEY: run
run:
	go run cmd/main.go

.PHONEY: fmt
fmt:
	gofmt -s -w .

.PHONEY: vet
vet:
	go vet ./cmd
	go vet ./internal/**
	go vet ./web

.PHONEY: clean
clean:
	rm -rf build

.PHONEY: img
img:
	docker build -t go-web .

.PHONEY: rm-img
rm-img:
	docker rmi -f go-web

.PHONEY: air
air:
	air --build.cmd "go build -o build/go-web ./cmd/main.go" --build.bin "./build/go-web" --build.exclude_dir "static,build"