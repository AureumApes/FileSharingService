package commands

import "fmt"

func Help() {
	
	fmt.Println("Command\t\tEffect")
	fmt.Println("help/h\t\tthis\nls/list\t\tlist every file on the server\nget/g\t\tload the content of an file by its id\ne/exit\t\tCloses the Client")
}
