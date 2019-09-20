package main

import (
    "fmt"
    "regexp"
)

const text  = `hehe apricityhand@gmail.com@abc.com
email1 is abc@def.org 
email2 is    kkk@qq.com
`

func main(){
    re  := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
    //match := re.FindString(text)
    //fmt.Println(match)

    //match := re.FindAllString(text,-1)
    //fmt.Println(match)

    match := re.FindAllStringSubmatch(text,-1)
    for _,m := range match{
        fmt.Println(m)
    }


}
