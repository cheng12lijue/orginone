// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappkeysecret"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsMarketAppKeySecret is the model entity for the AsMarketAppKeySecret schema.
type AsMarketAppKeySecret struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// AppID holds the value of the "app_id" field.
	// app_id主键
	AppID int64 `json:"appId"`
	// AppKey holds the value of the "app_key" field.
	// 应用id,随机生成
	AppKey string `json:"appKey"`
	// AppSecret holds the value of the "app_secret" field.
	// 应用加密密钥,随机生成
	AppSecret string `json:"appSecret"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除
	IsDeleted int64 `json:"isDeleted"`
	// Status holds the value of the "status" field.
	// 状态
	Status int64 `json:"status"`
	// CreateUser holds the value of the "create_user" field.
	// 创建用户
	CreateUser int64 `json:"createUser"`
	// UpdateUser holds the value of the "update_user" field.
	// 更新用户
	UpdateUser int64 `json:"updateUser"`
	// CreateTime holds the value of the "create_time" field.
	// 创建时间
	CreateTime date.DateTime `json:"createTime"`
	// UpdateTime holds the value of the "update_time" field.
	// 更新时间
	UpdateTime date.DateTime `json:"updateTime"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AsMarketAppKeySecretQuery when eager-loading is set.
	Edges AsMarketAppKeySecretEdges `json:"edges"`
}

// AsMarketAppKeySecretEdges holds the relations/edges for other nodes in the graph.
type AsMarketAppKeySecretEdges struct {
	// Appx holds the value of the appx edge.
	Appx *AsMarketApp `json:"appx"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AppxOrErr returns the Appx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketAppKeySecretEdges) AppxOrErr() (*AsMarketApp, error) {
	if e.loadedTypes[0] {
		if e.Appx == nil {
			// The edge appx was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asmarketapp.Label}
		}
		return e.Appx, nil
	}
	return nil, &NotLoadedError{edge: "appx"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsMarketAppKeySecret) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asmarketappkeysecret.FieldID, asmarketappkeysecret.FieldAppID, asmarketappkeysecret.FieldIsDeleted, asmarketappkeysecret.FieldStatus, asmarketappkeysecret.FieldCreateUser, asmarketappkeysecret.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asmarketappkeysecret.FieldAppKey, asmarketappkeysecret.FieldAppSecret:
			values[i] = new(sql.NullString)
		case asmarketappkeysecret.FieldCreateTime, asmarketappkeysecret.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsMarketAppKeySecret", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsMarketAppKeySecret fields.
func (amaks *AsMarketAppKeySecret) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asmarketappkeysecret.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			amaks.ID = int64(value.Int64)
		case asmarketappkeysecret.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				amaks.AppID = value.Int64
			}
		case asmarketappkeysecret.FieldAppKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_key", values[i])
			} else if value.Valid {
				amaks.AppKey = value.String
			}
		case asmarketappkeysecret.FieldAppSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_secret", values[i])
			} else if value.Valid {
				amaks.AppSecret = value.String
			}
		case asmarketappkeysecret.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				amaks.IsDeleted = value.Int64
			}
		case asmarketappkeysecret.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				amaks.Status = value.Int64
			}
		case asmarketappkeysecret.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				amaks.CreateUser = value.Int64
			}
		case asmarketappkeysecret.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				amaks.UpdateUser = value.Int64
			}
		case asmarketappkeysecret.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				amaks.CreateTime = date.DateTime(value.Time)
			}
		case asmarketappkeysecret.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				amaks.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryAppx queries the "appx" edge of the AsMarketAppKeySecret entity.
func (amaks *AsMarketAppKeySecret) QueryAppx() *AsMarketAppQuery {
	return (&AsMarketAppKeySecretClient{config: amaks.config}).QueryAppx(amaks)
}

// Update returns a builder for updating this AsMarketAppKeySecret.
// Note that you need to call AsMarketAppKeySecret.Unwrap() before calling this method if this AsMarketAppKeySecret
// was returned from a transaction, and the transaction was committed or rolled back.
func (amaks *AsMarketAppKeySecret) Update() *AsMarketAppKeySecretUpdateOne {
	return (&AsMarketAppKeySecretClient{config: amaks.config}).UpdateOne(amaks)
}

// Unwrap unwraps the AsMarketAppKeySecret entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (amaks *AsMarketAppKeySecret) Unwrap() *AsMarketAppKeySecret {
	tx, ok := amaks.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsMarketAppKeySecret is not a transactional entity")
	}
	amaks.config.driver = tx.drv
	return amaks
}

// String implements the fmt.Stringer.
func (amaks *AsMarketAppKeySecret) String() string {
	var builder strings.Builder
	builder.WriteString("AsMarketAppKeySecret(")
	builder.WriteString(fmt.Sprintf("id=%v", amaks.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", amaks.AppID))
	builder.WriteString(", app_key=")
	builder.WriteString(amaks.AppKey)
	builder.WriteString(", app_secret=")
	builder.WriteString(amaks.AppSecret)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", amaks.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", amaks.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", amaks.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", amaks.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", amaks.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", amaks.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsMarketAppKeySecrets is a parsable slice of AsMarketAppKeySecret.
type AsMarketAppKeySecrets []*AsMarketAppKeySecret

func (amaks AsMarketAppKeySecrets) config(cfg config) {
	for _i := range amaks {
		amaks[_i].config = cfg
	}
}