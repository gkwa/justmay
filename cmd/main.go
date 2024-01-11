package main

import (
	"os"

	"github.com/taylormonacelli/justmay"
)

func main() {
	code := justmay.Execute()
	os.Exit(code)
}
