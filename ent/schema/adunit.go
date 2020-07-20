package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Adunit holds the schema definition for the Adunit entity.
type Adunit struct {
	ent.Schema
}

// Fields of the Adunit.
func (Adunit) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknow"),
	}
}

// Edges of the Adunit.
func (Adunit) Edges() []ent.Edge {
	return nil
}
