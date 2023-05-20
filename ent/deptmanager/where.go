// Code generated by ent, DO NOT EDIT.

package deptmanager

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sonkibon/compare-go-orm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLTE(FieldID, id))
}

// FromDate applies equality check predicate on the "from_date" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldFromDate, v))
}

// ToDate applies equality check predicate on the "to_date" field. It's identical to ToDateEQ.
func ToDate(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldToDate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldUpdatedAt, v))
}

// FromDateEQ applies the EQ predicate on the "from_date" field.
func FromDateEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldFromDate, v))
}

// FromDateNEQ applies the NEQ predicate on the "from_date" field.
func FromDateNEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNEQ(FieldFromDate, v))
}

// FromDateIn applies the In predicate on the "from_date" field.
func FromDateIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldIn(FieldFromDate, vs...))
}

// FromDateNotIn applies the NotIn predicate on the "from_date" field.
func FromDateNotIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNotIn(FieldFromDate, vs...))
}

// FromDateGT applies the GT predicate on the "from_date" field.
func FromDateGT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGT(FieldFromDate, v))
}

// FromDateGTE applies the GTE predicate on the "from_date" field.
func FromDateGTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGTE(FieldFromDate, v))
}

// FromDateLT applies the LT predicate on the "from_date" field.
func FromDateLT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLT(FieldFromDate, v))
}

// FromDateLTE applies the LTE predicate on the "from_date" field.
func FromDateLTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLTE(FieldFromDate, v))
}

// ToDateEQ applies the EQ predicate on the "to_date" field.
func ToDateEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldToDate, v))
}

// ToDateNEQ applies the NEQ predicate on the "to_date" field.
func ToDateNEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNEQ(FieldToDate, v))
}

// ToDateIn applies the In predicate on the "to_date" field.
func ToDateIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldIn(FieldToDate, vs...))
}

// ToDateNotIn applies the NotIn predicate on the "to_date" field.
func ToDateNotIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNotIn(FieldToDate, vs...))
}

// ToDateGT applies the GT predicate on the "to_date" field.
func ToDateGT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGT(FieldToDate, v))
}

// ToDateGTE applies the GTE predicate on the "to_date" field.
func ToDateGTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGTE(FieldToDate, v))
}

// ToDateLT applies the LT predicate on the "to_date" field.
func ToDateLT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLT(FieldToDate, v))
}

// ToDateLTE applies the LTE predicate on the "to_date" field.
func ToDateLTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLTE(FieldToDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeptManager {
	return predicate.DeptManager(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
		step := newEmployeeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
		step := newDepartmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeptManager) predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeptManager) predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
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
func Not(p predicate.DeptManager) predicate.DeptManager {
	return predicate.DeptManager(func(s *sql.Selector) {
		p(s.Not())
	})
}