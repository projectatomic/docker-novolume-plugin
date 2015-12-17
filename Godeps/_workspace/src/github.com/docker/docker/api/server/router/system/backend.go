package system

import (
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/docker/api/types"
	"github.com/runcom/docker-novolume-plugin/Godeps/_workspace/src/github.com/docker/docker/api/types/filters"
)

// Backend is the methods that need to be implemented to provide
// system specific functionality.
type Backend interface {
	SystemInfo() (*types.Info, error)
	SystemVersion() types.Version
	SubscribeToEvents(since, sinceNano int64, ef filters.Args) ([]*jsonmessage.JSONMessage, chan interface{})
	UnsubscribeFromEvents(chan interface{})
	AuthenticateToRegistry(authConfig *types.AuthConfig) (string, error)
}
