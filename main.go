package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"time"
)

func main() {
	fmt.Println("test1")
	handle, err := pcapgo.NewEthernetHandle("enp0s3")
	if err != nil {
		panic(err)
	}

	var eth layers.Ethernet
	var ip4 layers.IPv4
	var ip6 layers.IPv6
	var tcp layers.TCP

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &ip6, &tcp)
	decoded := []gopacket.LayerType {}
	go func() {
		fmt.Println("goroutine")
		for {
			data, info, _ := handle.ZeroCopyReadPacketData()
			parser.DecodeLayers(data, &decoded)
			fmt.Println(info)
			for _, layerType := range decoded {
				switch layerType {

				case layers.LayerTypeIPv6:
					fmt.Println(ip6.SrcIP, ip6.DstIP)

				case layers.LayerTypeIPv4:
					fmt.Println(ip4.SrcIP, ip4.DstIP)
				}
			}
		}
	} ()


	time.Sleep(10 * time.Second)

}