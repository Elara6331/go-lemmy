package main

import (
	"context"
	"log"

	"go.elara.ws/go-lemmy"
	"go.elara.ws/go-lemmy/types"
)

func main() {
	ctx := context.Background()

	c, err := lemmy.NewWebSocket("https://lemmygrad.ml")
	if err != nil {
		panic(err)
	}

	err = c.ClientLogin(ctx, types.Login{
		UsernameOrEmail: "user@example.com",
		Password:        `TestPwd`,
	})
	if err != nil {
		panic(err)
	}

	// If nil is passed as data, go-lemmy will just send
	// the auth token, which is all that's needed for
	// the UserJoin operation.
	c.Request(types.UserOperationUserJoin, nil)

	// Subscribe to all communities
	c.Request(types.UserOperationCommunityJoin, types.CommunityJoin{
		CommunityID: 0,
	})

	go handleErrors(c)
	handleResponses(c)
}

func handleResponses(c *lemmy.WSClient) {
	for res := range c.Responses() {
		if res.IsOneOf(types.UserOperationCRUDCreateComment) {
			var data types.CommentResponse
			err := lemmy.DecodeResponse(res.Data, &data)
			if err != nil {
				log.Println("Error decoding response:", err)
				continue
			}

			err = c.Request(types.UserOperationCreateCommentLike, types.CreateCommentLike{
				CommentID: data.CommentView.Comment.ID,
			})
			if err != nil {
				log.Println("Error decoding response:", err)
				continue
			}
		}
	}
}

func handleErrors(c *lemmy.WSClient) {
	for err := range c.Errors() {
		log.Println("Error decoding response:", err)
	}
}
