package main

import (
	api "favorite/kitex_gen/api/favorite"
	"log"

	"github.com/karry-almond/packages/favoriteDB"
)

func main() {
	favoriteDB.Init()
	svr := api.NewServer(new(FavoriteImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
