package main

import (
	api "api/internal/perfin-api"
)

func main() {

	api.NewAPIService().Start()
}
