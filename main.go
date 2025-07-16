package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var CLI struct {
	UserName UserCmd     `cmd help:"Sets user's name" required:"true"`
	RepoName RepoNameCmd `cmd help:"Pulls a specific repository"`
}

type UserCmd struct {
	UserName string `arg:"" required:"true" help:"Name of the user"`
	Token    string `help:"Github token" required:"true" env:"GITHUB_TOKEN"`
	All      bool   `help:"List all repositories for user. Use RepoName command to list single repo" short:"a"`
}

type RepoNameCmd struct {
	RepoName string `arg:"" required:"false" help:"Name of a single repository to pull"`
}

func (u *UserCmd) Run() error {
	fmt.Printf("User name set to: %s\n", u.UserName)
	// github client connection
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: u.Token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	return nil
}

func (r *RepoNameCmd) Run() error {
	fmt.Printf("Pulling repository: %s\n", r.RepoName)
	return nil
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("RepoFetcher"),
		kong.Description("Application to pull repositories by name"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			NoExpandSubcommands: true,
		}),
	)
	err := ctx.Run()
	if err != nil {
		log.Fatalf("Error running command: %v", err)
		fmt.Println(err)
	}

}
