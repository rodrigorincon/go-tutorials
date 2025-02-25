package main

import (
    "io"
    "net/http"
    "bytes"
    "fmt"
)

func main(){
  // do the http call (download the file)
  resp, err := http.Get("https://www.google.com")
  if err != nil {
    fmt.Println("fail to connect with URL:", err.Error())
    return
  }
  defer resp.Body.Close()

  // check if http code = 200
  if resp.StatusCode != http.StatusOK {
    fmt.Println("Fail in http call. Status:", resp.StatusCode)
    return
  }

  bodyStream, err := io.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("fail to read the response body:", err.Error())
    return
  }

  textToSearch := []byte("<input class=\"lst Ucigb\"")
  textFound := bytes.Contains(bodyStream, textToSearch)

  fmt.Println("Text founded: ", textFound)
}