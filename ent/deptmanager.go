// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sonkibon/compare-go-orm/ent/department"
	"github.com/sonkibon/compare-go-orm/ent/deptmanager"
	"github.com/sonkibon/compare-go-orm/ent/employee"
)

// DeptManager is the model entity for the DeptManager schema.
type DeptManager struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FromDate holds the value of the "from_date" field.
	FromDate time.Time `json:"from_date,omitempty"`
	// ToDate holds the value of the "to_date" field.
	ToDate time.Time `json:"to_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeptManagerQuery when eager-loading is set.
	Edges                   DeptManagerEdges `json:"edges"`
	department_dept_manager *int
	employee_dept_manager   *int
	selectValues            sql.SelectValues
}

// DeptManagerEdges holds the relations/edges for other nodes in the graph.
type DeptManagerEdges struct {
	// Employee holds the value of the employee edge.
	Employee *Employee `json:"employee,omitempty"`
	// Department holds the value of the department edge.
	Department *Department `json:"department,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeptManagerEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[0] {
		if e.Employee == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeptManagerEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[1] {
		if e.Department == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeptManager) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deptmanager.FieldID:
			values[i] = new(sql.NullInt64)
		case deptmanager.FieldFromDate, deptmanager.FieldToDate, deptmanager.FieldCreatedAt, deptmanager.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deptmanager.ForeignKeys[0]: // department_dept_manager
			values[i] = new(sql.NullInt64)
		case deptmanager.ForeignKeys[1]: // employee_dept_manager
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeptManager fields.
func (dm *DeptManager) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deptmanager.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dm.ID = int(value.Int64)
		case deptmanager.FieldFromDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from_date", values[i])
			} else if value.Valid {
				dm.FromDate = value.Time
			}
		case deptmanager.FieldToDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field to_date", values[i])
			} else if value.Valid {
				dm.ToDate = value.Time
			}
		case deptmanager.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dm.CreatedAt = value.Time
			}
		case deptmanager.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dm.UpdatedAt = value.Time
			}
		case deptmanager.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field department_dept_manager", value)
			} else if value.Valid {
				dm.department_dept_manager = new(int)
				*dm.department_dept_manager = int(value.Int64)
			}
		case deptmanager.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employee_dept_manager", value)
			} else if value.Valid {
				dm.employee_dept_manager = new(int)
				*dm.employee_dept_manager = int(value.Int64)
			}
		default:
			dm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DeptManager.
// This includes values selected through modifiers, order, etc.
func (dm *DeptManager) Value(name string) (ent.Value, error) {
	return dm.selectValues.Get(name)
}

// QueryEmployee queries the "employee" edge of the DeptManager entity.
func (dm *DeptManager) QueryEmployee() *EmployeeQuery {
	return NewDeptManagerClient(dm.config).QueryEmployee(dm)
}

// QueryDepartment queries the "department" edge of the DeptManager entity.
func (dm *DeptManager) QueryDepartment() *DepartmentQuery {
	return NewDeptManagerClient(dm.config).QueryDepartment(dm)
}

// Update returns a builder for updating this DeptManager.
// Note that you need to call DeptManager.Unwrap() before calling this method if this DeptManager
// was returned from a transaction, and the transaction was committed or rolled back.
func (dm *DeptManager) Update() *DeptManagerUpdateOne {
	return NewDeptManagerClient(dm.config).UpdateOne(dm)
}

// Unwrap unwraps the DeptManager entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dm *DeptManager) Unwrap() *DeptManager {
	_tx, ok := dm.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeptManager is not a transactional entity")
	}
	dm.config.driver = _tx.drv
	return dm
}

// String implements the fmt.Stringer.
func (dm *DeptManager) String() string {
	var builder strings.Builder
	builder.WriteString("DeptManager(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dm.ID))
	builder.WriteString("from_date=")
	builder.WriteString(dm.FromDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("to_date=")
	builder.WriteString(dm.ToDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(dm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dm.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DeptManagers is a parsable slice of DeptManager.
type DeptManagers []*DeptManager
