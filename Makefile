APPNAME = wasp

build:
	go build -o ${APPNAME} .

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s" -v -o ${APPNAME}-linux-amd64 .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s" -v -o ${APPNAME}-darwin-amd64 .

build-all: build-mac build-linux

all: setup
	build
	install

setup:  
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
	go get github.com/parnurzeal/gorequest

install: build
	sudo install -d /usr/local/bin
	sudo install -c ${APPNAME} /usr/local/bin/${APPNAME}

uninstall:
	sudo rm /usr/local/bin/${APPNAME}
