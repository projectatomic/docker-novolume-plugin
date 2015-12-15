Docker No volumes Plugin
=
Building:
-
```sh
$ export GOPATH=~ # optional if you already have this
$ mkdir -p ~/src/github.com/runcom # optional, from now on I'm assuming GOPATH=~
$ cd ~/src/github.com/runcom && git clone https://github.com/runcom/docker-novolume-plugin
$ cd docker-novolume-plugin
$ make
```
Installing:
-
Either:
```
sudo make install
```
Or:
```sh
$ systemctl enable docker-novolume-plugin
```
Running:
-
Specify --authz-plugin=docker-novolume-plugin docker daemon $OPTIONS (/etc/sysconfig/docker)

The plugin must be started before `docker` (done automatically via systemd unit file).
If you're not using the systemd unit file:
```sh
$ docker-novolume-plugin &
```
Before (re)starting `docker` make sure you add `--authz-plugin=docker-novolume-plugin` to
the `docker daemon` command line flags (either in the systemd unit file or manually).
