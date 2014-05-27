/* My First Go at Go! */
package main

import "fmt"
import "net"

func main() {
	// Get interfaces
	interfaces, err := net.Interfaces()

	// Interate over interfaces
	for index := 0; index < len(interfaces); index++ {
		if err == nil {
			// Grab the current interface address
			address, othererr := interfaces[index].Addrs()
			if othererr == nil {
				// Print out Address Information
				fmt.Printf("Interface: %s\tAddress: %s\n", interfaces[index].Name, address)
			}
		}
	}
}
