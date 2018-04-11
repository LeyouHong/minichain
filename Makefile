all: minimd

minimd:
		@echo "Building minimd to cmd/minimd/minimd"
			@go build -o cmd/minimd/minimd cmd/minimd/main.go
