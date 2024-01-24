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

func createFile(name string) {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Impossible de créer le fichier:", err)
			return
		}
		defer file.Close()
		fmt.Println("Fichier créé:", filePath)
	} else {
		fmt.Println("Le fichier existe déjà.")
	}
}

func readFile(name string) {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Impossible de lire le fichier:", err)
		return
	}
	fmt.Println("Contenu du fichier:", string(data))
}

func updateTextFile(name string, data string) {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		fmt.Println("Impossible de mettre à jour le fichier:", err)
		return
	}
	fmt.Println("Fichier mis à jour.")
}

func updateNameFile(oldName string, newName string) {
	oldPath := "C:\\GoEstiamProjet\\src\\data\\" + oldName
	newPath := "C:\\GoEstiamProjet\\src\\data\\" + newName
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("Impossible de mettre à jour le nom fichier:", err)
		return
	}
	fmt.Println("Nom fichier mis à jour.")
}

func deleteFile(name string) {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Impossible de supprimer le fichier:", err)
		return
	}
	fmt.Println("Fichier supprimé.")
}
