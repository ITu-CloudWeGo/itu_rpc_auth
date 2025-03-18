package main

import (
	auth_service "github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service/authservice"
	"log"
)

func main() {
	svr := auth_service.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
