Docker No volumes Plugin
=
_In order to use this plugin you need to be running at least Docker 1.10 which
has support for authorization plugins._

When a volume in provisioned via the `VOLUME` instruction in a Dockerfile or via
`docker run -v volumename`, host's storage space is used. This could lead to an
unexpected out of space issue which could bring down everything.
There are situations where this is not an accepted behavior. PAAS, for instance,
can't allow their users to run their own images without the risk of filling the
entire storage space on a server. One solution to this is to deny users from running
images with volumes. This way the only storage a user gets can be limited and PAAS
can assign quota to it.

This plugin solves this issue by disallowing starting a container with local volumes defined.
In particular, the plugin will block `docker run` with:

- `--volumes-from`
- images that have `VOLUME`(s) defined
- volumes early provisioned with `docker volume` command

The only thing allowed will be just bind mounts.

Building
-
```sh
$ export GOPATH=~ # optional if you already have this
$ mkdir -p ~/src/github.com/runcom # optional, from now on I'm assuming GOPATH=~
$ cd ~/src/github.com/runcom && git clone https://github.com/runcom/docker-novolume-plugin
$ cd docker-novolume-plugin
$ make
```
Installing
-
```sh
$ sudo make install
$ systemctl enable docker-novolume-plugin
```
Running
-
Specify `--authorization-plugin=docker-novolume-plugin` in the `docker daemon` command line
flags (either in the systemd unit file or `/etc/sysconfig/docker` under `$OPTIONS`
or when manually starting the daemon)
The plugin must be started before `docker` (done automatically via systemd unit file).
If you're not using the systemd unit file:
```sh
$ docker-novolume-plugin &
```
Just restart `docker` and you're good to go!
License
-
MIT
