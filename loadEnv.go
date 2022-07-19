package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)


func loadEnv() (EnvHolder){
	file, ferr := os.Open(".env")
	if ferr != nil {
		log.Fatalf("Could not find .env file, details: %s", ferr.Error())
	}

	scanner := bufio.NewScanner(file)
	counter := 0
	var_holder := EnvHolder{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		split_line := strings.Split(line, "=")
		if len(split_line) != 2 {
			log.Printf("Invalid format for line %d of file: \"%s\"", counter, line)
			continue
		}
		loaded_variable := EnvVar{
			name: split_line[0],
			val: split_line[1],
		}
		var_holder.vars = append(var_holder.vars, loaded_variable)
		counter++
	}
	return var_holder
}