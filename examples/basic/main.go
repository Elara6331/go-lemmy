package main

import (
	"context"

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

	_, err = c.SaveUserSettings(ctx, lemmy.SaveUserSettings{
		BotAccount: lemmy.NewOptional(true),
	})
	if err != nil {
		panic(err)
	}
}
