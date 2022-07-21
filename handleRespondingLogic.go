package main

import "log"

func handleRespondingLogic(newUpdate DatabaseUpdate) {
	if newUpdate.Type == "bot_command" {
		command_name := newUpdate.Text[newUpdate.Offset:newUpdate.Length]
		switch command_name {
		case "/lembrete":
			log.Print("Lembrete!")
		}
	}
}