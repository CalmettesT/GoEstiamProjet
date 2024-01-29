package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"projet/dossiers"
	"projet/fichiers"
)

type LocalManager struct{}

type OnlineManager struct {
	ServerURL string
}

const path = "C:\\GoEstiamProjet\\src\\data\\"


func (lfm *LocalManager) CreateFolder(name string) (string, error) {
	// Logique locale pour créer un dossier
	return dossiers.CreateFolder(name, path)
}


func (ofm *OnlineManager) CreateFolder(name string) (string, error) {
	// Création de la requête
	requestBody, err := json.Marshal(map[string]string{"name": name})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(ofm.ServerURL+"/dossiers/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("serveur a retourné une erreur : %s", body)
	}

	// Lire la réponse
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


func (lfm *LocalManager) CreateFile(name string, content string) error {
	err := fichiers.CreateFile(name, content, path)
	if err != nil {
		return err
	}
	return nil
}


func (ofm *OnlineManager) CreateFile(name string, content string) error {
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
	response, err := http.Post(ofm.ServerURL+"/fichiers/", "application/json", bytes.NewBuffer(requestBody))
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
