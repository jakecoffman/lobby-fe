package main

import (
    "net/http/httputil"
    "log"
    "net/url"
    "github.com/gin-gonic/gin"
    "os/exec"
)

func main() {
    cmd := exec.Command("polymer", "serve")
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    defer cmd.Process.Kill()

    fe, _ := url.Parse("http://localhost:8080")
    be, _ := url.Parse("http://localhost:8444")

    fep := httputil.NewSingleHostReverseProxy(fe)
    bep := httputil.NewSingleHostReverseProxy(be)

    r := gin.Default()
    r.Any("/api/*path", func(ctx *gin.Context) {
        ctx.Request.URL.Path = ctx.Param("path")
        bep.ServeHTTP(ctx.Writer, ctx.Request)
    })
    r.NoRoute(func(ctx *gin.Context) {
        fep.ServeHTTP(ctx.Writer, ctx.Request)
    })
    log.Println("FRONTEND:", fe)
    log.Println("BACKEND:", be)
    log.Println("SERVING http://localhost:8111")
    r.Run("0.0.0.0:8111")
}
