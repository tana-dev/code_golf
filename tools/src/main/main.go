package main

import (
        "net/http"
        "net/http/cookiejar"
        "io/ioutil"
        "bytes"
)

func main() {
        cookieJar, _ := cookiejar.New(nil)

        client := &http.Client {
                Jar: cookieJar,
        }

        // Login to swipe.to
        var str = []byte("next=&csrfmiddlewaretoken=QzQDjeUISAKL7rnX0SmGRNm6dfPvuWkq&username=nekodanshaku27%40gmail.com&password=shun1127Pass")
        req, _ := http.NewRequest("POST", "https://bitbucket.org/account/signin/", bytes.NewBuffer(str))
        req.Header.Set("Referer", "https://bitbucket.org/account/signin/")
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

        res, _ := client.Do(req)
        body, _ := ioutil.ReadAll(res.Body)
        println(string(body))

        // Create New Doc
        // req2, _ := http.NewRequest("POST", "https://www.swipe.to/edit/create", nil)
        // req2.Header.Set("Referer", "https://www.swipe.to/home")
        // res2, _ := client.Do(req2)
        // body, _ := ioutil.ReadAll(res2.Body)
        // defer res2.Body.Close()
        //
        // println(string(body))
}
