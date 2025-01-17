// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	asNodeButtonFieldNames          = builder.RawFieldNames(&AsNodeButton{})
	asNodeButtonRows                = strings.Join(asNodeButtonFieldNames, ",")
	asNodeButtonRowsExpectAutoSet   = strings.Join(stringx.Remove(asNodeButtonFieldNames, "`create_time`", "`update_time`"), ",")
	asNodeButtonRowsWithPlaceHolder = strings.Join(stringx.Remove(asNodeButtonFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsNodeButtonIdPrefix = "cache:asset:asNodeButton:id:"
)

type (
	asNodeButtonModel interface {
		Insert(ctx context.Context, data *AsNodeButton) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsNodeButton, error)
		Update(ctx context.Context, data *AsNodeButton) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsNodeButtonModel struct {
		sqlc.CachedConn
		table string
	}

	AsNodeButton struct {
		Id       string         `db:"id"`        // 主键id
		NodeId   sql.NullString `db:"node_id"`   // as_proc_node表的id字段，不是那张表的node_id字段
		ButtonId sql.NullString `db:"button_id"` // 配置的按钮id
		Action   sql.NullString `db:"action"`    // 访问的接口如果没有指定的话，就是as_button表中的默认值
	}
)

func newAsNodeButtonModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsNodeButtonModel {
	return &defaultAsNodeButtonModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_node_button`",
	}
}

func (m *defaultAsNodeButtonModel) Insert(ctx context.Context, data *AsNodeButton) (sql.Result, error) {
	assetAsNodeButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsNodeButtonIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, asNodeButtonRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.NodeId, data.ButtonId, data.Action)
	}, assetAsNodeButtonIdKey)
	return ret, err
}

func (m *defaultAsNodeButtonModel) FindOne(ctx context.Context, id string) (*AsNodeButton, error) {
	assetAsNodeButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsNodeButtonIdPrefix, id)
	var resp AsNodeButton
	err := m.QueryRowCtx(ctx, &resp, assetAsNodeButtonIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asNodeButtonRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAsNodeButtonModel) Update(ctx context.Context, data *AsNodeButton) error {
	assetAsNodeButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsNodeButtonIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asNodeButtonRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.NodeId, data.ButtonId, data.Action, data.Id)
	}, assetAsNodeButtonIdKey)
	return err
}

func (m *defaultAsNodeButtonModel) Delete(ctx context.Context, id string) error {
	assetAsNodeButtonIdKey := fmt.Sprintf("%s%v", cacheAssetAsNodeButtonIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsNodeButtonIdKey)
	return err
}

func (m *defaultAsNodeButtonModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsNodeButtonIdPrefix, primary)
}

func (m *defaultAsNodeButtonModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asNodeButtonRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsNodeButtonModel) tableName() string {
	return m.table
}
