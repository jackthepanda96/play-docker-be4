package models

import "fmt"

type User struct {
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type MemoryUserModel struct {
	data []User
}

func NewUserModel() *MemoryUserModel {
	return &MemoryUserModel{data: []User{}}
}

type UserModel interface {
	GetAll() ([]User, error)
	Insert(user User) (User, error)
}

func (mu *MemoryUserModel) GetAll() ([]User, error) {
	return mu.data, nil
}

func (mu *MemoryUserModel) Insert(user User) (User, error) {
	mu.data = append(mu.data, user)
	fmt.Println(mu.data)
	return user, nil
}
