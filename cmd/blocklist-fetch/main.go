package main

import (
  "io"
  "log"
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
  content, err := ioutil.ReadALL(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  return string(content)
}
