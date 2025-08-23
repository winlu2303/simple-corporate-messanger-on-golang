package main
import (
  "bufio"
  "fmt"
  "log"
  "net"
)
func handleConnection(conn net.Conn) {
  defer conn.Close()

  fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())
  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    message := scanner.Text()
    fmt.Printf("Delivered message: %s\n", message)
    broadcast(message, conn)
    }
  }
func broadcast(message string, origin net.Conn){
  //Here logic of distribution of messages for every linked clients
  fmt.Printf("Distribution of messages: %s\n", message)
}

  
  
