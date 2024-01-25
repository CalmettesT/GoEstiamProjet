package main

import (
	"fmt"
	"os"
	"projet/dossiers"
	"projet/fichiers"
)

const path = "C:\\GoEstiamProjet\\src\\data\\"

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
						err := dossiers.CreateFolder(os.Args[3], path)
						if err != nil {
							fmt.Println("Erreur :", err)
						}
					} else {
						fmt.Println("Le chemin que vous avez saisi est vide.")
					}

				case "read":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Lire le dossier
						err := dossiers.ReadFolder(os.Args[3], path)
						if err != nil {
							fmt.Println("Erreur :", err)
						}
					} else {
						fmt.Println("Le chemin que vous avez saisi est vide.")
					}

				case "rename":
					// Vérifie si le nombre d'argument est supérieur à 4
					if len(os.Args) > 4 {
						// Met à jour le nom du dossier
						err := dossiers.RenameFolder(os.Args[3], os.Args[4], path)
						if err != nil {
							fmt.Println("Erreur :", err)
						}
					} else {
						fmt.Println("Le chemin que vous avez saisi est vide.")
					}

				case "delete":
					// Vérifie si le nombre d'argument est supérieur à 3
					if len(os.Args) > 3 {
						// Delete un dossier
						err := dossiers.DeleteFolder(os.Args[3], path)
						if err != nil {
							fmt.Println("Erreur :", err)
						}
					} else {
						fmt.Println("Le chemin que vous avez saisi est vide.")
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
						fichiers.CreateFile(os.Args[3], path)
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
			if len(os.Args) > 2 {
				switch os.Args[2] {
				case "dir":
					// Commande dir
					command := []string{"create", "read", "rename", "delete"}
					if len(command) > 0 {
						fmt.Println("Voici les sous-commandes disponibles pour la commande dir:")
						for _, entry := range command {
							fmt.Println("-", entry)
						}
					}

				case "file":
					// Commande file
					command := []string{""}
					if len(command) > 0 {
						fmt.Println("Voici les sous-commandes disponibles pour la commande file:")
						for _, entry := range command {
							fmt.Println("-", entry)
						}
					}

				default:
					fmt.Println("Aucune commande ne correspond à votre saisie.")

				}
			} else if len(os.Args) == 2 {
				// Commande de base
				command := []string{"dir", "file"}
				if len(command) > 0 {
					for _, entry := range command {
						fmt.Println(entry)
					}
				}
			}

		default:
			fmt.Println("Aucune commande ne correspond à votre saisie.")
		}
	} else {
		fmt.Println("Exécutez la commande \"-help\" si vous n'êtes pas familier avec les instructions disponibles.")
	}
}
