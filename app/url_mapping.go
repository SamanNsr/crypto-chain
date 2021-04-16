package app

import (
	"github.com/SamanNsr/bookstore_users-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
