package core

import (
	"time"

	"github.com/pion/rtp/v2"
)

type data struct {
	rtp   *rtp.Packet
	nalus [][]byte
	pts   time.Duration
}
