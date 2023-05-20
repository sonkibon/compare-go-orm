// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/sonkibon/compare-go-orm/ent/department"
	"github.com/sonkibon/compare-go-orm/ent/deptmanager"
	"github.com/sonkibon/compare-go-orm/ent/employee"
	"github.com/sonkibon/compare-go-orm/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescCreatedAt is the schema descriptor for created_at field.
	departmentDescCreatedAt := departmentFields[1].Descriptor()
	// department.DefaultCreatedAt holds the default value on creation for the created_at field.
	department.DefaultCreatedAt = departmentDescCreatedAt.Default.(func() time.Time)
	// departmentDescUpdatedAt is the schema descriptor for updated_at field.
	departmentDescUpdatedAt := departmentFields[2].Descriptor()
	// department.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	department.DefaultUpdatedAt = departmentDescUpdatedAt.Default.(func() time.Time)
	deptmanagerFields := schema.DeptManager{}.Fields()
	_ = deptmanagerFields
	// deptmanagerDescCreatedAt is the schema descriptor for created_at field.
	deptmanagerDescCreatedAt := deptmanagerFields[2].Descriptor()
	// deptmanager.DefaultCreatedAt holds the default value on creation for the created_at field.
	deptmanager.DefaultCreatedAt = deptmanagerDescCreatedAt.Default.(func() time.Time)
	// deptmanagerDescUpdatedAt is the schema descriptor for updated_at field.
	deptmanagerDescUpdatedAt := deptmanagerFields[3].Descriptor()
	// deptmanager.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deptmanager.DefaultUpdatedAt = deptmanagerDescUpdatedAt.Default.(func() time.Time)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescCreatedAt is the schema descriptor for created_at field.
	employeeDescCreatedAt := employeeFields[5].Descriptor()
	// employee.DefaultCreatedAt holds the default value on creation for the created_at field.
	employee.DefaultCreatedAt = employeeDescCreatedAt.Default.(func() time.Time)
	// employeeDescUpdatedAt is the schema descriptor for updated_at field.
	employeeDescUpdatedAt := employeeFields[6].Descriptor()
	// employee.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	employee.DefaultUpdatedAt = employeeDescUpdatedAt.Default.(func() time.Time)
}
