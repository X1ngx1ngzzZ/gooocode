package main

import (
    "github.com/urfave/cli"
    "os"
    "log"
)

func main(){
    err:=cli.NewApp().Run(os.Args)
    if err!=nil{
        log.Fatal(err)
    }
}