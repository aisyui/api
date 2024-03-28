package schema

import (
	"time"
	"math/rand"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}
var url = "https://card.syui.ai"

var card int
var super string
var skill string
var cp int

func (Card) Fields() []ent.Field {
	return []ent.Field{  

		field.String("password").
		NotEmpty().
		Immutable().
		Sensitive(),

		field.Int("card").
		DefaultFunc(func() int {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(20)
			if a == 1 {
				card = rand.Intn(3) + 123
			} else {
				card = 0
			}
			return card
		}).
		Optional(),

		field.String("skill").
		DefaultFunc(func() string {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(30)
			if a == 1 {
				skill = "critical"
			} else {
				skill = "normal"
			}
			if card == 0 {
				skill = "normal"
			}
			return skill
		}).
		Optional(),

		field.String("status").
		//Immutable().
		DefaultFunc(func() string {
			rand.Seed(time.Now().UnixNano())
			var a = rand.Intn(40)
			if a == 1 {
				super = "super"
			} else {
				super = "normal"
			}
			if card == 0 {
				super = "normal"
			}
			return super
		}).
		Optional(),

		field.String("token").
		Optional().
		Sensitive(),

		field.Int("cp").
		//Immutable().
		DefaultFunc(func() int {
			rand.Seed(time.Now().UnixNano())
			var cp = 1 + rand.Intn(200)
			if cp == 2 {
				cp = 50 + rand.Intn(300)
			}
			if card >= 1 {
				cp = 200 + cp + rand.Intn(500)
			}
			if super == "super" {
				cp = 400 + cp + rand.Intn(700)
			}

			if skill == "critical" {
				cp = 400 + cp + rand.Intn(700)
			}

			if card == 22 {
				cp = 800 + cp + rand.Intn(1500)
			}

			if card == 25 {
				cp = 0
			}

			return cp
		}).
		Optional(),

		field.String("url").
		Default(url).
		Optional(),

		field.Int("count").
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

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Unique().
			Required(),

	}
}
