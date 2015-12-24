package main

import (
	"net"
	"os"
	"time"
	"fmt"
	"os/signal"
	"./system"
	"github.com/golang/protobuf/proto"
	"log"
	"encoding/binary"
	"./protocol"
)

func main() {
	ch := make(chan []byte)

	go sampleData(ch)
	go sendData(ch)

	handleCtrlC()
}

func sampleData(metricsChan chan []byte) {
	tick := time.Tick(2000 * time.Millisecond)
	for now := range tick {
		fmt.Printf("Start sample data at %v\n", now)

		metricsMessage := new(protocol.MetricsMessage)
		for _, pid := range system.GetAllPids() {
			if pid > 0 {
				metric := &protocol.Metrics{
					Pid: proto.Uint32(uint32(pid)),
					Cpu:  proto.Float32(0.32),
				}
//				log.Println("PID: ", uint32(pid))

//				log.Printf("Pid=%d, CPU=%f\n", *metric.Pid, *metric.Cpu)

				metricsMessage.Metricsmessage = append(metricsMessage.Metricsmessage, metric);
			}
		}

//		log.Printf("Pid=%d, CPU=%f\n", metricsMessage.GetMetricsmessage()[0].Pid, metricsMessage.GetMetricsmessage()[0].Cpu)
		data, err := proto.Marshal(metricsMessage)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		metricsChan <- data
	}
}

func sendData(metricsChan chan []byte) {
	conn, err := net.Dial("tcp", "localhost:3333")
	if err != nil {
		log.Fatal("Failed opening socket")
	}

	for item := range metricsChan {
		len := len(item)

		msgLen := 4 + len
		msg := make([]byte, msgLen, msgLen)
		copy(msg[4:], item)
		binary.BigEndian.PutUint32(msg, uint32(len))

		fmt.Printf("Sending metrics, length (without header)=%d\n", len)
		_, err = conn.Write(msg)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

	}

	conn.Close()
}

func handleCtrlC() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		println("Terminate program", sig)
		os.Exit(0)
	}

}
