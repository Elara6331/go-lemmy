package main

import (
	"context"
	"log"

	"go.elara.ws/go-lemmy"
)

func main() {
	ctx := context.Background()

	c, err := lemmy.New("https://lemmy.ml")
	if err != nil {
		panic(err)
	}

	// Log in to lemmy.ml
	err = c.ClientLogin(ctx, lemmy.Login{
		UsernameOrEmail: "user@example.com",
		Password:        `TestPwd`,
	})
	if err != nil {
		panic(err)
	}

	// Get the linux community to get its ID.
	gcr, err := c.Community(ctx, lemmy.GetCommunity{
		Name: lemmy.NewOptional("linux"),
	})
	if err != nil {
		panic(err)
	}

	// Create a Hello World post in the linux community.
	pr, err := c.CreatePost(ctx, lemmy.CreatePost{
		CommunityID: gcr.CommunityView.Community.ID,
		Name:        "Hello, World!",
		Body:        lemmy.NewOptional("This is an example post"),
	})

	log.Println("Created post:", pr.PostView.Post.ID)
}
