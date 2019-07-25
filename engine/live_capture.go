package engine

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

func Capture() {


	var (
		device       string = "en5"
		snapshotLen int32  = 1024
		promiscuous  bool   = false
		err          error
		timeout      time.Duration = 5 * time.Second
		handle       *pcap.Handle
	)



	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	count := 0
	menu := 2


	if menu == 1 {
		packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSrc.Packets() {
			count++
			fmt.Println(count)
			if packet.NetworkLayer() != nil {
				fmt.Println(packet.NetworkLayer().NetworkFlow())
			}

		}

	} else if menu == 2 {
		var eth layers.Ethernet
		var ip4 layers.IPv4
		var ip6 layers.IPv6
		var tcp layers.TCP


		parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &ip6, &tcp)
		decoded := []gopacket.LayerType {}
		for {
			data, _, _ := handle.ZeroCopyReadPacketData()
			parser.DecodeLayers(data, &decoded)

			for _, layerType := range decoded {
				switch layerType {

				case layers.LayerTypeIPv6:
					fmt.Println(ip6.SrcIP, ip6.DstIP)
					count++
					fmt.Println(count)
				case layers.LayerTypeIPv4:
					fmt.Println(ip4.SrcIP, ip4.DstIP)
					count++
					fmt.Println(count)
				}
			}
		}
	}




}