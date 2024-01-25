package dossiers

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	name := "TestCreateFolder"
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	err := CreateFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(path)
}

func TestCreateFolderAccent(t *testing.T) {
	name := "TestCreateFolderAccentééééééè"
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	err := CreateFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}

	if _, err := os.Stat(path); err == nil {
		t.Error("Le dossier a été créé avec des accents")
	}

	name = RemoveAccents(name)
	pathVerif := "C:\\GoEstiamProjet\\src\\data\\" + name
	os.RemoveAll(pathVerif)
}

func TestReadFolder(t *testing.T) {
	name1 := "TestReadFolder"
	name2 := "TestReadFolder\\Alire"
	path1 := "C:\\GoEstiamProjet\\src\\data\\" + name1
	path2 := "C:\\GoEstiamProjet\\src\\data\\" + name2

	os.Mkdir(path1, 0755)
	os.Mkdir(path2, 0755)

	err := ReadFolder(name1)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(path1)
}

func TestReadFolderNotExist(t *testing.T) {
	name := "TestReadFolderNotExist"

	err := ReadFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}
}

func TestUpdateFolder(t *testing.T) {
	name1 := "TestUpdateFolder1"
	name2 := "TestUpdateFolder2"
	path1 := "C:\\GoEstiamProjet\\src\\data\\" + name1
	path2 := "C:\\GoEstiamProjet\\src\\data\\" + name2

	os.Mkdir(path1, 0755)

	err := UpdateFolder(name1, name2)
	if err != nil {
		t.Error("Erreur :", err)
	}

	os.RemoveAll(path2)
}

func TestUpdateFolderAccent(t *testing.T) {
	name1 := "TestUpdateFolderAccent"
	name2 := "JeSuisàllémàngérù2"
	path1 := "C:\\GoEstiamProjet\\src\\data\\" + name1
	path2 := "C:\\GoEstiamProjet\\src\\data\\" + name2

	os.Mkdir(path1, 0755)

	err2 := UpdateFolder(name1, name2)
	if err2 != nil {
		t.Error("Erreur :", err2)
	}

	if _, err3 := os.Stat(path2); err3 == nil {
		t.Error("Le dossier a été renommé avec des accents")
	}

	os.RemoveAll(path2)
}

func TestDelete(t *testing.T) {
	name := "TestDelete"

	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	os.Mkdir(path, 0755)

	err2 := DeleteFolder(name)
	if err2 != nil {
		t.Error("Erreur :", err2)
	}

	os.RemoveAll(path)
}

func TestDeleteNotExist(t *testing.T) {
	name := "TestDeleteNotExist"

	err := DeleteFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}
}
