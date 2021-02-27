// Package stdout is a packet handler for stdout.
package stdout

import (
	"github.com/pd0mz/go-aprs"
	log "github.com/sirupsen/logrus"
)

// Stdout helps fulfill the Handler interface contract.
type Stdout struct{}

const id = "4967ade5-7a97-416f-86bf-6e2ae8a5e581"
const name = "Stdout"

// Id defines the Id of this handler.
func (s Stdout) Id() string {
	return id
}

// Name defines the name of this handler.
func (s Stdout) Name() string {
	return name
}

// Enabled specifies that this handler is always enabled.
func (s Stdout) Enabled() bool {
	return true
}

// Start does not do anything in this implementation.
func (s Stdout) Start() {}

// Handle prints packet information to stdout via the log package.
func (s Stdout) Handle(p *aprs.Packet) {
	log.Printf("%s -> %s (%f, %f) %s", p.Src.Call, p.Dst.Call,
		p.Position.Latitude, p.Position.Longitude, p.Comment)
}

// Stop does not do anything in this implementation.
func (s Stdout) Stop() {}
