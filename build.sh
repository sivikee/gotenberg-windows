go build -o gotenberg.exe -ldflags "-X 'github.com/gotenberg/gotenberg/v8/cmd.Version=$GOTENBERG_VERSION'" cmd/gotenberg/main.go