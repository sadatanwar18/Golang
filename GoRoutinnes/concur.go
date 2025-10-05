package main

import(
	"fmt"
	"time"
	"sync"
)

/*
What is a Process?
- a process is simply an instance of a program in execution
eg - when we open chrome : it becomes one process
	- open vs code : aother process

process consist of:
1- text section - compiled code fo program(read only)
2- data section - global and static variables
3 - heap - dynamicallyy allocated memory
4 - stack - Function call stack contains local variables and return addresses.
5 - program counter - keeps the track of instruction being executed
6 - registers - cpu register used by the process
7 - processid (pid) - unique identifier
8 - open file descritors - list of files the process can access

In short, a process is not just code — it’s code + data + execution context + resources.


*/

/*
Goroutine:
its a lightweight concurrent thread managed by go runtime

*/
func say(s string){
	for i := 0; i< 3; i++{
		fmt.Println(s, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func worker(id int, wg *sync.WaitGroup){
	// what is waitGroup
	// wiatgroup is a synchronization mechanism form GO;s sync package
	// used to Wait for a collection of goroutines to finish executing
	// think of it as a counter that keeps track of how many goroutines are still runnig=ng

	// waitGroup has 3 key methods
	// Add(n)  increase the couter by n(you are telling how many goroutines to wait for)
	// Done()  Decrease the counter by 1
	// Wait() Blocks(pauses) the main goryitne until the counter beocmes 0


	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("WOrker %d done\n",  id)
}


// Race condition
// happens when two or more goroutines access same shared data at the same time
// and at least one of them modifies it, causing unpredictable results



/*
Channels - 
A channel in Go is a pipe (or communication medium) that allows one goroutine to send data and another
 goroutine to receive data in a synchronized and safe way.


*/
func main(){

	//  var wg sync.WaitGroup
	// for i := 1; i <= 3; i++{
	// 	wg.Add(1)
	// 	go worker(i, &wg)
	// }
	// wg.Wait()
	// fmt.Println("All worker done")

	// this will casue race condition

	// to fix race condition we use mutex
	// var mu sync.Mutex
	// couter := 0
	// for i:= 0; i < 1000; i++{
	// 	wg.Add(1)
	// 	go func ()  {
	// 		mu.Lock()
	// 		couter++
	// 		mu.Unlock()
	// 		wg.Done()
	// 	}()
	// }

	// wg.Wait()
	// fmt.Println("Final counter: ", couter)



	// channels

	// ch := make(chan int)  // channel initialization
	// ch <- 10  // send value 10 into channel
	// value := <-ch  // recevive value form channe;
	// ch := make(chan string)

	// // sender goroutine
	// go func ()  {
	// 	ch <- "Hello from goroutine"
	// }()

	// // Receiver (main goroutine)
	// msg := <-ch
	// fmt.Println(msg)

	// buffered vs unbuffered channels
	// unbufferd channel
	// - no buffer : each send must have a corresponding recieve at that time
	// - sender and reciver must synchronixe directly

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)


}