package main

import (
    "os"
    "net/http"
    "io"
    "gmt"
)

func main(){
  // Create the file
  file, err := os.Create("downloaded-pdf.pdf")
  if err != nil  {
    fmt.Println("File can't be created:", err.Error())
    return
  }
  defer file.Close()

  // do the http call (download the file)
  resp, err := http.Get("")
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // check if http code = 200
  if resp.StatusCode != http.StatusOK {
    fmt.Println("Fail in http call. Status:", resp.StatusCode)
    return
  }

  // copy the stream data from the response body to the file
  _, err = io.Copy(file, resp.Body)
  if err != nil  {
    fmt.Println("Fail copying data to the file:", err.Error())
    return
  }
}