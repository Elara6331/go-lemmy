// This example shows how to save a login token so the user doesn't have to log in every time
// IMPORTANT: This is not a secure way to save your token. This is an informational example ONLY!
package main

import (
	"context"
	"os"

	"go.elara.ws/go-lemmy/v4"
)

const (
	filename  = "token.txt"
	serverURL = "https://lemmy.ml"
)

func main() {
	token, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c *lemmy.Client
	if len(token) == 0 {
		c, err = lemmy.NewWithLogin(serverURL, &lemmy.Login{
			UsernameOrEmail: "user@example.com",
			Password:        `TestPwd`,
			StayLoggedIn:    lemmy.NewOptBool(true),
		})
		if err != nil {
			panic(err)
		}

		tok, err := c.Token()
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filename, []byte(tok), 0o644)
		if err != nil {
			panic(err)
		}
	} else {
		c, err = lemmy.NewClient(serverURL, lemmy.Token(token))
		if err != nil {
			panic(err)
		}
	}

	user, err := c.GetMyUser(context.Background())
	if err != nil {
		panic(err)
	}

	println("username:", user.LocalUserView.Person.Name)
}
