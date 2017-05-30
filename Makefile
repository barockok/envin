.SILENT :
.PHONY : envin clean fmt

TAG:=`git describe --abbrev=0 --tags`
LDFLAGS:=-X main.buildVersion=$(TAG)

all: envin

deps:
	go get github.com/robfig/glock
	glock sync -n < GLOCKFILE

dockerize:
	echo "Building envin"
	go install -ldflags "$(LDFLAGS)"

dist-clean:
	rm -rf dist
	rm -f envin-alpine-linux-*.tar.gz
	rm -f envin-linux-*.tar.gz
dist: deps dist-clean
	mkdir -p dist/alpine-linux/amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -a -tags netgo -installsuffix netgo -o dist/alpine-linux/amd64/envin
	mkdir -p dist/linux/amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux/amd64/envin
	mkdir -p dist/linux/armel && GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "$(LDFLAGS)" -o dist/linux/armel/envin
	mkdir -p dist/linux/armhf && GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "$(LDFLAGS)" -o dist/linux/armhf/envin

release: dist
	tar -cvzf envin-alpine-linux-amd64-$(TAG).tar.gz -C dist/alpine-linux/amd64 envin
	tar -cvzf envin-linux-amd64-$(TAG).tar.gz -C dist/linux/amd64 envin
	tar -cvzf envin-linux-armel-$(TAG).tar.gz -C dist/linux/armel envin
	tar -cvzf envin-linux-armhf-$(TAG).tar.gz -C dist/linux/armhf envin
