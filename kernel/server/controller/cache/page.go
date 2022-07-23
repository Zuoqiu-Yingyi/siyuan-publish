package cache

import (
	"publish/server/controller"
)

type page struct {
	controller.Page
}

var (
	P *page
)

func init() {
	P = &page{}
}
