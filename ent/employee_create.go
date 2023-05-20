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
)

// EmployeeCreate is the builder for creating a Employee entity.
type EmployeeCreate struct {
	config
	mutation *EmployeeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBirthDate sets the "birth_date" field.
func (ec *EmployeeCreate) SetBirthDate(t time.Time) *EmployeeCreate {
	ec.mutation.SetBirthDate(t)
	return ec
}

// SetFirstName sets the "first_name" field.
func (ec *EmployeeCreate) SetFirstName(s string) *EmployeeCreate {
	ec.mutation.SetFirstName(s)
	return ec
}

// SetLastName sets the "last_name" field.
func (ec *EmployeeCreate) SetLastName(s string) *EmployeeCreate {
	ec.mutation.SetLastName(s)
	return ec
}

// SetGender sets the "gender" field.
func (ec *EmployeeCreate) SetGender(e employee.Gender) *EmployeeCreate {
	ec.mutation.SetGender(e)
	return ec
}

// SetHireDate sets the "hire_date" field.
func (ec *EmployeeCreate) SetHireDate(t time.Time) *EmployeeCreate {
	ec.mutation.SetHireDate(t)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EmployeeCreate) SetCreatedAt(t time.Time) *EmployeeCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EmployeeCreate) SetNillableCreatedAt(t *time.Time) *EmployeeCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EmployeeCreate) SetUpdatedAt(t time.Time) *EmployeeCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EmployeeCreate) SetNillableUpdatedAt(t *time.Time) *EmployeeCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetDeletedAt sets the "deleted_at" field.
func (ec *EmployeeCreate) SetDeletedAt(t time.Time) *EmployeeCreate {
	ec.mutation.SetDeletedAt(t)
	return ec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ec *EmployeeCreate) SetNillableDeletedAt(t *time.Time) *EmployeeCreate {
	if t != nil {
		ec.SetDeletedAt(*t)
	}
	return ec
}

// AddDeptManagerIDs adds the "dept_manager" edge to the DeptManager entity by IDs.
func (ec *EmployeeCreate) AddDeptManagerIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddDeptManagerIDs(ids...)
	return ec
}

// AddDeptManager adds the "dept_manager" edges to the DeptManager entity.
func (ec *EmployeeCreate) AddDeptManager(d ...*DeptManager) *EmployeeCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ec.AddDeptManagerIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (ec *EmployeeCreate) Mutation() *EmployeeMutation {
	return ec.mutation
}

// Save creates the Employee in the database.
func (ec *EmployeeCreate) Save(ctx context.Context) (*Employee, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeCreate) SaveX(ctx context.Context) *Employee {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EmployeeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EmployeeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EmployeeCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := employee.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := employee.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EmployeeCreate) check() error {
	if _, ok := ec.mutation.BirthDate(); !ok {
		return &ValidationError{Name: "birth_date", err: errors.New(`ent: missing required field "Employee.birth_date"`)}
	}
	if _, ok := ec.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Employee.first_name"`)}
	}
	if _, ok := ec.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Employee.last_name"`)}
	}
	if _, ok := ec.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Employee.gender"`)}
	}
	if v, ok := ec.mutation.Gender(); ok {
		if err := employee.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Employee.gender": %w`, err)}
		}
	}
	if _, ok := ec.mutation.HireDate(); !ok {
		return &ValidationError{Name: "hire_date", err: errors.New(`ent: missing required field "Employee.hire_date"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Employee.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Employee.updated_at"`)}
	}
	return nil
}

func (ec *EmployeeCreate) sqlSave(ctx context.Context) (*Employee, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EmployeeCreate) createSpec() (*Employee, *sqlgraph.CreateSpec) {
	var (
		_node = &Employee{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(employee.Table, sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ec.conflict
	if value, ok := ec.mutation.BirthDate(); ok {
		_spec.SetField(employee.FieldBirthDate, field.TypeTime, value)
		_node.BirthDate = value
	}
	if value, ok := ec.mutation.FirstName(); ok {
		_spec.SetField(employee.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := ec.mutation.LastName(); ok {
		_spec.SetField(employee.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := ec.mutation.Gender(); ok {
		_spec.SetField(employee.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := ec.mutation.HireDate(); ok {
		_spec.SetField(employee.FieldHireDate, field.TypeTime, value)
		_node.HireDate = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(employee.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(employee.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.DeletedAt(); ok {
		_spec.SetField(employee.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := ec.mutation.DeptManagerIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Employee.Create().
//		SetBirthDate(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EmployeeUpsert) {
//			SetBirthDate(v+v).
//		}).
//		Exec(ctx)
func (ec *EmployeeCreate) OnConflict(opts ...sql.ConflictOption) *EmployeeUpsertOne {
	ec.conflict = opts
	return &EmployeeUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Employee.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EmployeeCreate) OnConflictColumns(columns ...string) *EmployeeUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EmployeeUpsertOne{
		create: ec,
	}
}

type (
	// EmployeeUpsertOne is the builder for "upsert"-ing
	//  one Employee node.
	EmployeeUpsertOne struct {
		create *EmployeeCreate
	}

	// EmployeeUpsert is the "OnConflict" setter.
	EmployeeUpsert struct {
		*sql.UpdateSet
	}
)

// SetBirthDate sets the "birth_date" field.
func (u *EmployeeUpsert) SetBirthDate(v time.Time) *EmployeeUpsert {
	u.Set(employee.FieldBirthDate, v)
	return u
}

// UpdateBirthDate sets the "birth_date" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateBirthDate() *EmployeeUpsert {
	u.SetExcluded(employee.FieldBirthDate)
	return u
}

// SetFirstName sets the "first_name" field.
func (u *EmployeeUpsert) SetFirstName(v string) *EmployeeUpsert {
	u.Set(employee.FieldFirstName, v)
	return u
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateFirstName() *EmployeeUpsert {
	u.SetExcluded(employee.FieldFirstName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *EmployeeUpsert) SetLastName(v string) *EmployeeUpsert {
	u.Set(employee.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateLastName() *EmployeeUpsert {
	u.SetExcluded(employee.FieldLastName)
	return u
}

// SetGender sets the "gender" field.
func (u *EmployeeUpsert) SetGender(v employee.Gender) *EmployeeUpsert {
	u.Set(employee.FieldGender, v)
	return u
}

// UpdateGender sets the "gender" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateGender() *EmployeeUpsert {
	u.SetExcluded(employee.FieldGender)
	return u
}

// SetHireDate sets the "hire_date" field.
func (u *EmployeeUpsert) SetHireDate(v time.Time) *EmployeeUpsert {
	u.Set(employee.FieldHireDate, v)
	return u
}

// UpdateHireDate sets the "hire_date" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateHireDate() *EmployeeUpsert {
	u.SetExcluded(employee.FieldHireDate)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *EmployeeUpsert) SetCreatedAt(v time.Time) *EmployeeUpsert {
	u.Set(employee.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateCreatedAt() *EmployeeUpsert {
	u.SetExcluded(employee.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EmployeeUpsert) SetUpdatedAt(v time.Time) *EmployeeUpsert {
	u.Set(employee.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateUpdatedAt() *EmployeeUpsert {
	u.SetExcluded(employee.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EmployeeUpsert) SetDeletedAt(v time.Time) *EmployeeUpsert {
	u.Set(employee.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EmployeeUpsert) UpdateDeletedAt() *EmployeeUpsert {
	u.SetExcluded(employee.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EmployeeUpsert) ClearDeletedAt() *EmployeeUpsert {
	u.SetNull(employee.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Employee.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EmployeeUpsertOne) UpdateNewValues() *EmployeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Employee.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EmployeeUpsertOne) Ignore() *EmployeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EmployeeUpsertOne) DoNothing() *EmployeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EmployeeCreate.OnConflict
// documentation for more info.
func (u *EmployeeUpsertOne) Update(set func(*EmployeeUpsert)) *EmployeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EmployeeUpsert{UpdateSet: update})
	}))
	return u
}

// SetBirthDate sets the "birth_date" field.
func (u *EmployeeUpsertOne) SetBirthDate(v time.Time) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetBirthDate(v)
	})
}

// UpdateBirthDate sets the "birth_date" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateBirthDate() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateBirthDate()
	})
}

// SetFirstName sets the "first_name" field.
func (u *EmployeeUpsertOne) SetFirstName(v string) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateFirstName() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *EmployeeUpsertOne) SetLastName(v string) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateLastName() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateLastName()
	})
}

// SetGender sets the "gender" field.
func (u *EmployeeUpsertOne) SetGender(v employee.Gender) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetGender(v)
	})
}

// UpdateGender sets the "gender" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateGender() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateGender()
	})
}

// SetHireDate sets the "hire_date" field.
func (u *EmployeeUpsertOne) SetHireDate(v time.Time) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetHireDate(v)
	})
}

// UpdateHireDate sets the "hire_date" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateHireDate() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateHireDate()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *EmployeeUpsertOne) SetCreatedAt(v time.Time) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateCreatedAt() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EmployeeUpsertOne) SetUpdatedAt(v time.Time) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateUpdatedAt() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EmployeeUpsertOne) SetDeletedAt(v time.Time) *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EmployeeUpsertOne) UpdateDeletedAt() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EmployeeUpsertOne) ClearDeletedAt() *EmployeeUpsertOne {
	return u.Update(func(s *EmployeeUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *EmployeeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EmployeeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EmployeeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EmployeeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EmployeeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EmployeeCreateBulk is the builder for creating many Employee entities in bulk.
type EmployeeCreateBulk struct {
	config
	builders []*EmployeeCreate
	conflict []sql.ConflictOption
}

// Save creates the Employee entities in the database.
func (ecb *EmployeeCreateBulk) Save(ctx context.Context) ([]*Employee, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Employee, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmployeeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EmployeeCreateBulk) SaveX(ctx context.Context) []*Employee {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EmployeeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EmployeeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Employee.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EmployeeUpsert) {
//			SetBirthDate(v+v).
//		}).
//		Exec(ctx)
func (ecb *EmployeeCreateBulk) OnConflict(opts ...sql.ConflictOption) *EmployeeUpsertBulk {
	ecb.conflict = opts
	return &EmployeeUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Employee.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EmployeeCreateBulk) OnConflictColumns(columns ...string) *EmployeeUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EmployeeUpsertBulk{
		create: ecb,
	}
}

// EmployeeUpsertBulk is the builder for "upsert"-ing
// a bulk of Employee nodes.
type EmployeeUpsertBulk struct {
	create *EmployeeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Employee.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EmployeeUpsertBulk) UpdateNewValues() *EmployeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Employee.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EmployeeUpsertBulk) Ignore() *EmployeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EmployeeUpsertBulk) DoNothing() *EmployeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EmployeeCreateBulk.OnConflict
// documentation for more info.
func (u *EmployeeUpsertBulk) Update(set func(*EmployeeUpsert)) *EmployeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EmployeeUpsert{UpdateSet: update})
	}))
	return u
}

// SetBirthDate sets the "birth_date" field.
func (u *EmployeeUpsertBulk) SetBirthDate(v time.Time) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetBirthDate(v)
	})
}

// UpdateBirthDate sets the "birth_date" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateBirthDate() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateBirthDate()
	})
}

// SetFirstName sets the "first_name" field.
func (u *EmployeeUpsertBulk) SetFirstName(v string) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateFirstName() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *EmployeeUpsertBulk) SetLastName(v string) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateLastName() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateLastName()
	})
}

// SetGender sets the "gender" field.
func (u *EmployeeUpsertBulk) SetGender(v employee.Gender) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetGender(v)
	})
}

// UpdateGender sets the "gender" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateGender() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateGender()
	})
}

// SetHireDate sets the "hire_date" field.
func (u *EmployeeUpsertBulk) SetHireDate(v time.Time) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetHireDate(v)
	})
}

// UpdateHireDate sets the "hire_date" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateHireDate() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateHireDate()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *EmployeeUpsertBulk) SetCreatedAt(v time.Time) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateCreatedAt() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EmployeeUpsertBulk) SetUpdatedAt(v time.Time) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateUpdatedAt() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EmployeeUpsertBulk) SetDeletedAt(v time.Time) *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EmployeeUpsertBulk) UpdateDeletedAt() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EmployeeUpsertBulk) ClearDeletedAt() *EmployeeUpsertBulk {
	return u.Update(func(s *EmployeeUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *EmployeeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EmployeeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EmployeeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EmployeeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
