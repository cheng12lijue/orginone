// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/aslayer"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsLayer is the model entity for the AsLayer schema.
type AsLayer struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// Layer holds the value of the "layer" field.
	// 层级
	Layer int64 `json:"layer"`
	// Width holds the value of the "width" field.
	// 宽度
	Width int64 `json:"width"`
	// GroupID holds the value of the "group_id" field.
	// 宽度
	GroupID int64 `json:"groupId"`
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
	// The values are being populated by the AsLayerQuery when eager-loading is set.
	Edges AsLayerEdges `json:"edges"`
}

// AsLayerEdges holds the relations/edges for other nodes in the graph.
type AsLayerEdges struct {
	// Group holds the value of the group edge.
	Group *AsAllGroup `json:"group"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsLayerEdges) GroupOrErr() (*AsAllGroup, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// The edge group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asallgroup.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsLayer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case aslayer.FieldID, aslayer.FieldLayer, aslayer.FieldWidth, aslayer.FieldGroupID, aslayer.FieldIsDeleted, aslayer.FieldStatus, aslayer.FieldCreateUser, aslayer.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case aslayer.FieldCreateTime, aslayer.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsLayer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsLayer fields.
func (al *AsLayer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case aslayer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			al.ID = int64(value.Int64)
		case aslayer.FieldLayer:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field layer", values[i])
			} else if value.Valid {
				al.Layer = value.Int64
			}
		case aslayer.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				al.Width = value.Int64
			}
		case aslayer.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				al.GroupID = value.Int64
			}
		case aslayer.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				al.IsDeleted = value.Int64
			}
		case aslayer.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				al.Status = value.Int64
			}
		case aslayer.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				al.CreateUser = value.Int64
			}
		case aslayer.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				al.UpdateUser = value.Int64
			}
		case aslayer.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				al.CreateTime = date.DateTime(value.Time)
			}
		case aslayer.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				al.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the AsLayer entity.
func (al *AsLayer) QueryGroup() *AsAllGroupQuery {
	return (&AsLayerClient{config: al.config}).QueryGroup(al)
}

// Update returns a builder for updating this AsLayer.
// Note that you need to call AsLayer.Unwrap() before calling this method if this AsLayer
// was returned from a transaction, and the transaction was committed or rolled back.
func (al *AsLayer) Update() *AsLayerUpdateOne {
	return (&AsLayerClient{config: al.config}).UpdateOne(al)
}

// Unwrap unwraps the AsLayer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (al *AsLayer) Unwrap() *AsLayer {
	tx, ok := al.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsLayer is not a transactional entity")
	}
	al.config.driver = tx.drv
	return al
}

// String implements the fmt.Stringer.
func (al *AsLayer) String() string {
	var builder strings.Builder
	builder.WriteString("AsLayer(")
	builder.WriteString(fmt.Sprintf("id=%v", al.ID))
	builder.WriteString(", layer=")
	builder.WriteString(fmt.Sprintf("%v", al.Layer))
	builder.WriteString(", width=")
	builder.WriteString(fmt.Sprintf("%v", al.Width))
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", al.GroupID))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", al.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", al.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", al.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", al.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", al.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", al.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsLayers is a parsable slice of AsLayer.
type AsLayers []*AsLayer

func (al AsLayers) config(cfg config) {
	for _i := range al {
		al[_i].config = cfg
	}
}