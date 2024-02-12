package helper

import (
	"context"
	"encoding/base64"
	"fmt"
	"gocroot/config"
	"io"
	"mime/multipart"

	"github.com/google/go-github/v59/github"

	"golang.org/x/oauth2"
)

func GithubUpload(fileHeader *multipart.FileHeader) (err error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file content: %w", err)
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
	_, _, err = client.Repositories.CreateFile(ctx, config.GitHubOwner, config.GitHubRepo, fileHeader.Filename, opts)
	if err != nil {
		return err
	}
	return
}
