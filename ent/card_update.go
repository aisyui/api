// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"t/ent/card"
	"t/ent/predicate"
	"t/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetSkill sets the "skill" field.
func (cu *CardUpdate) SetSkill(s string) *CardUpdate {
	cu.mutation.SetSkill(s)
	return cu
}

// SetNillableSkill sets the "skill" field if the given value is not nil.
func (cu *CardUpdate) SetNillableSkill(s *string) *CardUpdate {
	if s != nil {
		cu.SetSkill(*s)
	}
	return cu
}

// ClearSkill clears the value of the "skill" field.
func (cu *CardUpdate) ClearSkill() *CardUpdate {
	cu.mutation.ClearSkill()
	return cu
}

// SetStatus sets the "status" field.
func (cu *CardUpdate) SetStatus(s string) *CardUpdate {
	cu.mutation.SetStatus(s)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CardUpdate) SetNillableStatus(s *string) *CardUpdate {
	if s != nil {
		cu.SetStatus(*s)
	}
	return cu
}

// ClearStatus clears the value of the "status" field.
func (cu *CardUpdate) ClearStatus() *CardUpdate {
	cu.mutation.ClearStatus()
	return cu
}

// SetToken sets the "token" field.
func (cu *CardUpdate) SetToken(s string) *CardUpdate {
	cu.mutation.SetToken(s)
	return cu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (cu *CardUpdate) SetNillableToken(s *string) *CardUpdate {
	if s != nil {
		cu.SetToken(*s)
	}
	return cu
}

// ClearToken clears the value of the "token" field.
func (cu *CardUpdate) ClearToken() *CardUpdate {
	cu.mutation.ClearToken()
	return cu
}

// SetCp sets the "cp" field.
func (cu *CardUpdate) SetCp(i int) *CardUpdate {
	cu.mutation.ResetCp()
	cu.mutation.SetCp(i)
	return cu
}

// SetNillableCp sets the "cp" field if the given value is not nil.
func (cu *CardUpdate) SetNillableCp(i *int) *CardUpdate {
	if i != nil {
		cu.SetCp(*i)
	}
	return cu
}

// AddCp adds i to the "cp" field.
func (cu *CardUpdate) AddCp(i int) *CardUpdate {
	cu.mutation.AddCp(i)
	return cu
}

// ClearCp clears the value of the "cp" field.
func (cu *CardUpdate) ClearCp() *CardUpdate {
	cu.mutation.ClearCp()
	return cu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *CardUpdate) SetOwnerID(id int) *CardUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *CardUpdate) SetOwner(u *User) *CardUpdate {
	return cu.SetOwnerID(u.ID)
}

// Mutation returns the CardMutation object of the builder.
func (cu *CardUpdate) Mutation() *CardMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *CardUpdate) ClearOwner() *CardUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CardMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CardUpdate) check() error {
	if _, ok := cu.mutation.OwnerID(); cu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Card.owner"`)
	}
	return nil
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.CardCleared() {
		_spec.ClearField(card.FieldCard, field.TypeInt)
	}
	if value, ok := cu.mutation.Skill(); ok {
		_spec.SetField(card.FieldSkill, field.TypeString, value)
	}
	if cu.mutation.SkillCleared() {
		_spec.ClearField(card.FieldSkill, field.TypeString)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(card.FieldStatus, field.TypeString, value)
	}
	if cu.mutation.StatusCleared() {
		_spec.ClearField(card.FieldStatus, field.TypeString)
	}
	if value, ok := cu.mutation.Token(); ok {
		_spec.SetField(card.FieldToken, field.TypeString, value)
	}
	if cu.mutation.TokenCleared() {
		_spec.ClearField(card.FieldToken, field.TypeString)
	}
	if value, ok := cu.mutation.Cp(); ok {
		_spec.SetField(card.FieldCp, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedCp(); ok {
		_spec.AddField(card.FieldCp, field.TypeInt, value)
	}
	if cu.mutation.CpCleared() {
		_spec.ClearField(card.FieldCp, field.TypeInt)
	}
	if cu.mutation.URLCleared() {
		_spec.ClearField(card.FieldURL, field.TypeString)
	}
	if cu.mutation.CreatedAtCleared() {
		_spec.ClearField(card.FieldCreatedAt, field.TypeTime)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetSkill sets the "skill" field.
func (cuo *CardUpdateOne) SetSkill(s string) *CardUpdateOne {
	cuo.mutation.SetSkill(s)
	return cuo
}

// SetNillableSkill sets the "skill" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableSkill(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetSkill(*s)
	}
	return cuo
}

// ClearSkill clears the value of the "skill" field.
func (cuo *CardUpdateOne) ClearSkill() *CardUpdateOne {
	cuo.mutation.ClearSkill()
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CardUpdateOne) SetStatus(s string) *CardUpdateOne {
	cuo.mutation.SetStatus(s)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableStatus(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetStatus(*s)
	}
	return cuo
}

// ClearStatus clears the value of the "status" field.
func (cuo *CardUpdateOne) ClearStatus() *CardUpdateOne {
	cuo.mutation.ClearStatus()
	return cuo
}

// SetToken sets the "token" field.
func (cuo *CardUpdateOne) SetToken(s string) *CardUpdateOne {
	cuo.mutation.SetToken(s)
	return cuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableToken(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetToken(*s)
	}
	return cuo
}

// ClearToken clears the value of the "token" field.
func (cuo *CardUpdateOne) ClearToken() *CardUpdateOne {
	cuo.mutation.ClearToken()
	return cuo
}

// SetCp sets the "cp" field.
func (cuo *CardUpdateOne) SetCp(i int) *CardUpdateOne {
	cuo.mutation.ResetCp()
	cuo.mutation.SetCp(i)
	return cuo
}

// SetNillableCp sets the "cp" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableCp(i *int) *CardUpdateOne {
	if i != nil {
		cuo.SetCp(*i)
	}
	return cuo
}

// AddCp adds i to the "cp" field.
func (cuo *CardUpdateOne) AddCp(i int) *CardUpdateOne {
	cuo.mutation.AddCp(i)
	return cuo
}

// ClearCp clears the value of the "cp" field.
func (cuo *CardUpdateOne) ClearCp() *CardUpdateOne {
	cuo.mutation.ClearCp()
	return cuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *CardUpdateOne) SetOwnerID(id int) *CardUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// Mutation returns the CardMutation object of the builder.
func (cuo *CardUpdateOne) Mutation() *CardMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *CardUpdateOne) ClearOwner() *CardUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Where appends a list predicates to the CardUpdate builder.
func (cuo *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Card entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks[*Card, CardMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CardUpdateOne) check() error {
	if _, ok := cuo.mutation.OwnerID(); cuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Card.owner"`)
	}
	return nil
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.CardCleared() {
		_spec.ClearField(card.FieldCard, field.TypeInt)
	}
	if value, ok := cuo.mutation.Skill(); ok {
		_spec.SetField(card.FieldSkill, field.TypeString, value)
	}
	if cuo.mutation.SkillCleared() {
		_spec.ClearField(card.FieldSkill, field.TypeString)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(card.FieldStatus, field.TypeString, value)
	}
	if cuo.mutation.StatusCleared() {
		_spec.ClearField(card.FieldStatus, field.TypeString)
	}
	if value, ok := cuo.mutation.Token(); ok {
		_spec.SetField(card.FieldToken, field.TypeString, value)
	}
	if cuo.mutation.TokenCleared() {
		_spec.ClearField(card.FieldToken, field.TypeString)
	}
	if value, ok := cuo.mutation.Cp(); ok {
		_spec.SetField(card.FieldCp, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedCp(); ok {
		_spec.AddField(card.FieldCp, field.TypeInt, value)
	}
	if cuo.mutation.CpCleared() {
		_spec.ClearField(card.FieldCp, field.TypeInt)
	}
	if cuo.mutation.URLCleared() {
		_spec.ClearField(card.FieldURL, field.TypeString)
	}
	if cuo.mutation.CreatedAtCleared() {
		_spec.ClearField(card.FieldCreatedAt, field.TypeTime)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Card{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
