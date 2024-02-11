package config

import "os"

var PublicKey string = os.Getenv("PUBLICKEY")

var GitHubAccessToken string = os.Getenv("GitHubAccessToken")
var GitHubOwner string = os.Getenv("GitHubOwner")
var GitHubRepo string = os.Getenv("GitHubRepo")
var GitHubAuthorName string = os.Getenv("GitHubAuthorName")
var GitHubAuthorEmail string = os.Getenv("GitHubAuthorEmail")
