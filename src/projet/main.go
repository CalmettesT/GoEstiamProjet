package main

import (
	"fmt"
	"os"
	"projet/dossiers"
	"projet/fichiers"
)

func main() {
	// Dossier
	if len(os.Args) > 1 {
		switch os.Args[1] {

		case "dir":
			// Vérifie si le nombre d'argument est supérieur à 2
			if len(os.Args) > 2 {

				switch os.Args[2] {
				case "create":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Créer un dossier
						dossiers.CreateFolder(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "read":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Lire le dossier
						dossiers.ReadFolder(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "update":
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						dossiers.UpdateFolder(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "delete":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Delete un dossier
						dossiers.DeleteFolder(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				default:
					fmt.Println("Aucune commande ne correspond à votre saisie.")
				}
			} else {
				fmt.Println("Il est nécessaire de saisir un argument à la suite de la commande dir.")
			}

		case "file":
			// Vérifie si le nombre d'argument est supérieur à 2
			if len(os.Args) > 2 {

				switch os.Args[2] {
				case "create":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 4 {
						// Créer un fichier
						fichiers.CreateFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "read":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Lire le fichier
						fichiers.ReadFile(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "updatename":
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						fichiers.UpdateNameFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "updatetext":
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						fichiers.UpdateTextFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				case "delete":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Delete un dossier
						fichiers.DeleteFile(os.Args[3])
					} else {
						fmt.Println("Le chemin est vide.")
					}

				default:
					fmt.Println("Aucune commande ne correspond à votre saisie.")
				}

			} else {
				fmt.Println("Il est nécessaire de saisir un argument à la suite de la commande file.")
			}

		case "-help":
			// Liste de toutes les commandes disponibles
			fmt.Println("C'est en cours")
		default:
			fmt.Println("Aucune commande ne correspond à votre saisie.")
		}
	} else {
		fmt.Println("Exécutez la commande \"-help\" si vous n'êtes pas familier avec les instructions disponibles.")
	}
}
