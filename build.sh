# https://wails.io/docs/reference/cli#platforms

wails build -platform darwin/universal -o ikun-grpc_v0.1_darwin

wails build -platform windows/amd64 -o ikun-grpc_v0.1_windows_amd64
wails build -platform windows/arm64 -o ikun-grpc_v0.1_windows_arm64

wails build -platform linux/amd64 -o ikun-grpc_v0.1_linux_amd64
wails build -platform linux/arm64 -o ikun-grpc_v0.1_linux_amd64

