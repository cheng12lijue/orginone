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
	actCmmnRuCaseInstFieldNames          = builder.RawFieldNames(&ActCmmnRuCaseInst{})
	actCmmnRuCaseInstRows                = strings.Join(actCmmnRuCaseInstFieldNames, ",")
	actCmmnRuCaseInstRowsExpectAutoSet   = strings.Join(stringx.Remove(actCmmnRuCaseInstFieldNames, "`create_time`", "`update_time`"), ",")
	actCmmnRuCaseInstRowsWithPlaceHolder = strings.Join(stringx.Remove(actCmmnRuCaseInstFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActCmmnRuCaseInstIDPrefix = "cache:asset:actCmmnRuCaseInst:iD:"
)

type (
	actCmmnRuCaseInstModel interface {
		Insert(ctx context.Context, data *ActCmmnRuCaseInst) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActCmmnRuCaseInst, error)
		Update(ctx context.Context, data *ActCmmnRuCaseInst) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActCmmnRuCaseInstModel struct {
		sqlc.CachedConn
		table string
	}

	ActCmmnRuCaseInst struct {
		ID             string         `db:"ID_"`
		REV            int64          `db:"REV_"`
		BUSINESSKEY    sql.NullString `db:"BUSINESS_KEY_"`
		NAME           sql.NullString `db:"NAME_"`
		PARENTID       sql.NullString `db:"PARENT_ID_"`
		CASEDEFID      sql.NullString `db:"CASE_DEF_ID_"`
		STATE          sql.NullString `db:"STATE_"`
		STARTTIME      sql.NullTime   `db:"START_TIME_"`
		STARTUSERID    sql.NullString `db:"START_USER_ID_"`
		CALLBACKID     sql.NullString `db:"CALLBACK_ID_"`
		CALLBACKTYPE   sql.NullString `db:"CALLBACK_TYPE_"`
		TENANTID       string         `db:"TENANT_ID_"`
		LOCKTIME       sql.NullTime   `db:"LOCK_TIME_"`
		ISCOMPLETEABLE byte           `db:"IS_COMPLETEABLE_"`
	}
)

func newActCmmnRuCaseInstModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActCmmnRuCaseInstModel {
	return &defaultActCmmnRuCaseInstModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_cmmn_ru_case_inst`",
	}
}

func (m *defaultActCmmnRuCaseInstModel) Insert(ctx context.Context, data *ActCmmnRuCaseInst) (sql.Result, error) {
	assetActCmmnRuCaseInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuCaseInstIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actCmmnRuCaseInstRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.BUSINESSKEY, data.NAME, data.PARENTID, data.CASEDEFID, data.STATE, data.STARTTIME, data.STARTUSERID, data.CALLBACKID, data.CALLBACKTYPE, data.TENANTID, data.LOCKTIME, data.ISCOMPLETEABLE)
	}, assetActCmmnRuCaseInstIDKey)
	return ret, err
}

func (m *defaultActCmmnRuCaseInstModel) FindOne(ctx context.Context, iD string) (*ActCmmnRuCaseInst, error) {
	assetActCmmnRuCaseInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuCaseInstIDPrefix, iD)
	var resp ActCmmnRuCaseInst
	err := m.QueryRowCtx(ctx, &resp, assetActCmmnRuCaseInstIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnRuCaseInstRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, iD)
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

func (m *defaultActCmmnRuCaseInstModel) Update(ctx context.Context, data *ActCmmnRuCaseInst) error {
	assetActCmmnRuCaseInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuCaseInstIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actCmmnRuCaseInstRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.BUSINESSKEY, data.NAME, data.PARENTID, data.CASEDEFID, data.STATE, data.STARTTIME, data.STARTUSERID, data.CALLBACKID, data.CALLBACKTYPE, data.TENANTID, data.LOCKTIME, data.ISCOMPLETEABLE, data.ID)
	}, assetActCmmnRuCaseInstIDKey)
	return err
}

func (m *defaultActCmmnRuCaseInstModel) Delete(ctx context.Context, iD string) error {
	assetActCmmnRuCaseInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuCaseInstIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActCmmnRuCaseInstIDKey)
	return err
}

func (m *defaultActCmmnRuCaseInstModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActCmmnRuCaseInstIDPrefix, primary)
}

func (m *defaultActCmmnRuCaseInstModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnRuCaseInstRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActCmmnRuCaseInstModel) tableName() string {
	return m.table
}