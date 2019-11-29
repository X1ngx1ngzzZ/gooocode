package main

import (
    "github.com/urfave/cli"
	"os"
	"log"
	"fmt"
)

func main(){
	app := cli.NewApp()
	app.Name="xxx"
	app.Usage="i dont know how to use"
	app.Action=func (c *cli.Context)error{       //error is an interface,In this interface,it just include one method (ERROR),return the method of string,error is used to return the wrong message in string.
		fmt.Println("output something")
		return nil 
	}
	err:=app.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}