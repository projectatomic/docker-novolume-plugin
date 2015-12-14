package main

import (
	_ "net/http/pprof"

	"github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/auth/htpasswd"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/auth/silly"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/auth/token"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/proxy"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/azure"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/filesystem"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/gcs"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/inmemory"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/middleware/cloudfront"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/oss"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/s3"
	_ "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver/swift"
)

func main() {
	registry.Cmd.Execute()
}
