```go
package main

import "fmt"

func main() {

    // This is the problematic code. It uses an uninitialized map.
    var m map[string]int
    m["key"] = 10 // This will panic

    fmt.Println(m["key"])
}