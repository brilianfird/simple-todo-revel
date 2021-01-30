package controllers

import (
	"SimpleToDoList/app/model"
	"SimpleToDoList/app/service"
	"github.com/revel/revel"
	"net/http"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) InsertToDo(t *model.ToDoRestWeb) revel.Result {
	c.Validation.Required(t)
	c.Validation.Required(t.Title)
	c.Validation.Required(t.Description)

	if c.Validation.HasErrors() {
		revel.AppLog.Error("validation error {}", c.Validation.Errors)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(model.Response{StatusCode: http.StatusBadRequest, Error: c.Validation.Errors})
	}

	res := service.CreateToDo(t.Title, t.Description)

	return c.RenderJSON(model.Response{
		StatusCode: 200,
		Error:      nil,
		Data:       res,
	})
}

func (c App) GetAllToDo() revel.Result {
	todos := service.GetAllToDo()
	return c.RenderJSON(model.Response{
		StatusCode: 200,
		Error: nil,
		Data: todos,
	})
}
