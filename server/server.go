package server

import (
	"fmt"
	"net/http"

	"github.com/cjaewon/deploysarus/utils/config"
	"github.com/cjaewon/deploysarus/utils/logger.go"
)

// Listen starts server for catch hooks
func Listen() error {
	platform := config.GetString("platform")
	address := config.GetString("server.address")
	port := config.GetInt("server.port")
	path := config.GetString("server.path")

	if platform == "github" {
		http.HandleFunc(path, githubHandler())
	}

	logger.Printlnf("Receiver Server is available at http://%s:%d", address, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil)
	return err
}
