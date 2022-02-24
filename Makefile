GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

.PHONY: all
all: http-shell-darwin_amd64 http-shell-darwin_arm64 http-shell-linux_amd64 http-shell-linux_arm64 http-shell-windows_amd64

.PHONY: http-shell-darwin_amd64
http-shell-darwin_amd64:
	GOOS=darwin  GOARCH=amd64 $(MAKE) http-shell

.PHONY: http-shell-darwin_arm64
http-shell-darwin_arm64:
	GOOS=darwin  GOARCH=arm64 $(MAKE) http-shell

.PHONY: http-shell-linux_amd64
http-shell-linux_amd64:
	GOOS=linux   GOARCH=amd64 $(MAKE) http-shell

.PHONY: http-shell-linux_arm64
http-shell-linux_arm64:
	GOOS=linux   GOARCH=arm64 $(MAKE) http-shell

.PHONY: http-shell-windows_amd64
http-shell-windows_amd64:
	GOOS=windows GOARCH=amd64 EXTENSION=.exe $(MAKE) http-shell

.PHONY: http-shell
http-shell:
	CGO_ENABLED=0 go build -trimpath -o ./bin/http-shell_$(GOOS)_$(GOARCH)$(EXTENSION) .