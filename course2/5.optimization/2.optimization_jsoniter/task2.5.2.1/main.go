package main

import (
	"encoding/json"
	"fmt"

	jsointer "github.com/json-iterator/go"
	"github.com/mailru/easyjson"

	"student.vkusvill.ru/vovangy/go-course/course2/5.optimization/2.optimization_jsoniter/task2.5.2.1/model"
)

type MarshalUnMarshaler interface {
	Marshal(user *model.User) ([]byte, error)
	Unmarshal([]byte) (model.User, error)
}

type StandartJson struct{}

func (s *StandartJson) Marshal(user *model.User) ([]byte, error) {
	return json.Marshal(user)
}

func (s *StandartJson) Unmarshal(data []byte) (model.User, error) {
	var user model.User
	err := json.Unmarshal(data, &user)
	return user, err
}

type EasyJson struct{}

func (s *EasyJson) Marshal(user *model.User) ([]byte, error) {
	return easyjson.Marshal(user)
}

func (s *EasyJson) Unmarshal(data []byte) (model.User, error) {
	var user model.User
	err := easyjson.Unmarshal(data, &user)
	return user, err
}

type Jsointer struct{}

func (s *Jsointer) Marshal(user *model.User) ([]byte, error) {
	my := jsointer.ConfigCompatibleWithStandardLibrary
	return my.Marshal(user)
}

func (s *Jsointer) Unmarshal(data []byte) (model.User, error) {
	var user model.User
	my := jsointer.ConfigCompatibleWithStandardLibrary
	err := my.Unmarshal(data, &user)
	return user, err
}

func GenerateUser(count int) []model.User {
	users := make([]model.User, count)
	for i := 0; i < count; i++ {
		users[i] = model.User{
			ID:       i,
			Username: fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("pass%d", i),
			Age:      i,
			Email:    fmt.Sprintf("email%d", i),
		}
	}
	return users
}
