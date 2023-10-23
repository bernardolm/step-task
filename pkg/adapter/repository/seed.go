package repository

import (
	"strings"

	"github.com/jaswdr/faker"
	"github.com/k0kubun/pp"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

var fake = faker.New()

var newUser = func() *model.User {
	return &model.User{
		ID:   fake.Int64Between(100, 900),
		Name: fake.Person().Name(),
	}
}

var newTask = func() *model.Task {
	return &model.Task{
		ID:          fake.Int64Between(100, 900),
		Description: strings.Join(fake.Lorem().Words(10), " "),
	}
}

func Seed(ur contracts.UserRepository, tr contracts.TaskRepository) error {
	u1 := newUser()
	if err := ur.Create(nil, u1); err != nil {
		return err
	}

	pp.Println(u1)

	t1 := newTask()
	t1.User = *u1
	if err := tr.Create(nil, t1); err != nil {
		return err
	}

	t2 := newTask()
	t2.User = *u1
	t2.Task = t1
	if err := tr.Create(nil, t2); err != nil {
		return err
	}

	t3 := newTask()
	t3.User = *u1
	t3.Task = t2
	if err := tr.Create(nil, t3); err != nil {
		return err
	}

	t4 := newTask()
	t4.User = *u1
	t4.Task = t3
	if err := tr.Create(nil, t4); err != nil {
		return err
	}

	tasks, err := tr.FindAll(nil)
	if err != nil {
		return err
	}

	pp.Println(tasks)

	return nil
}
