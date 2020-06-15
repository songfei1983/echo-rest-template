package adunit

import (
	"net/http"

	"github.com/songfei1983/go-api-server/internal/server"
)

type AdUnit struct {
	Code  string
	Sizes []Size
	Bids  []Bid
}

type Size struct {
	Height int
	Width  int
}

type Bid struct {
	Name string
}

type Handler struct{}

func InitHandler() *Handler {
	h := new(Handler)
	return h
}

func List(c *server.Context) error {
	fake := &AdUnit{
		Code:  "hoge",
		Sizes: []Size{{100, 120}},
		Bids:  []Bid{{Name: "hoge"}},
	}
	return c.JSON(http.StatusOK, fake)
}
