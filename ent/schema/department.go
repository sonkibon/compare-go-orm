package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("dept_name").SchemaType(map[string]string{dialect.MySQL: "varchar(40)"}).Unique(),
		field.Time("created_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now),
		field.Time("updated_at").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dept_manager", DeptManager.Type),
	}
}
