package fichiers

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func containsNoSpecificChars(s string) bool {
	// Retourne `false` si `s` contient au moins un des caractères dans `chars`
	chars := "\\/:*?\"<>|"

	return strings.ContainsAny(s, chars)
}

func CreateFile(name string, text string, path string) error {
	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(name) {
		return errors.New("la chaîne contient au moins un caractère bloquant")
	}

	filePath := path + name

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

func ReadFile(name string, path string) (string, error) {
	filePath := path + name
	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", errors.New("ce fichier n'existe pas")

	}
	fmt.Println("Contenu du fichier:", string(data))
	return string(data), nil
}

func UpdateTextFile(name string, data string, path string) error {
	filePath := path + name

	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("aucun fichier n'existe avec ce nom")
		}
		// Autres erreurs de système de fichiers
		return errors.New("impossible de vérifier l'existence du fichier")
	}

	// Vérifier si le chemin est un dossier
	if info.IsDir() {
		return errors.New("le chemin correspond à un dossier, pas à un fichier")
	}

	err = os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return errors.New("impossible de mettre à jour le fichier")
	}

	fmt.Println("Fichier mis à jour.")

	return nil
}

func UpdateNameFile(oldName string, newName string, path string) error {
	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(newName) {
		return errors.New("la chaîne contient au moins un caractère bloquant")
	}
	oldPath := path + oldName
	newPath := path + newName
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return errors.New("impossible de mettre à jour le nom fichier")

	}
	fmt.Println("Nom fichier mis à jour.")
	return nil
}

func DeleteFile(name string, path string) error {
	filePath := path + name
	err := os.Remove(filePath)
	if err != nil {
		return errors.New("impossible de supprimer le fichier")

	}
	fmt.Println("Fichier supprimé.")
	return nil
}
