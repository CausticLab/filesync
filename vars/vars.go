package vars

import (
  "os"
  "fmt"
  "log"
  "strconv"
  "io/ioutil"
  simplejson "github.com/bitly/go-simplejson"
)

type ConfigVars struct{
  Mode string
  Ip string
  Port int
  Monitors map[string]interface{}
}

func Args() []string {
  ret := []string{}
  if len(os.Args) >= 1 {
    for i := 1; i < len(os.Args); i++ {
      ret = append(ret, os.Args[i])
    }
  }
  return ret
}

func GetConfig() ConfigVars{
  var config ConfigVars
  var configFile string

  input := Args()
  if len(input) >= 1 {
    configFile = input[0]
  }

  if(configFile != ""){
    fmt.Println("HERE ", configFile)
    b, err := ioutil.ReadFile(configFile)
    if err != nil {
      fmt.Println(configFile, " not found")
      //return
    } else {
      json, _ := simplejson.NewJson(b)
      config.Mode = json.Get("mode").MustString()
      config.Ip = json.Get("ip").MustString()
      config.Port = json.Get("port").MustInt()
      config.Monitors = json.Get("monitors").MustMap()
    }
  } else {
    config.Mode     = os.Getenv("FILESYNC_MODE")
    config.Ip       = os.Getenv("FILESYNC_IP")
    config.Port, _  = strconv.Atoi(os.Getenv("FILESYNC_PORT"))

    config.Monitors = make(map[string]interface{})
    config.Monitors["default"] = os.Getenv("FILESYNC_PATH")
  }

  if(config.Mode == ""){
    config.Mode = "server"
  }

  if(config.Ip == ""){
    config.Ip = "0.0.0.0"
  }

  if(config.Port <= 0){
    config.Port = 6776
  }

  if(len(config.Monitors) == 0){
    log.Println("No paths to monitor - defaulting to /share")

    config.Monitors = make(map[string]interface{})
    config.Monitors["default"] = "/share"

    /*
    var err error

    if _, err := os.Stat("/share"); os.IsNotExist(err) {
      os.Mkdir("/share", os.ModePerm)
    }

    if err != nil {
      log.Fatal("Could not create default /share directory")
    } else {
      log.Println("Created /share directory")
    }
    */
  }


  return config
}
















