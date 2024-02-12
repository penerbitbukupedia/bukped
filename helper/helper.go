package helper

import (
	"context"
	"gocroot/config"
	"io"
	"mime/multipart"

	"github.com/google/go-github/v59/github"

	"golang.org/x/oauth2"
)

func GithubUpload(fileHeader *multipart.FileHeader) (content *github.RepositoryContentResponse, response *github.Response, err error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer file.Close()
	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return
	}

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
		Content: fileContent,
		Branch:  github.String("main"),
		Author: &github.CommitAuthor{
			Name:  github.String(config.GitHubAuthorName),
			Email: github.String(config.GitHubAuthorEmail),
		},
	}

	// Membuat permintaan untuk mengunggah file
	content, response, err = client.Repositories.CreateFile(ctx, config.GitHubOwner, config.GitHubRepo, fileHeader.Filename, opts)
	if err != nil {
		currentContent, _, _, _ := client.Repositories.GetContents(ctx, config.GitHubOwner, config.GitHubRepo, fileHeader.Filename, nil)
		opts.SHA = github.String(currentContent.GetSHA())
		content, response, err = client.Repositories.UpdateFile(ctx, config.GitHubOwner, config.GitHubRepo, fileHeader.Filename, opts)
		return
	}

	return
}
