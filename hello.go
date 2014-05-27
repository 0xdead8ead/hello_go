package main

import "fmt"
import "net"

/*

My First Go App

*/

func main() {
	fmt.Printf("Hello, world.\n")
	//addresses, err := net.InterfaceAddrs()
	interfaces, err := net.Interfaces()

	for index := 0; index < len(interfaces); index++ {
		if err == nil {

			address, othererr := interfaces[index].Addrs()
			if othererr == nil {
				fmt.Printf("\nInterface: %s\tAddress: %s", interfaces[index].Name, address)
			}
		}
	}
	fmt.Printf("%s", err)
}
