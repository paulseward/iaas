// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  "fmt"
  "net/http"

  "bufio"
  "os"
  "time"
  "math/rand"

  "google.golang.org/appengine"
)

func main() {
  http.HandleFunc("/", handle)
  appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {

  insults, err := readLines("./data/insults.txt")

  if err != nil {
    fmt.Fprintln(w, "Something went wrong")
  }

  random := rand.New(rand.NewSource(time.Now().Unix()))
  fmt.Fprintln(w, insults[random.Intn(len(insults))])
}


func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
