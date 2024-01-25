package dossiers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RemoveAccents(s string) string {
	// Transformer la chaîne en Forme de Décomposition Canonique (NFD)
	t := transform.Chain(norm.NFD)
	s, _, _ = transform.String(t, s)

	// Filtrer les caractères non-espacés (supprimer les accents)
	s = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) { // Mn: marque non-espacée (accents, etc.)
			return -1
		}
		return r
	}, s)

	return s
}

func CreateFolder(name string) error {
	name = RemoveAccents(name)

	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	// Si dossier existe
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0755)

		if err != nil {
			return errors.New("le dossier ne s'est pas créé")
		}

		return nil
	} else {
		return errors.New("le dossier existe déjà")
	}
}

func ReadFolder(name string) error {
	name = RemoveAccents(name)

	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return errors.New("le dossier n'existe pas")
	} else {
		valeurs, err := os.ReadDir(path)

		if err != nil {
			return errors.New("la lecture du dossier n'a pas fonctionnée")
		}

		if len(valeurs) > 0 {
			for _, entry := range valeurs {
				fmt.Println(entry.Name())
			}
		} else {
			fmt.Println("Le dossier ne contient aucune donnée.")
		}

		return nil
	}
}

func UpdateFolder(oldName string, newName string) error {
	oldName = RemoveAccents(oldName)
	newName = RemoveAccents(newName)

	oldPath := "C:\\GoEstiamProjet\\src\\data\\" + oldName
	newPath := "C:\\GoEstiamProjet\\src\\data\\" + newName

	if _, err := os.Stat(oldPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("le dossier n'existe pas")
	} else {

		err := os.Rename(oldPath, newPath)
		if err != nil {
			return errors.New("la mise à jour du dossier n'a pas fonctionnée")
		}
		return nil
	}
}

func DeleteFolder(name string) error {
	name = RemoveAccents(name)

	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return errors.New("le dossier n'existe pas")
	} else {
		err := os.RemoveAll(path)
		if err != nil {
			return errors.New("la tentative de suppression du dossier et de son contenu a échoué")
		} else {
			fmt.Println("Le dossier, ainsi que toutes les données qu'il contenait, ont été intégralement supprimés.")
		}
		return nil
	}
}
