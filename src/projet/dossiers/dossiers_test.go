package dossiers

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	name := "MonDossierTest"

	err := CreateFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}
}

func TestCreateFolderAccent(t *testing.T) {
	name := "MonDossierTestéééééééé"
	path := "C:\\GoEstiamProjet\\src\\data\\" + name

	err1 := CreateFolder(name)
	if err1 != nil {
		t.Error("Erreur :", err1)
	}

	if _, err2 := os.Stat(path); err2 == nil {
		t.Error("Le dossier a été créé avec des accents")
	}
}

func TestReadFolder(t *testing.T) {
	name1 := "MonDossier"
	name2 := "MonDossier\\Alire"

	err1 := CreateFolder(name1)
	if err1 != nil {
		t.Error("Erreur :", err1)
	}

	err2 := CreateFolder(name2)
	if err2 != nil {
		t.Error("Erreur :", err2)
	}

	err3 := ReadFolder(name1)
	if err3 != nil {
		t.Error("Erreur :", err3)
	}
}

func TestReadFolderNotExist(t *testing.T) {
	name := "JeNeSuisPasLa"

	err := ReadFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}
}

func TestUpdateFolder(t *testing.T) {
	name1 := "JeSuis1"
	name2 := "JeSuis2"

	err1 := CreateFolder(name1)
	if err1 != nil {
		t.Error("Erreur :", err1)
	}

	err2 := UpdateFolder(name1, name2)
	if err2 != nil {
		t.Error("Erreur :", err2)
	}
}

func TestUpdateFolderAccent(t *testing.T) {
	name1 := "JeSuisLa"
	name2 := "JeSuisàllémàngérù2"
	path := "C:\\GoEstiamProjet\\src\\data\\" + name2

	err1 := CreateFolder(name1)
	if err1 != nil {
		t.Error("Erreur :", err1)
	}

	err2 := UpdateFolder(name1, name2)
	if err2 != nil {
		t.Error("Erreur :", err2)
	}

	if _, err3 := os.Stat(path); err3 == nil {
		t.Error("Le dossier a été renommé avec des accents")
	}
}

func TestDelete(t *testing.T) {
	name := "JeSuisLaAvecMonDossier"

	err1 := CreateFolder(name)
	if err1 != nil {
		t.Error("Erreur :", err1)
	}

	err2 := DeleteFolder(name)
	if err2 != nil {
		t.Error("Erreur :", err2)
	}
}

func TestDeleteNotExist(t *testing.T) {
	name := "BonSupprimeMOI"

	err := DeleteFolder(name)
	if err != nil {
		t.Error("Erreur :", err)
	}
}
