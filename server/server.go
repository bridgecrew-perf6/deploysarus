package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cjaewon/deploysarus/utils/config"
)

// Listen starts server for catch hooks
func Listen() error {
	// platform := config.GetString("platform")
	address := config.GetString("setting.server.address")
	port := ":" + strconv.Itoa(config.GetInt("setting.server.port"))
	path := config.GetString("setting.server.path")

	fmt.Println(address, port, path)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

	})

	if err := http.ListenAndServe(address+port, nil); err != nil {
		return err
	}

	return nil
}
