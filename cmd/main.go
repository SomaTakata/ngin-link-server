package main

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/di"
)

func main() {
	router := di.Wire()
	router.Run()
}
