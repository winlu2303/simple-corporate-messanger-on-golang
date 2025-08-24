package main
import (
    "crypto/bcrypt"
    "encoding/json"
    "net/http"
)
type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        SendError(w, http.StatusBadRequest, "Invalid request")
        return
    }
    if !isValidUsername(req.Username) {
        SendError(w, http.StatusBadRequest, "Invalid user name")
        return
    }
    if !isValidPassword(req.Password) {
        SendError(w, http.StatusBadRequest, "Invalid password")
        return
    }
    if err := RegisterUser(req.Username, req.Password); err != nil {
        SendError(w, http.StatusInternalServerError, err.Error())
        return
    }
    SendResponse(w, http.StatusCreated, "The user has been successfully registered") 
} 

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        SendError(w, http.StatusBadRequest, "Invalid request")
        return
    }
    user, err := LoginUser(req.Username, req.Password)
    if err != nil {
        SendError(w, http.StatusUnauthorized, "Invalid user credentials")
        return
    }
    
    SendResponse(w, http.StatusOK, "Successful authorization") 
} 

func isValidUsername(username string) bool {
    //Simple user name validation
    return len(username) >= 3 && len(username) <= 32 
} 

func isValidPassword(password string) bool {
    //Simple password validation
    return len(password) >= 8
}
