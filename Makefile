export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

all:
	go build  -o docker-novolume-plugin .

install:
	cp docker-novolume-plugin /usr/local/bin/docker-novolume-plugin

clean:
	rm docker-novolume-plugin
