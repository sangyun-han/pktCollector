package engine

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/afpacket"
	"github.com/google/gopacket/layers"
	"os"
	"sync/atomic"
)

type Sniffer struct {
	config InterfacesConfig
	state atomic.Value
	filter string
}

type snifferHandle interface {
	gopacket.PacketDataSource
	LinkType() layers.LinkType
	Close()
}

type afpacketHandle struct {
	TPacket *afpacket.TPacket
}

func OpenAFPacket(filter string, cfg *InterfacesConfig) (*afpacketHandle, error) {
	var frameSize int
	var blockSize int
	var numBlocks int
	var err error

	// No need in now
	if cfg.DefaultOpt {
		frameSize, blockSize, numBlocks, err = afpacketComputeSize(cfg.BufferSizeMb, cfg.SnapLen, os.Getpagesize())
		if err != nil {
			return nil, err
		}
	} else {
		frameSize = afpacket.DefaultFrameSize
		blockSize = afpacket.DefaultBlockSize
		numBlocks = afpacket.DefaultNumBlocks
	}

	handle := &afpacketHandle{}

	//timeout := 1000 * time.Millisecond
	if cfg.Device == "any" {
		handle.TPacket, err = afpacket.NewTPacket(
			afpacket.OptFrameSize(frameSize),
			afpacket.OptBlockSize(blockSize),
			afpacket.OptNumBlocks(numBlocks))
	} else {
		handle.TPacket, err = afpacket.NewTPacket(
			afpacket.OptInterface(cfg.Device),
			afpacket.OptFrameSize(frameSize),
			afpacket.OptBlockSize(blockSize),
			afpacket.OptNumBlocks(numBlocks))
	}

	if err != nil {
		return nil, err
	}

	return handle, err
}

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