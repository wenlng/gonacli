build:
	go build -ldflags="-s -w" gonacli.go
	$(if $(shell command -v upx), upx gonacli)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o gonacli-darwin gonacli.go
	$(if $(shell command -v upx), upx gonacli-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o gonacli.exe gonacli.go
	$(if $(shell command -v upx), upx gonacli.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o gonacli-linux gonacli.go
	$(if $(shell command -v upx), upx gonacli-linux)