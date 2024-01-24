package main

import (
	"errors"
	"fmt"
	"os"
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
						createFolder(os.Args[3])
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
		}
	} else {
		fmt.Println("Exécutez la commande \"-help\" si vous n'êtes pas familier avec les instructions disponibles.")
	}
}

func createFolder(name string) {
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, error := os.Stat(path); errors.Is(error, os.ErrNotExist) {
		error := os.Mkdir(path, 0755)

		if error != nil {
			fmt.Println("Une erreur est apparue.")
			return
		} else {
			fmt.Println("Le dossier a bien été créé.")
		}
	} else {
		fmt.Println("Le dossier existe déjà.")
	}

}

func readFolder(name string) {
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, error := os.Stat(path); errors.Is(error, os.ErrNotExist) {
		fmt.Println("Le dossier n'existe pas.")
	} else {
		valeurs, error := os.ReadDir(path)

		if error != nil {
			fmt.Println(error)
			return
		}

		if len(valeurs) > 1 {
			for _, entry := range valeurs {
				fmt.Println(entry.Name())
			}
		} else {
			fmt.Println("Le dossier ne contient aucune donnée.")
		}
	}
}

func updateFolder(oldName string, newName string) {
	oldPath := "C:\\GoEstiamProjet\\src\\data\\" + oldName
	newPath := "C:\\GoEstiamProjet\\src\\data\\" + newName

	if _, error := os.Stat(oldPath); errors.Is(error, os.ErrNotExist) {
		fmt.Println("Le dossier n'existe pas.")
	} else {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func deleteFolder(name string) {
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, error := os.Stat(path); errors.Is(error, os.ErrNotExist) {
		fmt.Println("Le dossier n'existe pas.")
	} else {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Le dossier, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")
		}
	}
}
