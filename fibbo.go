package main

import "fmt"

func main(){
    var a1,a2=0,1
    fmt.Println(a1)
    fmt.Println(a2)
    for i:=0; i<5; i++{
    fmt.Println(a1+a2)
    tmp:=a2
    a2=a1+a2
    a1=tmp
   }
}