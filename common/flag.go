// Author : TeamsSix
// Blog : teamssix.com
// WeChat Official Accounts: TeamsSix

package common

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Flag() (string, string, string, string, bool, bool) {
	logo()
	var username string
	var password string
	var cloneuser string
	var deleteusername string
	var check bool
	var onlycreate bool
	var version bool
	flag.StringVar(&username, "u", "", "Set username, If the username does not end with a $ sign, a $ sign will be added automatically")
	flag.StringVar(&password, "p", "", "Set password")
	flag.StringVar(&deleteusername, "d", "", "Set delete username, If the username does not end with a $ sign, a $ sign will be added automatically")
	flag.BoolVar(&check, "c", false, "Check the hidden accounts of the current system")
	flag.StringVar(&cloneuser, "cu", "Administrator", "Set clone user")
	flag.BoolVar(&onlycreate, "oc", false, "Only create hidden users, do not clone users by modifying the registry")
	flag.BoolVar(&version, "v", false, "View version")
	flag.Parse()

	CheckUserName(username)
	CheckUserName(deleteusername)
	CheckUserName(cloneuser)
	if version {
		fmt.Println("[+] Current Version: v0.2")
		os.Exit(3)
	}
	if check {
		return "", "", "", "", false, check
	} else if username == "" && password == "" && deleteusername == "" {
		flag.Usage()
		os.Exit(3)
	} else if username == "" && deleteusername == "" {
		fmt.Println("[-] Username cannot be empty. Enter -h to view the help information")
		os.Exit(3)
	} else if password == "" && deleteusername == "" {
		fmt.Println("[-] Password cannot be empty. Enter -h to view the help information")
		os.Exit(3)
	}
	if username != "" {
		if string(string(username[len(username)-1:len(username)])) != "$" {
			username = username + "$"
		}
	}
	if deleteusername != "" {
		if string(deleteusername[len(deleteusername)-1:len(deleteusername)]) != "$" {
			deleteusername = deleteusername + "$"
		}
	}

	return username, password, deleteusername, cloneuser, onlycreate, check
}

func CheckUserName(username string) {
	WhiteList := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890$_"
	k := 0
	for _, i := range strings.Split(WhiteList, "") {
		for _, j := range strings.Split(username, "") {
			if j == i {
				k++
			}
		}
	}
	if k != len(strings.Split(username, "")) {
		fmt.Println("[-] Invalid user name.")
		os.Exit(3)
	}
}

func logo() {
	fmt.Println(`
╔═╗┬─┐┌─┐┌─┐┌┬┐┌─┐  ╦ ╦┬┌┬┐┌┬┐┌─┐┌┐┌  ╔═╗┌─┐┌─┐┌─┐┬ ┬┌┐┌┌┬┐
║  ├┬┘├┤ ├─┤ │ ├┤   ╠═╣│ ││ ││├┤ │││  ╠═╣│  │  │ ││ ││││ │ 
╚═╝┴└─└─┘┴ ┴ ┴ └─┘  ╩ ╩┴─┴┘─┴┘└─┘┘└┘  ╩ ╩└─┘└─┘└─┘└─┘┘└┘ ┴  v0.2
                       Team: WgpSec
                     Author: TeamsSix
                    Blog: teamssix.com                    
             WeChat Official Accounts: TeamsSix            
   Project Address: github.com/wgpsec/CreateHiddenAccount  
`)
	fmt.Println("[!] 请勿将工具用于非法用途，开发人员不承担任何责任，也不对任何滥用或损坏负责。")
	fmt.Println("[!] Do not use the tool for illegal purposes, the developer is not responsible, nor responsible for any misuse or damage.\n")
}
