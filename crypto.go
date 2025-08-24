package main
import (
    "crypto/aes"
    "crypto/cipher"
    "errors"
)
const key = "thisisasecretkey123456789" 
//It must be 16, 24 or 32 bytes 
func EncryptMessage(plaintext string) ([]byte, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return nil, err
    }
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    rest := ciphertext[aes.BlockSize:]
    
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    stream := cipher.NewCBCEncrypter(block, iv)
    stream.CryptBlocks(rest, []byte(plaintext))
       return ciphertext, nil
}

func DecryptMessage(ciphertext []byte) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    if len(ciphertext) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    rest := ciphertext[aes.BlockSize:]
    
    stream := cipher.NewCBCDecrypter(block, iv)
    stream.CryptBlocks(rest, rest)
       return string(rest), nil
}
