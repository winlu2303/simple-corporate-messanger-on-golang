package main
import (
  "fmt"
  "log"
  "net"
)
func main() {
  fmt.Println("Start CM")
  //Go Server
  go runServer()
  //Link Client
  runClient()
}
func runServer() {
  //Create listener
  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
      log.Fatal(err)
  }
  defer listener.Close()

  for {
      conn, err := listener.Accept()
    if err != nil {
      log.Println(err)
      continue
    }
    go handleConnection(conn)
  }
}
  
