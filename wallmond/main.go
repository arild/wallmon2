package main

import (
    "net"
    "os"
    "time"
    "fmt"
    "os/signal"
    "./system"
)

type DataItem struct {
    cpu float32
}

func main() {

    pids := system.GetAllPids()
    for elem, _ := range pids {
        fmt.Println(elem)
    }
    return

    ch := make(chan *DataItem)

    go sampleData(ch)
    go sendData(ch)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    for sig := range c {
        println("Terminate program", sig)
        os.Exit(0)
    }


}

func sampleData(ch chan *DataItem) {

    tick := time.Tick(1000 * time.Millisecond)
    for now := range tick {
        fmt.Printf("%v\n", now)
        //monitor := new(Monitor)
        //println(monitor.getCpu())
        dataItem := new(DataItem)
        dataItem.cpu = 50.2

        ch <- dataItem
    }
    println("test")
}

func sendData(ch chan *DataItem) {
    
    for item := range ch {
        println(item.cpu)
    }

    conn, err := net.Dial("tcp", "localhost:3333")
    if err != nil {
        // handle error
    }

    _, err = conn.Write([]byte("Hello"))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    reply := make([]byte, 1024)

    _, err = conn.Read(reply)
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    println("reply from server=", string(reply))

    conn.Close()

}