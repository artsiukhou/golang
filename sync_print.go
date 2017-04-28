package main

import (
    "log"
    "time"
)

func printOdd(checkCh chan int, oddCh chan int, evenCh chan int) {
    for i := 1; ; i += 2 {
        time.Sleep(100 * time.Millisecond)
        <- oddCh
        checkCh <- i
        evenCh <- 1
    }
}

func printEven(checkCh chan int, oddCh chan int, evenCh chan int) {
    for i := 2; ; i += 2 {
        <- evenCh
        checkCh <- i
        oddCh <- 1
    }
}

func check(checkCh chan int) {
    prevVal := 0
    for {
        val := <- checkCh
        if val != prevVal + 1 {
            log.Fatalln("ERROR")
        }
        log.Println(val)
        prevVal = val
    }
}

func main() {
    oddCh := make(chan int)
    evenCh := make(chan int)
    checkCh := make(chan int)

    go check(checkCh)
    go printOdd(checkCh, oddCh, evenCh)
    go printEven(checkCh, oddCh, evenCh)

    oddCh <- 1
    time.Sleep(5 * time.Second)
}
