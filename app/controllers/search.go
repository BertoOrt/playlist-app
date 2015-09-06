package controllers

import ("github.com/revel/revel")

type Search struct {
	*revel.Controller
}

func (c Search) Index() revel.Result {
	return c.Render()
}
