/*
Andrew Meijer V00805554
Please refer to the description file for information about this code.
this code is intended to simulate the pattern in vector_clocks_example.png
*/

package main

import ("fmt")

const num_threads = 3
/*
Hard-coded channels between threads
*/
var ab = make(chan [num_threads]int)
var ac = make(chan [num_threads]int)
var ba = make(chan [num_threads]int)
var bc = make(chan [num_threads]int)
var ca = make(chan [num_threads]int)
var cb = make(chan [num_threads]int)
var done = make(chan bool)

func thread_A () {
    TS := [num_threads]int{0,0,0}
    /*Receive a message from B*/
    TS[0] = TS[0] + 1 //increment local counter
    MTS := <- ba
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from B to A. TS Updated: ", TS)

    /*Send a message to B*/
    TS[0] = TS[0] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from A to B.")
    ab <- TS

    /*Receive a message from C*/
    TS[0] = TS[0] + 1 //increment local counter
    MTS = <- ca
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from C to A. TS Updated: ", TS)

    /*Receive a message from C*/
    TS[0] = TS[0] + 1 //increment local counter
    MTS = <- ca
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from C to A. TS Updated: ", TS)

    /*All done.*/
    done <- true
}

func thread_B () {
    TS := [num_threads]int{0,0,0}
    /*Receive a message from C*/
    TS[1] = TS[1] + 1 //increment local counter
    MTS := <- cb
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from C to B. TS Updated: ", TS)

    /*Send a message to A*/
    TS[1] = TS[1] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from B to A.")
    ba <- TS

    /*Send a message to C*/
    TS[1] = TS[1] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from B to C.")
    bc <- TS

    /*Receive a message from A*/
    TS[1] = TS[1] + 1 //increment local counter
    MTS = <- ab
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from A to B. TS Updated: ", TS)

    /*Send a message to C*/
    TS[1] = TS[1] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from B to C.")
    bc <- TS

}

func thread_C () {
    TS := [num_threads]int{0,0,0}
    /*Send a message to B*/
    TS[2] = TS[2] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from C to B.")
    cb <- TS

    /*Receive a message from B*/
    TS[2] = TS[2] + 1 //increment local counter
    MTS := <- bc
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from B to C. TS Updated: ", TS)

    /*Send a message to A*/
    TS[2] = TS[2] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from C to A.")
    ca <- TS

    /*Receive a message from B*/
    TS[2] = TS[2] + 1 //increment local counter
    MTS = <- bc
    // The magic: TS[k] = max(TS[k], MTS[k]) for k = 1 to N
    for k := 0; k < num_threads; k++ {
        if TS[k] < MTS[k] {
            TS[k] = MTS[k]
        }
    }
    fmt.Println("Received message from B to C. TS Updated: ", TS)

    /*Send a mesasge to A*/
    TS[2] = TS[2] + 1 //increment local counter
    fmt.Println("Sending ", TS, " from C to A.")
    ca <- TS
}

func main () {
    go thread_C()
    go thread_B()
    go thread_A()
    <- done
}
