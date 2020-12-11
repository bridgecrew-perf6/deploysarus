package server

import (
	"net/http"
	"strconv"

	"github.com/cjaewon/deploysarus/utils/config"
)

// Listen starts server for catch hooks
func Listen() error {
	platform := config.GetString("platform")
	address := config.GetString("server.address")
	port := ":" + strconv.Itoa(config.GetInt("server.port"))
	path := config.GetString("server.path")

	if platform == "github" {
		http.HandleFunc(path, githubHandler())
	}

	if err := http.ListenAndServe(address+port, nil); err != nil {
		return err
	}

	return nil
}
