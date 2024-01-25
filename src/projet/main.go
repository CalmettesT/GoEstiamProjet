package main

import (
	"fmt"
	"os"
	"projet/fichiers"
)

func main() {

	// Dossier
	if len(os.Args) > 1 {
		if os.Args[1] == "dir" {
			// Vérifie si le nombre d'argument est supérieur à 2
			if len(os.Args) > 2 {

				// CRUD
				if os.Args[2] == "create" {
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Créer un dossier
						fichiers.CreateFile(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "read" {
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Lire le dossier
						readFolder(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "update" {
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						updateFolder(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "delete" {
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Delete un dossier
						deleteFolder(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else {
					fmt.Println("Aucune commande ne correspond à votre saisie.")
				}
			} else {
				fmt.Println("Il est nécessaire de saisir un argument à la suite de la commande dir.")
			}
		} else if os.Args[1] == "-help" {
			// Liste de toutes les commandes disponibles
			fmt.Println("C'est en cours")
		} else if os.Args[1] == "file" {
			// Vérifie si le nombre d'argument est supérieur à 2
			if len(os.Args) > 2 {

				// CRUD
				if os.Args[2] == "create" {
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Créer un fichier
						createFile(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "read" {
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Lire le fichier
						readFile(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "updatename" {
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						updateNameFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "updatetext" {
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						updateTextFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else if os.Args[2] == "delete" {
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Delete un dossier
						deleteFile(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}
				} else {
					fmt.Println("Aucune commande ne correspond à votre saisie.")
				}
			} else {
				fmt.Println("Il est nécessaire de saisir un argument à la suite de la commande dir.")
			}
		}
	} else {
		fmt.Println("Exécutez la commande \"-help\" si vous n'êtes pas familier avec les instructions disponibles.")
	}
}
