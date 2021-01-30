package repository

import (
	"context"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"SimpleToDoList/app/config"
	"SimpleToDoList/app/model"
)


func SaveToDoList(t model.ToDo) model.ToDo {
	_, err := config.Database.Collection("todo").InsertOne(context.TODO(), t)

	if err != nil {
		panic(err)
	}

	return t
}

func GetAll() []model.ToDo {

	var todos []model.ToDo
	find, err := config.Database.Collection("todo").Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	defer find.Close(context.TODO())

	for find.Next(context.TODO()) {
		var t model.ToDo
		find.Decode(&t)
		todos = append(todos, t)
	}
	revel.AppLog.Info("asd")
	return todos
}