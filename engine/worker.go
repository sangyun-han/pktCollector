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
	count := 0
	for packetData := range data {
		w.parser.DecodeLayers(packetData, &w.decoded)
		for _, layerType := range w.decoded {
			switch layerType {
			case layers.LayerTypeTCP:
				fmt.Println("TCP", w.tcp)
				count++
				fmt.Println(count)

			//case layers.LayerTypeUDP:
			//	fmt.Println("UDP", w.udp.SrcPort, w.udp.DstPort)
			//	count++
			//	fmt.Println(count)
			//
			//case layers.LayerTypeIPv6:
			//	fmt.Println("IPv6 : ", w.ip6.SrcIP, w.ip6.DstIP)
			//	count++
			//	fmt.Println(count)
			//
			case layers.LayerTypeIPv4:
				fmt.Println("IPv4 : ", w.ip4.NetworkFlow())
				count++
				fmt.Println(count)

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

