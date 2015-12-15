package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/runcom/dkauthz"
)

const (
	defaultDockerHost   = "unix:///var/run/docker.sock"
	defaultPluginSocket = "/var/run/docker/plugins/docker-novolume-plugin.sock"
)

var (
	flDockerHost   = flag.String("host", defaultDockerHost, "Docker host the plugin connects to when inspecting")
	flPluginSocket = flag.String("socket", defaultPluginSocket, "Plugin's socket path")
	// TODO(runcom): add tls option to connect to docker?
	// TODO(runcom): add plugin tls option (need to learn more...)
)

func main() {
	flag.Parse()

	novolume, err := newPlugin(*flDockerHost)
	if err != nil {
		logrus.Fatal(err)
	}

	// TODO(runcom): parametrize this when the bin starts
	h := dkauthz.NewHandler(novolume)
	if err := h.ServeUnix("root", *flPluginSocket); err != nil {
		logrus.Fatal(err)
	}
}
