package main

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
)

var(
	NetIntarfaceExp = regexp.MustCompile(`"(.*?)"`)
)

type dnsList struct{
	Dns []dns
}

type dns struct{
	Name string `toml:"name"`
	IpAddr string `toml:"ip-addr"`
}

func readToml(path string) (string, error){
	ext := filepath.Ext(path)
	if ext != ".toml"{
		return "", errors.New(path + " is not .toml")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil{
		return "", err
	}
	return string(data), nil
}

// read the specified .toml and mapping the dnsList.
func readAndMap(path string) (*dnsList, error){
	tomlString, err := readToml(path)
	if err != nil{
		// if the dns file is noting, do noting.
		return nil, err
	}

	var d dnsList
	_, err = toml.Decode(tomlString, &d)
	return &d, nil
}

// read the specified .toml and look up the name.
func readAndLookup(path string, name string) (*dns, error){
	dnsList, err := readAndMap(path)
	if err != nil{
		return nil, err
	}

	// should use way of O(1).
	for _, d := range dnsList.Dns{
		if d.Name == name{
			return &d, nil
		}
	}

	return nil, errors.New("Not Found `" + name + "`")
}

func listDns(path string) (*dnsList, error){
	d, err := readAndMap(path)
	if err != nil{
		return nil, err
	}
	return d, nil
}

func parseInterface(s string) ([]string, error){
	return NetIntarfaceExp.FindAllString(s, -1), nil
}

func getInterfaces() ([]string, error){
	cmd := exec.Command("powershell", "netsh interface ip show config")
	b, err := cmd.Output()
	if err != nil{
		return nil, err
	}

	// TODO : check locale
	str, _, err := transform.String(japanese.ShiftJIS.NewDecoder(), string(b))
	if err != nil{
		return nil, err
	}

	interfaces, err := parseInterface(str)
	if err != nil{
		return nil, err
	}

	return interfaces, nil
}

func containsInterface(iname string) bool{
	infs, _ := getInterfaces()
	for _, inf := range infs{
		if iname == inf{
			return true
		}
	}
	return false
}

func enableDns(iname string, ipAddr string) error{
	// validate network interface
	// validate format of IPv4 address

	arg := `Start-Process netsh -ArgumentList 'interface ip set dns "` + iname + `" static "` + ipAddr + `" primary' -Verb runas`
	cmd := exec.Command("powershell", arg)
	if err := cmd.Run(); err != nil{
		return err
	}
	return nil
}

func disableDns(iname string) error{
	// validation network interface

	arg := `Start-Process netsh -ArgumentList 'interface ip set dns "` + iname + `" dhcp' -Verb runas`
	cmd := exec.Command("powershell", arg)
	if err := cmd.Run(); err != nil{
		return err
	}
	return nil
}