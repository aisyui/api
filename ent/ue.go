// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/ue"
	"api/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Ue is the model entity for the Ue schema.
type Ue struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Limit holds the value of the "limit" field.
	Limit bool `json:"limit,omitempty"`
	// LimitBoss holds the value of the "limit_boss" field.
	LimitBoss bool `json:"limit_boss,omitempty"`
	// LimitItem holds the value of the "limit_item" field.
	LimitItem bool `json:"limit_item,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Lv holds the value of the "lv" field.
	Lv int `json:"lv,omitempty"`
	// LvPoint holds the value of the "lv_point" field.
	LvPoint int `json:"lv_point,omitempty"`
	// Model holds the value of the "model" field.
	Model int `json:"model,omitempty"`
	// Sword holds the value of the "sword" field.
	Sword int `json:"sword,omitempty"`
	// Card holds the value of the "card" field.
	Card int `json:"card,omitempty"`
	// Mode holds the value of the "mode" field.
	Mode string `json:"mode,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"-"`
	// Cp holds the value of the "cp" field.
	Cp int `json:"cp,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// LocationX holds the value of the "location_x" field.
	LocationX int `json:"location_x,omitempty"`
	// LocationY holds the value of the "location_y" field.
	LocationY int `json:"location_y,omitempty"`
	// LocationZ holds the value of the "location_z" field.
	LocationZ int `json:"location_z,omitempty"`
	// LocationN holds the value of the "location_n" field.
	LocationN int `json:"location_n,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UeQuery when eager-loading is set.
	Edges        UeEdges `json:"edges"`
	user_ue      *int
	selectValues sql.SelectValues
}

// UeEdges holds the relations/edges for other nodes in the graph.
type UeEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UeEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ue.FieldLimit, ue.FieldLimitBoss, ue.FieldLimitItem:
			values[i] = new(sql.NullBool)
		case ue.FieldID, ue.FieldLv, ue.FieldLvPoint, ue.FieldModel, ue.FieldSword, ue.FieldCard, ue.FieldCp, ue.FieldCount, ue.FieldLocationX, ue.FieldLocationY, ue.FieldLocationZ, ue.FieldLocationN:
			values[i] = new(sql.NullInt64)
		case ue.FieldPassword, ue.FieldMode, ue.FieldToken, ue.FieldAuthor:
			values[i] = new(sql.NullString)
		case ue.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case ue.ForeignKeys[0]: // user_ue
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ue fields.
func (u *Ue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case ue.FieldLimit:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field limit", values[i])
			} else if value.Valid {
				u.Limit = value.Bool
			}
		case ue.FieldLimitBoss:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field limit_boss", values[i])
			} else if value.Valid {
				u.LimitBoss = value.Bool
			}
		case ue.FieldLimitItem:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field limit_item", values[i])
			} else if value.Valid {
				u.LimitItem = value.Bool
			}
		case ue.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case ue.FieldLv:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lv", values[i])
			} else if value.Valid {
				u.Lv = int(value.Int64)
			}
		case ue.FieldLvPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lv_point", values[i])
			} else if value.Valid {
				u.LvPoint = int(value.Int64)
			}
		case ue.FieldModel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				u.Model = int(value.Int64)
			}
		case ue.FieldSword:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sword", values[i])
			} else if value.Valid {
				u.Sword = int(value.Int64)
			}
		case ue.FieldCard:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field card", values[i])
			} else if value.Valid {
				u.Card = int(value.Int64)
			}
		case ue.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				u.Mode = value.String
			}
		case ue.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				u.Token = value.String
			}
		case ue.FieldCp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cp", values[i])
			} else if value.Valid {
				u.Cp = int(value.Int64)
			}
		case ue.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				u.Count = int(value.Int64)
			}
		case ue.FieldLocationX:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_x", values[i])
			} else if value.Valid {
				u.LocationX = int(value.Int64)
			}
		case ue.FieldLocationY:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_y", values[i])
			} else if value.Valid {
				u.LocationY = int(value.Int64)
			}
		case ue.FieldLocationZ:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_z", values[i])
			} else if value.Valid {
				u.LocationZ = int(value.Int64)
			}
		case ue.FieldLocationN:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_n", values[i])
			} else if value.Valid {
				u.LocationN = int(value.Int64)
			}
		case ue.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				u.Author = value.String
			}
		case ue.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case ue.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_ue", value)
			} else if value.Valid {
				u.user_ue = new(int)
				*u.user_ue = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Ue.
// This includes values selected through modifiers, order, etc.
func (u *Ue) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Ue entity.
func (u *Ue) QueryOwner() *UserQuery {
	return NewUeClient(u.config).QueryOwner(u)
}

// Update returns a builder for updating this Ue.
// Note that you need to call Ue.Unwrap() before calling this method if this Ue
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Ue) Update() *UeUpdateOne {
	return NewUeClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Ue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Ue) Unwrap() *Ue {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ue is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Ue) String() string {
	var builder strings.Builder
	builder.WriteString("Ue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("limit=")
	builder.WriteString(fmt.Sprintf("%v", u.Limit))
	builder.WriteString(", ")
	builder.WriteString("limit_boss=")
	builder.WriteString(fmt.Sprintf("%v", u.LimitBoss))
	builder.WriteString(", ")
	builder.WriteString("limit_item=")
	builder.WriteString(fmt.Sprintf("%v", u.LimitItem))
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("lv=")
	builder.WriteString(fmt.Sprintf("%v", u.Lv))
	builder.WriteString(", ")
	builder.WriteString("lv_point=")
	builder.WriteString(fmt.Sprintf("%v", u.LvPoint))
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(fmt.Sprintf("%v", u.Model))
	builder.WriteString(", ")
	builder.WriteString("sword=")
	builder.WriteString(fmt.Sprintf("%v", u.Sword))
	builder.WriteString(", ")
	builder.WriteString("card=")
	builder.WriteString(fmt.Sprintf("%v", u.Card))
	builder.WriteString(", ")
	builder.WriteString("mode=")
	builder.WriteString(u.Mode)
	builder.WriteString(", ")
	builder.WriteString("token=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("cp=")
	builder.WriteString(fmt.Sprintf("%v", u.Cp))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", u.Count))
	builder.WriteString(", ")
	builder.WriteString("location_x=")
	builder.WriteString(fmt.Sprintf("%v", u.LocationX))
	builder.WriteString(", ")
	builder.WriteString("location_y=")
	builder.WriteString(fmt.Sprintf("%v", u.LocationY))
	builder.WriteString(", ")
	builder.WriteString("location_z=")
	builder.WriteString(fmt.Sprintf("%v", u.LocationZ))
	builder.WriteString(", ")
	builder.WriteString("location_n=")
	builder.WriteString(fmt.Sprintf("%v", u.LocationN))
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(u.Author)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Ues is a parsable slice of Ue.
type Ues []*Ue
