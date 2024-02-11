package helper

import (
	"context"
	"encoding/base64"
	"fmt"
	"gocroot/config"
	"os"

	"github.com/google/go-github/v38/github"

	"golang.org/x/oauth2"
)

func GithubUpload(filePath string) {
	// Membaca file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Mengkodekan isi file ke base64
	encodedContent := base64.StdEncoding.EncodeToString(fileContent)

	// Konfigurasi koneksi ke GitHub menggunakan token akses
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GitHubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Membuat opsi untuk mengunggah file
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("Upload file"),
		Content: []byte(encodedContent),
		Branch:  github.String("main"),
		Author: &github.CommitAuthor{
			Name:  github.String(config.GitHubAuthorName),
			Email: github.String(config.GitHubAuthorEmail),
		},
	}

	// Membuat permintaan untuk mengunggah file
	_, _, err = client.Repositories.CreateFile(ctx, config.GitHubOwner, config.GitHubRepo, filePath, opts)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("File uploaded successfully.")
}
