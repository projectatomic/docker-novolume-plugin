.PHONY: all binary man install clean
export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

LIBDIR=${DESTDIR}/lib/systemd/system
BINDIR=${DESTDIR}/usr/lib/docker/
PREFIX ?= ${DESTDIR}/usr
MANINSTALLDIR=${PREFIX}/share/man

all: man binary

binary:
	go build  -o docker-novolume-plugin .

man:
	go-md2man -in man/docker-novolume-plugin.8.md -out docker-novolume-plugin.8

install:
	install -d -m 0755 ${LIBDIR}
	install -m 644 systemd/docker-novolume-plugin.service ${LIBDIR}
	install -d -m 0755 ${LIBDIR}
	install -m 644 systemd/docker-novolume-plugin.socket ${LIBDIR}
	install -d -m 0755 ${BINDIR}
	install -m 755 docker-novolume-plugin ${BINDIR}
	install -m 644 docker-novolume-plugin.8 ${MANINSTALLDIR}/man8/

clean:
	rm -f docker-novolume-plugin
