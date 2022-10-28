// Code generated by entc, DO NOT EDIT.

package aslayer

import (
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Layer applies equality check predicate on the "layer" field. It's identical to LayerEQ.
func Layer(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLayer), v))
	})
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreateUser applies equality check predicate on the "create_user" field. It's identical to CreateUserEQ.
func CreateUser(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUser), v))
	})
}

// UpdateUser applies equality check predicate on the "update_user" field. It's identical to UpdateUserEQ.
func UpdateUser(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateUser), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), vc))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), vc))
	})
}

// LayerEQ applies the EQ predicate on the "layer" field.
func LayerEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLayer), v))
	})
}

// LayerNEQ applies the NEQ predicate on the "layer" field.
func LayerNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLayer), v))
	})
}

// LayerIn applies the In predicate on the "layer" field.
func LayerIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLayer), v...))
	})
}

// LayerNotIn applies the NotIn predicate on the "layer" field.
func LayerNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLayer), v...))
	})
}

// LayerGT applies the GT predicate on the "layer" field.
func LayerGT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLayer), v))
	})
}

// LayerGTE applies the GTE predicate on the "layer" field.
func LayerGTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLayer), v))
	})
}

// LayerLT applies the LT predicate on the "layer" field.
func LayerLT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLayer), v))
	})
}

// LayerLTE applies the LTE predicate on the "layer" field.
func LayerLTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLayer), v))
	})
}

// LayerIsNil applies the IsNil predicate on the "layer" field.
func LayerIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLayer)))
	})
}

// LayerNotNil applies the NotNil predicate on the "layer" field.
func LayerNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLayer)))
	})
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWidth), v))
	})
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWidth), v...))
	})
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWidth), v...))
	})
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWidth), v))
	})
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWidth), v))
	})
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWidth), v))
	})
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWidth), v))
	})
}

// WidthIsNil applies the IsNil predicate on the "width" field.
func WidthIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWidth)))
	})
}

// WidthNotNil applies the NotNil predicate on the "width" field.
func WidthNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWidth)))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// GroupIDIsNil applies the IsNil predicate on the "group_id" field.
func GroupIDIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGroupID)))
	})
}

// GroupIDNotNil applies the NotNil predicate on the "group_id" field.
func GroupIDNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGroupID)))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedIn applies the In predicate on the "is_deleted" field.
func IsDeletedIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedNotIn applies the NotIn predicate on the "is_deleted" field.
func IsDeletedNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedGT applies the GT predicate on the "is_deleted" field.
func IsDeletedGT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedGTE applies the GTE predicate on the "is_deleted" field.
func IsDeletedGTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLT applies the LT predicate on the "is_deleted" field.
func IsDeletedLT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLTE applies the LTE predicate on the "is_deleted" field.
func IsDeletedLTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeleted), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// CreateUserEQ applies the EQ predicate on the "create_user" field.
func CreateUserEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUser), v))
	})
}

// CreateUserNEQ applies the NEQ predicate on the "create_user" field.
func CreateUserNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateUser), v))
	})
}

// CreateUserIn applies the In predicate on the "create_user" field.
func CreateUserIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateUser), v...))
	})
}

// CreateUserNotIn applies the NotIn predicate on the "create_user" field.
func CreateUserNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateUser), v...))
	})
}

// CreateUserGT applies the GT predicate on the "create_user" field.
func CreateUserGT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateUser), v))
	})
}

// CreateUserGTE applies the GTE predicate on the "create_user" field.
func CreateUserGTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateUser), v))
	})
}

// CreateUserLT applies the LT predicate on the "create_user" field.
func CreateUserLT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateUser), v))
	})
}

// CreateUserLTE applies the LTE predicate on the "create_user" field.
func CreateUserLTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateUser), v))
	})
}

// CreateUserIsNil applies the IsNil predicate on the "create_user" field.
func CreateUserIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateUser)))
	})
}

// CreateUserNotNil applies the NotNil predicate on the "create_user" field.
func CreateUserNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateUser)))
	})
}

// UpdateUserEQ applies the EQ predicate on the "update_user" field.
func UpdateUserEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserNEQ applies the NEQ predicate on the "update_user" field.
func UpdateUserNEQ(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserIn applies the In predicate on the "update_user" field.
func UpdateUserIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateUser), v...))
	})
}

// UpdateUserNotIn applies the NotIn predicate on the "update_user" field.
func UpdateUserNotIn(vs ...int64) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateUser), v...))
	})
}

// UpdateUserGT applies the GT predicate on the "update_user" field.
func UpdateUserGT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserGTE applies the GTE predicate on the "update_user" field.
func UpdateUserGTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserLT applies the LT predicate on the "update_user" field.
func UpdateUserLT(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserLTE applies the LTE predicate on the "update_user" field.
func UpdateUserLTE(v int64) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserIsNil applies the IsNil predicate on the "update_user" field.
func UpdateUserIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateUser)))
	})
}

// UpdateUserNotNil applies the NotNil predicate on the "update_user" field.
func UpdateUserNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateUser)))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...date.DateTime) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...date.DateTime) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...date.DateTime) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...date.DateTime) predicate.AsLayer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsLayer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v date.DateTime) predicate.AsLayer {
	vc := time.Time(v)
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.AsAllGroup) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AsLayer) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AsLayer) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AsLayer) predicate.AsLayer {
	return predicate.AsLayer(func(s *sql.Selector) {
		p(s.Not())
	})
}