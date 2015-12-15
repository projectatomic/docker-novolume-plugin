Building:
```
export GOPATH=~ # optional if you already have this
mkdir -p ~/src/github.com/runcom # optional, from now on I'm assuming GOPATH=~
cd ~/src/github.com/runcom && git clone https://github.com/runcom/docker-novolume-plugin
cd docker-novolume-plugin
make
sudo make install # optional
```
