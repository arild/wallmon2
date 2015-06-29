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
)

func main() {

    system.GetAllPids()

    ch := make(chan []byte)

    go sampleData(ch)
    go sendData(ch)

    handleCtrlC()
}

func handleCtrlC() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    for sig := range c {
        println("Terminate program", sig)
        os.Exit(0)
    }
}

func sampleData(ch chan []byte) {

    tick := time.Tick(2000 * time.Millisecond)
    for now := range tick {
        fmt.Printf("%v\n", now)

        metrics := &Metrics{
            Pid: proto.Uint32(1),
            Cpu:  proto.Float32(0.32),
        }

        data, err := proto.Marshal(metrics)
        if err != nil {
            log.Fatal("marshaling error: ", err)
        }

        fmt.Printf("Forwording sampled data")
        ch <- data
    }
    println("test")
}

func sendData(ch chan []byte) {
    conn, err := net.Dial("tcp", "localhost:3333")
    if err != nil {
        log.Fatal("Failed opening socket")
    }

    for item := range ch {
        fmt.Printf("Item on channel, length=%d", len(item))

        _, err = conn.Write(item)
        if err != nil {
            println("Write to server failed:", err.Error())
            os.Exit(1)
        }

    }

    conn.Close()

}