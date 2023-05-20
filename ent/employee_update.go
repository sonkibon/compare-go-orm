// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sonkibon/compare-go-orm/ent/deptmanager"
	"github.com/sonkibon/compare-go-orm/ent/employee"
	"github.com/sonkibon/compare-go-orm/ent/predicate"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetBirthDate sets the "birth_date" field.
func (eu *EmployeeUpdate) SetBirthDate(t time.Time) *EmployeeUpdate {
	eu.mutation.SetBirthDate(t)
	return eu
}

// SetFirstName sets the "first_name" field.
func (eu *EmployeeUpdate) SetFirstName(s string) *EmployeeUpdate {
	eu.mutation.SetFirstName(s)
	return eu
}

// SetLastName sets the "last_name" field.
func (eu *EmployeeUpdate) SetLastName(s string) *EmployeeUpdate {
	eu.mutation.SetLastName(s)
	return eu
}

// SetGender sets the "gender" field.
func (eu *EmployeeUpdate) SetGender(e employee.Gender) *EmployeeUpdate {
	eu.mutation.SetGender(e)
	return eu
}

// SetHireDate sets the "hire_date" field.
func (eu *EmployeeUpdate) SetHireDate(t time.Time) *EmployeeUpdate {
	eu.mutation.SetHireDate(t)
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EmployeeUpdate) SetCreatedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableCreatedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EmployeeUpdate) SetUpdatedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableUpdatedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetUpdatedAt(*t)
	}
	return eu
}

// SetDeletedAt sets the "deleted_at" field.
func (eu *EmployeeUpdate) SetDeletedAt(t time.Time) *EmployeeUpdate {
	eu.mutation.SetDeletedAt(t)
	return eu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eu *EmployeeUpdate) SetNillableDeletedAt(t *time.Time) *EmployeeUpdate {
	if t != nil {
		eu.SetDeletedAt(*t)
	}
	return eu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (eu *EmployeeUpdate) ClearDeletedAt() *EmployeeUpdate {
	eu.mutation.ClearDeletedAt()
	return eu
}

// AddDeptManagerIDs adds the "dept_manager" edge to the DeptManager entity by IDs.
func (eu *EmployeeUpdate) AddDeptManagerIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.AddDeptManagerIDs(ids...)
	return eu
}

// AddDeptManager adds the "dept_manager" edges to the DeptManager entity.
func (eu *EmployeeUpdate) AddDeptManager(d ...*DeptManager) *EmployeeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return eu.AddDeptManagerIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// ClearDeptManager clears all "dept_manager" edges to the DeptManager entity.
func (eu *EmployeeUpdate) ClearDeptManager() *EmployeeUpdate {
	eu.mutation.ClearDeptManager()
	return eu
}

// RemoveDeptManagerIDs removes the "dept_manager" edge to DeptManager entities by IDs.
func (eu *EmployeeUpdate) RemoveDeptManagerIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.RemoveDeptManagerIDs(ids...)
	return eu
}

// RemoveDeptManager removes "dept_manager" edges to DeptManager entities.
func (eu *EmployeeUpdate) RemoveDeptManager(d ...*DeptManager) *EmployeeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return eu.RemoveDeptManagerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmployeeUpdate) check() error {
	if v, ok := eu.mutation.Gender(); ok {
		if err := employee.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Employee.gender": %w`, err)}
		}
	}
	return nil
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.BirthDate(); ok {
		_spec.SetField(employee.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := eu.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
	}
	if value, ok := eu.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeEnum, value)
	}
	if value, ok := eu.mutation.HireDate(); ok {
		_spec.SetField(employee.FieldHireDate, field.TypeTime, value)
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.SetField(employee.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
	}
	if eu.mutation.DeletedAtCleared() {
		_spec.ClearField(employee.FieldDeletedAt, field.TypeTime)
	}
	if eu.mutation.DeptManagerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.DeptManagerTable,
			Columns: []string{employee.DeptManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedDeptManagerIDs(); len(nodes) > 0 && !eu.mutation.DeptManagerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.DeptManagerTable,
			Columns: []string{employee.DeptManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.DeptManagerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.DeptManagerTable,
			Columns: []string{employee.DeptManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetBirthDate sets the "birth_date" field.
func (euo *EmployeeUpdateOne) SetBirthDate(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetBirthDate(t)
	return euo
}

// SetFirstName sets the "first_name" field.
func (euo *EmployeeUpdateOne) SetFirstName(s string) *EmployeeUpdateOne {
	euo.mutation.SetFirstName(s)
	return euo
}

// SetLastName sets the "last_name" field.
func (euo *EmployeeUpdateOne) SetLastName(s string) *EmployeeUpdateOne {
	euo.mutation.SetLastName(s)
	return euo
}

// SetGender sets the "gender" field.
func (euo *EmployeeUpdateOne) SetGender(e employee.Gender) *EmployeeUpdateOne {
	euo.mutation.SetGender(e)
	return euo
}

// SetHireDate sets the "hire_date" field.
func (euo *EmployeeUpdateOne) SetHireDate(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetHireDate(t)
	return euo
}

// SetCreatedAt sets the "created_at" field.
func (euo *EmployeeUpdateOne) SetCreatedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableCreatedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EmployeeUpdateOne) SetUpdatedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableUpdatedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetUpdatedAt(*t)
	}
	return euo
}

// SetDeletedAt sets the "deleted_at" field.
func (euo *EmployeeUpdateOne) SetDeletedAt(t time.Time) *EmployeeUpdateOne {
	euo.mutation.SetDeletedAt(t)
	return euo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (euo *EmployeeUpdateOne) SetNillableDeletedAt(t *time.Time) *EmployeeUpdateOne {
	if t != nil {
		euo.SetDeletedAt(*t)
	}
	return euo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (euo *EmployeeUpdateOne) ClearDeletedAt() *EmployeeUpdateOne {
	euo.mutation.ClearDeletedAt()
	return euo
}

// AddDeptManagerIDs adds the "dept_manager" edge to the DeptManager entity by IDs.
func (euo *EmployeeUpdateOne) AddDeptManagerIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.AddDeptManagerIDs(ids...)
	return euo
}

// AddDeptManager adds the "dept_manager" edges to the DeptManager entity.
func (euo *EmployeeUpdateOne) AddDeptManager(d ...*DeptManager) *EmployeeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return euo.AddDeptManagerIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// ClearDeptManager clears all "dept_manager" edges to the DeptManager entity.
func (euo *EmployeeUpdateOne) ClearDeptManager() *EmployeeUpdateOne {
	euo.mutation.ClearDeptManager()
	return euo
}

// RemoveDeptManagerIDs removes the "dept_manager" edge to DeptManager entities by IDs.
func (euo *EmployeeUpdateOne) RemoveDeptManagerIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.RemoveDeptManagerIDs(ids...)
	return euo
}

// RemoveDeptManager removes "dept_manager" edges to DeptManager entities.
func (euo *EmployeeUpdateOne) RemoveDeptManager(d ...*DeptManager) *EmployeeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return euo.RemoveDeptManagerIDs(ids...)
}

// Where appends a list predicates to the EmployeeUpdate builder.
func (euo *EmployeeUpdateOne) Where(ps ...predicate.Employee) *EmployeeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmployeeUpdateOne) Select(field string, fields ...string) *EmployeeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Employee entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmployeeUpdateOne) check() error {
	if v, ok := euo.mutation.Gender(); ok {
		if err := employee.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Employee.gender": %w`, err)}
		}
	}
	return nil
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (_node *Employee, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(employee.Table, employee.Columns, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Employee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for _, f := range fields {
			if !employee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.BirthDate(); ok {
		_spec.SetField(employee.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := euo.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
	}
	if value, ok := euo.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeEnum, value)
	}
	if value, ok := euo.mutation.HireDate(); ok {
		_spec.SetField(employee.FieldHireDate, field.TypeTime, value)
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.SetField(employee.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
	}
	if euo.mutation.DeletedAtCleared() {
		_spec.ClearField(employee.FieldDeletedAt, field.TypeTime)
	}
	if euo.mutation.DeptManagerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.DeptManagerTable,
			Columns: []string{employee.DeptManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedDeptManagerIDs(); len(nodes) > 0 && !euo.mutation.DeptManagerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.DeptManagerTable,
			Columns: []string{employee.DeptManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.DeptManagerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.DeptManagerTable,
			Columns: []string{employee.DeptManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Employee{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}