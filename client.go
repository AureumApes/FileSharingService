package main

import (
	"FileSharingClient/commands"
	"fmt"
)

func main() {
	commands.Clear()
	Commands()
}

func Commands() {
	fmt.Print("AmiraxobaFS >> ")
	var cmd string
	fmt.Scanln(&cmd)
	switch cmd {
	case "ls", "list":
		commands.ListFiles()
		break
	case "clear", "c":
		commands.Clear()
		break
	case "get", "g":
		commands.GetFile()
		break
	case "h", "help":
		commands.Help()
	case "exit", "e":
		return
	default:
		commands.Clear()
		fmt.Println("Unknown command '" + cmd + "' use 'h' or 'help' for a list of commands")
		break
	}
	Commands()
}
