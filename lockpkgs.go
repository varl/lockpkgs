package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    fmt.Printf("Let's do stupid stuff.\n")

    file, err := os.Open("example_package.json")

    if err != nil {
        panic(err)
    }

    defer file.Close()

    // just decode all the shit into this shit
    var pkgs map[string]interface{}

    err = json.NewDecoder(file).Decode(&pkgs)

    if err != nil {
        panic(err)
    }

    // hey we got a lot of shit!
    for k, v := range pkgs {
        if k == "devDependencies"  || k == "dependencies" {
            fmt.Println("Value:", v)
            deps := v.(map[string]interface{})
            for module, version := range deps {
                deps[module] = "changed"
                fmt.Println("Module:", module, "Version:", version)
            }
            break
        }
    }

    enc := json.NewEncoder(os.Stdout)
    enc.Encode(pkgs)
}
