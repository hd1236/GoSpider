package main

import (
    "bufio"
    "fmt"
    "golang.org/x/net/html/charset"
    "golang.org/x/text/encoding"
    "golang.org/x/text/transform"
    "io"
    "io/ioutil"
    "net/http"
    "regexp"
)

func main() {
    resp, err := http.Get("http://www.zhenai.com/zhenghun")
    if err != nil{
       panic(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK{
       fmt.Println("resp error statusCode:",resp.StatusCode)
       return
    }


    e := determineEncoding(resp.Body)
    utf8Reader := transform.NewReader(resp.Body,
       e.NewDecoder())

    all, err := ioutil.ReadAll(utf8Reader)
    if err!= nil{
       panic(err)
    }

    //fmt.Printf("%s\n",all)
    
    printCityList(all)
}

func printCityList(content []byte) {
    re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
    match := re.FindAllSubmatch(content, -1)
    for _,m :=range match {
            fmt.Printf("City: %s,URL: %s\n",m[2],m[1])
    }
    fmt.Printf("Matches found : %d\n",len(match))
}

/**
获取网页charset
 */
func determineEncoding(r io.Reader) encoding.Encoding{
    bytes, err := bufio.NewReader(r).Peek(1024)
    if err != nil{
        panic(err)
    }
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}