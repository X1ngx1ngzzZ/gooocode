package main

import (
    "github.com/urfave/cli"
    "os"
	"log"
	"fmt"
)

func main(){

	/*
	app := cli.NewApp()			//APP message
	app.Name   = "xxx"
	app.Usage  = "diss zcy"
	app.Version = "0.0.1"                              
	*/
	app:=&cli.App{

	Name : "xxx",		
	Usage: "diss zcy",
	Version : "0.0.1", 


	//FLAG
	//app.Flags = []cli.Flag {              //every Flag has a cli.flag,
	Flags:[]cli.Flag{
    &cli.StringFlag{                          
      Name: "lang,l",                    //l is shorthand
      Value: "english",						
      Usage: "language for the greeting",
	},
	&cli.StringFlag{
		Name: "country,cou",                    //l is shorthand
		Value: "China",						
		Usage: "location for the greeting",
	},
   },


	//command
	Commands:[]*cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action:  func(c *cli.Context) error {
			  return nil
			},
		},

		{
			Name:    "change",
			Aliases: []string{"gai","huan","zhuan"},
			Usage:   "change a task on the list",
			Action:  func(c *cli.Context) error {
			  return nil
			},
		},
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
//Command 'gp' not found, but can be installed with:

//sudo apt install pari-gp
