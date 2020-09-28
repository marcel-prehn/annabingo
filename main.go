package main

import (
	"marcel.works/bingo-backend/app"
)

func main() {
	a := app.App{}
	a.Start(":8000")
}
