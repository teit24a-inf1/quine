package main

import (
	_ "embed"
	"fmt"
)

// Datei `main.go` als Variable einlesen.

//go:embed main.go
var src string

func main() {
	fmt.Println("Ich gebe meinen Quelltext aus, indem ich beim Compilieren meine eigene Quelldatei einlese.")
	fmt.Println("Das gleiche geht auch ohne Zugriff auf die Quelldatei.")

	// Quelltext ausgeben.
	fmt.Println(src)
}
