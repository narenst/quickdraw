package main

import (
    "github.com/hoisie/web"
    "encoding/json"
    "log"
)

func getVideos() string {
   video := new(Video)
   video.Url = "http://vimeo.com/69170991"
   video.Service = "vimeo"
   video.Climber.Name = "James Kassay"
   video.Climber.Age = 28
   video.Location = "Grampians, Australia"
   log.Println(video)

   json, err := json.Marshal(video)

   if err != nil {
      log.Println(err)
   }
   return string(json)
}

func hello(path string) string {
   return "Hola " + path
}

func main() {
    web.Get("/videos", getVideos)
    web.Run("0.0.0.0:9999")
}
