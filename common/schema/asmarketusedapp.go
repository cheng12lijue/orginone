// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketusedapp"
	"orginone/common/schema/asuser"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsMarketUsedApp is the model entity for the AsMarketUsedApp schema.
type AsMarketUsedApp struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// AppID holds the value of the "app_id" field.
	// 应用id
	AppID int64 `json:"appId"`
	// UserID holds the value of the "user_id" field.
	// 用户id
	UserID int64 `json:"userId"`
	// Sort holds the value of the "sort" field.
	// 排序
	Sort int64 `json:"sort"`
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
	// The values are being populated by the AsMarketUsedAppQuery when eager-loading is set.
	Edges AsMarketUsedAppEdges `json:"edges"`
}

// AsMarketUsedAppEdges holds the relations/edges for other nodes in the graph.
type AsMarketUsedAppEdges struct {
	// Appx holds the value of the appx edge.
	Appx *AsMarketApp `json:"appx"`
	// Userx holds the value of the userx edge.
	Userx *AsUser `json:"userx"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AppxOrErr returns the Appx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketUsedAppEdges) AppxOrErr() (*AsMarketApp, error) {
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

// UserxOrErr returns the Userx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketUsedAppEdges) UserxOrErr() (*AsUser, error) {
	if e.loadedTypes[1] {
		if e.Userx == nil {
			// The edge userx was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asuser.Label}
		}
		return e.Userx, nil
	}
	return nil, &NotLoadedError{edge: "userx"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsMarketUsedApp) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asmarketusedapp.FieldID, asmarketusedapp.FieldAppID, asmarketusedapp.FieldUserID, asmarketusedapp.FieldSort, asmarketusedapp.FieldIsDeleted, asmarketusedapp.FieldStatus, asmarketusedapp.FieldCreateUser, asmarketusedapp.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asmarketusedapp.FieldCreateTime, asmarketusedapp.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsMarketUsedApp", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsMarketUsedApp fields.
func (amua *AsMarketUsedApp) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asmarketusedapp.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			amua.ID = int64(value.Int64)
		case asmarketusedapp.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				amua.AppID = value.Int64
			}
		case asmarketusedapp.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				amua.UserID = value.Int64
			}
		case asmarketusedapp.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				amua.Sort = value.Int64
			}
		case asmarketusedapp.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				amua.IsDeleted = value.Int64
			}
		case asmarketusedapp.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				amua.Status = value.Int64
			}
		case asmarketusedapp.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				amua.CreateUser = value.Int64
			}
		case asmarketusedapp.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				amua.UpdateUser = value.Int64
			}
		case asmarketusedapp.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				amua.CreateTime = date.DateTime(value.Time)
			}
		case asmarketusedapp.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				amua.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryAppx queries the "appx" edge of the AsMarketUsedApp entity.
func (amua *AsMarketUsedApp) QueryAppx() *AsMarketAppQuery {
	return (&AsMarketUsedAppClient{config: amua.config}).QueryAppx(amua)
}

// QueryUserx queries the "userx" edge of the AsMarketUsedApp entity.
func (amua *AsMarketUsedApp) QueryUserx() *AsUserQuery {
	return (&AsMarketUsedAppClient{config: amua.config}).QueryUserx(amua)
}

// Update returns a builder for updating this AsMarketUsedApp.
// Note that you need to call AsMarketUsedApp.Unwrap() before calling this method if this AsMarketUsedApp
// was returned from a transaction, and the transaction was committed or rolled back.
func (amua *AsMarketUsedApp) Update() *AsMarketUsedAppUpdateOne {
	return (&AsMarketUsedAppClient{config: amua.config}).UpdateOne(amua)
}

// Unwrap unwraps the AsMarketUsedApp entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (amua *AsMarketUsedApp) Unwrap() *AsMarketUsedApp {
	tx, ok := amua.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsMarketUsedApp is not a transactional entity")
	}
	amua.config.driver = tx.drv
	return amua
}

// String implements the fmt.Stringer.
func (amua *AsMarketUsedApp) String() string {
	var builder strings.Builder
	builder.WriteString("AsMarketUsedApp(")
	builder.WriteString(fmt.Sprintf("id=%v", amua.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", amua.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", amua.UserID))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", amua.Sort))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", amua.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", amua.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", amua.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", amua.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", amua.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", amua.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsMarketUsedApps is a parsable slice of AsMarketUsedApp.
type AsMarketUsedApps []*AsMarketUsedApp

func (amua AsMarketUsedApps) config(cfg config) {
	for _i := range amua {
		amua[_i].config = cfg
	}
}
