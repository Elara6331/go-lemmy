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

	err = c.ClientLogin(ctx, lemmy.Login{
		UsernameOrEmail: "user@example.com",
		Password:        `TestPwd`,
	})
	if err != nil {
		panic(err)
	}

	cr, err := c.CreateComment(ctx, lemmy.CreateComment{
		PostID:  2,
		Content: "Hello from go-lemmy!",
	})
	if err != nil {
		panic(err)
	}

	cr2, err := c.CreateComment(ctx, lemmy.CreateComment{
		PostID:   2,
		ParentID: lemmy.NewOptional(cr.CommentView.Comment.ID),
		Content:  "Reply from go-lemmy",
	})
	if err != nil {
		panic(err)
	}

	log.Printf(
		"Created comment %d and replied to it with comment %d!\n",
		cr.CommentView.Comment.ID,
		cr2.CommentView.Comment.ID,
	)
}
