package main

import (
	"context"
	"log"

	"go.elara.ws/go-lemmy/v4"
)

const postID = 2

func main() {
	ctx := context.Background()

	c, err := lemmy.NewWithLogin("https://lemmy.ml", &lemmy.Login{
		UsernameOrEmail: "user@example.com",
		Password:        `TestPwd`,
	})
	if err != nil {
		panic(err)
	}

	cr, err := c.CreateComment(ctx, &lemmy.CreateComment{
		PostID:  postID,
		Content: "Hello from go-lemmy!",
	})
	if err != nil {
		panic(err)
	}

	cr2, err := c.CreateComment(ctx, &lemmy.CreateComment{
		PostID:   postID,
		ParentID: lemmy.NewOptCommentId(cr.CommentView.Comment.ID),
		Content:  "Reply from go-lemmy",
	})
	if err != nil {
		panic(err)
	}

	log.Printf(
		"Created comment %f and replied to it with comment %f!\n",
		cr.CommentView.Comment.ID,
		cr2.CommentView.Comment.ID,
	)
}
