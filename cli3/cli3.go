package main

import (
    "github.com/urfave/cli"
    "os"
	"log"
	"fmt"
)

func main(){
//	var ar *string
//	ar = "cqzx"
//ar := [...]*string{&8,&9}
	var name string
	app := cli.NewApp()			//APP message
	app.Name   = "xxx"
	app.Usage  = "diss zcy"
	app.Version = "0.0.1"                              
	// app.Authors =[]&{"cqzx","aaa@bb.com"}			//!!!!!!!need to slove
	                                        		   //app.Email  = "aaa@bb.com"     V2 is delete
	




/*	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"xxx",
			Value:"Chinese",
			Usage:"language",
		},
	}

*/
app.Flags = []cli.Flag {              //
    &cli.StringFlag{                          
      Name: "lang",
      Value: "english",
	  Usage: "language for the greeting",
	  Destination:&name,
    },
  }


	app.Action = func(c *cli.Context)error{
		name:= "ssssx"
		if c.NArg()>0{
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		  return nil
	   }
	

    err:=app.Run(os.Args)
    if err!=nil{
        log.Fatal(err)
    }
}