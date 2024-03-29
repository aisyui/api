// Code generated by ent, DO NOT EDIT.

package ue

import (
	"api/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldID, id))
}

// Limit applies equality check predicate on the "limit" field. It's identical to LimitEQ.
func Limit(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLimit, v))
}

// LimitBoss applies equality check predicate on the "limit_boss" field. It's identical to LimitBossEQ.
func LimitBoss(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLimitBoss, v))
}

// LimitItem applies equality check predicate on the "limit_item" field. It's identical to LimitItemEQ.
func LimitItem(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLimitItem, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldPassword, v))
}

// Lv applies equality check predicate on the "lv" field. It's identical to LvEQ.
func Lv(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLv, v))
}

// LvPoint applies equality check predicate on the "lv_point" field. It's identical to LvPointEQ.
func LvPoint(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLvPoint, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldModel, v))
}

// Sword applies equality check predicate on the "sword" field. It's identical to SwordEQ.
func Sword(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldSword, v))
}

// Card applies equality check predicate on the "card" field. It's identical to CardEQ.
func Card(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCard, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldMode, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldToken, v))
}

// Cp applies equality check predicate on the "cp" field. It's identical to CpEQ.
func Cp(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCp, v))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCount, v))
}

// LocationX applies equality check predicate on the "location_x" field. It's identical to LocationXEQ.
func LocationX(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationX, v))
}

// LocationY applies equality check predicate on the "location_y" field. It's identical to LocationYEQ.
func LocationY(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationY, v))
}

// LocationZ applies equality check predicate on the "location_z" field. It's identical to LocationZEQ.
func LocationZ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationZ, v))
}

// LocationN applies equality check predicate on the "location_n" field. It's identical to LocationNEQ.
func LocationN(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationN, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldAuthor, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCreatedAt, v))
}

// LimitEQ applies the EQ predicate on the "limit" field.
func LimitEQ(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLimit, v))
}

// LimitNEQ applies the NEQ predicate on the "limit" field.
func LimitNEQ(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLimit, v))
}

// LimitIsNil applies the IsNil predicate on the "limit" field.
func LimitIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLimit))
}

// LimitNotNil applies the NotNil predicate on the "limit" field.
func LimitNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLimit))
}

// LimitBossEQ applies the EQ predicate on the "limit_boss" field.
func LimitBossEQ(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLimitBoss, v))
}

// LimitBossNEQ applies the NEQ predicate on the "limit_boss" field.
func LimitBossNEQ(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLimitBoss, v))
}

// LimitBossIsNil applies the IsNil predicate on the "limit_boss" field.
func LimitBossIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLimitBoss))
}

// LimitBossNotNil applies the NotNil predicate on the "limit_boss" field.
func LimitBossNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLimitBoss))
}

// LimitItemEQ applies the EQ predicate on the "limit_item" field.
func LimitItemEQ(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLimitItem, v))
}

// LimitItemNEQ applies the NEQ predicate on the "limit_item" field.
func LimitItemNEQ(v bool) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLimitItem, v))
}

// LimitItemIsNil applies the IsNil predicate on the "limit_item" field.
func LimitItemIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLimitItem))
}

// LimitItemNotNil applies the NotNil predicate on the "limit_item" field.
func LimitItemNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLimitItem))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContainsFold(FieldPassword, v))
}

// LvEQ applies the EQ predicate on the "lv" field.
func LvEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLv, v))
}

// LvNEQ applies the NEQ predicate on the "lv" field.
func LvNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLv, v))
}

// LvIn applies the In predicate on the "lv" field.
func LvIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldLv, vs...))
}

// LvNotIn applies the NotIn predicate on the "lv" field.
func LvNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldLv, vs...))
}

// LvGT applies the GT predicate on the "lv" field.
func LvGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldLv, v))
}

// LvGTE applies the GTE predicate on the "lv" field.
func LvGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldLv, v))
}

// LvLT applies the LT predicate on the "lv" field.
func LvLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldLv, v))
}

// LvLTE applies the LTE predicate on the "lv" field.
func LvLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldLv, v))
}

// LvIsNil applies the IsNil predicate on the "lv" field.
func LvIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLv))
}

// LvNotNil applies the NotNil predicate on the "lv" field.
func LvNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLv))
}

// LvPointEQ applies the EQ predicate on the "lv_point" field.
func LvPointEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLvPoint, v))
}

// LvPointNEQ applies the NEQ predicate on the "lv_point" field.
func LvPointNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLvPoint, v))
}

// LvPointIn applies the In predicate on the "lv_point" field.
func LvPointIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldLvPoint, vs...))
}

// LvPointNotIn applies the NotIn predicate on the "lv_point" field.
func LvPointNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldLvPoint, vs...))
}

// LvPointGT applies the GT predicate on the "lv_point" field.
func LvPointGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldLvPoint, v))
}

// LvPointGTE applies the GTE predicate on the "lv_point" field.
func LvPointGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldLvPoint, v))
}

// LvPointLT applies the LT predicate on the "lv_point" field.
func LvPointLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldLvPoint, v))
}

// LvPointLTE applies the LTE predicate on the "lv_point" field.
func LvPointLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldLvPoint, v))
}

// LvPointIsNil applies the IsNil predicate on the "lv_point" field.
func LvPointIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLvPoint))
}

// LvPointNotNil applies the NotNil predicate on the "lv_point" field.
func LvPointNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLvPoint))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldModel, v))
}

// ModelIsNil applies the IsNil predicate on the "model" field.
func ModelIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldModel))
}

// ModelNotNil applies the NotNil predicate on the "model" field.
func ModelNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldModel))
}

// SwordEQ applies the EQ predicate on the "sword" field.
func SwordEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldSword, v))
}

// SwordNEQ applies the NEQ predicate on the "sword" field.
func SwordNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldSword, v))
}

// SwordIn applies the In predicate on the "sword" field.
func SwordIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldSword, vs...))
}

// SwordNotIn applies the NotIn predicate on the "sword" field.
func SwordNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldSword, vs...))
}

// SwordGT applies the GT predicate on the "sword" field.
func SwordGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldSword, v))
}

// SwordGTE applies the GTE predicate on the "sword" field.
func SwordGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldSword, v))
}

// SwordLT applies the LT predicate on the "sword" field.
func SwordLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldSword, v))
}

// SwordLTE applies the LTE predicate on the "sword" field.
func SwordLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldSword, v))
}

// SwordIsNil applies the IsNil predicate on the "sword" field.
func SwordIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldSword))
}

// SwordNotNil applies the NotNil predicate on the "sword" field.
func SwordNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldSword))
}

// CardEQ applies the EQ predicate on the "card" field.
func CardEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCard, v))
}

// CardNEQ applies the NEQ predicate on the "card" field.
func CardNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldCard, v))
}

// CardIn applies the In predicate on the "card" field.
func CardIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldCard, vs...))
}

// CardNotIn applies the NotIn predicate on the "card" field.
func CardNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldCard, vs...))
}

// CardGT applies the GT predicate on the "card" field.
func CardGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldCard, v))
}

// CardGTE applies the GTE predicate on the "card" field.
func CardGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldCard, v))
}

// CardLT applies the LT predicate on the "card" field.
func CardLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldCard, v))
}

// CardLTE applies the LTE predicate on the "card" field.
func CardLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldCard, v))
}

// CardIsNil applies the IsNil predicate on the "card" field.
func CardIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldCard))
}

// CardNotNil applies the NotNil predicate on the "card" field.
func CardNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldCard))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldMode, v))
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContains(FieldMode, v))
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasPrefix(FieldMode, v))
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasSuffix(FieldMode, v))
}

// ModeIsNil applies the IsNil predicate on the "mode" field.
func ModeIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldMode))
}

// ModeNotNil applies the NotNil predicate on the "mode" field.
func ModeNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldMode))
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEqualFold(FieldMode, v))
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContainsFold(FieldMode, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasSuffix(FieldToken, v))
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldToken))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContainsFold(FieldToken, v))
}

// CpEQ applies the EQ predicate on the "cp" field.
func CpEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCp, v))
}

// CpNEQ applies the NEQ predicate on the "cp" field.
func CpNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldCp, v))
}

// CpIn applies the In predicate on the "cp" field.
func CpIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldCp, vs...))
}

// CpNotIn applies the NotIn predicate on the "cp" field.
func CpNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldCp, vs...))
}

// CpGT applies the GT predicate on the "cp" field.
func CpGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldCp, v))
}

// CpGTE applies the GTE predicate on the "cp" field.
func CpGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldCp, v))
}

// CpLT applies the LT predicate on the "cp" field.
func CpLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldCp, v))
}

// CpLTE applies the LTE predicate on the "cp" field.
func CpLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldCp, v))
}

// CpIsNil applies the IsNil predicate on the "cp" field.
func CpIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldCp))
}

// CpNotNil applies the NotNil predicate on the "cp" field.
func CpNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldCp))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldCount, v))
}

// CountIsNil applies the IsNil predicate on the "count" field.
func CountIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldCount))
}

// CountNotNil applies the NotNil predicate on the "count" field.
func CountNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldCount))
}

// LocationXEQ applies the EQ predicate on the "location_x" field.
func LocationXEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationX, v))
}

// LocationXNEQ applies the NEQ predicate on the "location_x" field.
func LocationXNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLocationX, v))
}

// LocationXIn applies the In predicate on the "location_x" field.
func LocationXIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldLocationX, vs...))
}

// LocationXNotIn applies the NotIn predicate on the "location_x" field.
func LocationXNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldLocationX, vs...))
}

// LocationXGT applies the GT predicate on the "location_x" field.
func LocationXGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldLocationX, v))
}

// LocationXGTE applies the GTE predicate on the "location_x" field.
func LocationXGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldLocationX, v))
}

// LocationXLT applies the LT predicate on the "location_x" field.
func LocationXLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldLocationX, v))
}

// LocationXLTE applies the LTE predicate on the "location_x" field.
func LocationXLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldLocationX, v))
}

// LocationXIsNil applies the IsNil predicate on the "location_x" field.
func LocationXIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLocationX))
}

// LocationXNotNil applies the NotNil predicate on the "location_x" field.
func LocationXNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLocationX))
}

// LocationYEQ applies the EQ predicate on the "location_y" field.
func LocationYEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationY, v))
}

// LocationYNEQ applies the NEQ predicate on the "location_y" field.
func LocationYNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLocationY, v))
}

// LocationYIn applies the In predicate on the "location_y" field.
func LocationYIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldLocationY, vs...))
}

// LocationYNotIn applies the NotIn predicate on the "location_y" field.
func LocationYNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldLocationY, vs...))
}

// LocationYGT applies the GT predicate on the "location_y" field.
func LocationYGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldLocationY, v))
}

// LocationYGTE applies the GTE predicate on the "location_y" field.
func LocationYGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldLocationY, v))
}

// LocationYLT applies the LT predicate on the "location_y" field.
func LocationYLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldLocationY, v))
}

// LocationYLTE applies the LTE predicate on the "location_y" field.
func LocationYLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldLocationY, v))
}

// LocationYIsNil applies the IsNil predicate on the "location_y" field.
func LocationYIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLocationY))
}

// LocationYNotNil applies the NotNil predicate on the "location_y" field.
func LocationYNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLocationY))
}

// LocationZEQ applies the EQ predicate on the "location_z" field.
func LocationZEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationZ, v))
}

// LocationZNEQ applies the NEQ predicate on the "location_z" field.
func LocationZNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLocationZ, v))
}

// LocationZIn applies the In predicate on the "location_z" field.
func LocationZIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldLocationZ, vs...))
}

// LocationZNotIn applies the NotIn predicate on the "location_z" field.
func LocationZNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldLocationZ, vs...))
}

// LocationZGT applies the GT predicate on the "location_z" field.
func LocationZGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldLocationZ, v))
}

// LocationZGTE applies the GTE predicate on the "location_z" field.
func LocationZGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldLocationZ, v))
}

// LocationZLT applies the LT predicate on the "location_z" field.
func LocationZLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldLocationZ, v))
}

// LocationZLTE applies the LTE predicate on the "location_z" field.
func LocationZLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldLocationZ, v))
}

// LocationZIsNil applies the IsNil predicate on the "location_z" field.
func LocationZIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLocationZ))
}

// LocationZNotNil applies the NotNil predicate on the "location_z" field.
func LocationZNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLocationZ))
}

// LocationNEQ applies the EQ predicate on the "location_n" field.
func LocationNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldLocationN, v))
}

// LocationNNEQ applies the NEQ predicate on the "location_n" field.
func LocationNNEQ(v int) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldLocationN, v))
}

// LocationNIn applies the In predicate on the "location_n" field.
func LocationNIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldLocationN, vs...))
}

// LocationNNotIn applies the NotIn predicate on the "location_n" field.
func LocationNNotIn(vs ...int) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldLocationN, vs...))
}

// LocationNGT applies the GT predicate on the "location_n" field.
func LocationNGT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldLocationN, v))
}

// LocationNGTE applies the GTE predicate on the "location_n" field.
func LocationNGTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldLocationN, v))
}

// LocationNLT applies the LT predicate on the "location_n" field.
func LocationNLT(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldLocationN, v))
}

// LocationNLTE applies the LTE predicate on the "location_n" field.
func LocationNLTE(v int) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldLocationN, v))
}

// LocationNIsNil applies the IsNil predicate on the "location_n" field.
func LocationNIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldLocationN))
}

// LocationNNotNil applies the NotNil predicate on the "location_n" field.
func LocationNNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldLocationN))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Ue {
	return predicate.Ue(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorIsNil applies the IsNil predicate on the "author" field.
func AuthorIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldAuthor))
}

// AuthorNotNil applies the NotNil predicate on the "author" field.
func AuthorNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldAuthor))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Ue {
	return predicate.Ue(sql.FieldContainsFold(FieldAuthor, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Ue {
	return predicate.Ue(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Ue {
	return predicate.Ue(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Ue {
	return predicate.Ue(sql.FieldNotNull(FieldCreatedAt))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Ue {
	return predicate.Ue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Ue {
	return predicate.Ue(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ue) predicate.Ue {
	return predicate.Ue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ue) predicate.Ue {
	return predicate.Ue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ue) predicate.Ue {
	return predicate.Ue(func(s *sql.Selector) {
		p(s.Not())
	})
}
