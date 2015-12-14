package image

import "github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/docker/layer"

// Append appends a new diffID to rootfs
func (r *RootFS) Append(id layer.DiffID) {
	r.DiffIDs = append(r.DiffIDs, id)
}
