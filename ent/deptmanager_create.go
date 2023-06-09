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
	"github.com/sonkibon/compare-go-orm/ent/department"
	"github.com/sonkibon/compare-go-orm/ent/deptmanager"
	"github.com/sonkibon/compare-go-orm/ent/employee"
)

// DeptManagerCreate is the builder for creating a DeptManager entity.
type DeptManagerCreate struct {
	config
	mutation *DeptManagerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetFromDate sets the "from_date" field.
func (dmc *DeptManagerCreate) SetFromDate(t time.Time) *DeptManagerCreate {
	dmc.mutation.SetFromDate(t)
	return dmc
}

// SetToDate sets the "to_date" field.
func (dmc *DeptManagerCreate) SetToDate(t time.Time) *DeptManagerCreate {
	dmc.mutation.SetToDate(t)
	return dmc
}

// SetCreatedAt sets the "created_at" field.
func (dmc *DeptManagerCreate) SetCreatedAt(t time.Time) *DeptManagerCreate {
	dmc.mutation.SetCreatedAt(t)
	return dmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dmc *DeptManagerCreate) SetNillableCreatedAt(t *time.Time) *DeptManagerCreate {
	if t != nil {
		dmc.SetCreatedAt(*t)
	}
	return dmc
}

// SetUpdatedAt sets the "updated_at" field.
func (dmc *DeptManagerCreate) SetUpdatedAt(t time.Time) *DeptManagerCreate {
	dmc.mutation.SetUpdatedAt(t)
	return dmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dmc *DeptManagerCreate) SetNillableUpdatedAt(t *time.Time) *DeptManagerCreate {
	if t != nil {
		dmc.SetUpdatedAt(*t)
	}
	return dmc
}

// SetEmployeeID sets the "employee" edge to the Employee entity by ID.
func (dmc *DeptManagerCreate) SetEmployeeID(id int) *DeptManagerCreate {
	dmc.mutation.SetEmployeeID(id)
	return dmc
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (dmc *DeptManagerCreate) SetEmployee(e *Employee) *DeptManagerCreate {
	return dmc.SetEmployeeID(e.ID)
}

// SetDepartmentID sets the "department" edge to the Department entity by ID.
func (dmc *DeptManagerCreate) SetDepartmentID(id int) *DeptManagerCreate {
	dmc.mutation.SetDepartmentID(id)
	return dmc
}

// SetDepartment sets the "department" edge to the Department entity.
func (dmc *DeptManagerCreate) SetDepartment(d *Department) *DeptManagerCreate {
	return dmc.SetDepartmentID(d.ID)
}

// Mutation returns the DeptManagerMutation object of the builder.
func (dmc *DeptManagerCreate) Mutation() *DeptManagerMutation {
	return dmc.mutation
}

// Save creates the DeptManager in the database.
func (dmc *DeptManagerCreate) Save(ctx context.Context) (*DeptManager, error) {
	dmc.defaults()
	return withHooks(ctx, dmc.sqlSave, dmc.mutation, dmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dmc *DeptManagerCreate) SaveX(ctx context.Context) *DeptManager {
	v, err := dmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmc *DeptManagerCreate) Exec(ctx context.Context) error {
	_, err := dmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmc *DeptManagerCreate) ExecX(ctx context.Context) {
	if err := dmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dmc *DeptManagerCreate) defaults() {
	if _, ok := dmc.mutation.CreatedAt(); !ok {
		v := deptmanager.DefaultCreatedAt()
		dmc.mutation.SetCreatedAt(v)
	}
	if _, ok := dmc.mutation.UpdatedAt(); !ok {
		v := deptmanager.DefaultUpdatedAt()
		dmc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dmc *DeptManagerCreate) check() error {
	if _, ok := dmc.mutation.FromDate(); !ok {
		return &ValidationError{Name: "from_date", err: errors.New(`ent: missing required field "DeptManager.from_date"`)}
	}
	if _, ok := dmc.mutation.ToDate(); !ok {
		return &ValidationError{Name: "to_date", err: errors.New(`ent: missing required field "DeptManager.to_date"`)}
	}
	if _, ok := dmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeptManager.created_at"`)}
	}
	if _, ok := dmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeptManager.updated_at"`)}
	}
	if _, ok := dmc.mutation.EmployeeID(); !ok {
		return &ValidationError{Name: "employee", err: errors.New(`ent: missing required edge "DeptManager.employee"`)}
	}
	if _, ok := dmc.mutation.DepartmentID(); !ok {
		return &ValidationError{Name: "department", err: errors.New(`ent: missing required edge "DeptManager.department"`)}
	}
	return nil
}

func (dmc *DeptManagerCreate) sqlSave(ctx context.Context) (*DeptManager, error) {
	if err := dmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dmc.mutation.id = &_node.ID
	dmc.mutation.done = true
	return _node, nil
}

func (dmc *DeptManagerCreate) createSpec() (*DeptManager, *sqlgraph.CreateSpec) {
	var (
		_node = &DeptManager{config: dmc.config}
		_spec = sqlgraph.NewCreateSpec(deptmanager.Table, sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeInt))
	)
	_spec.OnConflict = dmc.conflict
	if value, ok := dmc.mutation.FromDate(); ok {
		_spec.SetField(deptmanager.FieldFromDate, field.TypeTime, value)
		_node.FromDate = value
	}
	if value, ok := dmc.mutation.ToDate(); ok {
		_spec.SetField(deptmanager.FieldToDate, field.TypeTime, value)
		_node.ToDate = value
	}
	if value, ok := dmc.mutation.CreatedAt(); ok {
		_spec.SetField(deptmanager.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dmc.mutation.UpdatedAt(); ok {
		_spec.SetField(deptmanager.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := dmc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deptmanager.EmployeeTable,
			Columns: []string{deptmanager.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.employee_dept_manager = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dmc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deptmanager.DepartmentTable,
			Columns: []string{deptmanager.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.department_dept_manager = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeptManager.Create().
//		SetFromDate(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeptManagerUpsert) {
//			SetFromDate(v+v).
//		}).
//		Exec(ctx)
func (dmc *DeptManagerCreate) OnConflict(opts ...sql.ConflictOption) *DeptManagerUpsertOne {
	dmc.conflict = opts
	return &DeptManagerUpsertOne{
		create: dmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeptManager.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmc *DeptManagerCreate) OnConflictColumns(columns ...string) *DeptManagerUpsertOne {
	dmc.conflict = append(dmc.conflict, sql.ConflictColumns(columns...))
	return &DeptManagerUpsertOne{
		create: dmc,
	}
}

type (
	// DeptManagerUpsertOne is the builder for "upsert"-ing
	//  one DeptManager node.
	DeptManagerUpsertOne struct {
		create *DeptManagerCreate
	}

	// DeptManagerUpsert is the "OnConflict" setter.
	DeptManagerUpsert struct {
		*sql.UpdateSet
	}
)

// SetFromDate sets the "from_date" field.
func (u *DeptManagerUpsert) SetFromDate(v time.Time) *DeptManagerUpsert {
	u.Set(deptmanager.FieldFromDate, v)
	return u
}

// UpdateFromDate sets the "from_date" field to the value that was provided on create.
func (u *DeptManagerUpsert) UpdateFromDate() *DeptManagerUpsert {
	u.SetExcluded(deptmanager.FieldFromDate)
	return u
}

// SetToDate sets the "to_date" field.
func (u *DeptManagerUpsert) SetToDate(v time.Time) *DeptManagerUpsert {
	u.Set(deptmanager.FieldToDate, v)
	return u
}

// UpdateToDate sets the "to_date" field to the value that was provided on create.
func (u *DeptManagerUpsert) UpdateToDate() *DeptManagerUpsert {
	u.SetExcluded(deptmanager.FieldToDate)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeptManagerUpsert) SetCreatedAt(v time.Time) *DeptManagerUpsert {
	u.Set(deptmanager.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeptManagerUpsert) UpdateCreatedAt() *DeptManagerUpsert {
	u.SetExcluded(deptmanager.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeptManagerUpsert) SetUpdatedAt(v time.Time) *DeptManagerUpsert {
	u.Set(deptmanager.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeptManagerUpsert) UpdateUpdatedAt() *DeptManagerUpsert {
	u.SetExcluded(deptmanager.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DeptManager.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DeptManagerUpsertOne) UpdateNewValues() *DeptManagerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeptManager.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeptManagerUpsertOne) Ignore() *DeptManagerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeptManagerUpsertOne) DoNothing() *DeptManagerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeptManagerCreate.OnConflict
// documentation for more info.
func (u *DeptManagerUpsertOne) Update(set func(*DeptManagerUpsert)) *DeptManagerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeptManagerUpsert{UpdateSet: update})
	}))
	return u
}

// SetFromDate sets the "from_date" field.
func (u *DeptManagerUpsertOne) SetFromDate(v time.Time) *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetFromDate(v)
	})
}

// UpdateFromDate sets the "from_date" field to the value that was provided on create.
func (u *DeptManagerUpsertOne) UpdateFromDate() *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateFromDate()
	})
}

// SetToDate sets the "to_date" field.
func (u *DeptManagerUpsertOne) SetToDate(v time.Time) *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetToDate(v)
	})
}

// UpdateToDate sets the "to_date" field to the value that was provided on create.
func (u *DeptManagerUpsertOne) UpdateToDate() *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateToDate()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeptManagerUpsertOne) SetCreatedAt(v time.Time) *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeptManagerUpsertOne) UpdateCreatedAt() *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeptManagerUpsertOne) SetUpdatedAt(v time.Time) *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeptManagerUpsertOne) UpdateUpdatedAt() *DeptManagerUpsertOne {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *DeptManagerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeptManagerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeptManagerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeptManagerUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeptManagerUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeptManagerCreateBulk is the builder for creating many DeptManager entities in bulk.
type DeptManagerCreateBulk struct {
	config
	builders []*DeptManagerCreate
	conflict []sql.ConflictOption
}

// Save creates the DeptManager entities in the database.
func (dmcb *DeptManagerCreateBulk) Save(ctx context.Context) ([]*DeptManager, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dmcb.builders))
	nodes := make([]*DeptManager, len(dmcb.builders))
	mutators := make([]Mutator, len(dmcb.builders))
	for i := range dmcb.builders {
		func(i int, root context.Context) {
			builder := dmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeptManagerMutation)
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
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmcb *DeptManagerCreateBulk) SaveX(ctx context.Context) []*DeptManager {
	v, err := dmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmcb *DeptManagerCreateBulk) Exec(ctx context.Context) error {
	_, err := dmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmcb *DeptManagerCreateBulk) ExecX(ctx context.Context) {
	if err := dmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeptManager.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeptManagerUpsert) {
//			SetFromDate(v+v).
//		}).
//		Exec(ctx)
func (dmcb *DeptManagerCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeptManagerUpsertBulk {
	dmcb.conflict = opts
	return &DeptManagerUpsertBulk{
		create: dmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeptManager.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmcb *DeptManagerCreateBulk) OnConflictColumns(columns ...string) *DeptManagerUpsertBulk {
	dmcb.conflict = append(dmcb.conflict, sql.ConflictColumns(columns...))
	return &DeptManagerUpsertBulk{
		create: dmcb,
	}
}

// DeptManagerUpsertBulk is the builder for "upsert"-ing
// a bulk of DeptManager nodes.
type DeptManagerUpsertBulk struct {
	create *DeptManagerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeptManager.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DeptManagerUpsertBulk) UpdateNewValues() *DeptManagerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeptManager.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeptManagerUpsertBulk) Ignore() *DeptManagerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeptManagerUpsertBulk) DoNothing() *DeptManagerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeptManagerCreateBulk.OnConflict
// documentation for more info.
func (u *DeptManagerUpsertBulk) Update(set func(*DeptManagerUpsert)) *DeptManagerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeptManagerUpsert{UpdateSet: update})
	}))
	return u
}

// SetFromDate sets the "from_date" field.
func (u *DeptManagerUpsertBulk) SetFromDate(v time.Time) *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetFromDate(v)
	})
}

// UpdateFromDate sets the "from_date" field to the value that was provided on create.
func (u *DeptManagerUpsertBulk) UpdateFromDate() *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateFromDate()
	})
}

// SetToDate sets the "to_date" field.
func (u *DeptManagerUpsertBulk) SetToDate(v time.Time) *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetToDate(v)
	})
}

// UpdateToDate sets the "to_date" field to the value that was provided on create.
func (u *DeptManagerUpsertBulk) UpdateToDate() *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateToDate()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeptManagerUpsertBulk) SetCreatedAt(v time.Time) *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeptManagerUpsertBulk) UpdateCreatedAt() *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeptManagerUpsertBulk) SetUpdatedAt(v time.Time) *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeptManagerUpsertBulk) UpdateUpdatedAt() *DeptManagerUpsertBulk {
	return u.Update(func(s *DeptManagerUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *DeptManagerUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeptManagerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeptManagerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeptManagerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
