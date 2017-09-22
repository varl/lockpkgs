package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Packages struct {
    DevDeps map[string]string `json:"devDependencies"`
    Deps map[string]string `json:"depedencies"`
}

func main() {
    fmt.Printf("Let's do stupid stuff.\n")

    file, err := os.Open("example_package.json")

    if err != nil {
        panic(err)
    }

    defer file.Close()

    var pkgs Packages

    err = json.NewDecoder(file).Decode(&pkgs)

    if err != nil {
        panic(err)
    }
}
