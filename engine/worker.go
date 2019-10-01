package engine

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Worker struct {
	parser		*gopacket.DecodingLayerParser
	data		[]byte
	decoded		[]gopacket.LayerType

	eth layers.Ethernet
	ip4 layers.IPv4
	ip6 layers.IPv6
	tcp layers.TCP
	udp layers.UDP
}


func NewWorker() *Worker {
	var worker = Worker{}

	worker.parser = gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &worker.eth, &worker.ip4, &worker.ip6, &worker.tcp, &worker.udp)
	worker.decoded = []gopacket.LayerType{}

	return &worker
}

func (w *Worker) Decode(data chan []byte) {
	// sample
	var stats Statistics
	for packetData := range data {
		w.parser.DecodeLayers(packetData, &w.decoded)
		for _, layerType := range w.decoded {
			switch layerType {
			case layers.LayerTypeTCP:
				stats = Statistics{
					Protocol:       6,
					IPPacketLength: w.ip4.Length,
					SrcIP:          w.ip4.SrcIP,
					DstIP:          w.ip4.DstIP,
					TTL:            w.ip4.TTL,
					IPChecksum:     w.ip4.Checksum,
					SrcPort:        uint16(w.tcp.SrcPort),
					DstPort:        uint16(w.tcp.DstPort),
					AckNum:         w.tcp.Ack,
					SeqNum:         w.tcp.Seq,
					CWR:            w.tcp.CWR,
					ECE:            w.tcp.ECE,
					SYN:            w.tcp.SYN,
					FIN:            w.tcp.FIN,
					URG:            w.tcp.URG,
					PSH:            w.tcp.PSH,
					RST:            w.tcp.RST,
					ACK:            w.tcp.ACK,
					TCPChecksum:    w.tcp.Checksum,
					WindowSize:     w.tcp.Window,
					DataOffset:     w.tcp.DataOffset,
				}
				fmt.Println(stats)

			case layers.LayerTypeUDP:
				stats = Statistics{
					Protocol:       17,
					IPPacketLength: w.ip4.Length,
					SrcIP:          w.ip4.SrcIP,
					DstIP:          w.ip4.DstIP,
					TTL:            w.ip4.TTL,
					IPChecksum:     w.ip4.Checksum,
					SrcPort:        uint16(w.tcp.SrcPort),
					DstPort:        uint16(w.tcp.DstPort),
					AckNum:         w.tcp.Ack,
					SeqNum:         w.tcp.Seq,
					CWR:            w.tcp.CWR,
					ECE:            w.tcp.ECE,
					SYN:            w.tcp.SYN,
					FIN:            w.tcp.FIN,
					URG:            w.tcp.URG,
					PSH:            w.tcp.PSH,
					RST:            w.tcp.RST,
					ACK:            w.tcp.ACK,
					TCPChecksum:    w.tcp.Checksum,
					WindowSize:     w.tcp.Window,
					DataOffset:     w.tcp.DataOffset,
				}
				fmt.Println(stats)

			case layers.LayerTypeIPv4:
				stats = Statistics{
					Protocol:       4,
					IPPacketLength: w.ip4.Length,
					SrcIP:          w.ip4.SrcIP,
					DstIP:          w.ip4.DstIP,
					TTL:            w.ip4.TTL,
					IPChecksum:     w.ip4.Checksum,
					SrcPort:        uint16(w.tcp.SrcPort),
					DstPort:        uint16(w.tcp.DstPort),
					AckNum:         w.tcp.Ack,
					SeqNum:         w.tcp.Seq,
					CWR:            w.tcp.CWR,
					ECE:            w.tcp.ECE,
					SYN:            w.tcp.SYN,
					FIN:            w.tcp.FIN,
					URG:            w.tcp.URG,
					PSH:            w.tcp.PSH,
					RST:            w.tcp.RST,
					ACK:            w.tcp.ACK,
					TCPChecksum:    w.tcp.Checksum,
					WindowSize:     w.tcp.Window,
					DataOffset:     w.tcp.DataOffset,
				}
				fmt.Println(stats)

			case layers.LayerTypeIPv6:
				stats = Statistics{
					Protocol:       41,
					IPPacketLength: w.ip4.Length,
					SrcIP:          w.ip4.SrcIP,
					DstIP:          w.ip4.DstIP,
					TTL:            w.ip4.TTL,
					IPChecksum:     w.ip4.Checksum,
					SrcPort:        uint16(w.tcp.SrcPort),
					DstPort:        uint16(w.tcp.DstPort),
					AckNum:         w.tcp.Ack,
					SeqNum:         w.tcp.Seq,
					CWR:            w.tcp.CWR,
					ECE:            w.tcp.ECE,
					SYN:            w.tcp.SYN,
					FIN:            w.tcp.FIN,
					URG:            w.tcp.URG,
					PSH:            w.tcp.PSH,
					RST:            w.tcp.RST,
					ACK:            w.tcp.ACK,
					TCPChecksum:    w.tcp.Checksum,
					WindowSize:     w.tcp.Window,
					DataOffset:     w.tcp.DataOffset,
				}
				fmt.Println(stats)

			}
		}
	}
}


func (w *Worker) Decode2(packetSource *gopacket.PacketSource) {
	count := 0
	for packet := range packetSource.Packets() {
		count++
		fmt.Println(count)
		if packet.NetworkLayer() != nil {
			fmt.Println(packet.TransportLayer(), "\n")
		}
	}
}

