package service

import (
	"github.com/revel/revel"
	"SimpleToDoList/app/model"
	"SimpleToDoList/app/repository"
	"time"
)

func CreateToDo(title string, description string) model.ToDoRestWeb {

	t := model.ToDo{
		Timestamp: time.Now(),
		Title: title,
		Cleared: false,
		Description: description,
	}

	return model.ToDoRestWeb(repository.SaveToDoList(t))
}

func GetAllToDo() []model.ToDoRestWeb {
	revel.AppLog.Info("hallo my boy2")
	all := repository.GetAll()

	var toReturn []model.ToDoRestWeb

	for _, val := range all {
		toReturn = append(toReturn, model.ToDoRestWeb(val))
	}

	return toReturn
}