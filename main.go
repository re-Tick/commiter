package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("/bin/bash", "commiter.sh")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error executing script: %s\n", err)
        return
    }
    fmt.Printf("Script output:\n%s\n", string(output))
}
