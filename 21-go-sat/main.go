package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/suborbital/sat/sat"
)

func main() {

	argsWithoutCaller := os.Args[1:]
	wasmModulePath := argsWithoutCaller[0]
	wasmModuleParameters := argsWithoutCaller[1:]

	wasmModuleConfig, _ := sat.ConfigFromRunnableArg(wasmModulePath)
	
	satFunction, _ := sat.New(wasmModuleConfig, nil)

	result, err := satFunction.Exec([]byte(strings.Join(wasmModuleParameters, ";")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🎉 " + string(result.Output))

}
