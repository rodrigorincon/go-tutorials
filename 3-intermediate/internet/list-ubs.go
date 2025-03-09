package main

import (
    "net/http"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "errors"
)

type Ubs struct {
  lat float64
  lon float64
  codMunic uint32
  codCnes uint32
  name string
  address string
  phone string
}

func convertToStruct(scanner *bufio.Scanner)[]Ubs{
  list := []Ubs{}

  // skip the header
  scanner.Scan()

  hasMore := scanner.Scan()
  for hasMore {
    line := scanner.Text()
    attributes := strings.Split(line, ",")

    lat, _ := strconv.ParseFloat(attributes[0], 32)
    lon, _ := strconv.ParseFloat(attributes[1], 32)
    codMunic, _ := strconv.ParseUint(attributes[2], 10, 32)
    codCnes, _ := strconv.ParseUint(attributes[3], 10, 32)

    ubs := Ubs{ 
      lat: lat,
      lon: lon,
      codMunic: uint32(codMunic),
      codCnes: uint32(codCnes),
      name: attributes[4],
      address: attributes[5]+" "+attributes[6]+" "+attributes[7],
      phone: attributes[8],
    }
    list = append(list, ubs)

    hasMore = scanner.Scan()
  }
  return list
}

func readFromUrl()(ubs []Ubs, err error){
  resp, err := http.Get("http://repositorio.dados.gov.br/saude/unidades-saude/unidade-basica-saude/ubs.csv")
  if err != nil {
    return []Ubs{}, err
  }
  defer resp.Body.Close()

  // check if http code = 200
  if resp.StatusCode != http.StatusOK {
    return []Ubs{}, err
  }

  var mySlice []byte
  scanner := bufio.NewScanner(resp.Body)
  scanner.Buffer(mySlice, 1024)

  ubs = convertToStruct(scanner)
  if len(ubs) == 0 {
    return []Ubs{}, errors.New("can't convert response.body to struct")
  }

  return ubs, nil
}

func main(){
  ubsList, err := readFromUrl()
  if err != nil {
    fmt.Println("Error trying to get the UBS: ", err.Error())
  }

  fmt.Println( ubsList[len(ubsList)-1] )
}