package main

import (
	"os"

	"github.com/Sceptyre/go-five9-scim/internal/httpserver"
	"github.com/Sceptyre/go-five9-scim/internal/sync"
	"github.com/Sceptyre/go-five9-scim/pkg/five9"
)

func main() {
	five9.Login(os.Getenv("FIVE9_USERNAME"), os.Getenv("FIVE9_PASSWORD"))

	sync.SyncFive9()
	go sync.Sync()

	httpserver.StartServer()
}
