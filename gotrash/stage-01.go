//Works, but requires the go interpreter to be installed on the client :/
package main

import (
"github.com/novalagung/golpal"
"io/ioutil"
"fmt"
)

func main() {
    cmdString, _ := ioutil.ReadFile("gotrash.go")

    output, err := golpal.New().ExecuteRaw(string(cmdString))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("result", "=>", output)
}
