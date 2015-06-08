https://www.youtube.com/watch?v=QDDwwePbDtw


CAR Hoare : CSP

c := make(chan bool) - makes unbuffered channel of booleans
c <- x - send on channel
<- c receive on channel
x <- c receive on channel and store in x
x, ok <- c - receive; ok is false if channel is closed and empty
                      x is zero in this case. closed channels never block.

Unbuffered Channels : Communication + Synchronization

#Signalling with Unbuffered Channels

##Wait for event

c := make(chan bool)

go func() {
    ... do func stuff
    close(c)
}()

.. do main stuff

<- c

##Coordinate multiple workers

func worker(start chan bool) {
    <- start
    // .. do worker stuff
}

func main() {
    start := make(chan bool)

    for i := 0; i < 100; i++) {
        close(start)
        // .. all workers running now
    }
}

##Listen to multiple channels

select {
    case x := <- somechan: 
        //
    case y, ok := <- someotherhcan:
        //
    case outputChan <- z:
        //...
    default:
        // nothing to do
}

##Listen Forever

for {
    select {

    }
}

##Terminate Workers
func worker(die chan bool) {
    for {
        select {
            // ... do stuff
        case <- die:
            return
    }
}

func main() {
    die := make(chan bool)
    for (i := 0; i < 100; i++) {
        go worker(die)
    }

    // do something
    close(die)
}

##Single Worker Death Verification
func worker(die chan bool) {
       for {
        select {
            // ... do stuff
        case <- die:
            // do termination
            die <- true
            return
    } 
}

func main() {
    die := make(chan bool)
    go worker(die)
    die <- true
    <- die
}

#Buffered Channels
##Closing
func main() {
    c := make(chan int, 3) // buffer of 3
    c <- 15
    c <- 32
    close(c)
    fmt.Printf("%d\n", <-c)
    fmt.Printf("%d\n", <-c)
    fmt.Printf("%d\n", <-c) // 0
}

#Generators
id := make(chan string, 3)

go func() {
    var counter int64 = 0
    for {
        id <- fmt.Sprintf("%x", counter)
        counter += 1
    }
}()

#Queues
idle := make(chan []byte, 5)
select {
    case b <- idle
    // recycle
    default:
    // let GC handle
}

#Nil Channels
##Blocking
func main() {
    var c chan bool
    <- c
}

##select ??

for {
    select {
    case x, ok := <- c1:
        if !ok {
            c1 = nil
        }
    }

    case x, ok := <- c2:
        if !ok {
            c2 = nil
        }
    }
    if c1 == nil && c2 == nil { 
        return
    }
}

##Timers
http://blog.golang.org/go-concurrency-patterns-timing-out-and

##Timeout
func worker(start chan bool) {
    for {
        timeout := time.After(30 * time.Seconds)
        select {

        }
    }
}

#Pool


#Cleanup

#Heartbeat

#Multiplexor

http://golang.org/doc/effective_go.html#chan_of_chan
http://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement

http://play.golang.org/p/8zwvSk4kjx

func worker(messages chan string) {
    for {
        var msg string
        messages <- msg
    }
}

func main() {
    ...
    for {
        msg := <- messages
        conn.Write([]byte(msg))
    }
}

#FirstOfN
type response struct {
    resp *http.Response
    url string
}

func get(url string, r chan response) {
    if resp, err := http.Get(url); err == nil {
        r <- response(resp, url)
    }
}

func main() {
    first := make(chan response)

    for _, url := range []string{"..", "..", ".."} {
        go get(url, first)
    }

    r := <- first
}

#Workers with Return Address

type work struct {
    url string
    resp chan *http.Response
}

func get(w chan work) {
    for {
        do := <- w
        resp, _ := http.Get(do.url)
        do.resp <- resp
    }
}

func main() {
    w := make(chan work)

    go getter(w)

    resp := make(chan * http.Response)

    w <- work{"http://cdnjs.cloudflare.com/jquery..", resp}

    r := <- resp
}

##Parallel Work
rob pike's talk

##Parallel Work with timeout
rob pike's talk

##map reduce

##Worker Pool
http://stackoverflow.com/questions/18405023/how-would-you-define-a-pool-of-goroutines-to-be-executed-at-once-in-golang

##Double Select
http://golang.org/test/chan/doubleselect.go

##Sieve
http://golang.org/test/chan/sieve2.go

##Data Streams 

[Squinting at Power Series](http://www.cs.bell-labs.com/who/rsc/thread/squint.pdf)
http://golang.org/test/chan/powser1.go
http://golang.org/test/chan/powser2.go

http://blog.golang.org/pipelines
http://talks.golang.org/2013/advconc.slide
[Power Series / Power Serious - Haskell](http://www.cs.dartmouth.edu/~doug/pearl.ps.gz)
