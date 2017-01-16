package main

import (
    "log"
    "os/exec"
    "os"
)

func main() {
    // start reverse proxy
    caddy := exec.Command("caddy", "-conf", "Caddyfile.txt")
    caddy.Stdout = os.Stdout
    caddy.Stderr = os.Stderr
    if err := caddy.Start(); err != nil {
        log.Fatal(err)
    }
    defer caddy.Process.Kill()
    caddy.Wait()
}
