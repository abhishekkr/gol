package gollb

import (
	"container/ring"
	"strings"
)

type LoadBalancer interface {
	Load(Service map[string][]string)
	LoadWithSeparator(map[string]string, string)
	GetBackend(string) string
}

type RoundRobin struct {
	// {"/abc": Ring["svc://b1", "svc://b2"], "/xyz": Ring[svc://"z1"]}
	ServiceRings map[string]*ring.Ring
}

func (lb *RoundRobin) Load(Services map[string][]string) {
	lb.ServiceRings = make(map[string]*ring.Ring, len(Services))
	for svc, backends := range Services {
		lb.ServiceRings[svc] = newBackendRing(backends)
	}
}

func (lb *RoundRobin) LoadWithSeparator(Services map[string]string, separator string) {
	lb.ServiceRings = make(map[string]*ring.Ring, len(Services))
	for svc, backends := range Services {
		lb.ServiceRings[svc] = newBackendRing(strings.Split(backends,
			separator))
	}
}

func (lb *RoundRobin) GetBackend(svc string) string {
	lb.ServiceRings[svc] = lb.ServiceRings[svc].Next()
	return lb.ServiceRings[svc].Value.(string)
}

func newBackendRing(backends []string) *ring.Ring {
	list := ring.New(len(backends))
	for _, backend := range backends {
		list.Value = backend
		list = list.Next()
	}
	return list
}
