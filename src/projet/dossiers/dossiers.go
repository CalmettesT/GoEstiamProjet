package dossiers

import (
	"errors"
	"fmt"
	"os"
	"projet/databases"
	"strings"
	"time"
)

const module = "Dir"

func logCommand(command, argument, statut string) {
	databases.ConnectDataBase()

	log := databases.LogData{
		DH:       time.Now(),
		MF:       module + " " + command,
		Argument: argument,
		Statut:   statut,
	}

	_, err := databases.AddLog(log)
	if err != nil {
		fmt.Println(err)
	}
}

func containsNoSpecificChars(s string) bool {
	// Retourne `false` si `s` contient au moins un des caractères dans `chars`
	chars := "\\/:*?\"<>|"

	return strings.ContainsAny(s, chars)
}

func CreateFolder(name, path string) error {
	command := "create"
	path = path + name

	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(name) {
		logCommand(command, name+" "+path, "La chaîne contient au moins un caractère bloquant.")
		return errors.New("la chaîne contient au moins un caractère bloquant")
	}

	// Vérifie si un dossier existe déjà avec ce nom
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		logCommand(command, name+" "+path, "Le dossier existe déjà.")
		return errors.New("le dossier existe déjà")
	}

	// Gère les autres erreurs lié à la vérification
	if err != nil && !os.IsNotExist(err) {
		logCommand(command, name+" "+path, "Erreur lors de la vérification de l'existence du dossier.")
		return errors.New("erreur lors de la vérification de l'existence du dossier")
	}

	// Création du dossier
	err = os.Mkdir(path, 0755)
	if err != nil {
		logCommand(command, name+" "+path, "Erreur lors de la création du dossier.")
		return errors.New("erreur lors de la création du dossier")
	} else {
		logCommand(command, name+" "+path, "Le dossier a bien été créé.")
		fmt.Println("Le dossier a bien été créé.")
	}

	return nil
}

func ReadFolder(name, path string) error {
	command := "read"
	path = path + name

	// Vérifie si le dossier existe
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			logCommand(command, name+" "+path, "Le dossier n'existe pas.")
			return errors.New("le dossier n'existe pas")
		}
		// Gérer les autres types d'erreurs lors de l'appel à os.Stat
		logCommand(command, name+" "+path, "Impossible de vérifier l'existence du dossier.")
		return errors.New("impossible de vérifier l'existence du dossier")
	}

	// Vérifie si le chemin est bien un dossier
	if !info.IsDir() {
		logCommand(command, name+" "+path, "Le chemin ne renvoi pas vers un dossier.")
		return errors.New("le chemin ne renvoi pas vers un dossier")
	}

	// Lire le contenu du dossier
	valeurs, err := os.ReadDir(path)
	if err != nil {
		logCommand(command, name+" "+path, "La lecture du dossier n'a pas fonctionnée.")
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

	logCommand(command, name+" "+path, "L'opération a bien fonctionnée.")
	return nil
}

func RenameFolder(oldName, newName, path string) error {
	oldPath := path + oldName
	newPath := path + newName
	command := "rename"

	// Si le nom du dossier contient des caractères bloquant
	if containsNoSpecificChars(newName) {
		logCommand(command, oldName+" "+newName+" "+path, "La chaîne contient au moins un caractère bloquant.")
		return errors.New("la chaîne contient au moins un caractère bloquant")
	}

	// Vérifie si le dossier existe
	info, err := os.Stat(oldPath)
	if err != nil {
		if os.IsNotExist(err) {
			logCommand(command, oldName+" "+newName+" "+path, "Le dossier n'existe pas.")
			return errors.New("le dossier n'existe pas")
		}
		// Gérer les autres types d'erreurs lors de l'appel à os.Stat
		logCommand(command, oldName+" "+newName+" "+path, "Impossible de vérifier l'existence du dossier.")
		return errors.New("impossible de vérifier l'existence du dossier")
	}

	// Vérifie si le chemin est bien un dossier
	if !info.IsDir() {
		logCommand(command, oldName+" "+newName+" "+path, "Le chemin ne renvoi pas vers un dossier.")
		return errors.New("le chemin ne renvoi pas vers un dossier")
	}

	// Vérifie si le dossier existe
	_, err = os.Stat(newPath)
	if err == nil {
		logCommand(command, oldName+" "+newName+" "+path, "Un dossier avec le nouveau nom existe déjà.")
		return errors.New("un dossier avec le nouveau nom existe déjà")
	}

	// Renommer le dossier
	err = os.Rename(oldPath, newPath)
	if err != nil {
		logCommand(command, oldName+" "+newName+" "+path, "La mise à jour du dossier n'a pas fonctionnée.")
		return errors.New("la mise à jour du dossier n'a pas fonctionnée")
	}

	logCommand(command, oldName+" "+newName+" "+path, "Le nom du dossier a bien été modifié, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")
	fmt.Println("Le nom du dossier a bien été modifié, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")

	return nil
}

func DeleteFolder(name, path string) error {
	path = path + name
	command := "delete"

	// Vérifie si le dossier existe
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		logCommand(command, name+" "+path, "Le dossier n'existe pas.")
		return errors.New("le dossier n'existe pas")
	}

	// Supprimer le dossier
	err := os.RemoveAll(path)
	if err != nil {
		logCommand(command, name+" "+path, "La tentative de suppression du dossier et de son contenu a échoué.")
		return errors.New("la tentative de suppression du dossier et de son contenu a échoué")
	}

	logCommand(command, name+" "+path, "Le dossier, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")
	fmt.Println("Le dossier, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")

	return nil
}
