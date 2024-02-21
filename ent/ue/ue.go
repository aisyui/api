// Code generated by ent, DO NOT EDIT.

package ue

import (
	"time"
)

const (
	// Label holds the string label denoting the ue type in the database.
	Label = "ue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLimit holds the string denoting the limit field in the database.
	FieldLimit = "limit"
	// FieldLimitBoss holds the string denoting the limit_boss field in the database.
	FieldLimitBoss = "limit_boss"
	// FieldLimitItem holds the string denoting the limit_item field in the database.
	FieldLimitItem = "limit_item"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldLv holds the string denoting the lv field in the database.
	FieldLv = "lv"
	// FieldLvPoint holds the string denoting the lv_point field in the database.
	FieldLvPoint = "lv_point"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldSword holds the string denoting the sword field in the database.
	FieldSword = "sword"
	// FieldCard holds the string denoting the card field in the database.
	FieldCard = "card"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldCp holds the string denoting the cp field in the database.
	FieldCp = "cp"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldLocationX holds the string denoting the location_x field in the database.
	FieldLocationX = "location_x"
	// FieldLocationY holds the string denoting the location_y field in the database.
	FieldLocationY = "location_y"
	// FieldLocationZ holds the string denoting the location_z field in the database.
	FieldLocationZ = "location_z"
	// FieldLocationN holds the string denoting the location_n field in the database.
	FieldLocationN = "location_n"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the ue in the database.
	Table = "ues"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "ues"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_ue"
)

// Columns holds all SQL columns for ue fields.
var Columns = []string{
	FieldID,
	FieldLimit,
	FieldLimitBoss,
	FieldLimitItem,
	FieldPassword,
	FieldLv,
	FieldLvPoint,
	FieldModel,
	FieldSword,
	FieldCard,
	FieldMode,
	FieldToken,
	FieldCp,
	FieldCount,
	FieldLocationX,
	FieldLocationY,
	FieldLocationZ,
	FieldLocationN,
	FieldAuthor,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ues"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_ue",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLimit holds the default value on creation for the "limit" field.
	DefaultLimit bool
	// DefaultLimitBoss holds the default value on creation for the "limit_boss" field.
	DefaultLimitBoss bool
	// DefaultLimitItem holds the default value on creation for the "limit_item" field.
	DefaultLimitItem bool
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)