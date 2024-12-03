run: build
	@./tmp/main

build:
	@go build -o ./tmp/main cmd/app/main.go

air:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build -o ./tmp/main cmd/app/main.go" \
	--build.delay "100" \
	--build.bin "./tmp/main" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go,html" \
	--build.stop_on_error "false"
