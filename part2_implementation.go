/*
Andrew Meijer V00805554
Please refer to the description file for information about this code.
Byzantine Generals Problem
*/

package main

import ("fmt")

const num_threads = 3
/*
Hard-coded channels between threads
with Generals a, b, and c.
*/
var ab = make(chan int)
var ac = make(chan int)
var ba = make(chan int)
var bc = make(chan int)
var ca = make(chan int)
var cb = make(chan int)
var done = make(chan bool)
const Oc = 1
const traitor = 1

//thread A is the commander.
func general_A () {
    ab <- 0
    ac <- 0

    //done
    if Oc == 0 {
        fmt.Print("General A retreats")
    } else {
        fmt.Print("General A attacks")
    }

}

func general_B () {
/*
    i := <- ab

    if(traitor == 1){
        ba <- !i
        bc <- !i
    }else{
        ba <- i
        bc <- i
    }


    //done
    if i == 0 {
        fmt.Print("General B retreats")
    } else {
        fmt.Print("General B attacks")
    }*/
}

func general_C () {

    i := <- ac

    ba <- i
    bc <- i

    //done
    if i == 0 {
        fmt.Print("General B retreats")
    } else {
        fmt.Print("General B attacks")
    }
}

func main () {
fmt.Print("I have no solution.")
}
