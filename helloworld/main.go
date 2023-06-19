package main

import "fmt"

func main() {
    student := map[string]int {
        "a": 20,
        "b": 30,
    }
    delete(student, "b")
    fmt.Println(student)
}