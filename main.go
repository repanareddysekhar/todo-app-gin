package main

import (
	"todo-app/Config"
	"todo-app/Routes"
)

func main() {
	Config.DatabaseConnection()

	r := Routes.SetupRouter()
	// running
	r.Run("localhost:8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
