package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetAll(userId, groupId int) ([]core.UserInputGetAll, error)
	Invite(id, groupId int, username string) error
	KickUser(id, gropId, kickUserId int) error
}

type Group interface {
	Create(userId int, group core.Group) (int, error)
	GetAll(userId int) ([]core.Group, error)
	GetById(userId, groupId int) (core.Group, error)
	Delete(userId, gropId int) error
	Update(userId, groupId int, input core.UpdateGroupInput) error
}

type Debt interface {
	GetAll(groupId int) ([]core.Debt, []core.Debt, error)
	Update(debt core.Debt) error
}

type Purchase interface {
	Create(purchase core.Purchase) (core.CreatePurchaseResponse, error)
	GetAll(groupId int) ([]core.Purchase, error)
	GetById(id int) (core.Purchase, error)
	Update(purchase core.Purchase) error
	Delete(purchase core.Purchase) error
}

type Service struct {
	Authorization
	User
	Group
	Debt
	Purchase
}

func NewService(store *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(store.Authorization),
		Group:         NewGroupService(store.Group, store.User),
		User:          NewUserService(store.User, store.Group),
		Purchase:      NewPurchaseService(store.Purchase, store.User),
		Debt:          NewDebtService(store.Debt),
	}
}
