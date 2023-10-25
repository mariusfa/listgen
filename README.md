# Listgen
A generic list utility package for Go. Heavily inspired by Kotlin.

## Overview
`listgen` provides a collection of utility functions and data structures to work with generic lists in Go. Built to leverage the new generics feature introduced in Go 1.18, genlist offers type-safe operations to manipulate and query lists.

## Features
- Filter: Filter elements based provided predicate function.
- Map: Map elements based on provided list and mapping function.
- Add: Add elements to the list.

## Installation
```bash
go get github.com/mariusfa/listgen
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/yourusername/genlist"
)

func main() {
    // Create a list
    list := genlist.List[int]{1, 2, 3, 4, 5}
    
    // Filter the list
    evens := list.Filter(func(i int) bool {
        return i%2 == 0
    })
    
    fmt.Println(evens) // Outputs: [2 4]
}
```

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

## License
`listgen` is licensed under the MIT License.