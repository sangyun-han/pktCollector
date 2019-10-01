package main

import (
	"fmt"
	"github.com/sangyun-han/pktCollector/engine"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)


func main() {

	config := engine.InterfacesConfig{
		Device:		"enp0s3",
		BpfFilter:	false,
		DefaultOpt:	false,
	}


	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var workerNum  = 10
	var channelBufferSize = 100
	var dataChannel = make(chan []byte, channelBufferSize)



	//afpacketHandle, err := afpacket.NewTPacket(afpacket.OptInterface("enp0s3"))
	afpacketHandle, err := engine.OpenAFPacket("", &config)


	if err != nil {
		log.Fatal(err)
	}
	defer afpacketHandle.TPacket.Close()

	for i := 0; i < workerNum; i += 1 {
		w := engine.NewWorker()
		go w.Decode(dataChannel)

	}

	go func() {
		for {
			//data, _, _ := handle.ZeroCopyReadPacketData()
			//dataChannel <- data
			//fmt.Println(data)
			pkt, _, err := afpacketHandle.TPacket.ZeroCopyReadPacketData()
			if err != nil {
				log.Fatal(err)
			}
			dataChannel <- pkt
		}
	}()


	ticker := time.NewTicker(1 * time.Second)
	//collector, err := engine.NewCollector()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	go func() {
		for t:= range ticker.C {
			fmt.Println("Tick at ", t)
			//collector.Collect()
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
}
