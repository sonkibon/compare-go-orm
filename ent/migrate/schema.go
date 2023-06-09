// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dept_name", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "varchar(40)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// DeptManagersColumns holds the columns for the "dept_managers" table.
	DeptManagersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "from_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "to_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "department_dept_manager", Type: field.TypeInt},
		{Name: "employee_dept_manager", Type: field.TypeInt},
	}
	// DeptManagersTable holds the schema information for the "dept_managers" table.
	DeptManagersTable = &schema.Table{
		Name:       "dept_managers",
		Columns:    DeptManagersColumns,
		PrimaryKey: []*schema.Column{DeptManagersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dept_managers_departments_dept_manager",
				Columns:    []*schema.Column{DeptManagersColumns[5]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "dept_managers_employees_dept_manager",
				Columns:    []*schema.Column{DeptManagersColumns[6]},
				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "birth_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "first_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(14)"}},
		{Name: "last_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(16)"}},
		{Name: "gender", Type: field.TypeEnum, Enums: []string{"M", "F"}},
		{Name: "hire_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DepartmentsTable,
		DeptManagersTable,
		EmployeesTable,
	}
)

func init() {
	DeptManagersTable.ForeignKeys[0].RefTable = DepartmentsTable
	DeptManagersTable.ForeignKeys[1].RefTable = EmployeesTable
}
