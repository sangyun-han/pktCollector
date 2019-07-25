package engine

type Worker struct {

}

//var eth layers.Ethernet
//var ip4 layers.IPv4
//var ip6 layers.IPv6
//var tcp layers.TCP

//parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &ip6, &tcp)

func (worker Worker) Decode(data chan []byte) {

}