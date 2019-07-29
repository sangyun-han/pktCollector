package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/sangyun-han/pktCollector/engine"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
//var memprofile = flag.String("memprofile", "", "write mem profile to `file`")


func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var workerNum  = 1
	var channelBufferSize = 100
	var dataChannel = make(chan []byte, channelBufferSize)

	handle, err := pcap.OpenLive("en5", 65536, false, 10 * time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()


	menu := 1

	for i := 0; i < workerNum; i += 1 {
		w := engine.NewWorker()

		if menu == 1 {

			go w.Decode(dataChannel)
		} else {

		}

	}

	//go func() {
	//	for {
	//		data, _, _ := handle.ZeroCopyReadPacketData()
	//		dataChannel <- data
	//		//fmt.Println(data)
	//	}
	//}()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	wo := engine.NewWorker()
	go wo.Decode2(packetSource)

	time.Sleep(1 * time.Second)
}
