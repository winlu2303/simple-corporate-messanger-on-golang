package main 
import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)
func runClient() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    go readMessages(conn)
    
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter your message: ")
        message, _ := reader.ReadString('\n')
        _, err := conn.Write([]byte(message))
        if err != nil {
            log.Println(err)
            return
        }
    }
}

func readMessages(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println("Message received:", scanner.Text())
    }
}
