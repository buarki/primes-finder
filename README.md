# primes finder

## WebAssembly

This project was created to check the state of [WebAssembly](https://webassembly.org/) development tools now, in the beginning of 2024, specifically with Go programing language.

So far, I was able to put this project to run by digging into the [provided documentation](https://github.com/golang/go/wiki/WebAssembly#getting-started), but it was not a straightforward process, especially in terms of passing complex objects from Go to Javascript level. I'll need to study more it :)

## Sieve of Eratosthenes

The [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes) is an algorithm very useful for finding prime numbers up to a given limit. For instance, if you ask for all the prime numbers up to 12 you'll have 2, 3, 5, 7 and 11.

The ideia is very simple:
- we create a list of numbers with the size of the given limit + 1 (just to be able to represent the numbers using indexes);
- we assume that all numbers in such a list are truly prime numbers;
- as we know that 0 and 1 are not prime numbers we mark them as false;
- then, we iterate from 2 up to the square root of the given limit checking if the current number can divide the limit, if so we mark all multiples of the current number as false;
- and then we repeat this process;

The implemention of this algorithm can be found at the [find_primes.go](cmd/wasm//find_primes.go), such implementation has a time complexity of **O(n log log n)** and space complexity of **O(n)**.

## How this project works?

This project has two main parts:
- the WebAssembly binary [written in go](/cmd/wasm/main.go);
- the [Javascript client](public/index.html) that uses the WebAssembly binary;

### The Go part

The [Go part](/cmd/wasm/main.go) is a simple program that makes a Go function called **findPrimes** available at the Javascript level when running in a WebAssembly environment. This is done bellow line:

```go
js.Global().Set("findPrimes", findPrimes())
```

It also defines simple channel to prevent the program from exiting immediately. This is typical in WebAssembly programs where the Go runtime does not automatically wait for asynchronous tasks. This is done by bellow line:

```go
<-make(chan bool)
```

To build this project the prepared command **gowasm* at the [Makefile](./Makefile) can be used, and behind the scenes it does bellow:

```go
GOARCH=wasm GOOS=js go build -o public/main.wasm cmd/wasm/*.go
```

And another important file is the **wasm_exec.js**, already provided by Go, thus we just need to copy it. The Makefile also has the command **wasm_exec** to do it out of the box, but behind the scenes it runs this:

```go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" public/wasm_exec.js
```

With files **wasm_exec.js** and **main.wasm** available we can load them on the Javascript client to call the function **findPrimes**.

Everytime the button **Find** is pressed a function is called to collect the limit given by user to validate it and if valid to execute **findPrimes**. With the found prime numbers available, they are rendered as an unordered list.

### The Javascript part

This one is relativelly simple, [it just an HTML file with inline Javascript](public/index.html) to receive the limit to be used to perform the Sieve of Eratosthenes and call the function **findPrimes** provided by Go.

## Running Locally

This repository has a [Makefile](./Makefile) with a `local` command that will be in charge of building needed stuff and starting the server on port 8080, to use just run bellow command:

```sh
make local
```

## The deploy to production using Vercel

Vercel now [supports Go functions](https://vercel.com/docs/functions/serverless-functions/runtimes/go), and due to the easy integration this project is deployed [there](https://primes-finder.vercel.app/).

There's just one small note on this: in order to make the files [public/main.wasm](./public/main.wasm) and [/public/wasm_exec.js](/public/wasm_exec.js) available from [public/index.html](public/index.html) we need to strip the prefix **/public**, due to that the script [adjust-index.sh](./adjust-index.sh) was built and is configured to be run on Vercel during the release.
