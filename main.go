package main

import (
	"AuthInGo/app"
)

func main() {
cfg:= app.NewConfig(":8080")
app:= app.NewApplication(cfg)

app.Run()
}