package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeptManager holds the schema definition for the DeptManager entity.
type DeptManager struct {
	ent.Schema
}

// Fields of the DeptManager.
func (DeptManager) Fields() []ent.Field {
	return []ent.Field{
		field.Time("from_date").SchemaType(map[string]string{dialect.MySQL: "date"}),
		field.Time("to_date").SchemaType(map[string]string{dialect.MySQL: "date"}),
		field.Time("created_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now),
		field.Time("updated_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now),
	}
}

// Edges of the DeptManager.
func (DeptManager) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).Ref("dept_manager").Unique().Required(),
		edge.From("department", Department.Type).Ref("dept_manager").Unique().Required(),
	}
}
