package engine

import (
	"fmt"
	"github.com/elastic/beats/packetbeat/config"
)

func openAFPacket(filter string, cfg *config.InterfacesConfig) {}

func afpacketComputeSize(targetSizeMb int, snaplen int, pageSize int) (frameSize int, blockSize int, numBlocks int, err error) {
	if snaplen < pageSize {
		frameSize = pageSize / (pageSize / snaplen)
	} else {
		frameSize = (snaplen / pageSize + 1) * pageSize
	}

	blockSize = frameSize * 128
	numBlocks = (targetSizeMb * 1024 * 1024) / blockSize

	if numBlocks == 0 {
		return 0, 0, 0, fmt.Errorf("Buffer size is too small")
	}

	return frameSize, blockSize, numBlocks, nil
}