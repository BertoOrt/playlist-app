package controllers

import ("github.com/revel/revel"; "playlist-app/app/routes")

type Search struct {
	*revel.Controller
}

func (c Search) Index() revel.Result {
	return c.Render(routes.Search.Index())
}
