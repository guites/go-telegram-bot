package main

import "log"

// EnvHolder holds all variables from .env file into memory
type EnvHolder struct {
	vars []EnvVar
}

// EnvVar is a key => value pair representing one variable from the .env file
type EnvVar struct {
	name string
	val string
}

// Returns the value for a given variable loaded from .env file by name
func (x EnvHolder) getVar(var_name string) (string){
	for _, envVariable := range(x.vars) {
		if envVariable.name == var_name {
			return envVariable.val
		}
	}
	log.Fatalf("Variable %s not found in .env file.", var_name)
	return ""
}