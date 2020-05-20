package main

import (
	"log"
	"fmt"
	"io/ioutil"
)

var (
	TargFile = "./src/db.sql"
	DestFile = "./src/rawsql.go"

	Template = "package main\nvar RawSQL = " + "`" + "%s"+"`"+"\n"
)

func main() {

	file, err := ioutil.ReadFile(TargFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	compiled := fmt.Sprintf(Template, string(file))

	err = ioutil.WriteFile(DestFile, []byte(compiled), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}


}