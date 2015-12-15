export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

LIBDIR=${DESTDIR}/lib/systemd/system
BINDIR=${DESTDIR}/usr/lib/docker/

all:
	go build  -o docker-novolume-plugin .

install:
	install -d -m 0755 ${LIBDIR}
	install -m 644 systemd/docker-novolume-plugin.service ${LIBDIR}
	install -d -m 0755 ${BINDIR}
	install -m 755 docker-novolume-plugin ${BINDIR}

clean:
	rm docker-novolume-plugin *~
