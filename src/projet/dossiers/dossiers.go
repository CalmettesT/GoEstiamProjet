package dossiers

import (
	"errors"
	"fmt"
	"os"
)

func CreateFolder(name string) {
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

func ReadFolder(name string) {
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

func UpdateFolder(oldName string, newName string) {
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

func DeleteFolder(name string) {
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
