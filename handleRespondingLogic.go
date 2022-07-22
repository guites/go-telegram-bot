package main

import (
	"database/sql"
	"log"
)

func getCommandsFromDb(db *sql.DB) ([]DatabaseCommand){
	rows, err := db.Query("SELECT * FROM commands ORDER BY Id desc;");
	
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	commands := make([]DatabaseCommand, 0)
	for rows.Next() {
		dbCommand := DatabaseCommand{}
		err = rows.Scan(&dbCommand.Id, &dbCommand.Name, &dbCommand.Callback)
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, dbCommand)
	}
	return commands

}

func handleRespondingLogic(db *sql.DB, newUpdate DatabaseUpdate) {

	registeredCommands := getCommandsFromDb(db)
	log.Print(registeredCommands)
	if newUpdate.Type == "bot_command" {
		command_name := newUpdate.Text[newUpdate.Offset:newUpdate.Length]
		switch command_name {
		case "/lembrete":
			log.Print("Lembrete!")
		default:
			log.Print("Comando não encontrado!")
		}
	}
}