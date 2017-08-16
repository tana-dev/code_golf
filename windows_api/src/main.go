package main

import (
  "fmt"
  "net/http"
  // "net/url"
  // "os"
  "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "Hello, World")

  // pathを取るにはr.URL.Pathで受け取る
  path := r.URL.Path
  // path := r.URL

  // stringsパッケージのReplaceで/を空白に変えておく
  path = strings.Replace(path, "/", " ", -1)

  // fmt.Println(path)

  // ファイル存在チェック
//  _, err := os.Stat()
//  if err != nil {
//    fmt.Println(err)
//  }

  // ak, _ := url.QueryUnescape(path)
  // ak := Query()path.Encode()
  // fmt.Fprintf(w, "Hello!%s", path)
  fmt.Fprintf(w, path)
  // fmt.Fprintf(w, ak)
}

func main() {
  // 引数取得

  // ファイル存在チェック
//  _, err := os.Stat()
//  if err != nil {
//    fmt.Println(err)
//  }

  http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
  http.ListenAndServe(":8080", nil)
}
