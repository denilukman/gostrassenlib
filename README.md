gostrassenlib
=======

Description
-----------

Go library for matrix multiplication using [Strassen Algorithm](https://en.wikipedia.org/wiki/Strassen_algorithm)

Usage
------------

This package can be installed with the go get command:

    go get github.com/sumpahpemuda/gostrassenlib

### Example

```
package main

import (
    "fmt"
    "github.com/sumpahpemuda/gostrassenlib"
)

func main() {
    B := [][]int{
        {1, 2},
        {3, 4},
        {5, 6},
        {7, 8},
    }

    A := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
    }

    C := gostrassenlib.Multiply(A, B)

    for i := 0; i < len(C); i++ {
        for j := 0; j < len(C[0]); j++ {
            fmt.Print(C[i][j], " ")
        }
        fmt.Println()
    }

}
```


ToDo
----

* Generic data type (float, etc)
* Add test/benchmark
* Add other todo lists (hmm this seems recursive)
