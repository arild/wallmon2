package main

import (
    "fmt"
    "net"
    "os"
    "log"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()

    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        fmt.Print("New connection")
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    defer conn.Close()
    for {
        buf := make([]byte, 1024 * 10)
        n, err := conn.Read(buf)
        log.Printf("Received %d bytes", n)
        if err != nil {
            fmt.Println("Error reading:", err.Error())
        }
    }
}