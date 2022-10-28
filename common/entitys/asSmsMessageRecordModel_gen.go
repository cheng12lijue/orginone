// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	asSmsMessageRecordFieldNames          = builder.RawFieldNames(&AsSmsMessageRecord{})
	asSmsMessageRecordRows                = strings.Join(asSmsMessageRecordFieldNames, ",")
	asSmsMessageRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(asSmsMessageRecordFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asSmsMessageRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(asSmsMessageRecordFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsSmsMessageRecordIdPrefix = "cache:asset:asSmsMessageRecord:id:"
)

type (
	asSmsMessageRecordModel interface {
		Insert(ctx context.Context, data *AsSmsMessageRecord) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsSmsMessageRecord, error)
		Update(ctx context.Context, data *AsSmsMessageRecord) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsSmsMessageRecordModel struct {
		sqlc.CachedConn
		table string
	}

	AsSmsMessageRecord struct {
		Id               int64          `db:"id"`                 // 短信消息历史记录主键
		SmsMessageId     sql.NullInt64  `db:"sms_message_id"`     // 短信消息id
		SendPhoneNumbers sql.NullString `db:"send_phone_numbers"` // 发送短信手机号
		SendTime         sql.NullString `db:"send_time"`          // 发送时间
		CreateTime       time.Time      `db:"create_time"`        // 创建时间
		UpdateTime       time.Time      `db:"update_time"`        // 修改时间
		IsDeleted        int64          `db:"is_deleted"`         // 是否已删除
		Status           sql.NullInt64  `db:"status"`             // 状态
		CreateUser       sql.NullInt64  `db:"create_user"`        // 创建用户
		UpdateUser       sql.NullInt64  `db:"update_user"`        // 修改用户
	}
)

func newAsSmsMessageRecordModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsSmsMessageRecordModel {
	return &defaultAsSmsMessageRecordModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_sms_message_record`",
	}
}

func (m *defaultAsSmsMessageRecordModel) Insert(ctx context.Context, data *AsSmsMessageRecord) (sql.Result, error) {
	assetAsSmsMessageRecordIdKey := fmt.Sprintf("%s%v", cacheAssetAsSmsMessageRecordIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, asSmsMessageRecordRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.SmsMessageId, data.SendPhoneNumbers, data.SendTime, data.IsDeleted, data.Status, data.CreateUser, data.UpdateUser)
	}, assetAsSmsMessageRecordIdKey)
	return ret, err
}

func (m *defaultAsSmsMessageRecordModel) FindOne(ctx context.Context, id int64) (*AsSmsMessageRecord, error) {
	assetAsSmsMessageRecordIdKey := fmt.Sprintf("%s%v", cacheAssetAsSmsMessageRecordIdPrefix, id)
	var resp AsSmsMessageRecord
	err := m.QueryRowCtx(ctx, &resp, assetAsSmsMessageRecordIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asSmsMessageRecordRows, m.table)
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

func (m *defaultAsSmsMessageRecordModel) Update(ctx context.Context, data *AsSmsMessageRecord) error {
	assetAsSmsMessageRecordIdKey := fmt.Sprintf("%s%v", cacheAssetAsSmsMessageRecordIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asSmsMessageRecordRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.SmsMessageId, data.SendPhoneNumbers, data.SendTime, data.IsDeleted, data.Status, data.CreateUser, data.UpdateUser, data.Id)
	}, assetAsSmsMessageRecordIdKey)
	return err
}

func (m *defaultAsSmsMessageRecordModel) Delete(ctx context.Context, id int64) error {
	assetAsSmsMessageRecordIdKey := fmt.Sprintf("%s%v", cacheAssetAsSmsMessageRecordIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsSmsMessageRecordIdKey)
	return err
}

func (m *defaultAsSmsMessageRecordModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsSmsMessageRecordIdPrefix, primary)
}

func (m *defaultAsSmsMessageRecordModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asSmsMessageRecordRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsSmsMessageRecordModel) tableName() string {
	return m.table
}