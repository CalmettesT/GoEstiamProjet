package dossiers

import (
	"errors"
	"fmt"
	"os"
)

<<<<<<< Updated upstream
func CreateFolder(name string) {
=======
func CreateFolder(name string) error {
>>>>>>> Stashed changes
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0755)

		if err != nil {
			return fmt.Errorf("le dossier ne s'est pas créé: %v", err)
		}

		return nil
	} else {
		return errors.New("le dossier existe déjà")
	}
}

<<<<<<< Updated upstream
func ReadFolder(name string) {
=======
func ReadFolder(name string) error {
>>>>>>> Stashed changes
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return errors.New("le dossier n'existe pas")
	} else {
		valeurs, err := os.ReadDir(path)

		if err != nil {
			return errors.New("la lecture du dossier n'a pas fonctionnée")
		}

		if len(valeurs) > 1 {
			for _, entry := range valeurs {
				fmt.Println(entry.Name())
			}
		} else {
			fmt.Println("Le dossier ne contient aucune donnée.")
		}

		return nil
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
