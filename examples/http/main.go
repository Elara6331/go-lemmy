package main

import (
	"context"

	"go.arsenm.dev/go-lemmy"
	"go.arsenm.dev/go-lemmy/types"
)

func main() {
	ctx := context.Background()

	c, err := lemmy.New("https://lemmygrad.ml")
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

	_, err = c.SaveUserSettings(ctx, types.SaveUserSettings{
		BotAccount: types.NewOptional(true),
	})
	if err != nil {
		panic(err)
	}
}
