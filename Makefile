run:
	go run -ldflags "-X github.com/andrewozarko/mariejean/cmd/cli.Version=1.0.1-beta" ./cmd/mj/main.go version
build:
	go build -o mj -ldflags "-X github.com/andrewozarko/mariejean/cmd/cli.Version=1.0.1-beta" ./cmd/mj/main.go
build-windows:
	GOOS=windows GOARCH=386 go build -o mj.exe -ldflags "-X github.com/andrewozarko/mariejean/cmd/cli.Version=1.0.1-beta" ./cmd/mj/main.go
install:
	go mod tidy
	make build
	sudo mv mj /usr/local/bin/mj
