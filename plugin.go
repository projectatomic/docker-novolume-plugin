package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	dockerapi "github.com/docker/docker/api"
	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/go-plugins-helpers/authorization"
)

func newPlugin(dockerHost, certPath string, tlsVerify bool) (*novolume, error) {
	var transport *http.Transport
	if certPath != "" {
		tlsc := &tls.Config{}

		cert, err := tls.LoadX509KeyPair(filepath.Join(certPath, "cert.pem"), filepath.Join(certPath, "key.pem"))
		if err != nil {
			return nil, fmt.Errorf("Error loading x509 key pair: %s", err)
		}

		tlsc.Certificates = append(tlsc.Certificates, cert)
		tlsc.InsecureSkipVerify = !tlsVerify
		transport = &http.Transport{
			TLSClientConfig: tlsc,
		}
	}

	client, err := dockerclient.NewClient(dockerHost, dockerapi.DefaultVersion.String(), transport, nil)
	if err != nil {
		return nil, err
	}
	return &novolume{client: client}, nil
}

var (
	startRegExp = regexp.MustCompile(`/containers/(.*)/start$`)
)

type novolume struct {
	client *dockerclient.Client
}

func (p *novolume) AuthZReq(req authorization.Request) authorization.Response {
	if req.RequestMethod == "POST" && startRegExp.MatchString(req.RequestURI) {
		// this is deprecated in docker, remove once hostConfig is dropped to
		// being available at start time
		if req.RequestBody != nil {
			type vfrom struct {
				VolumesFrom []string
			}
			vf := &vfrom{}
			if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(vf); err != nil {
				return authorization.Response{Err: err.Error()}
			}
			if len(vf.VolumesFrom) > 0 {
				goto noallow
			}
		}
		res := startRegExp.FindStringSubmatch(req.RequestURI)
		if len(res) < 1 {
			return authorization.Response{Err: "unable to find container name"}
		}
		container, err := p.client.ContainerInspect(res[1])
		if err != nil {
			return authorization.Response{Err: err.Error()}
		}
		bindDests := []string{}
		for _, m := range container.Mounts {
			if m.Driver != "" {
				goto noallow
			}
			bindDests = append(bindDests, m.Destination)
		}
		image, _, err := p.client.ImageInspectWithRaw(container.Image, false)
		if err != nil {
			return authorization.Response{Err: err.Error()}
		}
		if len(bindDests) == 0 && len(image.Config.Volumes) > 0 {
			goto noallow
		}
		if len(image.Config.Volumes) > 0 {
			for _, bd := range bindDests {
				if _, ok := image.Config.Volumes[bd]; !ok {
					goto noallow
				}
			}
		}
		if len(container.HostConfig.VolumesFrom) > 0 {
			goto noallow
		}
		// TODO(runcom): FROM scratch ?!?!
	}
	return authorization.Response{Allow: true}

noallow:
	return authorization.Response{Msg: "volumes are not allowed"}
}

func (p *novolume) AuthZRes(req authorization.Request) authorization.Response {
	return authorization.Response{Allow: true}
}
