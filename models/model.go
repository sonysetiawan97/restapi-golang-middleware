package models

import (
	"fmt"

	"example.com/v1/config"

	// _ ...
	_ "github.com/go-sql-driver/mysql"
)

// GetAllTodos ...
func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateATodo ...
func CreateATodo(todo *Todo) (err error) {
	todo.Title = "prefix_" + todo.Title + "_suffix"
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetATodo ...
func GetATodo(todo *Todo, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateATodo ...
func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Where("id = ?", id).Save(todo)
	return nil
}

// DeleteATodo ...
func DeleteATodo(todo *Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

// GetUser ...
func GetUser(user *User) (err error) {
	if err = config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error; err != nil {
		return err
	}
	return nil
}

// RegisterUser ...
func RegisterUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
