```go
package main

import "fmt"

func main() {

    // Initialize the map before use
    m := make(map[string]int)
    m["key"] = 10

    fmt.Println(m["key"])
}
```