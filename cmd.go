package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

const(
	DefaultDnsPath = "dns.toml"
)


func CmdAddDns(c *cli.Context) error{
	// validate duplicated name...
	return errors.New("[Add] Not Implemented")
}

func CmdRemoveDns(c *cli.Context) error{
	// if name doesn't exist...
	return errors.New("[Remove] Not Implemented")
}

func CmdGetInterfaces(c *cli.Context) error{
	interfaces, err := getInterfaces()
	if err != nil{
		return err
	}

	fmt.Println("Network Interfaces:")
	for _, i := range interfaces{
		fmt.Println("  " + i)
	}
	return nil
}

func CmdEnableDns(c *cli.Context) error{
	iname := c.String("interface")
	var ipAddr string

	if c.String("dnsName") != ""{
		d, err := readAndLookup(DefaultDnsPath, "staging")
		if err != nil{
			return err
		}
		ipAddr = d.IpAddr
	}else{
		ipAddr = c.String("dnsAddr")
	}

	err := enableDns(iname, ipAddr)
	if err != nil{
		return err
	}
	return nil
}

func CmdDisableDns(c *cli.Context) error{
	err := disableDns("Wi-Fi")
	if err != nil{
		return err
	}
	return nil
}

func CmdListDns(c *cli.Context) error{
	d, err := listDns(DefaultDnsPath)
	if err != nil{
		return err
	}

	for _, s := range d.Dns{
		fmt.Println("[" + s.Name + "] " + s.IpAddr)
	}
	return nil
}