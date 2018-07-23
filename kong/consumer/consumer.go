package consumer

import (
    "fmt"
    "os"
    "encoding/json"
    "github.com/imroc/req"
    "io/ioutil"
    "time"
    "errors"
)

var (
  kong_admin_enpoint = os.Getenv("KONG_ADMIN_API_ENDPOINT")
  timeout time.Duration = 5
  errConsumerAlreadyExist = errors.New("Consumer already exist.\n")
  errUnexpectedStatusCode = errors.New("Unexpected status code.\n")
  err_return = errors.New("default error.\n")
)

type consumer struct {
  CreatedAt int64  `json:"created_at"`
  Username  string `json:"username"`
  Id        string `json:"id"`
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

  r, err_req := req.Post(url, param)

  if(err_req != nil){

      fmt.Println("Error while querying Kong API: ", err_req)
      err_return = err_req

  } else {

    resp := r.Response()

    if(resp.StatusCode == 201){

      fmt.Printf("Consumer %s successfully created.\n", consumer_name)
      body, _ := ioutil.ReadAll(resp.Body)
      err_json := json.Unmarshal(body, &c)

      if(err_json != nil){

          fmt.Println("whoops:", err_json)
          err_return = err_json

      }
    } else if (resp.StatusCode == 409){

      fmt.Printf("Consumer %s already created.\n", consumer_name)
      err_return = errConsumerAlreadyExist

    } else {
      fmt.Printf("Error unexpected status code %v.\n", resp.StatusCode)
      err_return = errUnexpectedStatusCode
    }
  }

  return c, err_return
}

func GetConsumer(consumer_name string) (*consumer, error) {
  var c = new(consumer)

  url := kong_admin_enpoint + "/consumers/" + consumer_name
  req.SetTimeout(timeout * time.Second)

  r, err_req := req.Get(url)

  if(err_req != nil){

      fmt.Println("Error while querying Kong API: ", err_req)
      err_return = err_req

  } else {

    resp := r.Response()
    if(resp.StatusCode == 200) {

      body, _ := ioutil.ReadAll(resp.Body)
      err_json := json.Unmarshal(body, &c)

      if(err_json != nil){

          fmt.Println("whoops:", err_json)
          err_return = err_json

      }

    } else {
      fmt.Printf("Error unexpected status code %v.\n", resp.StatusCode)
      err_return = errUnexpectedStatusCode
    }

  }


  return c, err_return
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