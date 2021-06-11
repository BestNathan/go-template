package main

import (
	"fmt"
	"go-template/internal/app/example"
	"go-template/internal/app/example2"
	"go-template/internal/modules/config"
	"go-template/internal/modules/healthy"
	"go-template/internal/pkg/http"
	"go-template/internal/pkg/server"
)

func main() {

	_ = config.CM()

	// router
	r := http.NewRouter()
	r.Route(healthy.HM())

	// servers
	em := example.New()
	em2 := example2.New()

	// run
	if err := server.Serve(em, em2); err != nil {
		fmt.Printf("shutdown due to err: %v\n", err)
		return
	}

	fmt.Printf("shutdown")
}
