package main 
import (
  "log"
  "os"
  "time"
  )

var logger*log.logger
func init()[
  file, err := os.OpenFile("messenger.log", os.OCREATE | os.OWRONLY | os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("Errror creattion log file: %v", err)
    }
  defer file.Close()
  logger = log.New(
    file,  //Where to write
    "[Messenger]", //Prefix for all messages
    log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile, //Flags of formatting
    )
  }
  //Special log func
  func LogInfo(message string) {
    logger.Printf("INFO: %s", message)
    }

  func LogError(message string, err error) {
    logger.Printf("ERROR: %s-%v", message, err)
    }

  func LogWarning(message string) {
    logger.Printf("WARNING: %s", message)
    }

  func LogDebug(message string) {
    logger.Printf("DEBUG: %s", message)
    }

  func LogClientConnect(addr string) {
    LogInfo(fmt.Sprintf("New connection %s", addr))
    }

  func LogClientDisconnect(addr string) {
    LogInfo(fmt.Sprintf("Disconnection user %s", addr))
    }

  func LogMessage(from, to, content string) {
    LogInfo(fmt.Sprintf("Message from %s to %s: [hide content]", from, to))
    }

  func LogAuthError(username string, err error) {
    LogError(fmt.Sprintf("User authorization error %s", username), err)
    }

  func LogSuccess(operation string) {
    LogInfo(fmt.Sprintf("Successfuly complete: %s", operation))
    }

  //Integartion log to main code
  func WebsocketHandler (w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
      LogError("Connect update error", err)
      return
      }
    defer conn.Clodse()
    LogClientConnection(conn.RemoteAddr().String())
    for {
      _, message, err := conn.ReadMessage()
      if err != nil {
        LogError("Read message error", err)
        break
        }
      LogMessage("user", "room", string(message))
      }
    LogClientDisconnect(conn.RemoteAddr().String())
    }

  func LoginHandler(w http.ResponseWriter, r *http.Request) {
    user, err := LoginUser(username, password)
    if err != nil {
      LogAuthError(esurname, err)
      SendError(w, http.StatusUnauthorized, "Incorrect data")
      return
      }
    LogSuccess(fmt.Sprintf("Successful user authorization %s", username))
    }
  
