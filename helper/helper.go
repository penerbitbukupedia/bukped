package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gocroot/config"
	"gocroot/model"
	"io"
	"net/http"
	"os"
)

func UploadGithub(filePath string) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	encodedContent := base64.StdEncoding.EncodeToString(fileContent)

	createFileRequest := model.CreateFileRequest{
		Message:     "Upload file",
		Content:     encodedContent,
		Branch:      "main",
		AuthorName:  config.GitHubAuthorName,
		AuthorEmail: config.GitHubAuthorEmail,
	}

	jsonData, err := json.Marshal(createFileRequest)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", config.GitHubOwner, config.GitHubRepo, filePath)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Set("Authorization", "token "+config.GitHubAccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(responseBody))
}
