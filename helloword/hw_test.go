package main

import (
    "testing"
    "strconv"
)

func TestAsdfksddf(t * testing.T){

}


func BenchmarkTesddt001(b * testing.B){
    b.StopTimer()

    b.StartTimer()

    imap := map[string]int{}
    for  i:=0; i<b.N; i++{
        imap[strconv.FormatInt(int64(i), 10)] = i
    }
}