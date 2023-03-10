// Code generated by ent, DO NOT EDIT.

package userchannel

import (
	"slack-application/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldUserID, v))
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldChannelID, v))
}

// RoleType applies equality check predicate on the "role_type" field. It's identical to RoleTypeEQ.
func RoleType(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldRoleType, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLTE(FieldUserID, v))
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldChannelID, v))
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNEQ(FieldChannelID, v))
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldIn(FieldChannelID, vs...))
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNotIn(FieldChannelID, vs...))
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGT(FieldChannelID, v))
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGTE(FieldChannelID, v))
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLT(FieldChannelID, v))
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v int64) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLTE(FieldChannelID, v))
}

// RoleTypeEQ applies the EQ predicate on the "role_type" field.
func RoleTypeEQ(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEQ(FieldRoleType, v))
}

// RoleTypeNEQ applies the NEQ predicate on the "role_type" field.
func RoleTypeNEQ(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNEQ(FieldRoleType, v))
}

// RoleTypeIn applies the In predicate on the "role_type" field.
func RoleTypeIn(vs ...string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldIn(FieldRoleType, vs...))
}

// RoleTypeNotIn applies the NotIn predicate on the "role_type" field.
func RoleTypeNotIn(vs ...string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldNotIn(FieldRoleType, vs...))
}

// RoleTypeGT applies the GT predicate on the "role_type" field.
func RoleTypeGT(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGT(FieldRoleType, v))
}

// RoleTypeGTE applies the GTE predicate on the "role_type" field.
func RoleTypeGTE(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldGTE(FieldRoleType, v))
}

// RoleTypeLT applies the LT predicate on the "role_type" field.
func RoleTypeLT(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLT(FieldRoleType, v))
}

// RoleTypeLTE applies the LTE predicate on the "role_type" field.
func RoleTypeLTE(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldLTE(FieldRoleType, v))
}

// RoleTypeContains applies the Contains predicate on the "role_type" field.
func RoleTypeContains(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldContains(FieldRoleType, v))
}

// RoleTypeHasPrefix applies the HasPrefix predicate on the "role_type" field.
func RoleTypeHasPrefix(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldHasPrefix(FieldRoleType, v))
}

// RoleTypeHasSuffix applies the HasSuffix predicate on the "role_type" field.
func RoleTypeHasSuffix(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldHasSuffix(FieldRoleType, v))
}

// RoleTypeEqualFold applies the EqualFold predicate on the "role_type" field.
func RoleTypeEqualFold(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldEqualFold(FieldRoleType, v))
}

// RoleTypeContainsFold applies the ContainsFold predicate on the "role_type" field.
func RoleTypeContainsFold(v string) predicate.UserChannel {
	return predicate.UserChannel(sql.FieldContainsFold(FieldRoleType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserChannel) predicate.UserChannel {
	return predicate.UserChannel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserChannel) predicate.UserChannel {
	return predicate.UserChannel(func(s *sql.Selector) {
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
func Not(p predicate.UserChannel) predicate.UserChannel {
	return predicate.UserChannel(func(s *sql.Selector) {
		p(s.Not())
	})
}
