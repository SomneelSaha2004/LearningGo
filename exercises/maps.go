//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayStatus(m map[string]int) {
	servers := map[string]int{
		"Offline    ": 0,
		"Maintenance": 0,
		"Retired    ": 0,
		"Online":      0,
	}
	for _, value := range m {
		switch value {
		case 0:
			servers["Online"]++
		case 1:
			servers["Offline"]++
		case 2:
			servers["Maintenence"]++
		case 3:
			servers["Retired"]++
		}
	}
	fmt.Println("No of servers :", len(m))
	for key, value := range servers {
		fmt.Println(key, "->", value)
	}
	fmt.Println()
}
func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	m := make(map[string]int)
	for _, val := range servers {
		m[val] = 0
	}
	displayStatus(m)
	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	//  - call display server info function
	//  - change server status of all servers to `Maintenance`
	//  - call display server info function
	m["darkstar"] = 3
	m["aiur"] = 1
	displayStatus(m)
	for key, _ := range m {
		m[key] = 2
	}
	displayStatus(m)
}
