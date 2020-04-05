package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
  log.Print("sys-hac-states: received a request")
  hac_url := os.Getenv("HAC_URL")
  hac_path := "/api/states"
  hac_token := os.Getenv("HAC_TOKEN")

  if hac_url == "" || hac_token == "" {
    fmt.Fprintln(w,"Home Assistant Core configuration missing!\nPlease specify ENV.HAC_URL and HAC_TOKEN")
  } else {
    req, err := http.NewRequest("GET", hac_url + hac_path, nil)
    req.Header.Add("Authorization", "Bearer " + hac_token)

    client := &http.Client{}
    response, clientErr := client.Do(req)
    if clientErr != nil {
        log.Println("Error on response.\n[ERRO] -", clientErr)
    }

    responseData, dataErr := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(dataErr)
    }

    responseString := string(responseData)
    fmt.Fprint(w, responseString)
  }
}

func main() {
  log.Print("sys-hac-states: starting server...")

  http.HandleFunc("/", handler)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Printf("sys-hac-states: listening on port %s", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
