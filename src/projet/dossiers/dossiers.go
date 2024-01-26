package dossiers

import (
	"errors"
	"fmt"
	"os"
	"projet/database"
	"strings"
)

func logCommand() {
	database.ConnectDataBase()
}

func containsNoSpecificChars(s string) bool {
	// Retourne `false` si `s` contient au moins un des caractères dans `chars`
	chars := "\\/:*?\"<>|"

	return strings.ContainsAny(s, chars)
}

func CreateFolder(name, path string) error {
	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(name) {
		return errors.New("la chaîne contient au moins un caractère bloquant")
	}
	path = path + name

	// Vérifie si un dossier existe déjà avec ce nom
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return errors.New("le dossier existe déjà")
	}

	// Gère les autres erreurs lié à la vérification
	if err != nil && !os.IsNotExist(err) {
		return errors.New("erreur lors de la vérification de l'existence du dossier")
	}

	// Création du dossier
	err = os.Mkdir(path, 0755)
	if err != nil {
		return errors.New("erreur lors de la création du dossier")
	} else {
		fmt.Println("Le dossier a bien été créé.")
	}

	return nil
}

func ReadFolder(name, path string) error {
	path = path + name

	// Vérifie si le dossier existe
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("le dossier n'existe pas")
		}
		// Gérer les autres types d'erreurs lors de l'appel à os.Stat
		return errors.New("impossible de vérifier l'existence du dossier")
	}

	// Vérifie si le chemin est bien un dossier
	if !info.IsDir() {
		return errors.New("le chemin ne renvoi pas vers un dossier")
	}

	// Lire le contenu du dossier
	valeurs, err := os.ReadDir(path)
	if err != nil {
		return errors.New("la lecture du dossier n'a pas fonctionnée")
	}

	// Afficher le contenu du dossier
	if len(valeurs) > 0 {
		for _, entry := range valeurs {
			fmt.Println(entry.Name())
		}
	} else {
		fmt.Println("Le dossier ne contient aucune donnée.")
	}

	return nil
}

func RenameFolder(oldName, newName, path string) error {
	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(newName) {
		return errors.New("la chaîne contient au moins un caractère bloquant")
	}

	oldPath := path + oldName
	newPath := path + newName

	// Vérifie si le dossier existe
	info, err := os.Stat(oldPath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("le dossier n'existe pas")
		}
		// Gérer les autres types d'erreurs lors de l'appel à os.Stat
		return errors.New("impossible de vérifier l'existence du dossier")
	}

	// Vérifie si le chemin est bien un dossier
	if !info.IsDir() {
		return errors.New("le chemin ne renvoi pas vers un dossier")
	}

	// Vérifie si le dossier existe
	_, err = os.Stat(newPath)
	if err == nil {
		return errors.New("un dossier avec le nouveau nom existe déjà")
	}

	// Renommer le dossier
	err = os.Rename(oldPath, newPath)
	if err != nil {
		return errors.New("la mise à jour du dossier n'a pas fonctionnée")
	}

	fmt.Println("Le nom du dossier a bien été modifié, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")

	return nil
}

func DeleteFolder(name, path string) error {
	path = path + name

	// Vérifie si le dossier existe
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return errors.New("le dossier n'existe pas")
	}

	// Supprimer le dossier
	err := os.RemoveAll(path)
	if err != nil {
		return errors.New("la tentative de suppression du dossier et de son contenu a échoué")
	}

	fmt.Println("Le dossier, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")

	return nil
}
