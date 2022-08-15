package main

import (
	"database/sql"
	"fmt"
	"log"

	"dynamicaller"
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

	if newUpdate.Type == "bot_command" {
		command_name := newUpdate.Text[newUpdate.Offset:newUpdate.Length]
		log.Print("Received command:", command_name)
		for i := range registeredCommands {
			if registeredCommands[i].Name == command_name {
				log.Print("Running command ", command_name)
				dynamicaller := new(dynamicaller.Dynamicaller)
				args := newUpdate.Text[newUpdate.Length:]
				res := dynamicaller.DynamicCall(dynamicaller, registeredCommands[i].Callback, args, newUpdate.ChatId)
				if err, ok := res[2].Interface().(error); ok && err != nil {
					fmt.Println(err)
				}
				break
			}
		}
	}
}