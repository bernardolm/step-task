package repository

import (
	"context"
	"strings"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/bernardolm/step-task/pkg/domain/model"
	"github.com/jaswdr/faker"
	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
)

var fake = faker.New()

var newUser = func() *model.User {
	m := model.User{}
	m.Name = fake.Person().Name()
	return &m
}

var newTask = func() *model.Task {
	m := model.Task{}
	m.Description = strings.Join(fake.Lorem().Words(10), " ")
	return &m
}

func Seed(ctx context.Context, sr contract.StateRepository,
	tr contract.TaskRepository, ur contract.UserRepository,
) error {
	log.Debug("seeding...")

	s1 := model.State{Label: fake.RandomStringWithLength(5)}
	if err := sr.Create(ctx, &s1); err != nil {
		return err
	}
	log.Debug(pp.Sprintln("new state: ", s1))

	if v, err := sr.FindAll(ctx); err != nil {
		return err
	} else {
		log.Debug(pp.Sprintln("states: ", v))
	}

	//-------------------------------------------------------------------------

	u1 := newUser()
	if err := ur.Create(ctx, u1); err != nil {
		return err
	}
	log.Debug(pp.Sprintln("new user: ", u1))

	if v, err := ur.FindAll(ctx); err != nil {
		return err
	} else {
		log.Debug(pp.Sprintln("users: ", v))
	}

	//-------------------------------------------------------------------------

	t1 := newTask()
	// t1.User = *u1
	t1.UserID = u1.ID
	// t1.State = s1
	if err := tr.Create(ctx, t1); err != nil {
		return err
	}
	log.Debug(pp.Sprintln("new task: ", t1))

	t2 := newTask()
	// t2.User = *u1
	t2.UserID = u1.ID
	// t2.State = s1
	if err := tr.Create(ctx, t2); err != nil {
		return err
	}
	log.Debug(pp.Sprintln("new task: ", t2))

	t3 := newTask()
	// t3.User = *u1
	t3.UserID = u1.ID
	if err := tr.Create(ctx, t3); err != nil {
		return err
	}
	log.Debug(pp.Sprintln("new task: ", t3))

	t4 := newTask()
	// t4.User = *u1
	t4.UserID = u1.ID
	// t4.State = s1
	if err := tr.Create(ctx, t4); err != nil {
		return err
	}
	log.Debug(pp.Sprintln("new task: ", t4))

	if v, err := tr.FindAll(ctx); err != nil {
		return err
	} else {
		log.Debug(pp.Sprintln("tasks: ", v))
	}

	return nil
}
