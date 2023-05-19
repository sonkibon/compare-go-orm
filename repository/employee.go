package repository

import (
	"context"

	"github.com/sonkibon/compare-go-orm/model"
)

type Employee interface {
	Find(ctx context.Context, empNo int) (*model.Employee, error)
	Select(ctx context.Context) ([]*model.Employee, error)
	Insert(ctx context.Context, employee model.Employee) error
	Update(ctx context.Context, employee model.Employee) error
	Delete(ctx context.Context, empNo int, hardDelete bool) error
	Upsert(ctx context.Context, employee model.Employee) error
}
