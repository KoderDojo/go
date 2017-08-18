# Log

## August 17, 2017

Day 29 - Pushed my Todos WebAPI to Heroku and also used Postgres to store the tasks. Lots of learning and coding but nothing pushed to GitHub.

### Learned

* Heroku as a host for Go web applications.
* Postgres as a datastore for Go Applications using *lib/pq*.
* Use git to update and deploy Go Application on Heroku.


## August 16, 2017

Day 28 - Finished developing a Todos WebAPI in Go. Definitely easier to do with a framework, but it is good experience to do it manually.


## August 15, 2017

Day 27 - Continued developing a Todos WebAPI in Go. Created a TCP Server example just to post some code.

### Code Examples:

* [TCP Server](../practice/play/TCPServer/main.go)


## August 14, 2017

Day 26 - Started developing a Todos WebAPI in Go.


## August 13, 2017

Day 25 - No coding. Watched the Go courses on Udemy.


## August 12, 2017

Day 24 - Reading *Building Microservices with Go*.

### Code Examples:

* [Timeout Handler](../language/net-http/03-TimeoutHandler/main.go)
* [Redirect Handler](../language/net-http/04-RedirectHandler/main.go)
* [NotFound Handler](../language/net-http/05-NotFoundHandler/main.go)
* [Creating Handlers and Using Context](../practice/building-microservices/CreatingHandlersUsingContext/main.go)


## August 11, 2017

Day 23 - Reading *Building Microservices with Go*.

### Code Examples:

* [JSON Pretty Print](../language/encoding-json/PrettyPrint/main.go)
* [DefaultServeMux](../language/net-http/02-DefaultServeMux/main.go)
* [TCP Server](../language/net/01-TCP-Server/main.go)
* [Set Deadline on Connection](../language/net/02-Conn-Set-Deadline/main.go)


## August 10, 2017

Day 22 - Purchased *Building Micrososervices with Go*. Read the first section, *Introduction to Microservices*. Great review of the *net/http* and *encoding/json* packages I just reviewed.


## August 9, 2017

Day 21 - Re-reviewed the *net/http* package, but this time writing a bunch of code for practice. I will officially write the code tomorrow and push it to GitHub. Wish I had my whiteboard. Didn't bring it on the road to Utah. Not very portable anyway :)

### Re-Reviewed:

* *net/http* package
* ListenAndServe
* Handler Type
* HandleFunc
* HandlerFunc
* ServeMux
* DefaultServeMux


## August 8, 2017

Dat 20 - Reviewed JSON.

### Code Examples:

* [Encode](../language/encoding-json/Encode/main.go)
* [Decode](../language/encoding-json/Decode/main.go)
* [Marshal](../language/encoding-json/Marshal/main.go)
* [Unmarshal](../language/encoding-json/Unmarshal/main.go)


## August 7, 2017

Day 19 - Reviewed the net/http package in the Udemy Go Web Development Course. No coding. I just enjoyed re-watching the videos.

### Reviewed:

* *net/http* package
* ListenAndServe
* Handler Type
* HandleFunc
* HandlerFunc
* ServeMux
* DefaultServeMux


## August 6, 2017

Day 18 - Started building a static website generator. Wrote numerous examples of using the [text/template](../language/text-template) package in Go.

### Code Examples

* [Basic](../language/text-template/01-Basic)
* [Struct](../language/text-template/02-Struct)
* [Map with Struct](../language/text-template/03-Map-With-Struct)
* [Slice Using Range](../language/text-template/04-Slice-Using-Range)
* [Read and Write Files](../language/text-template/05-Read-Write-Files)
* [Read Data as JSON from File](../language/text-template/06-Read-Data-As-Json-From-File)


## August 5, 2017

Day 17 - Watched my Udemy Web Development Course and coded in JetBrains Gogland.

### Reviewed:

* *net/http* package
* Handler Interface - ServeHTTP(w ResponseWriter, r *Request)
* http.ListenAndServe
* *text/template* package
* Request Structure Fields

Wrote a lot of practice code, but pushed none of it to GitHub. I'll be better about that in the future so I can get a better feel for what I am doing each day.


## August 4, 2017

Day 16 - Watched several presentations on YouTube about common concurrency patterns and advanced concurrency patterns. Added several resources.


## August 3, 2017

Day 15 - Reviewed Channels and common patterns. Watched a bit more of the Udemy Go Course as well as read a few articles. No code today.


## August 2, 2017

Day 14 - Coded a number of goroutines for practice. Used WaitGroup and Channels.

### Code

* [Hello World](../language/goroutines/01-Hello-World/main.go)
* [Hello World with Sleep](../language/goroutines/02-Hello-World-Sleep/main.go)
* [Run Concurrently](../language/goroutines/03-Run-Concurrently/main.go)
* [WaitGroup](../language/goroutines/04-WaitGroup/main.go)
* [Channels](../language/goroutines/05-Channels/main.go)


## August 1, 2017

Day 13 - Reviewed *text/template* and *html/template* packages a well as created a **TCP Server** and started using the *net/http* package. Also attended the monthly Go Meetup in Lehi, Utah.

* [Attended Utah Go User Group](https://www.meetup.com/utahgophers)
* Reviewed *text/template* package
* Reviewed *html/template* pacakge
* Created TCP Server and Client
* Used: **Dial**, **Listen**, **Accept**, **Close** as well as *bufio* package

Learned about creating DSL's at the Utah Go User Group.


## July 31, 2017

Day 12 - I solved 2 more programming challenges on Firecode.io that I had done in C as well, and I learned more about **benchmarking Go code** as well as using the *text/template* package.

### Code Written with Tests:

* [No Twins Allowed](../practice/firecodeio/no-twins-allowed.go)
* [Is it a Palindrome](../practice/firecodeio/palindrome-tester.go)
* [Benchmarking Selection Sort](../algorithms/sort/selection-sort.go)

Really enjoying the **golang-newbies** slack channel.


## July 30, 2017

Day 11 - Focused on writing code and tests.

### Code Written with Tests:

* [Reverse a String](../practice/firecodeio/reverse-a-string.go)
* [Even or Odd (linked list)](../practice/firecodeio/even-or-odd.go)
* [Selection Sort](../algorithms/sort/selection-sort.go)

I also wrote a **benchmark** for the Selection Sort, but I have all kinds of questions and research I need to do about how to benchmark an *in-place* sorting algorithm that also requires some initial set-up.

Firecode.io doesn't support Go, but I took the C challenges I did today in Firecode and solved a couple in Go, too.


## July 29, 2017

Day 10 - Switched gears and started looking at the "text/template" package in Go, which is super cool!

### Reviewed:

* Parsing and executing templates
* Passing data to templates
* Adding variables in templates
* Adding functions in templates
* Nesting templates


## July 28, 2017

Day 9 - Continued looking at concurrency and testing, and peeked at error handling.

### Reviewed:

* Mutex
* Atomicity
* Pipline pattern with Channels
* Handling errors and logging to file.
* [Table-Driven Tests](../notes/table-driven-tests.md).
* Benchmarking Code.


## July 27, 2017

Day 8 - Learned how to use the **testing** package in Go to unit test. Also practiced with **ginkgo** and **gomega**. Received some guidance on idiomatic testing on Gophers golang-newbies slack channel.

### Reviewed:

* Testing package
* Ginkgo and gomega
* Codewars Challenge: [Two Oldest Ages](../practice/codewars/kyu-7-two-oldest-ages.go)


## July 26, 2017

Day 7 - Wow! Big day! Learned interfaces, concurrency, and parallelism as well as did a couple of programming challenges.

### Reviewed:

* Interfaces
* Goroutines
* WaitGroup
* Channels
* Using range with channels
* Self-invoking anonymous functions

Two really easy programming challenges, but always learn something about Go: [Even or Odd](../practice/codewars/kyu-8-even-or-odd.go) and [Multiply](../practice/codewars/kyu-8-multiply.go)


## July 25, 2017

Day 6 - Learning Structs and Functions. I also started solving programming challenges on Codewars using Go.

### Reviewed:

* Structs
* Functions
* Codewars Programming Challenge: [String Repeat](../practice/codewars/kyu-8-string-repeat.go)

Unlike C, a **struct can be a receiver of a function**, which is really cool! Through the simple Codewars challenge, I learned about the **strings** package in Go, which has a lot of handy functions!

Although I simply labeled the review **Structs** and **Functions**, there was a lot to learn.


## July 24, 2017

Day 5 - Learning data structures. Although I mentioned using Array on Day 4, I am officially learning it and realizing I was technically using a Slice :)

### Reviewed:

* Array
* Slice ( a list )
* Map ( a dictionary )
* Make ( to create reference types )
* Supported functions: len, cap, range, etc.


## July 23, 2017

Day 4 - Wrote a [bubble sort algorithm](../algorithms/sort/bubble-sort.go). First time using arrays in Go, which I haven't reviewed yet and may not be using idiomatically. Used a lot of language features discussed thus far:

* Arrays
* `For` Loop with just a condition.
* `For` Loop with initialization, condition, and post statement.
* Shorthand variable initialization.
* `If` statement.
* `Break` statement.


## July 22, 2017

Day 3 - Reviewed Control Flow. Lots of interesting differences to the `For` Loop and `Switch` Statement. I will mention this in my notes at another time.

### Reviewed:

* For Loop.
* Conditions, Break, and Continue.
* Switch Statements.
* If Statements.


## July 19, 2017

Day 2 - Learning the Go language fundamentals. Learning C this summer paid off, because Go has many similarities!

### Reviewed:

* Packages and member visibility.
* Scope.
* Different ways to create variables.
* Constants.
* Blank Identifier ( similar to Python ).
* Pointers ( similar to C ).
* Used **go fmt**, which I did learn the other day as well.


## July 18, 2017

Day 1 - At the point where I can develop with Go!

* Installed Go.
* Created Go workspace.
* Installed Golang Plugin in JetBrains Webstorm.
* Installed JetBrains Gogland.
* Created [Hello World](../practice/play/hello-world) Program.
* Learned **go run**, **go build**, **go install**

