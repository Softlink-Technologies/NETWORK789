package main

import (
    "fmt"
    "net"
)

func main() {
    // Listen for incoming connections on port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()

    fmt.Println("Listening on :8080")

    for {
        // Wait for a connection
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err.Error())
            continue
        }

        // Handle the connection in a new goroutine
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    // Make a buffer to hold incoming data
    buffer := make([]byte, 1024)
    // Read the incoming connection into the buffer
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
        return
    }

    // Send a response back to the client
    response := "Hello, world!\n"
    conn.Write([]byte(response))

    // Close the connection when we're done
    conn.Close()
}
