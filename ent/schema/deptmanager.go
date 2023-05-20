package schema

import "entgo.io/ent"

// DeptManager holds the schema definition for the DeptManager entity.
type DeptManager struct {
	ent.Schema
}

// Fields of the DeptManager.
func (DeptManager) Fields() []ent.Field {
	return nil
}

// Edges of the DeptManager.
func (DeptManager) Edges() []ent.Edge {
	return nil
}
