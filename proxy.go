package main

import (
    "log"
    "os/exec"
    "os"
    "github.com/gin-gonic/gin"
    "strings"
    "html/template"
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

    // start serving web app. Using this instead of polymer serve since
    // polymer doesn't support multiple apps.
    //tmpl := template.Must(template.ParseFiles("./index.html"))

    r := gin.Default()
    //r.SetHTMLTemplate(tmpl)
    r.GET("/spyfall/*path", func(ctx *gin.Context) {
        tmpl := template.Must(template.ParseFiles("./index.html"))
        tmpl.Execute(ctx.Writer, map[string]interface{}{
            "title": "Spyfall",
            "src": "spyfall/game-spyfall",
            "tag": template.HTML("<game-spyfall></game-spyfall>"),
        })
    })
    r.GET("/speed/*path", func(ctx *gin.Context) {
        tmpl := template.Must(template.ParseFiles("./index.html"))
        tmpl.Execute(ctx.Writer, map[string]interface{}{
            "title": "Speed",
            "src": "speed/game-speed",
            "tag": template.HTML("<game-speed></game-speed>"),
        })
    })
    r.NoRoute(func (ctx *gin.Context) {
        path := ctx.Request.URL.Path

        // is it an asset?
        if strings.Contains(path, ".") {
            ctx.File("." + path)
            return
        }

        // otherwise serve main application (which will 404 if the route doesn't exist)
        tmpl := template.Must(template.ParseFiles("./index.html"))
        tmpl.Execute(ctx.Writer,  map[string]interface{}{
            "title": "Boardless",
            "src": "my-app",
            "tag": template.HTML("<my-app></my-app>"),
        })
    })
    r.Run("0.0.0.0:8888")
}
