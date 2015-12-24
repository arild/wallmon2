package main

import (
    "fmt"
    "net"
    "os"
    "log"
	"io"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"./protocol"
	"net/http"
	"github.com/gorilla/websocket"
)

func main() {
	metricChan := make(chan []byte)

	go listenMetricData(metricChan)

	//	http.HandleFunc("/", handler)
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		echoHandler(w, r, metricChan)
	})
	http.Handle("/", http.FileServer(http.Dir("./src/web")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request, metricsChan chan []byte) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	log.Println("echo handler")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for metric := range metricsChan {
		err = conn.WriteMessage(websocket.BinaryMessage, metric);
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}
	}
}

func listenMetricData(metricsChan chan []byte) {
	l, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("New connection")
		go handleRequest(conn, metricsChan)
	}
}

func handleRequest(conn net.Conn, metricsChan chan []byte) {
    defer conn.Close()
    for {
		header := socketRecv(conn, 4)
		msgLen := binary.BigEndian.Uint32(header)
//		log.Printf("Header size %d\n", msgLen)

		payload := socketRecv(conn, msgLen)
		metricsChan <- payload
		metrics := &protocol.MetricsMessage{}
		proto.Unmarshal(payload, metrics)

//		log.Printf("Num processes=%d, Pid=%d, CPU=%f\n", len(metrics.GetMetricsmessage()), *metrics.GetMetricsmessage()[0].Pid, *metrics.GetMetricsmessage()[0].Cpu)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func socketRecv(conn net.Conn, n uint32) []byte {
	buf := make([]byte, n)
	_, err := io.ReadFull(conn, buf);
	if err != nil {
		log.Fatal("Error reading:", err.Error(), " Terminating connection.")
	}
	return buf
}