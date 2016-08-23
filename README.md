# deref
Provides convenience functions to extract the value from a pointer, returning the appropriate zero value if the pointer is nil. 

Example:
```Go
package main

import (
        "fmt"

        "github.com/AlexSteele/deref"
)

func main() {
        var i *int
        fmt.Println(deref.Int(i)) // Prints '0'
}
```
