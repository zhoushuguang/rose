package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	getByUserIDSQL  = "SELECT id, user_name, create_time, update_time FROM user WHERE id = ?"
	getByUserIDsSQL = "SELECT id, user_name, create_time, update_time FROM user WHERE id IN (%s)"
	insertUserSQL   = "INSERT INTO user(user_name) VALUES(?)"
)

type UserModel struct {
	ID         int64     `json:"id"`
	UserName   string    `json:"user_name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (db *DB) InsertUser(ctx context.Context, userName string) (int64, error) {
	result, err := db.db.Exec(ctx, insertUserSQL, userName)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *DB) GetUserByID(ctx context.Context, userId int64) (*UserModel, error) {
	row := db.db.QueryRow(ctx, getByUserIDSQL, userId)
	ret := &UserModel{}
	err := row.Scan(&ret.ID, &ret.UserName, &ret.CreateTime, &ret.UpdateTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ret, err
}

func (db *DB) GetUserByIDs(ctx context.Context, userIds []string) ([]*UserModel, error) {
	ids := strings.Join(userIds, ",")
	sqlStr := fmt.Sprintf(getByUserIDsSQL, ids)
	rows, err := db.db.Query(ctx, sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := make([]*UserModel, 0, len(userIds))
	for rows.Next() {
		user := &UserModel{}
		err = rows.Scan(&user.ID, &user.UserName, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
		ret = append(ret, user)
	}
	return ret, nil
}
