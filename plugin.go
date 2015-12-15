package main

import (
	"regexp"

	docker "github.com/docker/docker/api/client/lib"
	"github.com/runcom/dkauthz"
)

func newPlugin(dockerHost string) (*novolume, error) {
	client, err := docker.NewClient(dockerHost, nil, nil)
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
		if req.RequestBody != nil {
			// TODO(runcom): this means an hostConfig was provided at start
			// which is currently deprecated. Until it's removed, if volumes are
			// found reply with do not allow
			// FIXME(runcom)
			resp(false)
		}
		res := startRegExp.FindStringSubmatch(req.RequestURI)
		if len(res) < 1 {
			return resp(false)
		}

		container, err := p.client.ContainerInspect(res[1])
		if err != nil {
			return resp(err)
		}
		image, _, err := p.client.ImageInspectWithRaw(container.Image, false)
		if err != nil {
			return resp(err)
		}
		if len(image.Config.Volumes) > 0 {
			return resp(newResponse(false, "volumes are not allowed", ""))
		}
		for _, m := range container.Mounts {
			if m.Driver != "" {
				return resp(newResponse(false, "volumes are not allowed", ""))
			}
		}
	}
	return resp(newResponse(true, "", ""))
}

func (p *novolume) AuthZRes(req dkauthz.Request) dkauthz.Response {
	return resp(newResponse(true, "", ""))
}

// TODO(runcom): awful...
func newResponse(allow bool, msg string, err string) dkauthz.Response {
	res := dkauthz.Response{}
	res.Allow = allow
	res.Msg = msg
	res.Err = err
	return res
}

func resp(r interface{}) dkauthz.Response {
	switch t := r.(type) {
	case error:
		return dkauthz.Response{Err: t.Error()}
	case dkauthz.Response:
		return t
	default:
		return dkauthz.Response{Err: "bad value writing response"}
	}
}
