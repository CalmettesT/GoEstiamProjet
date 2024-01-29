package main

import (
	"fmt"
	"log"
	"os"
	"projet/databases"
	"projet/dossiers"
	"projet/fichiers"
	"projet/server"
)

const path = "C:\\GoEstiamProjet\\src\\data\\"

const config = "offline"

type Command interface {
	Execute(args []string) error
}

type Manager interface {
	CreateFolder(name string) error
	ReadFolder(name string) error
	RenameFolder(oldName, newName string) error
	DeleteFolder(name string) error
}

// Implémentation Offline
type OfflineManager struct{}

func (fm OfflineManager) CreateFolder(name string) error {
	// Créer un dossier
	_, err := dossiers.CreateFolder(os.Args[3], path)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	return nil
}

func (fm OfflineManager) ReadFolder(name string) error {
	// Lire le dossier
	_, err := dossiers.ReadFolder(os.Args[3], path)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	return nil
}

func (fm OfflineManager) ReadFolder(name string) error {
	// Lire le dossier
	_, err := dossiers.ReadFolder(os.Args[3], path)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	return nil
}

// Implémentation Online
type OnlineManager struct{}

func (fm OnlineManager) CreateFolder(name string) error {
	// implémentation online de la création de dossier
	return nil
}

type DirCommand struct{}

func (c DirCommand) Execute(args []string) error {
	// Vérifie si le nombre d'argument est supérieur à 2
	if len(os.Args) > 2 {

		switch os.Args[2] {
		case "create":
			// Vérifie si le nombre d'argument est supérieur à 3
			if len(os.Args) > 3 {

			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "read":
			// Vérifie si le nombre d'argument est supérieur à 3
			if len(os.Args) > 3 {

			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "rename":
			// Vérifie si le nombre d'argument est supérieur à 4
			if len(os.Args) > 4 {
				// Met à jour le nom du dossier
				_, err := dossiers.RenameFolder(os.Args[3], os.Args[4], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "delete":
			// Vérifie si le nombre d'argument est supérieur à 3
			if len(os.Args) > 3 {
				// Delete un dossier
				err := dossiers.DeleteFolder(os.Args[3], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		default:
			fmt.Println("Aucune commande ne correspond à votre saisie.")
		}
	} else {
		fmt.Println("Il est nécessaire de saisir un argument à la suite de la commande dir.")
	}

	return nil
}

type FileCommand struct{}

func (c FileCommand) Execute(args []string) error {
	// Vérifie si le nombre d'argument est supérieur à 2
	if len(os.Args) > 2 {

		switch os.Args[2] {
		case "create":
			// Vérifie si le nombre d'argument est supérieur à 3
			if len(os.Args) > 4 {
				// Créer un fichier
				err := fichiers.CreateFile(os.Args[3], os.Args[4], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "read":
			// Vérifie si le nombre d'argument est supérieur à 3
			if len(os.Args) > 3 {
				// Lire le fichier
				_, err := fichiers.ReadFile(os.Args[3], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "rename":
			// Vérifie si le nombre d'argument est supérieur à 4
			if len(os.Args) > 4 {
				// Met à jour le nom du dossier
				err := fichiers.UpdateNameFile(os.Args[3], os.Args[4], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "updatetext":
			// Vérifie si le nombre d'argument est supérieur à 4
			if len(os.Args) > 4 {
				// Met à jour le nom du dossier
				err := fichiers.UpdateTextFile(os.Args[3], os.Args[4], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le nombre d'arguments fournis n'est pas valide.")
			}

		case "delete":
			// Vérifie si le nombre d'argument est supérieur à 3
			if len(os.Args) > 3 {
				// Delete un dossier
				err := fichiers.DeleteFile(os.Args[3], path)
				if err != nil {
					fmt.Println("Erreur :", err)
				}
			} else {
				fmt.Println("Le chemin est vide.")
			}

		default:
			fmt.Println("Aucune commande ne correspond à votre saisie.")
		}

	} else {
		fmt.Println("Il est nécessaire de saisir un argument à la suite de la commande file.")
	}

	return nil
}

type ServerCommand struct{}

func (c ServerCommand) Execute(args []string) error {
	server.ServerStart()
	return nil
}

type HistoriqueCommand struct{}

func (c HistoriqueCommand) Execute(args []string) error {
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

type HelpCommande struct{}

func (c HelpCommande) Execute(args []string) error {
	// Liste de toutes les commandes disponibles
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "dir":
			// Commande dir
			command := []string{"create", "read", "rename", "delete"}
			if len(command) > 0 {
				fmt.Println("Voici les sous-commandes disponibles pour la commande dir:")
				for _, entry := range command {
					fmt.Println("-", entry)
				}
			}

		case "file":
			// Commande file
			command := [...]string{"create", "read", "rename", "updatetext", "delete"}
			if len(command) > 0 {
				fmt.Println("Voici les sous-commandes disponibles pour la commande file:")
				for _, entry := range command {
					fmt.Println("-", entry)
				}
			}

		default:
			fmt.Println("Aucune commande ne correspond à votre saisie.")

		}
	} else if len(os.Args) == 2 {
		// Commande de base
		command := [...]string{"dir", "file"}
		if len(command) > 0 {
			for _, entry := range command {
				fmt.Println(entry)
			}
		}
	}
	return nil
}

func DispatchCommand(commandName string, args []string) error {
	var cmd Command

	switch commandName {

	case "dir":
		cmd = DirCommand{}
	case "file":
		cmd = FileCommand{}
	case "server":
		cmd = ServerCommand{}

	case "hist":
		cmd = HistoriqueCommand{}
	case "help":
		cmd = HelpCommande{}

	default:
		return fmt.Errorf("commande inconnue: %s", commandName)
	}

	return cmd.Execute(args)
}

func main() {
	if len(os.Args) > 1 {
		commandName := os.Args[1]
		err := DispatchCommand(commandName, os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Exécutez la commande \"help\" pour obtenir de l'aide.")
	}
}
