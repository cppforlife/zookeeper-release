package main

import (
  "fmt"
  "os"
  "strings"
  "time"

  "github.com/samuel/go-zookeeper/zk"
)

// See article: https://mmcgrana.github.io/2014/05/getting-started-with-zookeeper-and-go.html
func main() {
  zksStr := os.Getenv("ZOOKEEPER_SERVERS")
  zks := strings.Split(zksStr, ",")

  conn, _, err := zk.Connect(zks, time.Second)
  panicIfErr(err, "connect")

  defer conn.Close()

  keyPath := "/bosh_smoke_tests"

  {
    _, err := conn.Create(keyPath, []byte("data"), int32(0), zk.WorldACL(zk.PermAll))
    panicIfErr(err, "create")

    fmt.Printf("Successfully created value\n")
  }

  {
    data, stat, err := conn.Get(keyPath)
    panicIfErr(err, "get")

    if string(data) == "data" {
      fmt.Printf("Successfully retrieved created value\n")
    } else {
      panic(fmt.Errorf("Expected value match created value"))
    }

    stat, err = conn.Set(keyPath, []byte("newdata"), stat.Version)
    panicIfErr(err, "set")

    data, _, err = conn.Get(keyPath)
    panicIfErr(err, "get")

    if string(data) == "newdata" {
      fmt.Printf("Successfully set new value\n")
    } else {
      panic(fmt.Errorf("Expected value match changed value"))
    }
  }

  {
    err := conn.Delete(keyPath, -1)
    panicIfErr(err, "delete")

    exists, _, err := conn.Exists(keyPath)
    panicIfErr(err, "exists")

    if exists == false {
      fmt.Printf("Successfully deleted value\n")
    } else {
      panic(fmt.Errorf("Expected value to not exist"))
    }  
  }
}

func panicIfErr(err error, cause string) {
  if err != nil {
    panic(fmt.Errorf("%s: %s", cause, err))
  }
}
