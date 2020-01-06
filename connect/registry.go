package connect

import (
	"github.com/micro/go-micro/registry"
	consulRegistry "github.com/micro/go-plugins/registry/consul"
	"os"
)

func NewConsulRegistry() registry.Registry {
	return consulRegistry.NewRegistry(
		registry.Addrs(os.Getenv("CONSUL_ADDR")),
	)
}