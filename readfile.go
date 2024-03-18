package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
* Ouverte de mon argument os.Args[1] = s
*
* Lecture du fichier ligne par ligne;
* Création de mon graph()
* Si la ligne contient un espace == Création d'un sommet
* Si la ligne contient un "-" == Création d'une arrête
 */
func readFile(s string) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println("Erreur lecture du fichier :", err)
		return
	}
	defer file.Close()

	graph := graph()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, " ") {
			vertex := strings.Fields(line)
			room := vertex[0]
			graph.addVertex(room)
		} else if strings.Contains(line, "-") {
			edges := strings.Split(line, "-")
			extremite1 := edges[0]
			extremite2 := edges[1]
			graph.addEdges(extremite1, extremite2)
		}
	}
	graph.print()
}
