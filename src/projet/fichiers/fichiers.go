package fichiers

import (
	"errors"
	"fmt"
	"os"
)

func CreateFile(name string, text string) error {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name

	// Vérifie si le fichier existe déjà
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return errors.New("impossible de créer le fichier")
		}
		defer file.Close()

		// Si du texte est fourni, écrivez-le dans le fichier
		if text != "" {
			errorWrite := os.WriteFile(filePath, []byte(text), 0644)
			if errorWrite != nil {
				return errors.New("le fichier ne s'est pas remplie")
			}
		}

		fmt.Println("Fichier créé")
		return nil
	} else {
		return errors.New("le fichier existe déjà")
	}
}

func ReadFile(name string) (string, error) {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", errors.New("ce fichier n'existe pas")

	}
	fmt.Println("Contenu du fichier:", string(data))
	return string(data), nil
}

func UpdateTextFile(name string, data string) error {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return errors.New("impossible de mettre à jour le fichier")
	}
	fmt.Println("Fichier mis à jour.")
	return nil
}

func UpdateNameFile(oldName string, newName string) error {
	oldPath := "C:\\GoEstiamProjet\\src\\data\\" + oldName
	newPath := "C:\\GoEstiamProjet\\src\\data\\" + newName
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return errors.New("impossible de mettre à jour le nom fichier")

	}
	fmt.Println("Nom fichier mis à jour.")
	return nil
}

func DeleteFile(name string) error {
	filePath := "C:\\GoEstiamProjet\\src\\data\\" + name
	err := os.Remove(filePath)
	if err != nil {
		return errors.New("impossible de supprimer le fichier")

	}
	fmt.Println("Fichier supprimé.")
	return nil
}
