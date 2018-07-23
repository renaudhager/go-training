package consumer

import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/imroc/req"
    "io/ioutil"
    "time"
)

var (
  kong_admin_enpoint = os.Getenv("KONG_ADMIN_API_ENDPOINT")
  timeout time.Duration = 5
)

type consumer struct {
  createdAt int64  `json:"created_at"`
  Username  string `json:"username"`
  id        string `json:"id"`
}

type consumers struct {
  Total int `json:"total"`
  Consumers  []consumer
}

func CreateConsumer(consumer_name string) (*consumer, error) {
  var c = new(consumer)

  url := kong_admin_enpoint + "/consumers"
  req.SetTimeout(timeout * time.Second)
  param := req.Param{
    "username": consumer_name ,
  }

  r, _ := req.Post(url, param)
  resp := r.Response()
  body, _ := ioutil.ReadAll(resp.Body)
  err := json.Unmarshal(body, &c)
  if(err != nil){
      fmt.Println("whoops:", err)
  }

  return c, err
}

func ListConsumers() (*consumers, error) {
  var cs = new(consumers)
  url := kong_admin_enpoint + "/consumers"

  r, _ := req.Get(url)
  resp := r.Response()
  body, _ := ioutil.ReadAll(resp.Body)

  fmt.Println(string(body))

  err := json.Unmarshal(body, &cs)
  if(err != nil){
      fmt.Println("whoops:", err)
  }

  return cs, err

}