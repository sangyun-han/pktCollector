package engine

// Under IPv4
type Statistics struct {
	Protocol uint8 // IPProtocol (ICMP, IGMP, TCP, UDP)
	IPPacketLength uint16
	SrcIP []byte // net.IP
	DstIP []byte // net.IP
	TTL uint8
	IPChecksum uint16

	SrcPort uint16 // layers.TCPPort
	DstPort uint16 // layers.TCPPort
	AckNum uint32
	SeqNum uint32
	CWR bool
	ECE bool
	SYN bool
	FIN bool
	URG bool
	PSH bool
	RST bool
	ACK bool
	TCPChecksum uint16
	WindowSize uint16
	DataOffset uint8
}