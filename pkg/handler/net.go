package trigger

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/net"
)

type NetHandler struct {
	bytesPerSecThreshhold uint64
}

func NewNetHandler(mbpsThreshhold float64) *NetHandler {
	return &NetHandler{
		bytesPerSecThreshhold: uint64(mbpsThreshhold * 1024 * 1024),
	}
}

func (n *NetHandler) Passed() bool {
	initialStats, err := net.IOCounters(true)
	if err != nil {
		log.Error().Err(err).Msg("Error getting initial network stats")
		return false
	}
	time.Sleep(1 * time.Second)

	newStats, err := net.IOCounters(true)
	if err != nil {
		log.Error().Err(err).Msg("Error getting new network stats")
		return false
	}

	result := false

	validInterfaces := []string{}

	// 计算网络速度
	for i, initialStat := range initialStats {
		newStat := newStats[i]
		bytesSentPerSec := newStat.BytesSent - initialStat.BytesSent
		bytesRecvPerSec := newStat.BytesRecv - initialStat.BytesRecv
		if bytesSentPerSec > n.bytesPerSecThreshhold {
			validInterfaces = append(validInterfaces, fmt.Sprintf("%v sent", initialStat.Name))
		}
		if bytesRecvPerSec > n.bytesPerSecThreshhold {
			validInterfaces = append(validInterfaces, fmt.Sprintf("%v recv", initialStat.Name))
		}
	}

	if len(validInterfaces) > 0 {
		log.Info().Strs("interfaces", validInterfaces).Msg("Network activity detected")
		result = true
	}

	return result
}
