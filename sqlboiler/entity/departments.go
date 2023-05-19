// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Department is an object representing the database table.
type Department struct {
	DeptNo    string    `boil:"dept_no" json:"dept_no" toml:"dept_no" yaml:"dept_no"`
	DeptName  string    `boil:"dept_name" json:"dept_name" toml:"dept_name" yaml:"dept_name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *departmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L departmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DepartmentColumns = struct {
	DeptNo    string
	DeptName  string
	CreatedAt string
	UpdatedAt string
}{
	DeptNo:    "dept_no",
	DeptName:  "dept_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var DepartmentTableColumns = struct {
	DeptNo    string
	DeptName  string
	CreatedAt string
	UpdatedAt string
}{
	DeptNo:    "departments.dept_no",
	DeptName:  "departments.dept_name",
	CreatedAt: "departments.created_at",
	UpdatedAt: "departments.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var DepartmentWhere = struct {
	DeptNo    whereHelperstring
	DeptName  whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	DeptNo:    whereHelperstring{field: "`departments`.`dept_no`"},
	DeptName:  whereHelperstring{field: "`departments`.`dept_name`"},
	CreatedAt: whereHelpertime_Time{field: "`departments`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`departments`.`updated_at`"},
}

// DepartmentRels is where relationship names are stored.
var DepartmentRels = struct {
	DeptNoDeptEmps     string
	DeptNoDeptManagers string
}{
	DeptNoDeptEmps:     "DeptNoDeptEmps",
	DeptNoDeptManagers: "DeptNoDeptManagers",
}

// departmentR is where relationships are stored.
type departmentR struct {
	DeptNoDeptEmps     DeptEmpSlice     `boil:"DeptNoDeptEmps" json:"DeptNoDeptEmps" toml:"DeptNoDeptEmps" yaml:"DeptNoDeptEmps"`
	DeptNoDeptManagers DeptManagerSlice `boil:"DeptNoDeptManagers" json:"DeptNoDeptManagers" toml:"DeptNoDeptManagers" yaml:"DeptNoDeptManagers"`
}

// NewStruct creates a new relationship struct
func (*departmentR) NewStruct() *departmentR {
	return &departmentR{}
}

func (r *departmentR) GetDeptNoDeptEmps() DeptEmpSlice {
	if r == nil {
		return nil
	}
	return r.DeptNoDeptEmps
}

func (r *departmentR) GetDeptNoDeptManagers() DeptManagerSlice {
	if r == nil {
		return nil
	}
	return r.DeptNoDeptManagers
}

// departmentL is where Load methods for each relationship are stored.
type departmentL struct{}

var (
	departmentAllColumns            = []string{"dept_no", "dept_name", "created_at", "updated_at"}
	departmentColumnsWithoutDefault = []string{"dept_no", "dept_name", "created_at", "updated_at"}
	departmentColumnsWithDefault    = []string{}
	departmentPrimaryKeyColumns     = []string{"dept_no"}
	departmentGeneratedColumns      = []string{}
)

type (
	// DepartmentSlice is an alias for a slice of pointers to Department.
	// This should almost always be used instead of []Department.
	DepartmentSlice []*Department
	// DepartmentHook is the signature for custom Department hook methods
	DepartmentHook func(context.Context, boil.ContextExecutor, *Department) error

	departmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	departmentType                 = reflect.TypeOf(&Department{})
	departmentMapping              = queries.MakeStructMapping(departmentType)
	departmentPrimaryKeyMapping, _ = queries.BindMapping(departmentType, departmentMapping, departmentPrimaryKeyColumns)
	departmentInsertCacheMut       sync.RWMutex
	departmentInsertCache          = make(map[string]insertCache)
	departmentUpdateCacheMut       sync.RWMutex
	departmentUpdateCache          = make(map[string]updateCache)
	departmentUpsertCacheMut       sync.RWMutex
	departmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var departmentAfterSelectHooks []DepartmentHook

var departmentBeforeInsertHooks []DepartmentHook
var departmentAfterInsertHooks []DepartmentHook

var departmentBeforeUpdateHooks []DepartmentHook
var departmentAfterUpdateHooks []DepartmentHook

var departmentBeforeDeleteHooks []DepartmentHook
var departmentAfterDeleteHooks []DepartmentHook

var departmentBeforeUpsertHooks []DepartmentHook
var departmentAfterUpsertHooks []DepartmentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Department) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Department) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Department) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Department) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Department) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Department) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Department) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Department) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Department) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range departmentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDepartmentHook registers your hook function for all future operations.
func AddDepartmentHook(hookPoint boil.HookPoint, departmentHook DepartmentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		departmentAfterSelectHooks = append(departmentAfterSelectHooks, departmentHook)
	case boil.BeforeInsertHook:
		departmentBeforeInsertHooks = append(departmentBeforeInsertHooks, departmentHook)
	case boil.AfterInsertHook:
		departmentAfterInsertHooks = append(departmentAfterInsertHooks, departmentHook)
	case boil.BeforeUpdateHook:
		departmentBeforeUpdateHooks = append(departmentBeforeUpdateHooks, departmentHook)
	case boil.AfterUpdateHook:
		departmentAfterUpdateHooks = append(departmentAfterUpdateHooks, departmentHook)
	case boil.BeforeDeleteHook:
		departmentBeforeDeleteHooks = append(departmentBeforeDeleteHooks, departmentHook)
	case boil.AfterDeleteHook:
		departmentAfterDeleteHooks = append(departmentAfterDeleteHooks, departmentHook)
	case boil.BeforeUpsertHook:
		departmentBeforeUpsertHooks = append(departmentBeforeUpsertHooks, departmentHook)
	case boil.AfterUpsertHook:
		departmentAfterUpsertHooks = append(departmentAfterUpsertHooks, departmentHook)
	}
}

// One returns a single department record from the query.
func (q departmentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Department, error) {
	o := &Department{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for departments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Department records from the query.
func (q departmentQuery) All(ctx context.Context, exec boil.ContextExecutor) (DepartmentSlice, error) {
	var o []*Department

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Department slice")
	}

	if len(departmentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Department records in the query.
func (q departmentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count departments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q departmentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if departments exists")
	}

	return count > 0, nil
}

// DeptNoDeptEmps retrieves all the dept_emp's DeptEmps with an executor via dept_no column.
func (o *Department) DeptNoDeptEmps(mods ...qm.QueryMod) deptEmpQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`dept_emp`.`dept_no`=?", o.DeptNo),
	)

	return DeptEmps(queryMods...)
}

// DeptNoDeptManagers retrieves all the dept_manager's DeptManagers with an executor via dept_no column.
func (o *Department) DeptNoDeptManagers(mods ...qm.QueryMod) deptManagerQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`dept_manager`.`dept_no`=?", o.DeptNo),
	)

	return DeptManagers(queryMods...)
}

// LoadDeptNoDeptEmps allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (departmentL) LoadDeptNoDeptEmps(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDepartment interface{}, mods queries.Applicator) error {
	var slice []*Department
	var object *Department

	if singular {
		var ok bool
		object, ok = maybeDepartment.(*Department)
		if !ok {
			object = new(Department)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDepartment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDepartment))
			}
		}
	} else {
		s, ok := maybeDepartment.(*[]*Department)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDepartment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDepartment))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &departmentR{}
		}
		args = append(args, object.DeptNo)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &departmentR{}
			}

			for _, a := range args {
				if a == obj.DeptNo {
					continue Outer
				}
			}

			args = append(args, obj.DeptNo)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`dept_emp`),
		qm.WhereIn(`dept_emp.dept_no in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load dept_emp")
	}

	var resultSlice []*DeptEmp
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice dept_emp")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on dept_emp")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dept_emp")
	}

	if len(deptEmpAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DeptNoDeptEmps = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &deptEmpR{}
			}
			foreign.R.DeptNoDepartment = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DeptNo == foreign.DeptNo {
				local.R.DeptNoDeptEmps = append(local.R.DeptNoDeptEmps, foreign)
				if foreign.R == nil {
					foreign.R = &deptEmpR{}
				}
				foreign.R.DeptNoDepartment = local
				break
			}
		}
	}

	return nil
}

// LoadDeptNoDeptManagers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (departmentL) LoadDeptNoDeptManagers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDepartment interface{}, mods queries.Applicator) error {
	var slice []*Department
	var object *Department

	if singular {
		var ok bool
		object, ok = maybeDepartment.(*Department)
		if !ok {
			object = new(Department)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDepartment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDepartment))
			}
		}
	} else {
		s, ok := maybeDepartment.(*[]*Department)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDepartment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDepartment))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &departmentR{}
		}
		args = append(args, object.DeptNo)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &departmentR{}
			}

			for _, a := range args {
				if a == obj.DeptNo {
					continue Outer
				}
			}

			args = append(args, obj.DeptNo)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`dept_manager`),
		qm.WhereIn(`dept_manager.dept_no in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load dept_manager")
	}

	var resultSlice []*DeptManager
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice dept_manager")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on dept_manager")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dept_manager")
	}

	if len(deptManagerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DeptNoDeptManagers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &deptManagerR{}
			}
			foreign.R.DeptNoDepartment = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DeptNo == foreign.DeptNo {
				local.R.DeptNoDeptManagers = append(local.R.DeptNoDeptManagers, foreign)
				if foreign.R == nil {
					foreign.R = &deptManagerR{}
				}
				foreign.R.DeptNoDepartment = local
				break
			}
		}
	}

	return nil
}

// AddDeptNoDeptEmps adds the given related objects to the existing relationships
// of the department, optionally inserting them as new records.
// Appends related to o.R.DeptNoDeptEmps.
// Sets related.R.DeptNoDepartment appropriately.
func (o *Department) AddDeptNoDeptEmps(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DeptEmp) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DeptNo = o.DeptNo
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `dept_emp` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"dept_no"}),
				strmangle.WhereClause("`", "`", 0, deptEmpPrimaryKeyColumns),
			)
			values := []interface{}{o.DeptNo, rel.EmpNo, rel.DeptNo}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DeptNo = o.DeptNo
		}
	}

	if o.R == nil {
		o.R = &departmentR{
			DeptNoDeptEmps: related,
		}
	} else {
		o.R.DeptNoDeptEmps = append(o.R.DeptNoDeptEmps, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &deptEmpR{
				DeptNoDepartment: o,
			}
		} else {
			rel.R.DeptNoDepartment = o
		}
	}
	return nil
}

// AddDeptNoDeptManagers adds the given related objects to the existing relationships
// of the department, optionally inserting them as new records.
// Appends related to o.R.DeptNoDeptManagers.
// Sets related.R.DeptNoDepartment appropriately.
func (o *Department) AddDeptNoDeptManagers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DeptManager) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DeptNo = o.DeptNo
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `dept_manager` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"dept_no"}),
				strmangle.WhereClause("`", "`", 0, deptManagerPrimaryKeyColumns),
			)
			values := []interface{}{o.DeptNo, rel.EmpNo, rel.DeptNo}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DeptNo = o.DeptNo
		}
	}

	if o.R == nil {
		o.R = &departmentR{
			DeptNoDeptManagers: related,
		}
	} else {
		o.R.DeptNoDeptManagers = append(o.R.DeptNoDeptManagers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &deptManagerR{
				DeptNoDepartment: o,
			}
		} else {
			rel.R.DeptNoDepartment = o
		}
	}
	return nil
}

// Departments retrieves all the records using an executor.
func Departments(mods ...qm.QueryMod) departmentQuery {
	mods = append(mods, qm.From("`departments`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`departments`.*"})
	}

	return departmentQuery{q}
}

// FindDepartment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDepartment(ctx context.Context, exec boil.ContextExecutor, deptNo string, selectCols ...string) (*Department, error) {
	departmentObj := &Department{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `departments` where `dept_no`=?", sel,
	)

	q := queries.Raw(query, deptNo)

	err := q.Bind(ctx, exec, departmentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from departments")
	}

	if err = departmentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return departmentObj, err
	}

	return departmentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Department) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no departments provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(departmentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	departmentInsertCacheMut.RLock()
	cache, cached := departmentInsertCache[key]
	departmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			departmentAllColumns,
			departmentColumnsWithDefault,
			departmentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(departmentType, departmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(departmentType, departmentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `departments` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `departments` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `departments` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, departmentPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into departments")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.DeptNo,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for departments")
	}

CacheNoHooks:
	if !cached {
		departmentInsertCacheMut.Lock()
		departmentInsertCache[key] = cache
		departmentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Department.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Department) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	departmentUpdateCacheMut.RLock()
	cache, cached := departmentUpdateCache[key]
	departmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			departmentAllColumns,
			departmentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update departments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `departments` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, departmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(departmentType, departmentMapping, append(wl, departmentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update departments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for departments")
	}

	if !cached {
		departmentUpdateCacheMut.Lock()
		departmentUpdateCache[key] = cache
		departmentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q departmentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for departments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for departments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DepartmentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), departmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `departments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, departmentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in department slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all department")
	}
	return rowsAff, nil
}

var mySQLDepartmentUniqueColumns = []string{
	"dept_no",
	"dept_name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Department) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no departments provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(departmentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDepartmentUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	departmentUpsertCacheMut.RLock()
	cache, cached := departmentUpsertCache[key]
	departmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			departmentAllColumns,
			departmentColumnsWithDefault,
			departmentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			departmentAllColumns,
			departmentPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("entity: unable to upsert departments, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`departments`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `departments` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(departmentType, departmentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(departmentType, departmentMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert for departments")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(departmentType, departmentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "entity: unable to retrieve unique values for departments")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "entity: unable to populate default values for departments")
	}

CacheNoHooks:
	if !cached {
		departmentUpsertCacheMut.Lock()
		departmentUpsertCache[key] = cache
		departmentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Department record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Department) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Department provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), departmentPrimaryKeyMapping)
	sql := "DELETE FROM `departments` WHERE `dept_no`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from departments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for departments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q departmentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no departmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from departments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for departments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DepartmentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(departmentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), departmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `departments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, departmentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from department slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for departments")
	}

	if len(departmentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Department) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDepartment(ctx, exec, o.DeptNo)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DepartmentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DepartmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), departmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `departments`.* FROM `departments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, departmentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in DepartmentSlice")
	}

	*o = slice

	return nil
}

// DepartmentExists checks if the Department row exists.
func DepartmentExists(ctx context.Context, exec boil.ContextExecutor, deptNo string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `departments` where `dept_no`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, deptNo)
	}
	row := exec.QueryRowContext(ctx, sql, deptNo)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if departments exists")
	}

	return exists, nil
}

// Exists checks if the Department row exists.
func (o *Department) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DepartmentExists(ctx, exec, o.DeptNo)
}