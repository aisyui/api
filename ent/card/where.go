// Code generated by ent, DO NOT EDIT.

package card

import (
	"t/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldID, id))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldPassword, v))
}

// Card applies equality check predicate on the "card" field. It's identical to CardEQ.
func Card(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCard, v))
}

// Skill applies equality check predicate on the "skill" field. It's identical to SkillEQ.
func Skill(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldSkill, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldStatus, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldToken, v))
}

// Cp applies equality check predicate on the "cp" field. It's identical to CpEQ.
func Cp(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCp, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCreatedAt, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldPassword, v))
}

// CardEQ applies the EQ predicate on the "card" field.
func CardEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCard, v))
}

// CardNEQ applies the NEQ predicate on the "card" field.
func CardNEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCard, v))
}

// CardIn applies the In predicate on the "card" field.
func CardIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCard, vs...))
}

// CardNotIn applies the NotIn predicate on the "card" field.
func CardNotIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCard, vs...))
}

// CardGT applies the GT predicate on the "card" field.
func CardGT(v int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCard, v))
}

// CardGTE applies the GTE predicate on the "card" field.
func CardGTE(v int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCard, v))
}

// CardLT applies the LT predicate on the "card" field.
func CardLT(v int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCard, v))
}

// CardLTE applies the LTE predicate on the "card" field.
func CardLTE(v int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCard, v))
}

// CardIsNil applies the IsNil predicate on the "card" field.
func CardIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldCard))
}

// CardNotNil applies the NotNil predicate on the "card" field.
func CardNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldCard))
}

// SkillEQ applies the EQ predicate on the "skill" field.
func SkillEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldSkill, v))
}

// SkillNEQ applies the NEQ predicate on the "skill" field.
func SkillNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldSkill, v))
}

// SkillIn applies the In predicate on the "skill" field.
func SkillIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldSkill, vs...))
}

// SkillNotIn applies the NotIn predicate on the "skill" field.
func SkillNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldSkill, vs...))
}

// SkillGT applies the GT predicate on the "skill" field.
func SkillGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldSkill, v))
}

// SkillGTE applies the GTE predicate on the "skill" field.
func SkillGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldSkill, v))
}

// SkillLT applies the LT predicate on the "skill" field.
func SkillLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldSkill, v))
}

// SkillLTE applies the LTE predicate on the "skill" field.
func SkillLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldSkill, v))
}

// SkillContains applies the Contains predicate on the "skill" field.
func SkillContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldSkill, v))
}

// SkillHasPrefix applies the HasPrefix predicate on the "skill" field.
func SkillHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldSkill, v))
}

// SkillHasSuffix applies the HasSuffix predicate on the "skill" field.
func SkillHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldSkill, v))
}

// SkillIsNil applies the IsNil predicate on the "skill" field.
func SkillIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldSkill))
}

// SkillNotNil applies the NotNil predicate on the "skill" field.
func SkillNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldSkill))
}

// SkillEqualFold applies the EqualFold predicate on the "skill" field.
func SkillEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldSkill, v))
}

// SkillContainsFold applies the ContainsFold predicate on the "skill" field.
func SkillContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldSkill, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldStatus, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldToken, v))
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldToken))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldToken, v))
}

// CpEQ applies the EQ predicate on the "cp" field.
func CpEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCp, v))
}

// CpNEQ applies the NEQ predicate on the "cp" field.
func CpNEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCp, v))
}

// CpIn applies the In predicate on the "cp" field.
func CpIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCp, vs...))
}

// CpNotIn applies the NotIn predicate on the "cp" field.
func CpNotIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCp, vs...))
}

// CpGT applies the GT predicate on the "cp" field.
func CpGT(v int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCp, v))
}

// CpGTE applies the GTE predicate on the "cp" field.
func CpGTE(v int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCp, v))
}

// CpLT applies the LT predicate on the "cp" field.
func CpLT(v int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCp, v))
}

// CpLTE applies the LTE predicate on the "cp" field.
func CpLTE(v int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCp, v))
}

// CpIsNil applies the IsNil predicate on the "cp" field.
func CpIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldCp))
}

// CpNotNil applies the NotNil predicate on the "cp" field.
func CpNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldCp))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldURL, v))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldURL))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldCreatedAt))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
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
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		p(s.Not())
	})
}
