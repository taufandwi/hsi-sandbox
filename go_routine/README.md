## Goroutine 
* In simple terms: A goroutine is a function that is capable of running concurrently (seemingly at the same time) with other functions.
* Think of it as a very lightweight, cheap "assistant" that you can create to do a job for you without blocking your main work.
* Key Characteristics:
   - Lightweight: Goroutines are much lighter than traditional threads. You can create thousands of them without consuming a lot of memory.
   - Managed by Go Runtime: The Go runtime handles the scheduling of goroutines, so you don't have to worry about it.
   - Concurrent: Goroutines can run in parallel on multiple CPU cores, making your programs faster and more efficient.
  * How to Create a Goroutine:
     - Use the `go` keyword followed by a function call. This tells Go to run that function in a new goroutine.
     - Example:
     ```go
     package main
        import (
            "fmt"
            "time"
        )
        func sayHello() {
            for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond) // Simulate work
                fmt.Println("Hello from goroutine!")
            }
        }
        func main() {
            go sayHello() // Start a new goroutine
            for i := 0; i < 5; i++ {
                time.Sleep(150 * time.Millisecond) // Simulate work
                fmt.Println("Hello from main!")
            }
            time.Sleep(1 * time.Second) // Wait for goroutine to finish
        }
        ```

## Channel
**The Analogy: A Magical Conveyor Belt** 🚚 <br>
Imagine your two chefs (your goroutines) from the last lesson. What if the salad chef needs to give the chopped vegetables to the soup chef? They can't just throw them across the room. They need a safe way to pass them.<br> 

A channel is like a magical conveyor belt between them : 
* You can put an item (data) on one end of the belt. This is called sending. 
* Someone else can take the item off the other end. This is called receiving. 
* It's safe: Only one person can take an item at a time. No one can steal an item or grab a half-finished one. 
* It's typed: If it's a "vegetable" conveyor belt, you can only put vegetables on it, not soup pots.

This conveyor belt is the primary way goroutines communicate and synchronize their work.

---

## What is a Channel?
A **channel** is a typed "pipe" that connects concurrent goroutines. You can send values into the channel from one goroutine and receive those values in another.

The Syntax:
* Create a channel: myChannel := make(chan int) (This creates a channel that can only transport integers int).
* Send a value: myChannel <- 10 (The arrow <- points the data into the channel). 
* Receive a value: myValue := <-myChannel (The arrow <- points the data out of the channel and into a variable).

Now, let's look at the different kinds of conveyor belts (channels) you can have.

---
## Types of Channels
1. **Unbuffered Channels**:
   - Unbuffered Channels (The Direct Hand-off)
   - This is the default and most basic type of channel. 
   - Think of it as a conveyor belt with zero space. It's just a direct hand-off from one chef to another. 
   - If a sender (Chef A) wants to put something on the belt, they must wait until a receiver (Chef B) is there, ready to take it immediately. 🤝
   - If a receiver (Chef B) arrives at the belt to take something, they must wait until a sender (Chef A) arrives to give it to them. 
   - This act of waiting is called blocking. This is incredibly powerful because it forces synchronization. The sender and receiver are guaranteed to be in sync at the moment of exchange.
   - Example :
   ```go
   package main

   import (
   "fmt"
   "time"
   )

   func main() {
      // Create a new unbuffered channel for strings
      // make(chan string)
      messages := make(chan string)

      // Start a new goroutine (an anonymous function in this case)
	  go func() {
		    fmt.Println("Goroutine: Waiting to send a message...")
		    time.Sleep(2 * time.Second) // Pretend to do some work
		    messages <- "Ping!" // Send message. This will BLOCK until main is ready to receive.
		    fmt.Println("Goroutine: Message sent!")
	  }()

	  fmt.Println("Main: Waiting to receive a message...")
	  msg := <-messages // Receive message. This BLOCKS main until the goroutine sends something.
	  fmt.Println("Main: Received message:", msg)
   }
   ```         
    - Output :
    ```
   Main: Waiting to receive a message...
   Goroutine: Waiting to send a message...
   // ... two-second pause ...
   Goroutine: Message sent!
   Main: Received message: Ping!
    ```
   - Notice how main waited at msg := <-messages until the goroutine was ready to send. This is synchronization in action!<br><br>
   
2. **Buffered Channels**:
   - This is like a conveyor belt that has a limited amount of space on it.
   - You create it by adding a capacity number: make(chan string, 2) creates a channel that can hold 2 strings.
   - Sending: A sender can put an item on the belt without waiting, as long as there is empty space on the belt. If the belt is full, the sender will block and wait for a spot to open up. ✅
   - Receiving: A receiver can take an item off the belt if there's anything on it. If the belt is empty, the receiver will block and wait for a sender to put something on it. ❌
   - Example
   ```go
   package main

   import "fmt"
   
   func main() {
      // Create a buffered channel with a capacity of 2
      messages := make(chan string, 2)

       // We can send two values without blocking because there is space.
       messages <- "Buffered"
       messages <- "Channel"
       fmt.Println("Sent two messages without blocking.")
   
       // If we tried to send a third one here, it would block!
       // messages <- "Third" // This would cause a deadlock or wait for a receiver.
   
       // Now we can receive the two values.
       fmt.Println(<-messages)
       fmt.Println(<-messages)
    }
   ```
   - Output :
   ```
   Sent two messages without blocking.
   Buffered
   Channel
   ```
   - Notice how we could send two messages without blocking because the channel had space. If we tried to send a third message without receiving one first, it would block and wait for space to open up.<br><br>

3. **Directional Channels**:
   - Sometimes you want to be very clear about how a channel should be used. You can specify a channel's direction (send-only or receive-only). This is mostly used as function parameters to make your code safer and more readable.
   - Send-only: chan<- int (You can only send to this channel).
   - Receive-only: <-chan int (You can only receive from this channel).
   - Example
   - This example shows a "producer" that only sends data and a "consumer" that only receives it.
   ```go
   package main

   import "fmt"
   
   // ping function only ACCEPTS a channel for SENDING.
   // The `chan<- string` syntax means it's a send-only channel.
   func ping(pings chan<- string, msg string) {
        pings <- msg
   }
   
   // pong function ACCEPTS one channel for RECEIVING and a second for SENDING.
   // `<-chan string` is a receive-only channel.
   func pong(pings <-chan string, pongs chan<- string) {
        msg := <-pings // Receive from pings
        pongs <- msg   // Send to pongs
   }
   
   func main() {
        pings := make(chan string, 1)
        pongs := make(chan string, 1)
   
        ping(pings, "passed message")
        pong(pings, pongs)
   
       fmt.Println(<-pongs)
   }
   ```
   - Output :
    ```
    passed message
     ```
    - Notice how the ping function can only send messages to the pings channel, and the pong function can only receive from pings and send to pongs. This makes it clear what each function is supposed to do with the channels.
    - If you tried to receive from the pings channel inside the ping function, the Go compiler would give you an error! This helps prevent bugs.<br><br>

4. **Closing Channels**:
   - When you're done sending values on a channel, you can close it using the built-in close function: close(myChannel).
   - Closing a channel is like telling everyone, "I'm done sending messages. No more will be coming."
   - After a channel is closed, you can still receive values from it until it's empty, but you cannot send any more values to it. If you try to send to a closed channel, your program will panic (crash).
   - You can check if a channel is closed when receiving by using the "comma ok" idiom: value, ok := <-myChannel. If ok is false, the channel is closed and empty.
   - Example
   ```go
   package main

   import "fmt"
   
   func main() {
      messages := make(chan string, 2)
   
      messages <- "Hello"
      messages <- "World"
      close(messages) // Close the channel when done sending
   
      // Receiving from the channel until it's empty
      for msg := range messages {
         fmt.Println(msg)
      }
   
      // Trying to receive again after the channel is closed
      msg, ok := <-messages
      if !ok {
         fmt.Println("Channel is closed and empty.")
      } else {
         fmt.Println(msg)
      }
   }
   ```
   - Output :
   ```
   Hello
   World
   Channel is closed and empty.
   ```
   - Notice how we closed the channel after sending our messages. The for loop received all messages until the channel was empty. The final receive attempt showed that the channel was closed and empty.<br><br>
   - other example
   ```go
   package main

   import "fmt"
   
   func main() {
        jobs := make(chan int, 3)
   
       // A goroutine to produce work
       go func() {
           for i := 1; i <= 3; i++ {
               fmt.Println("Sending job", i)
               jobs <- i
           }
           // We are done sending jobs, so we close the channel.
           close(jobs)
           fmt.Println("Sent all jobs and closed channel.")
	   }()
   
       // In the main goroutine, we receive work.
       // This `for range` loop will receive values from the `jobs`
       // channel until it is closed.
       fmt.Println("Waiting for jobs...")
       for job := range jobs {
           fmt.Println("Received job", job)
       }
       fmt.Println("All done!")
   }
   ```