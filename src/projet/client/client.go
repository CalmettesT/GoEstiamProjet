package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ServerURL = "http://localhost:8080"

type Folder struct {
	Name string `json:"name"`
}

func CreateFolder(name string) (string, error) {
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
	fmt.Println(respBody.FolderPath)
	return respBody.FolderPath, nil
}

func ReadFolder(name string) ([]string, error) {
    // Envoyer une requête GET pour lire le contenu du dossier
    resp, err := http.Get(ServerURL + "/dossiers/" + name)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Vérifier si le serveur a renvoyé une erreur
    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("serveur a retourné une erreur : %s", body)
    }

    // Structure pour décoder la réponse du serveur
    var respBody []string
	//Utilisez json.Unmarshal lorsque vous avez déjà les données JSON dans une variable (comme un slice d'octets).
	//Utilisez json.NewDecoder lorsque vous lisez des données JSON à partir d'un flux (comme un corps de réponse HTTP ou un fichier).
    if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
        return nil, err
    }
	fmt.Println(respBody)
    return respBody, nil
}


func RenameFolder(oldName string, newName string) error {
	type FolderRenameRequest struct {
		NewName string `json:"newName"`
	}

	requestData := FolderRenameRequest{
		NewName: newName,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, ServerURL+"/dossiers/"+oldName, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le dossier a bien été renommé.")
	}

	return nil
}

func DeleteFolder(name string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, ServerURL+"/dossiers/"+name, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le dossier a bien été supprimé.")
	}

	return nil
}

func CreateFile(name string, content string) (string, error) {
	type FileCreationRequest struct {
		FileName    string `json:"name"`
		FileContent string `json:"content"`
	}

	requestData := FileCreationRequest{
		FileName:    name,
		FileContent: content,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(ServerURL+"/fichiers/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le fichier a bien été créé.")
	}

	var respBody struct {
		FilePath string `json:"filePath"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}

	return respBody.FilePath, nil
}

func ReadFile(name string) (string, error) {
	resp, err := http.Get(ServerURL + "/fichiers/" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le fichier a bien été lu")
	}

	var respBody struct {
		FileContent string `json:"content"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}
	fmt.Println(respBody.FileContent)
	return respBody.FileContent, nil
}

func UpdateTextFile(name string, content string) error {
	type FileUpdateRequest struct {
		FileContent string `json:"content"`
	}

	requestData := FileUpdateRequest{
		FileContent: content,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, ServerURL+"/fichiers/update/"+name, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le fichier a bien été mis a jour")
	}

	return nil
}

func UpdateNameFile(oldName string, newName string) error {
	type FileRenameRequest struct {
		NewName string `json:"newName"`
	}

	requestData := FileRenameRequest{
		NewName: newName,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, ServerURL+"/fichiers/rename/"+oldName, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le nom du fichier a bien été mis a jour")
	}

	return nil
}

func DeleteFile(name string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, ServerURL+"/fichiers/"+name, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("serveur a retourné une erreur : %s", body)
	} else {
		fmt.Println("Le fichier a bien été mis a jour")
	}

	return nil
}

func Hist() ([]LogEntry, error) {

	resp, err := http.Get(ServerURL + "/divers/hist")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("serveur a retourné une erreur : %s", body)
	}

	var log []LogEntry
	bodyBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &log)

	return log, nil
}

// LogEntry représente une entrée d'historique
type LogEntry struct {
	ID       int64
	DH       time.Time
	MF       string
	Argument string
	Statut   string
}
