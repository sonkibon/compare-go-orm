package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.Time("birth_date").SchemaType(map[string]string{dialect.MySQL: "date"}),
		field.String("first_name").SchemaType(map[string]string{dialect.MySQL: "varchar(14)"}),
		field.String("last_name").SchemaType(map[string]string{dialect.MySQL: "varchar(16)"}),
		field.Enum("gender").Values("M", "F"),
		field.Time("hire_date").SchemaType(map[string]string{dialect.MySQL: "date"}),
		field.Time("created_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now),
		field.Time("updated_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now),
		field.Time("deleted_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dept_manager", DeptManager.Type),
	}
}
