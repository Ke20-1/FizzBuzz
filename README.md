# FizzBuzz

This Service exposes a API service of Fizzbuzz algorithm


### Installation

FizzBuzz uses the Go Modules coming from the 1.11 version of the language to deal with dependencies.

```shell
$ git clone https://github.com/Ke20-1/FizzBuzz.git
```

### Development environment

As you can, on the root of the project there are one file: `Dockerfile`.
As implied by their name, if you run the following commands, you create an environment of Docker containers:

```shell
$ docker build -t my-go-fb .
$ docker run -it -p 8080:8080 my-go-fb
```

### Tests

```shell
$ cd tests
$ go build .
$ ./tests
```
