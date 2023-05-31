APP_NAME=dnat
APP_VERSION ?= dev
INSTALL_PATH=/usr/local/bin
CMD_PATH=github.com/go-bai/go-dnat/cmd

all: build

build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X '${CMD_PATH}.Version=${APP_VERSION}'" -o ${APP_NAME}-linux-amd64

build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X '${CMD_PATH}.Version=${APP_VERSION}'" -o ${APP_NAME}-linux-arm64

build:
	go build -ldflags="-s -w -X '${CMD_PATH}.Version=${APP_VERSION}'" -o ${APP_NAME}

install: build
	chmod +x ${APP_NAME}
	mv ${APP_NAME} ${INSTALL_PATH}

release-linux: build-linux-amd64 build-linux-arm64

clean:
	rm -f ${APP_NAME}*