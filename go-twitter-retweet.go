package main

import (
        "encoding/json"
        "fmt"
        "github.com/ChimeraCoder/anaconda"
        "log"
        "os"
        "strconv"
)

func main() {
  anaconda.SetConsumerKey(os.Args[1])
  anaconda.SetConsumerSecret(os.Args[2])
  api := anaconda.NewTwitterApi(os.Args[3], os.Args[4])
  id := os.Args[5]

  rt_id, err := strconv.ParseInt(id, 10, 64)

  if err != nil {
    log.Fatalf("%s", err)
  }

  rt, err := api.Retweet(rt_id, true)

  if err != nil {
    log.Fatalf("%s", err)
  }

  result_json,err := json.Marshal(rt)

  if err != nil {
    fmt.Println(err)
  }

  tweet_json := "{\"response\":" + string(result_json) + "}"

  fmt.Println(tweet_json)

}
