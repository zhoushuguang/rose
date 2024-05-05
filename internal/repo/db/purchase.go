package db

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	getPurchaseWithPageSQL = "SELECT id, user_id, price, create_time, update_time FROM purchase"
)

type PurchaseModel struct {
	ID         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	Price      float64   `json:"price"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (db *DB) GetPurchaseWithPage(ctx context.Context, userId int64, pageSize, pageNo int) ([]*PurchaseModel, int, error) {
	var query string
	var args []interface{}

	if userId > 0 {
		query = fmt.Sprintf("%s WHERE user_id = ? ORDER BY update_time DESC LIMIT ? OFFSET ?", getPurchaseWithPageSQL)
		args = append(args, userId)
	} else {
		query = fmt.Sprintf("%s ORDER BY update_time DESC LIMIT ? OFFSET ?", getPurchaseWithPageSQL)
	}
	args = append(args, pageSize, (pageNo-1)*pageSize)

	// 执行查询
	rows, err := db.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// 读取结果
	var purchases []*PurchaseModel
	var purchase PurchaseModel
	for rows.Next() {
		err := rows.Scan(&purchase.ID, &purchase.UserId, &purchase.Price, &purchase.CreateTime, &purchase.UpdateTime)
		if err != nil {
			return nil, 0, err
		}
		purchases = append(purchases, &purchase)
	}

	// 获取总条数
	totalRows, err := db.db.QueryContext(ctx, fmt.Sprintf("SELECT count(*) FROM purchase %s", strings.TrimPrefix(query, getPurchaseWithPageSQL)), args...)
	if err != nil {
		return nil, 0, err
	}
	defer totalRows.Close()

	var totalCount int
	for totalRows.Next() {
		if err := totalRows.Scan(&totalCount); err != nil {
			return nil, 0, err
		}
	}

	return purchases, totalCount, nil
}
