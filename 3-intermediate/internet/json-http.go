package main

import (
    "io"
    "net/http"
    "fmt"
    "encoding/json"
)

type Unit struct{ // if the var was not Caitalized, the unmarshal will not work
  Code int8 `json:"codigo_tipo_unidade"`
  Desc string `json:"descricao_tipo_unidade"`
}

type UnitType struct {
  List []Unit `json:"tipos_unidade"`
}

func main(){
  // do the http call
  resp, err := http.Get("https://apidadosabertos.saude.gov.br/cnes/tipounidades")
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

  bodyBytes, err := io.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("fail to read the response body:", err.Error())
    return
  }

  unitTypes := UnitType{}
  json.Unmarshal( bodyBytes, &unitTypes)
  for _, unit := range(unitTypes.List) {
    fmt.Println("Code: ", unit.Code, "Description: ", unit.Desc)
  }
}