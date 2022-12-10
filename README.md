# Go-Lemmy

Go bindings to the [Lemmy](https://join-lemmy.org) API

Example:

```go
ctx := context.Background()

c, err := lemmy.New("https://lemmy.ml")
if err != nil {
    panic(err)
}

err = c.Login(ctx, types.Login{
    UsernameOrEmail: "email@example.com",
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
```