package main

import (
	"bytes"
	"encoding/json"
	"regexp"

	dockerapi "github.com/docker/docker/api"
	docker "github.com/docker/docker/api/client/lib"
	"github.com/runcom/dkauthz"
)

func newPlugin(dockerHost string) (*novolume, error) {
	client, err := docker.NewClient(dockerHost, string(dockerapi.DefaultVersion), nil, nil)
	if err != nil {
		return nil, err
	}
	return &novolume{client: client}, nil
}

var (
	startRegExp = regexp.MustCompile(`/containers/(.*)/start$`)
)

type novolume struct {
	client *docker.Client
}

func (p *novolume) AuthZReq(req dkauthz.Request) dkauthz.Response {
	if req.RequestMethod == "POST" && startRegExp.MatchString(req.RequestURI) {
		// this is deprecated in docker, remove once hostConfig is dropped to
		// being available at start time
		if req.RequestBody != nil {
			type vfrom struct {
				VolumesFrom []string
			}
			vf := &vfrom{}
			if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(vf); err != nil {
				return dkauthz.Response{Err: err.Error()}
			}
			if len(vf.VolumesFrom) > 0 {
				goto noallow
			}
		}
		res := startRegExp.FindStringSubmatch(req.RequestURI)
		if len(res) < 1 {
			return dkauthz.Response{Err: "unable to find container name"}
		}

		container, err := p.client.ContainerInspect(res[1])
		if err != nil {
			return dkauthz.Response{Err: err.Error()}
		}
		image, _, err := p.client.ImageInspectWithRaw(container.Image, false)
		if err != nil {
			return dkauthz.Response{Err: err.Error()}
		}
		if len(image.Config.Volumes) > 0 {
			goto noallow
		}
		for _, m := range container.Mounts {
			if m.Driver != "" {
				goto noallow
			}
		}
		if len(container.HostConfig.VolumesFrom) > 0 {
			goto noallow
		}
	}
	return dkauthz.Response{Allow: true}

noallow:
	return dkauthz.Response{Msg: "volumes are not allowed"}
}

func (p *novolume) AuthZRes(req dkauthz.Request) dkauthz.Response {
	return dkauthz.Response{Allow: true}
}
