package main

import (
	"testing"
)

func TestListDns_FileIsNotToml(t *testing.T){
	_, err := listDns("nothing.yaml")
	if err == nil{
		t.Fatal("")
	}
}

func TestListDns_FileNothing(t *testing.T){
	_, err := listDns("nothing.toml")
	if err == nil{
		t.Fatal("")
	}
}

func TestListDns(t *testing.T){
	_, err := listDns("dns.toml")
	if err != nil{
		t.Fatal(err)
	}
}

func TestGetInterfaces_JPN(t *testing.T) {
	_, err := getInterfaces()
	if err != nil{
		t.Fatal(err)
	}
}

func TestGetInterfaces_ENG(t *testing.T) {
	_, err := getInterfaces()
	if err != nil{
		t.Fatal(err)
	}
}

