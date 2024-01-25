package fichiers

import (
	"fmt"
	"os"
)

func CreateFile(name string) {
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
