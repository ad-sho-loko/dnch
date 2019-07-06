package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var commands = []cli.Command{
	{
		Name:   "add",
		Usage:  "Add the new DNS Server",
		Action: CmdAddDns,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dnsName",
				Usage: "DNS name to add",
			},
			cli.StringFlag{
				Name:  "ip",
				Usage: "IP Address for DNS ",
			},
		},
	},
	{
		Name:    "remove",
		Usage:   "Delete the registerd DNS Server",
		Action:  CmdRemoveDns,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dnsName",
				Usage: "DNS name to delete",
			},
		},
	},
	{
		Name:    "list",
		Usage:   "Show the list of registerd DNS servers",
		Action:  CmdListDns,
	},
	{
		Name:    "interface",
		Usage:   "Show the current list of network interface",
		Action:  CmdGetInterfaces,
	},
	{
		Name:    "serve",
		Usage:   "Serve the api server",
		Action:  CmdServe,
	},
	{
		Name:    "enable",
		Usage:   "Enable the specified DNS server",
		Action:  CmdEnableDns,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "interface",
				Usage: "Network interface name to apply DNS.",
			},
			cli.StringFlag{
				Name:  "dnsName",
				Usage: "DNS name to identify",
			},
			cli.StringFlag{
				Name:  "dnsAddr",
				Usage: "DNS IP address ",
			},
		},
	},
	{
		Name:    "disable",
		Usage:   "Disable the DNS setting, it changes Default",
		Action:  CmdDisableDns,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "iname",
				Usage: "network interface to disable",
			},
		},
	},
}

func Run(args []string) error{
	app := cli.NewApp()
	app.Name = "dnch"
	app.Version = "0.0.1"
	app.Usage = "A Simple DNS Changer"
	app.Commands = commands
	return app.Run(args)
}

func main(){
	if err := Run(os.Args); err != nil{
		log.Fatal(err)
	}
}
