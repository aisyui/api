// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"t/ent/users"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Users is the model entity for the Users schema.
type Users struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// Chara holds the value of the "chara" field.
	Chara string `json:"chara,omitempty"`
	// Skill holds the value of the "skill" field.
	Skill int `json:"skill,omitempty"`
	// Hp holds the value of the "hp" field.
	Hp int `json:"hp,omitempty"`
	// Attack holds the value of the "attack" field.
	Attack int `json:"attack,omitempty"`
	// Defense holds the value of the "defense" field.
	Defense int `json:"defense,omitempty"`
	// Critical holds the value of the "critical" field.
	Critical int `json:"critical,omitempty"`
	// Battle holds the value of the "battle" field.
	Battle int `json:"battle,omitempty"`
	// Win holds the value of the "win" field.
	Win int `json:"win,omitempty"`
	// Day holds the value of the "day" field.
	Day int `json:"day,omitempty"`
	// Percentage holds the value of the "percentage" field.
	Percentage float64 `json:"percentage,omitempty"`
	// Limit holds the value of the "limit" field.
	Limit bool `json:"limit,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Next holds the value of the "next" field.
	Next string `json:"next,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Users) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case users.FieldLimit:
			values[i] = new(sql.NullBool)
		case users.FieldPercentage:
			values[i] = new(sql.NullFloat64)
		case users.FieldID, users.FieldSkill, users.FieldHp, users.FieldAttack, users.FieldDefense, users.FieldCritical, users.FieldBattle, users.FieldWin, users.FieldDay:
			values[i] = new(sql.NullInt64)
		case users.FieldUser, users.FieldChara, users.FieldStatus, users.FieldComment, users.FieldNext, users.FieldURL:
			values[i] = new(sql.NullString)
		case users.FieldCreatedAt, users.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Users", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Users fields.
func (u *Users) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case users.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				u.User = value.String
			}
		case users.FieldChara:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chara", values[i])
			} else if value.Valid {
				u.Chara = value.String
			}
		case users.FieldSkill:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field skill", values[i])
			} else if value.Valid {
				u.Skill = int(value.Int64)
			}
		case users.FieldHp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hp", values[i])
			} else if value.Valid {
				u.Hp = int(value.Int64)
			}
		case users.FieldAttack:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attack", values[i])
			} else if value.Valid {
				u.Attack = int(value.Int64)
			}
		case users.FieldDefense:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field defense", values[i])
			} else if value.Valid {
				u.Defense = int(value.Int64)
			}
		case users.FieldCritical:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field critical", values[i])
			} else if value.Valid {
				u.Critical = int(value.Int64)
			}
		case users.FieldBattle:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field battle", values[i])
			} else if value.Valid {
				u.Battle = int(value.Int64)
			}
		case users.FieldWin:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field win", values[i])
			} else if value.Valid {
				u.Win = int(value.Int64)
			}
		case users.FieldDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field day", values[i])
			} else if value.Valid {
				u.Day = int(value.Int64)
			}
		case users.FieldPercentage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field percentage", values[i])
			} else if value.Valid {
				u.Percentage = value.Float64
			}
		case users.FieldLimit:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field limit", values[i])
			} else if value.Valid {
				u.Limit = value.Bool
			}
		case users.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = value.String
			}
		case users.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				u.Comment = value.String
			}
		case users.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case users.FieldNext:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field next", values[i])
			} else if value.Valid {
				u.Next = value.String
			}
		case users.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case users.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				u.URL = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Users.
// Note that you need to call Users.Unwrap() before calling this method if this Users
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Users) Update() *UsersUpdateOne {
	return (&UsersClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Users entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Users) Unwrap() *Users {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Users is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Users) String() string {
	var builder strings.Builder
	builder.WriteString("Users(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", user=")
	builder.WriteString(u.User)
	builder.WriteString(", chara=")
	builder.WriteString(u.Chara)
	builder.WriteString(", skill=")
	builder.WriteString(fmt.Sprintf("%v", u.Skill))
	builder.WriteString(", hp=")
	builder.WriteString(fmt.Sprintf("%v", u.Hp))
	builder.WriteString(", attack=")
	builder.WriteString(fmt.Sprintf("%v", u.Attack))
	builder.WriteString(", defense=")
	builder.WriteString(fmt.Sprintf("%v", u.Defense))
	builder.WriteString(", critical=")
	builder.WriteString(fmt.Sprintf("%v", u.Critical))
	builder.WriteString(", battle=")
	builder.WriteString(fmt.Sprintf("%v", u.Battle))
	builder.WriteString(", win=")
	builder.WriteString(fmt.Sprintf("%v", u.Win))
	builder.WriteString(", day=")
	builder.WriteString(fmt.Sprintf("%v", u.Day))
	builder.WriteString(", percentage=")
	builder.WriteString(fmt.Sprintf("%v", u.Percentage))
	builder.WriteString(", limit=")
	builder.WriteString(fmt.Sprintf("%v", u.Limit))
	builder.WriteString(", status=")
	builder.WriteString(u.Status)
	builder.WriteString(", comment=")
	builder.WriteString(u.Comment)
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", next=")
	builder.WriteString(u.Next)
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", url=")
	builder.WriteString(u.URL)
	builder.WriteByte(')')
	return builder.String()
}

// UsersSlice is a parsable slice of Users.
type UsersSlice []*Users

func (u UsersSlice) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}