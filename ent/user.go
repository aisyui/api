// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"t/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Did holds the value of the "did" field.
	Did string `json:"did,omitempty"`
	// Member holds the value of the "member" field.
	Member bool `json:"member,omitempty"`
	// Book holds the value of the "book" field.
	Book bool `json:"book,omitempty"`
	// Manga holds the value of the "manga" field.
	Manga bool `json:"manga,omitempty"`
	// Badge holds the value of the "badge" field.
	Badge bool `json:"badge,omitempty"`
	// Bsky holds the value of the "bsky" field.
	Bsky bool `json:"bsky,omitempty"`
	// Mastodon holds the value of the "mastodon" field.
	Mastodon bool `json:"mastodon,omitempty"`
	// Delete holds the value of the "delete" field.
	Delete bool `json:"delete,omitempty"`
	// Handle holds the value of the "handle" field.
	Handle bool `json:"handle,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"-"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RaidAt holds the value of the "raid_at" field.
	RaidAt time.Time `json:"raid_at,omitempty"`
	// Luck holds the value of the "luck" field.
	Luck int `json:"luck,omitempty"`
	// LuckAt holds the value of the "luck_at" field.
	LuckAt time.Time `json:"luck_at,omitempty"`
	// Like holds the value of the "like" field.
	Like int `json:"like,omitempty"`
	// LikeRank holds the value of the "like_rank" field.
	LikeRank int `json:"like_rank,omitempty"`
	// LikeAt holds the value of the "like_at" field.
	LikeAt time.Time `json:"like_at,omitempty"`
	// Fav holds the value of the "fav" field.
	Fav int `json:"fav,omitempty"`
	// Ten holds the value of the "ten" field.
	Ten bool `json:"ten,omitempty"`
	// TenSu holds the value of the "ten_su" field.
	TenSu int `json:"ten_su,omitempty"`
	// TenKai holds the value of the "ten_kai" field.
	TenKai int `json:"ten_kai,omitempty"`
	// Aiten holds the value of the "aiten" field.
	Aiten int `json:"aiten,omitempty"`
	// TenCard holds the value of the "ten_card" field.
	TenCard string `json:"ten_card,omitempty"`
	// TenDelete holds the value of the "ten_delete" field.
	TenDelete string `json:"ten_delete,omitempty"`
	// TenPost holds the value of the "ten_post" field.
	TenPost string `json:"ten_post,omitempty"`
	// TenGet holds the value of the "ten_get" field.
	TenGet string `json:"ten_get,omitempty"`
	// TenAt holds the value of the "ten_at" field.
	TenAt time.Time `json:"ten_at,omitempty"`
	// Next holds the value of the "next" field.
	Next string `json:"next,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges       UserEdges `json:"edges"`
	group_users *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Card holds the value of the card edge.
	Card []*Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldMember, user.FieldBook, user.FieldManga, user.FieldBadge, user.FieldBsky, user.FieldMastodon, user.FieldDelete, user.FieldHandle, user.FieldTen:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldLuck, user.FieldLike, user.FieldLikeRank, user.FieldFav, user.FieldTenSu, user.FieldTenKai, user.FieldAiten:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldDid, user.FieldToken, user.FieldPassword, user.FieldTenCard, user.FieldTenDelete, user.FieldTenPost, user.FieldTenGet, user.FieldNext:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldRaidAt, user.FieldLuckAt, user.FieldLikeAt, user.FieldTenAt:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // group_users
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldDid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field did", values[i])
			} else if value.Valid {
				u.Did = value.String
			}
		case user.FieldMember:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field member", values[i])
			} else if value.Valid {
				u.Member = value.Bool
			}
		case user.FieldBook:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field book", values[i])
			} else if value.Valid {
				u.Book = value.Bool
			}
		case user.FieldManga:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field manga", values[i])
			} else if value.Valid {
				u.Manga = value.Bool
			}
		case user.FieldBadge:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field badge", values[i])
			} else if value.Valid {
				u.Badge = value.Bool
			}
		case user.FieldBsky:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field bsky", values[i])
			} else if value.Valid {
				u.Bsky = value.Bool
			}
		case user.FieldMastodon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field mastodon", values[i])
			} else if value.Valid {
				u.Mastodon = value.Bool
			}
		case user.FieldDelete:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				u.Delete = value.Bool
			}
		case user.FieldHandle:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field handle", values[i])
			} else if value.Valid {
				u.Handle = value.Bool
			}
		case user.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				u.Token = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldRaidAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field raid_at", values[i])
			} else if value.Valid {
				u.RaidAt = value.Time
			}
		case user.FieldLuck:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field luck", values[i])
			} else if value.Valid {
				u.Luck = int(value.Int64)
			}
		case user.FieldLuckAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field luck_at", values[i])
			} else if value.Valid {
				u.LuckAt = value.Time
			}
		case user.FieldLike:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like", values[i])
			} else if value.Valid {
				u.Like = int(value.Int64)
			}
		case user.FieldLikeRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like_rank", values[i])
			} else if value.Valid {
				u.LikeRank = int(value.Int64)
			}
		case user.FieldLikeAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field like_at", values[i])
			} else if value.Valid {
				u.LikeAt = value.Time
			}
		case user.FieldFav:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fav", values[i])
			} else if value.Valid {
				u.Fav = int(value.Int64)
			}
		case user.FieldTen:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ten", values[i])
			} else if value.Valid {
				u.Ten = value.Bool
			}
		case user.FieldTenSu:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ten_su", values[i])
			} else if value.Valid {
				u.TenSu = int(value.Int64)
			}
		case user.FieldTenKai:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ten_kai", values[i])
			} else if value.Valid {
				u.TenKai = int(value.Int64)
			}
		case user.FieldAiten:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field aiten", values[i])
			} else if value.Valid {
				u.Aiten = int(value.Int64)
			}
		case user.FieldTenCard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ten_card", values[i])
			} else if value.Valid {
				u.TenCard = value.String
			}
		case user.FieldTenDelete:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ten_delete", values[i])
			} else if value.Valid {
				u.TenDelete = value.String
			}
		case user.FieldTenPost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ten_post", values[i])
			} else if value.Valid {
				u.TenPost = value.String
			}
		case user.FieldTenGet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ten_get", values[i])
			} else if value.Valid {
				u.TenGet = value.String
			}
		case user.FieldTenAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ten_at", values[i])
			} else if value.Valid {
				u.TenAt = value.Time
			}
		case user.FieldNext:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field next", values[i])
			} else if value.Valid {
				u.Next = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_users", value)
			} else if value.Valid {
				u.group_users = new(int)
				*u.group_users = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCard queries the "card" edge of the User entity.
func (u *User) QueryCard() *CardQuery {
	return NewUserClient(u.config).QueryCard(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("did=")
	builder.WriteString(u.Did)
	builder.WriteString(", ")
	builder.WriteString("member=")
	builder.WriteString(fmt.Sprintf("%v", u.Member))
	builder.WriteString(", ")
	builder.WriteString("book=")
	builder.WriteString(fmt.Sprintf("%v", u.Book))
	builder.WriteString(", ")
	builder.WriteString("manga=")
	builder.WriteString(fmt.Sprintf("%v", u.Manga))
	builder.WriteString(", ")
	builder.WriteString("badge=")
	builder.WriteString(fmt.Sprintf("%v", u.Badge))
	builder.WriteString(", ")
	builder.WriteString("bsky=")
	builder.WriteString(fmt.Sprintf("%v", u.Bsky))
	builder.WriteString(", ")
	builder.WriteString("mastodon=")
	builder.WriteString(fmt.Sprintf("%v", u.Mastodon))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", u.Delete))
	builder.WriteString(", ")
	builder.WriteString("handle=")
	builder.WriteString(fmt.Sprintf("%v", u.Handle))
	builder.WriteString(", ")
	builder.WriteString("token=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("raid_at=")
	builder.WriteString(u.RaidAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("luck=")
	builder.WriteString(fmt.Sprintf("%v", u.Luck))
	builder.WriteString(", ")
	builder.WriteString("luck_at=")
	builder.WriteString(u.LuckAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("like=")
	builder.WriteString(fmt.Sprintf("%v", u.Like))
	builder.WriteString(", ")
	builder.WriteString("like_rank=")
	builder.WriteString(fmt.Sprintf("%v", u.LikeRank))
	builder.WriteString(", ")
	builder.WriteString("like_at=")
	builder.WriteString(u.LikeAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fav=")
	builder.WriteString(fmt.Sprintf("%v", u.Fav))
	builder.WriteString(", ")
	builder.WriteString("ten=")
	builder.WriteString(fmt.Sprintf("%v", u.Ten))
	builder.WriteString(", ")
	builder.WriteString("ten_su=")
	builder.WriteString(fmt.Sprintf("%v", u.TenSu))
	builder.WriteString(", ")
	builder.WriteString("ten_kai=")
	builder.WriteString(fmt.Sprintf("%v", u.TenKai))
	builder.WriteString(", ")
	builder.WriteString("aiten=")
	builder.WriteString(fmt.Sprintf("%v", u.Aiten))
	builder.WriteString(", ")
	builder.WriteString("ten_card=")
	builder.WriteString(u.TenCard)
	builder.WriteString(", ")
	builder.WriteString("ten_delete=")
	builder.WriteString(u.TenDelete)
	builder.WriteString(", ")
	builder.WriteString("ten_post=")
	builder.WriteString(u.TenPost)
	builder.WriteString(", ")
	builder.WriteString("ten_get=")
	builder.WriteString(u.TenGet)
	builder.WriteString(", ")
	builder.WriteString("ten_at=")
	builder.WriteString(u.TenAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("next=")
	builder.WriteString(u.Next)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
