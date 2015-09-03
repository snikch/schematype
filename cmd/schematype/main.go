package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/snikch/schematype"
)

var dir = flag.String("dir", "./", "The directory of schema .json files")
var pkg = flag.String("package", "main", "The name of the package")

func main() {
	flag.Parse()

	i := os.Stdin

	schema := schematype.Schema{}
	decoder := json.NewDecoder(i)
	err := decoder.Decode(&schema)
	if err != nil {
		log.Fatal("could not decode", err)
	}

	body, err := schema.TypeString(*name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("package "+*pkg, body)
}
