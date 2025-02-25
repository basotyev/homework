package main

import (
	"lesson13/internal"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	Сервер с базовой регистрацией в приложении через Bearer Token

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	internal.Run()
}
