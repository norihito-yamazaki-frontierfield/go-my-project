package main

import "fmt"

func main() {
    fmt.Println(Hello("world", ""))
    fmt.Println(Hello("", ""))
}

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
    finalName := name
    if finalName == "" {
        finalName = "World"
    }

    if language == spanish {
        return spanishHelloPrefix + finalName
    }
    if language == french {
        return frenchHelloPrefix + finalName
    }
    return englishHelloPrefix + finalName
}