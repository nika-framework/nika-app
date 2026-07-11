package main

import (
	"NikaApp/src"

	"github.com/nika-framework/nika"
	"github.com/nika-framework/nika/common/config"
	"github.com/nika-framework/nika/common/validator"
)

func main() {
	app := nika.NewApp()

	// Validation
	validator.Setup(app)
 
	//envPath := ".env"
	cfg := config.Setup(app, ".env") 
	 
	rootModule := src.NewAppModule()
	app.LoadModule(rootModule)

	port := cfg.GetString("PORT","3007")
	app.Listen(":" + port)
}
	