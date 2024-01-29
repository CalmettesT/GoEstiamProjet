package main

import (
	"fmt"
	"log"
	"os"
	"projet/databases"
	"projet/dossiers"
	"projet/fichiers"
)

const path = "C:\\GoEstiamProjet\\src\\data\\"

const config = "offline"

type Manager interface {
	CreateFolder(name string) error
	ReadFolder(name string) error
	RenameFolder(oldName, newName string) error
	DeleteFolder(name string) error
	CreateFile(name, text string) error
	ReadFile(name string) error
	RenameFile(oldName, newName string) error
	DeleteFile(name string) error
	UpdateText(name, text string) error
	Historic() error
}

// Implémentation Offline
type OfflineManager struct{}

func (fm OfflineManager) CreateFolder(name string) error {
	// Créer un dossier
	_, err := dossiers.CreateFolder(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la création du dossier :", err)
	}

	return nil
}

func (fm OfflineManager) ReadFolder(name string) error {
	// Lire le dossier
	_, err := dossiers.ReadFolder(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier :", err)
	}

	return nil
}

func (fm OfflineManager) RenameFolder(oldName, newName string) error {
	// Renommer un dossier
	_, err := dossiers.RenameFolder(oldName, newName, path)
	if err != nil {
		fmt.Println("Erreur lors du renommage du dossier :", err)
	}

	return nil
}

func (fm OfflineManager) DeleteFolder(name string) error {
	// Delete un dossier
	err := dossiers.DeleteFolder(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du dossier :", err)
	}

	return nil
}

func (fm OfflineManager) CreateFile(name, text string) error {
	// Créer un fichier
	err := fichiers.CreateFile(name, text, path)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
	}

	return nil
}

func (fm OfflineManager) ReadFile(name string) error {
	// Lire un fichier
	_, err := fichiers.ReadFile(name, path)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}

	return nil
}

func (fm OfflineManager) RenameFile(oldName, newName string) error {
	// Renommer un fichier
	err := fichiers.UpdateNameFile(oldName, newName, path)
	if err != nil {
		fmt.Println("Erreur lors du renommage du fichier :", err)
	}

	return nil
}

func (fm OfflineManager) UpdateText(name, text string) error {
	// Ajouter du texte dans un fichier
	err := fichiers.UpdateTextFile(name, text, path)
	if err != nil {
		fmt.Println("Erreur lors du renommage du fichier :", err)
	}

	return nil
}

func (fm OfflineManager) DeleteFile(name string) error {
	// Ajouter du texte dans un fichier
	err := fichiers.DeleteFile(name, path)
	if err != nil {
		fmt.Println("Erreur lors du renommage du fichier :", err)
	}

	return nil
}

func (fm OfflineManager) Historic() error {
	databases.ConnectDataBase()

	journaux, err := databases.LastJournal()
	if err != nil {
		log.Fatal(err)
	}

	if len(journaux) > 0 {
		fmt.Printf("Voici l'historique des 50 dernières commandes :\n\n")
		for _, entry := range journaux {
			fmt.Println(entry.ID, " | ", entry.DH, " | ", entry.MF, " | ", entry.Argument, " | ", entry.Statut)
		}
	}

	return nil
}

// // Implémentation Online
type OnlineManager struct{}

func (fm OnlineManager) CreateFolder(name string) error {
	return nil
}

func (fm OnlineManager) ReadFolder(name string) error {
	return nil
}

func (fm OnlineManager) RenameFolder(oldName, newName string) error {
	return nil
}

func (fm OnlineManager) DeleteFolder(name string) error {
	return nil
}

func (fm OnlineManager) CreateFile(name, text string) error {
	return nil
}

func (fm OnlineManager) ReadFile(name string) error {
	return nil
}

func (fm OnlineManager) RenameFile(oldName, newName string) error {
	return nil
}

func (fm OnlineManager) UpdateText(name, text string) error {
	return nil
}

func (fm OnlineManager) DeleteFile(name string) error {
	return nil
}

func (fm OnlineManager) Historic() error {
	return nil
}

func main() {
	if len(os.Args) > 1 {
		var manager Manager

		// Choix de l'implémentation basé sur la configuration
		if config == "offline" {
			manager = OfflineManager{}
		} else {
			manager = OnlineManager{}
		}

		commandName := os.Args[1]

		switch commandName {
		case "dir":
			if len(os.Args) > 2 {
				sousCommandName := os.Args[2]
				switch sousCommandName {
				case "create":
					if len(os.Args) > 3 {
						manager.CreateFolder(os.Args[3])
					} else {
						fmt.Println("Nom du dossier manquant")
					}

				case "read":
					if len(os.Args) > 3 {
						manager.ReadFolder(os.Args[3])
					} else {
						fmt.Println("Nom du dossier manquant")
					}

				case "rename":
					if len(os.Args) > 4 {
						manager.RenameFolder(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Nom des dossiers manquant")
					}

				case "delete":
					if len(os.Args) > 3 {
						manager.DeleteFolder(os.Args[3])
					} else {
						fmt.Println("Nom du dossier manquant")
					}
				default:
					fmt.Println("Commande inconnue")
				}
			} else {
				fmt.Println("Veuillez saisir une sous-commande.")
			}

		case "file":
			if len(os.Args) > 2 {
				sousCommandName := os.Args[2]
				switch sousCommandName {
				case "create":
					if len(os.Args) > 4 {
						manager.CreateFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le nom du fichier ou le texte est manquant")
					}

				case "read":
					if len(os.Args) > 3 {
						manager.ReadFile(os.Args[3])
					} else {
						fmt.Println("Nom du fichier manquant")
					}

				case "rename":
					if len(os.Args) > 4 {
						manager.RenameFile(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Nom des fichiers manquant")
					}

				case "update":
					if len(os.Args) > 4 {
						manager.UpdateText(os.Args[3], os.Args[4])
					} else {
						fmt.Println("Le nom du fichier ou le texte est manquant")
					}

				case "delete":
					if len(os.Args) > 3 {
						manager.DeleteFile(os.Args[3])
					} else {
						fmt.Println("Nom du fichier manquant")
					}
				// Ajoutez des cas pour les autres commandes ici...
				default:
					fmt.Println("Commande inconnue")
				}
			} else {
				fmt.Println("Veuillez saisir une sous-commande.")
			}

		case "help":

		case "hist":
			manager.Historic()

		default:
			fmt.Println("Commande inconnue")
		}

	} else {
		fmt.Println("Exécutez la commande \"help\" pour obtenir de l'aide.")
	}
}
