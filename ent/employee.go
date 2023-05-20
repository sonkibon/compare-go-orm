// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sonkibon/compare-go-orm/ent/employee"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate time.Time `json:"birth_date,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender employee.Gender `json:"gender,omitempty"`
	// HireDate holds the value of the "hire_date" field.
	HireDate time.Time `json:"hire_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeQuery when eager-loading is set.
	Edges        EmployeeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	// DeptManager holds the value of the dept_manager edge.
	DeptManager []*DeptManager `json:"dept_manager,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeptManagerOrErr returns the DeptManager value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) DeptManagerOrErr() ([]*DeptManager, error) {
	if e.loadedTypes[0] {
		return e.DeptManager, nil
	}
	return nil, &NotLoadedError{edge: "dept_manager"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			values[i] = new(sql.NullInt64)
		case employee.FieldFirstName, employee.FieldLastName, employee.FieldGender:
			values[i] = new(sql.NullString)
		case employee.FieldBirthDate, employee.FieldHireDate, employee.FieldCreatedAt, employee.FieldUpdatedAt, employee.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case employee.FieldBirthDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				e.BirthDate = value.Time
			}
		case employee.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				e.FirstName = value.String
			}
		case employee.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				e.LastName = value.String
			}
		case employee.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				e.Gender = employee.Gender(value.String)
			}
		case employee.FieldHireDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field hire_date", values[i])
			} else if value.Valid {
				e.HireDate = value.Time
			}
		case employee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case employee.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case employee.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				e.DeletedAt = value.Time
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Employee.
// This includes values selected through modifiers, order, etc.
func (e *Employee) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryDeptManager queries the "dept_manager" edge of the Employee entity.
func (e *Employee) QueryDeptManager() *DeptManagerQuery {
	return NewEmployeeClient(e.config).QueryDeptManager(e)
}

// Update returns a builder for updating this Employee.
// Note that you need to call Employee.Unwrap() before calling this method if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return NewEmployeeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Employee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("birth_date=")
	builder.WriteString(e.BirthDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(e.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(e.LastName)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", e.Gender))
	builder.WriteString(", ")
	builder.WriteString("hire_date=")
	builder.WriteString(e.HireDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(e.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee