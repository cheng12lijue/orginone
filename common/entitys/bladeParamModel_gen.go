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
	bladeParamFieldNames          = builder.RawFieldNames(&BladeParam{})
	bladeParamRows                = strings.Join(bladeParamFieldNames, ",")
	bladeParamRowsExpectAutoSet   = strings.Join(stringx.Remove(bladeParamFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bladeParamRowsWithPlaceHolder = strings.Join(stringx.Remove(bladeParamFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetBladeParamIdPrefix = "cache:asset:bladeParam:id:"
)

type (
	bladeParamModel interface {
		Insert(ctx context.Context, data *BladeParam) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BladeParam, error)
		Update(ctx context.Context, data *BladeParam) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBladeParamModel struct {
		sqlc.CachedConn
		table string
	}

	BladeParam struct {
		Id         int64          `db:"id"`          // 主键
		ParamName  sql.NullString `db:"param_name"`  // 参数名
		ParamKey   sql.NullString `db:"param_key"`   // 参数键
		ParamValue sql.NullString `db:"param_value"` // 参数值
		Remark     sql.NullString `db:"remark"`      // 备注
		CreateUser sql.NullInt64  `db:"create_user"` // 创建人
		CreateTime sql.NullTime   `db:"create_time"` // 创建时间
		UpdateUser sql.NullInt64  `db:"update_user"` // 修改人
		UpdateTime sql.NullTime   `db:"update_time"` // 修改时间
		Status     sql.NullInt64  `db:"status"`      // 状态
		IsDeleted  int64          `db:"is_deleted"`  // 是否已删除
	}
)

func newBladeParamModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBladeParamModel {
	return &defaultBladeParamModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`blade_param`",
	}
}

func (m *defaultBladeParamModel) Insert(ctx context.Context, data *BladeParam) (sql.Result, error) {
	assetBladeParamIdKey := fmt.Sprintf("%s%v", cacheAssetBladeParamIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, bladeParamRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ParamName, data.ParamKey, data.ParamValue, data.Remark, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted)
	}, assetBladeParamIdKey)
	return ret, err
}

func (m *defaultBladeParamModel) FindOne(ctx context.Context, id int64) (*BladeParam, error) {
	assetBladeParamIdKey := fmt.Sprintf("%s%v", cacheAssetBladeParamIdPrefix, id)
	var resp BladeParam
	err := m.QueryRowCtx(ctx, &resp, assetBladeParamIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeParamRows, m.table)
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

func (m *defaultBladeParamModel) Update(ctx context.Context, data *BladeParam) error {
	assetBladeParamIdKey := fmt.Sprintf("%s%v", cacheAssetBladeParamIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bladeParamRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ParamName, data.ParamKey, data.ParamValue, data.Remark, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.Id)
	}, assetBladeParamIdKey)
	return err
}

func (m *defaultBladeParamModel) Delete(ctx context.Context, id int64) error {
	assetBladeParamIdKey := fmt.Sprintf("%s%v", cacheAssetBladeParamIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetBladeParamIdKey)
	return err
}

func (m *defaultBladeParamModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetBladeParamIdPrefix, primary)
}

func (m *defaultBladeParamModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bladeParamRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBladeParamModel) tableName() string {
	return m.table
}
