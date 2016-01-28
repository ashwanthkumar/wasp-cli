APPNAME = wasp

build:
	go build -o ${APPNAME} .

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
