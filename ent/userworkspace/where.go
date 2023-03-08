// Code generated by ent, DO NOT EDIT.

package userworkspace

import (
	"slack-application/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldUserID, v))
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldWorkspaceID, v))
}

// RoleType applies equality check predicate on the "role_type" field. It's identical to RoleTypeEQ.
func RoleType(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldRoleType, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLTE(FieldUserID, v))
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldWorkspaceID, v))
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNEQ(FieldWorkspaceID, v))
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNotIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGT(FieldWorkspaceID, v))
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGTE(FieldWorkspaceID, v))
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLT(FieldWorkspaceID, v))
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v int64) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLTE(FieldWorkspaceID, v))
}

// RoleTypeEQ applies the EQ predicate on the "role_type" field.
func RoleTypeEQ(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEQ(FieldRoleType, v))
}

// RoleTypeNEQ applies the NEQ predicate on the "role_type" field.
func RoleTypeNEQ(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNEQ(FieldRoleType, v))
}

// RoleTypeIn applies the In predicate on the "role_type" field.
func RoleTypeIn(vs ...string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldIn(FieldRoleType, vs...))
}

// RoleTypeNotIn applies the NotIn predicate on the "role_type" field.
func RoleTypeNotIn(vs ...string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldNotIn(FieldRoleType, vs...))
}

// RoleTypeGT applies the GT predicate on the "role_type" field.
func RoleTypeGT(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGT(FieldRoleType, v))
}

// RoleTypeGTE applies the GTE predicate on the "role_type" field.
func RoleTypeGTE(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldGTE(FieldRoleType, v))
}

// RoleTypeLT applies the LT predicate on the "role_type" field.
func RoleTypeLT(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLT(FieldRoleType, v))
}

// RoleTypeLTE applies the LTE predicate on the "role_type" field.
func RoleTypeLTE(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldLTE(FieldRoleType, v))
}

// RoleTypeContains applies the Contains predicate on the "role_type" field.
func RoleTypeContains(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldContains(FieldRoleType, v))
}

// RoleTypeHasPrefix applies the HasPrefix predicate on the "role_type" field.
func RoleTypeHasPrefix(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldHasPrefix(FieldRoleType, v))
}

// RoleTypeHasSuffix applies the HasSuffix predicate on the "role_type" field.
func RoleTypeHasSuffix(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldHasSuffix(FieldRoleType, v))
}

// RoleTypeEqualFold applies the EqualFold predicate on the "role_type" field.
func RoleTypeEqualFold(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldEqualFold(FieldRoleType, v))
}

// RoleTypeContainsFold applies the ContainsFold predicate on the "role_type" field.
func RoleTypeContainsFold(v string) predicate.UserWorkspace {
	return predicate.UserWorkspace(sql.FieldContainsFold(FieldRoleType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserWorkspace) predicate.UserWorkspace {
	return predicate.UserWorkspace(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserWorkspace) predicate.UserWorkspace {
	return predicate.UserWorkspace(func(s *sql.Selector) {
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
func Not(p predicate.UserWorkspace) predicate.UserWorkspace {
	return predicate.UserWorkspace(func(s *sql.Selector) {
		p(s.Not())
	})
}