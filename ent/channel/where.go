// Code generated by ent, DO NOT EDIT.

package channel

import (
	"slack-application/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldID, id))
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldWorkspaceID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldName, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldCreatedBy, v))
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldWorkspaceID, v))
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldWorkspaceID, v))
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...int64) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...int64) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldWorkspaceID, v))
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldWorkspaceID, v))
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldWorkspaceID, v))
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldWorkspaceID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldName, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldCreatedBy, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
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
func Not(p predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		p(s.Not())
	})
}
