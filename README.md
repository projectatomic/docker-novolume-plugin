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
$ cp systemd/docker-novolume-plugin.service /lib/systemd/system/
$ systemctl enable docker-novolume-plugin
```
Running:
-
The plugin must be started before `docker` (done automatically via systemd unit file).
If you're not using the systemd unit file:
```sh
$ docker-novolume-plugin &
$ systemctl restart docker
```
