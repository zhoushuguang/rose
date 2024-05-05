package repo

import (
	"context"

	"github.com/zhoushuguang/rose/internal/conf"
	"github.com/zhoushuguang/rose/internal/repo/db"
)

type Repo struct {
	db *db.DB
}

func NewRepo(conf *conf.Conf) *Repo {
	return &Repo{
		db: db.NewDB(conf.DB),
	}
}

func (r *Repo) GetUserByID(ctx context.Context, userID int64) (*db.UserModel, error) {
	return r.db.GetUserByID(ctx, userID)
}

func (r *Repo) CreateUser(ctx context.Context, userName string) (int64, error) {
	return r.db.InsertUser(ctx, userName)
}

func (r *Repo) GetUserByIDs(ctx context.Context, userIDs []string) ([]*db.UserModel, error) {
	return r.db.GetUserByIDs(ctx, userIDs)
}

func (r *Repo) GetPurchaseWithPage(ctx context.Context, userId int64, pageSize, pageNo int) ([]*db.PurchaseModel, int, error) {
	return r.db.GetPurchaseWithPage(ctx, userId, pageSize, pageNo)
}
