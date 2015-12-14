export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

all:
	go build  -o docker-novolume-plugin .

install:
	cp runc /usr/local/bin/runc

clean:
	rm docker-novolume-plugin
