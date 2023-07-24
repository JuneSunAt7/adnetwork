package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

func local() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Network interface: %s\n", interf.Name)

		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("[+] Interface is work\n")
				fmt.Printf("\tIP: %v\n", ip)
			}
		}
	}
}
func ping() {
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && ping -n 3 www.google.com ")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func tracert() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && tracert -h 10 www.google.com ")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func ipconfigall() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && ipconfig/all ")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func ipconfigdns() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && ipconfig/displaydns ")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func getmac() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && getmac ")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func process() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && netstat -na")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func nethcheck() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && netsh wlan show drivers ")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func statinterface() {
	fmt.Printf("Wait please...")
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 &&  netstat -e")

	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
func fullcheck() {
	fmt.Printf("Wait please...")
	go nethcheck()
	go process()
	go getmac()
	ipconfigall()
	ipconfigdns()
	go tracert()
	ping()
	local()
	statinterface()
}
func main() {

	fmt.Printf("Choose option: \n")
	fmt.Printf("\n[1]\t netcheck(checking driwers for wlan)\n[2]\t process(checking connection in 80 port)\n")
	fmt.Printf("[3]\t getmac(get mac address)\n[4]\t ipconfigall(checking detailed info about adapters)\n")
	fmt.Printf("[5]\t ipconfigdns(checking dns list)\n[6]\t tracert(checking lost packets)\n")
	fmt.Printf("[7]\t ping(ping for google.com)\n[8]\t local(show all local network interfaces)\n")
	fmt.Printf("[9]\t statinter(statistics about interfaces)\n\n[99]\t full(full check)\n")

	var opt int
	fmt.Scan(&opt)

	switch opt {
	case 1:
		nethcheck()
	case 2:
		process()
	case 3:
		getmac()
	case 4:
		ipconfigall()
	case 5:
		ipconfigdns()
	case 6:
		tracert()
	case 7:
		ping()
	case 8:
		local()
	case 9:
		statinterface()
	case 99:
		fullcheck()
	}
}
