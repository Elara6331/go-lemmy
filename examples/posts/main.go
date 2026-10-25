package main

import (
	"context"
	"log"

	"go.elara.ws/go-lemmy/v4"
)

func main() {
	ctx := context.Background()

	c, err := lemmy.NewWithLogin("https://lemmy.ml", &lemmy.Login{
		UsernameOrEmail: "user@example.com",
		Password:        `TestPwd`,
	})
	if err != nil {
		panic(err)
	}

	// Get the linux community to get its ID.
	gcr, err := c.GetCommunity(ctx, lemmy.GetCommunityParams{
		Name: lemmy.NewOptString("linux"),
	})
	if err != nil {
		panic(err)
	}

	// Create a Hello World post in the linux community.
	pr, err := c.CreatePost(ctx, &lemmy.CreatePost{
		CommunityID: gcr.CommunityView.Community.ID,
		Name:        "Hello, World!",
		Body:        lemmy.NewOptString("This is an example post"),
	})

	log.Println("Created post:", pr.PostView.Post.ID)
}
