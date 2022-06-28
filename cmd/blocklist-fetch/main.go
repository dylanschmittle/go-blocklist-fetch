package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)

func main() {
  fmt.Println(getList("https://raw.githubusercontent.com/stamparm/ipsum/master/ipsum.txt"))
}

func getList(link string)(string) {
  res, err := http.Get(link)
  if err != nil {
    log.Fatal(err)
  }
  content, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  return string(content)
}
