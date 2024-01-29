package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ServerURL = "http://localhost:8080"


func CreateFolder(name string) (string, error) {
    // Structure représentant le corps de la requête de création de dossier
    type FolderCreationRequest struct {
        FolderName string `json:"name"`
    }

    requestData := FolderCreationRequest{
        FolderName: name,
    }

    requestBody, err := json.Marshal(requestData)
    if err != nil {
        return "", err
    }

    resp, err := http.Post(ServerURL+"/dossiers/", "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        body, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("serveur a retourné une erreur : %s", body)
    } else {
		fmt.Println("Le dossier a bien été créé.")
	}


    // Structure pour la réponse du serveur
    var respBody struct {
        FolderPath string `json:"folderPath"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
        return "", err
    }

    return respBody.FolderPath, nil
}


type FileCreationRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func CreateFile(name string, content string) error {
	// Préparer la requête avec la structure définie
	requestData := FileCreationRequest{
		Name:    name,
		Content: content,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Envoyer la requête POST au serveur
	response, err := http.Post(ServerURL+"/fichiers/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("server returned error: %s", body)
	}

	return nil
}
