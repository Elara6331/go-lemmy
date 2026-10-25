package main

import (
	"context"
	
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

	_, err = c.SaveUserSettings(ctx, &lemmy.SaveUserSettings{
		BotAccount: lemmy.NewOptBool(false),
	})
	if err != nil {
		panic(err)
	}
}
