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
$ mkdir -p ~/src/github.com/projectatomic # optional, from now on I'm assuming GOPATH=~
$ cd ~/src/github.com/projectatomic && git clone https://github.com/projectatomic/docker-novolume-plugin
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
flags (either in the systemd unit file or in `/etc/sysconfig/docker` under `$OPTIONS`
or when manually starting the daemon).
The plugin must be started before `docker` (done automatically via systemd unit file).
If you're not using the systemd unit file:
```sh
$ docker-novolume-plugin &
```
Just restart `docker` and you're good to go!
Systemd socket activation
-
The plugin can be socket activated by systemd. You just have to basically use the file provided
under `systemd/` (or installing via `make install`). This ensures the plugin gets activated
if for some reasons it's down.
How to test
-
```bash
$ sudo dnf install docker-novolume-plugin
$ sudo systemctl start docker-novolume-plugin
# edit /etc/sysconfig/docker and append --authorization-plugin=docker-novolume-plugin to OPTIONS
$ sudo systemctl restart docker
$ docker run -v /:/test fedora sh  # works
$ docker run -v /test fedora sh # blocked
$ docker volume create --name test
$ docker run -v test:/test fedora sh # blocked
$ docker build -t testimage - <<EOF
FROM fedora
VOLUME foo
EOF
$ docker run testimage sh # blocked
```
Future
-
Docker 1.11 will come with an Authentication infrastructure. Authorization plugins like
this one can leverage Authentication receiving the `username|group` of the user actually
doing the action in order to take more fine grained decisions.
We basically want to allow a particular user, say `dwalsh`, or group to run containers with
volumes while blocking everyone else. We'll bring this behavior introducing
a configuration file under `/etc/docker/plugins/auth/docker-novolume-plugin.conf` with
the following syntax (for the example above):
```toml
[docker-novolume-plugin]
  allow-user = ["dwalsh"]
  allow-group = []
```
License
-
MIT
