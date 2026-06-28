package main

import (
        "fmt"
        "log"
        "os"
        "goserver_recon/app"
)

func main() {
        fmt.Println("WebServer Informations:\n")

        aplicacao := app.Gerar()
        if erro := aplicacao.Run(os.Args); erro != nil { // <- corrige apicacao
                log.Fatal(erro)
        }
}
