package main

import (
    "log"
    "os/exec"
    "time"
    "os"
)

func main() {
    polymer := exec.Command("polymer", "serve")
    polymer.Stdout = os.Stdout
    polymer.Stderr = os.Stderr
    if err := polymer.Start(); err != nil {
        log.Fatal(err)
    }
    defer polymer.Process.Kill()

    caddy := exec.Command("caddy", "-conf", "Caddyfile.txt")
    caddy.Stdout = os.Stdout
    caddy.Stderr = os.Stderr
    if err := caddy.Start(); err != nil {
        log.Fatal(err)
    }
    defer caddy.Process.Kill()

    time.Sleep(24 * time.Hour)
}
