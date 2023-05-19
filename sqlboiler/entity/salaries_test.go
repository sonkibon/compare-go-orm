// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSalaries(t *testing.T) {
	t.Parallel()

	query := Salaries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSalariesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSalariesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Salaries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSalariesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SalarySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSalariesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SalaryExists(ctx, tx, o.EmpNo, o.FromDate)
	if err != nil {
		t.Errorf("Unable to check if Salary exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SalaryExists to return true, but got false.")
	}
}

func testSalariesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	salaryFound, err := FindSalary(ctx, tx, o.EmpNo, o.FromDate)
	if err != nil {
		t.Error(err)
	}

	if salaryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSalariesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Salaries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSalariesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Salaries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSalariesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	salaryOne := &Salary{}
	salaryTwo := &Salary{}
	if err = randomize.Struct(seed, salaryOne, salaryDBTypes, false, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}
	if err = randomize.Struct(seed, salaryTwo, salaryDBTypes, false, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = salaryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = salaryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Salaries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSalariesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	salaryOne := &Salary{}
	salaryTwo := &Salary{}
	if err = randomize.Struct(seed, salaryOne, salaryDBTypes, false, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}
	if err = randomize.Struct(seed, salaryTwo, salaryDBTypes, false, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = salaryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = salaryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func salaryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func salaryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Salary) error {
	*o = Salary{}
	return nil
}

func testSalariesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Salary{}
	o := &Salary{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, salaryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Salary object: %s", err)
	}

	AddSalaryHook(boil.BeforeInsertHook, salaryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	salaryBeforeInsertHooks = []SalaryHook{}

	AddSalaryHook(boil.AfterInsertHook, salaryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	salaryAfterInsertHooks = []SalaryHook{}

	AddSalaryHook(boil.AfterSelectHook, salaryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	salaryAfterSelectHooks = []SalaryHook{}

	AddSalaryHook(boil.BeforeUpdateHook, salaryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	salaryBeforeUpdateHooks = []SalaryHook{}

	AddSalaryHook(boil.AfterUpdateHook, salaryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	salaryAfterUpdateHooks = []SalaryHook{}

	AddSalaryHook(boil.BeforeDeleteHook, salaryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	salaryBeforeDeleteHooks = []SalaryHook{}

	AddSalaryHook(boil.AfterDeleteHook, salaryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	salaryAfterDeleteHooks = []SalaryHook{}

	AddSalaryHook(boil.BeforeUpsertHook, salaryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	salaryBeforeUpsertHooks = []SalaryHook{}

	AddSalaryHook(boil.AfterUpsertHook, salaryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	salaryAfterUpsertHooks = []SalaryHook{}
}

func testSalariesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSalariesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(salaryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSalaryToOneEmployeeUsingEmpNoEmployee(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Salary
	var foreign Employee

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, salaryDBTypes, false, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EmpNo = foreign.EmpNo
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.EmpNoEmployee().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.EmpNo != foreign.EmpNo {
		t.Errorf("want: %v, got %v", foreign.EmpNo, check.EmpNo)
	}

	ranAfterSelectHook := false
	AddEmployeeHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := SalarySlice{&local}
	if err = local.L.LoadEmpNoEmployee(ctx, tx, false, (*[]*Salary)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.EmpNoEmployee == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.EmpNoEmployee = nil
	if err = local.L.LoadEmpNoEmployee(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.EmpNoEmployee == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testSalaryToOneSetOpEmployeeUsingEmpNoEmployee(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Salary
	var b, c Employee

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, salaryDBTypes, false, strmangle.SetComplement(salaryPrimaryKeyColumns, salaryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, employeeDBTypes, false, strmangle.SetComplement(employeePrimaryKeyColumns, employeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, employeeDBTypes, false, strmangle.SetComplement(employeePrimaryKeyColumns, employeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Employee{&b, &c} {
		err = a.SetEmpNoEmployee(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.EmpNoEmployee != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EmpNoSalaries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EmpNo != x.EmpNo {
			t.Error("foreign key was wrong value", a.EmpNo)
		}

		if exists, err := SalaryExists(ctx, tx, a.EmpNo, a.FromDate); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testSalariesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSalariesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SalarySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSalariesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Salaries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	salaryDBTypes = map[string]string{`EmpNo`: `int`, `Salary`: `int`, `FromDate`: `date`, `ToDate`: `date`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_             = bytes.MinRead
)

func testSalariesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(salaryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(salaryAllColumns) == len(salaryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSalariesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(salaryAllColumns) == len(salaryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Salary{}
	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, salaryDBTypes, true, salaryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(salaryAllColumns, salaryPrimaryKeyColumns) {
		fields = salaryAllColumns
	} else {
		fields = strmangle.SetComplement(
			salaryAllColumns,
			salaryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SalarySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSalariesUpsert(t *testing.T) {
	t.Parallel()

	if len(salaryAllColumns) == len(salaryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLSalaryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Salary{}
	if err = randomize.Struct(seed, &o, salaryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Salary: %s", err)
	}

	count, err := Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, salaryDBTypes, false, salaryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Salary struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Salary: %s", err)
	}

	count, err = Salaries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
