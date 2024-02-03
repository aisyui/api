package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Ue struct {
	ent.Schema
}

func (Ue) Fields() []ent.Field {
	return []ent.Field{  

		field.Bool("limit").
		Default(false).
		Optional(),

		field.Bool("limit_boss").
		Default(false).
		Optional(),

		field.Bool("limit_item").
		Default(false).
		Optional(),

		field.String("password").
		NotEmpty().
		Immutable().
		Sensitive(),

		field.Int("lv").
		Optional(),

		field.Int("lv_point").
		Optional(),

		field.Int("model").
		Optional(),

		field.Int("sword").
		Optional(),

		field.Int("card").
		Optional(),

		field.String("mode").
		Optional(),

		field.String("token").
		Optional().
		Sensitive(),

		field.Int("cp").
		Optional(),

		field.Int("count").
		Optional(),

		field.Int("location_x").
		Optional(),

		field.Int("location_y").
		Optional(),

		field.Int("location_z").
		Optional(),

		field.Int("location_n").
		Optional(),

		field.String("author").
		Optional(),

		field.Time("created_at").
		Immutable().
		Optional().
		Default(func() time.Time {
			return time.Now().In(jst)
		}),
	}  
}

func (Ue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("ue").
			Unique().
			Required(),

	}
}
