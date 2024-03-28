// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/card"
	"api/ent/ue"
	"api/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetDid sets the "did" field.
func (uc *UserCreate) SetDid(s string) *UserCreate {
	uc.mutation.SetDid(s)
	return uc
}

// SetNillableDid sets the "did" field if the given value is not nil.
func (uc *UserCreate) SetNillableDid(s *string) *UserCreate {
	if s != nil {
		uc.SetDid(*s)
	}
	return uc
}

// SetMember sets the "member" field.
func (uc *UserCreate) SetMember(b bool) *UserCreate {
	uc.mutation.SetMember(b)
	return uc
}

// SetNillableMember sets the "member" field if the given value is not nil.
func (uc *UserCreate) SetNillableMember(b *bool) *UserCreate {
	if b != nil {
		uc.SetMember(*b)
	}
	return uc
}

// SetBook sets the "book" field.
func (uc *UserCreate) SetBook(b bool) *UserCreate {
	uc.mutation.SetBook(b)
	return uc
}

// SetNillableBook sets the "book" field if the given value is not nil.
func (uc *UserCreate) SetNillableBook(b *bool) *UserCreate {
	if b != nil {
		uc.SetBook(*b)
	}
	return uc
}

// SetManga sets the "manga" field.
func (uc *UserCreate) SetManga(b bool) *UserCreate {
	uc.mutation.SetManga(b)
	return uc
}

// SetNillableManga sets the "manga" field if the given value is not nil.
func (uc *UserCreate) SetNillableManga(b *bool) *UserCreate {
	if b != nil {
		uc.SetManga(*b)
	}
	return uc
}

// SetBadge sets the "badge" field.
func (uc *UserCreate) SetBadge(b bool) *UserCreate {
	uc.mutation.SetBadge(b)
	return uc
}

// SetNillableBadge sets the "badge" field if the given value is not nil.
func (uc *UserCreate) SetNillableBadge(b *bool) *UserCreate {
	if b != nil {
		uc.SetBadge(*b)
	}
	return uc
}

// SetBsky sets the "bsky" field.
func (uc *UserCreate) SetBsky(b bool) *UserCreate {
	uc.mutation.SetBsky(b)
	return uc
}

// SetNillableBsky sets the "bsky" field if the given value is not nil.
func (uc *UserCreate) SetNillableBsky(b *bool) *UserCreate {
	if b != nil {
		uc.SetBsky(*b)
	}
	return uc
}

// SetMastodon sets the "mastodon" field.
func (uc *UserCreate) SetMastodon(b bool) *UserCreate {
	uc.mutation.SetMastodon(b)
	return uc
}

// SetNillableMastodon sets the "mastodon" field if the given value is not nil.
func (uc *UserCreate) SetNillableMastodon(b *bool) *UserCreate {
	if b != nil {
		uc.SetMastodon(*b)
	}
	return uc
}

// SetDelete sets the "delete" field.
func (uc *UserCreate) SetDelete(b bool) *UserCreate {
	uc.mutation.SetDelete(b)
	return uc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (uc *UserCreate) SetNillableDelete(b *bool) *UserCreate {
	if b != nil {
		uc.SetDelete(*b)
	}
	return uc
}

// SetHandle sets the "handle" field.
func (uc *UserCreate) SetHandle(b bool) *UserCreate {
	uc.mutation.SetHandle(b)
	return uc
}

// SetNillableHandle sets the "handle" field if the given value is not nil.
func (uc *UserCreate) SetNillableHandle(b *bool) *UserCreate {
	if b != nil {
		uc.SetHandle(*b)
	}
	return uc
}

// SetToken sets the "token" field.
func (uc *UserCreate) SetToken(s string) *UserCreate {
	uc.mutation.SetToken(s)
	return uc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uc *UserCreate) SetNillableToken(s *string) *UserCreate {
	if s != nil {
		uc.SetToken(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetRaidAt sets the "raid_at" field.
func (uc *UserCreate) SetRaidAt(t time.Time) *UserCreate {
	uc.mutation.SetRaidAt(t)
	return uc
}

// SetNillableRaidAt sets the "raid_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableRaidAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetRaidAt(*t)
	}
	return uc
}

// SetServerAt sets the "server_at" field.
func (uc *UserCreate) SetServerAt(t time.Time) *UserCreate {
	uc.mutation.SetServerAt(t)
	return uc
}

// SetNillableServerAt sets the "server_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableServerAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetServerAt(*t)
	}
	return uc
}

// SetEggAt sets the "egg_at" field.
func (uc *UserCreate) SetEggAt(t time.Time) *UserCreate {
	uc.mutation.SetEggAt(t)
	return uc
}

// SetNillableEggAt sets the "egg_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableEggAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetEggAt(*t)
	}
	return uc
}

// SetLuck sets the "luck" field.
func (uc *UserCreate) SetLuck(i int) *UserCreate {
	uc.mutation.SetLuck(i)
	return uc
}

// SetNillableLuck sets the "luck" field if the given value is not nil.
func (uc *UserCreate) SetNillableLuck(i *int) *UserCreate {
	if i != nil {
		uc.SetLuck(*i)
	}
	return uc
}

// SetLuckAt sets the "luck_at" field.
func (uc *UserCreate) SetLuckAt(t time.Time) *UserCreate {
	uc.mutation.SetLuckAt(t)
	return uc
}

// SetNillableLuckAt sets the "luck_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableLuckAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLuckAt(*t)
	}
	return uc
}

// SetLike sets the "like" field.
func (uc *UserCreate) SetLike(i int) *UserCreate {
	uc.mutation.SetLike(i)
	return uc
}

// SetNillableLike sets the "like" field if the given value is not nil.
func (uc *UserCreate) SetNillableLike(i *int) *UserCreate {
	if i != nil {
		uc.SetLike(*i)
	}
	return uc
}

// SetLikeRank sets the "like_rank" field.
func (uc *UserCreate) SetLikeRank(i int) *UserCreate {
	uc.mutation.SetLikeRank(i)
	return uc
}

// SetNillableLikeRank sets the "like_rank" field if the given value is not nil.
func (uc *UserCreate) SetNillableLikeRank(i *int) *UserCreate {
	if i != nil {
		uc.SetLikeRank(*i)
	}
	return uc
}

// SetLikeAt sets the "like_at" field.
func (uc *UserCreate) SetLikeAt(t time.Time) *UserCreate {
	uc.mutation.SetLikeAt(t)
	return uc
}

// SetNillableLikeAt sets the "like_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableLikeAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLikeAt(*t)
	}
	return uc
}

// SetFav sets the "fav" field.
func (uc *UserCreate) SetFav(i int) *UserCreate {
	uc.mutation.SetFav(i)
	return uc
}

// SetNillableFav sets the "fav" field if the given value is not nil.
func (uc *UserCreate) SetNillableFav(i *int) *UserCreate {
	if i != nil {
		uc.SetFav(*i)
	}
	return uc
}

// SetTen sets the "ten" field.
func (uc *UserCreate) SetTen(b bool) *UserCreate {
	uc.mutation.SetTen(b)
	return uc
}

// SetNillableTen sets the "ten" field if the given value is not nil.
func (uc *UserCreate) SetNillableTen(b *bool) *UserCreate {
	if b != nil {
		uc.SetTen(*b)
	}
	return uc
}

// SetTenSu sets the "ten_su" field.
func (uc *UserCreate) SetTenSu(i int) *UserCreate {
	uc.mutation.SetTenSu(i)
	return uc
}

// SetNillableTenSu sets the "ten_su" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenSu(i *int) *UserCreate {
	if i != nil {
		uc.SetTenSu(*i)
	}
	return uc
}

// SetTenKai sets the "ten_kai" field.
func (uc *UserCreate) SetTenKai(i int) *UserCreate {
	uc.mutation.SetTenKai(i)
	return uc
}

// SetNillableTenKai sets the "ten_kai" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenKai(i *int) *UserCreate {
	if i != nil {
		uc.SetTenKai(*i)
	}
	return uc
}

// SetAiten sets the "aiten" field.
func (uc *UserCreate) SetAiten(i int) *UserCreate {
	uc.mutation.SetAiten(i)
	return uc
}

// SetNillableAiten sets the "aiten" field if the given value is not nil.
func (uc *UserCreate) SetNillableAiten(i *int) *UserCreate {
	if i != nil {
		uc.SetAiten(*i)
	}
	return uc
}

// SetTenCard sets the "ten_card" field.
func (uc *UserCreate) SetTenCard(s string) *UserCreate {
	uc.mutation.SetTenCard(s)
	return uc
}

// SetNillableTenCard sets the "ten_card" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenCard(s *string) *UserCreate {
	if s != nil {
		uc.SetTenCard(*s)
	}
	return uc
}

// SetTenDelete sets the "ten_delete" field.
func (uc *UserCreate) SetTenDelete(s string) *UserCreate {
	uc.mutation.SetTenDelete(s)
	return uc
}

// SetNillableTenDelete sets the "ten_delete" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenDelete(s *string) *UserCreate {
	if s != nil {
		uc.SetTenDelete(*s)
	}
	return uc
}

// SetTenPost sets the "ten_post" field.
func (uc *UserCreate) SetTenPost(s string) *UserCreate {
	uc.mutation.SetTenPost(s)
	return uc
}

// SetNillableTenPost sets the "ten_post" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenPost(s *string) *UserCreate {
	if s != nil {
		uc.SetTenPost(*s)
	}
	return uc
}

// SetTenGet sets the "ten_get" field.
func (uc *UserCreate) SetTenGet(s string) *UserCreate {
	uc.mutation.SetTenGet(s)
	return uc
}

// SetNillableTenGet sets the "ten_get" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenGet(s *string) *UserCreate {
	if s != nil {
		uc.SetTenGet(*s)
	}
	return uc
}

// SetTenAt sets the "ten_at" field.
func (uc *UserCreate) SetTenAt(t time.Time) *UserCreate {
	uc.mutation.SetTenAt(t)
	return uc
}

// SetNillableTenAt sets the "ten_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableTenAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetTenAt(*t)
	}
	return uc
}

// SetNext sets the "next" field.
func (uc *UserCreate) SetNext(s string) *UserCreate {
	uc.mutation.SetNext(s)
	return uc
}

// SetNillableNext sets the "next" field if the given value is not nil.
func (uc *UserCreate) SetNillableNext(s *string) *UserCreate {
	if s != nil {
		uc.SetNext(*s)
	}
	return uc
}

// SetRoom sets the "room" field.
func (uc *UserCreate) SetRoom(i int) *UserCreate {
	uc.mutation.SetRoom(i)
	return uc
}

// SetNillableRoom sets the "room" field if the given value is not nil.
func (uc *UserCreate) SetNillableRoom(i *int) *UserCreate {
	if i != nil {
		uc.SetRoom(*i)
	}
	return uc
}

// SetModel sets the "model" field.
func (uc *UserCreate) SetModel(b bool) *UserCreate {
	uc.mutation.SetModel(b)
	return uc
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (uc *UserCreate) SetNillableModel(b *bool) *UserCreate {
	if b != nil {
		uc.SetModel(*b)
	}
	return uc
}

// SetModelAt sets the "model_at" field.
func (uc *UserCreate) SetModelAt(t time.Time) *UserCreate {
	uc.mutation.SetModelAt(t)
	return uc
}

// SetNillableModelAt sets the "model_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetModelAt(*t)
	}
	return uc
}

// SetModelAttack sets the "model_attack" field.
func (uc *UserCreate) SetModelAttack(i int) *UserCreate {
	uc.mutation.SetModelAttack(i)
	return uc
}

// SetNillableModelAttack sets the "model_attack" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelAttack(i *int) *UserCreate {
	if i != nil {
		uc.SetModelAttack(*i)
	}
	return uc
}

// SetModelLimit sets the "model_limit" field.
func (uc *UserCreate) SetModelLimit(i int) *UserCreate {
	uc.mutation.SetModelLimit(i)
	return uc
}

// SetNillableModelLimit sets the "model_limit" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelLimit(i *int) *UserCreate {
	if i != nil {
		uc.SetModelLimit(*i)
	}
	return uc
}

// SetModelSkill sets the "model_skill" field.
func (uc *UserCreate) SetModelSkill(i int) *UserCreate {
	uc.mutation.SetModelSkill(i)
	return uc
}

// SetNillableModelSkill sets the "model_skill" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelSkill(i *int) *UserCreate {
	if i != nil {
		uc.SetModelSkill(*i)
	}
	return uc
}

// SetModelMode sets the "model_mode" field.
func (uc *UserCreate) SetModelMode(i int) *UserCreate {
	uc.mutation.SetModelMode(i)
	return uc
}

// SetNillableModelMode sets the "model_mode" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelMode(i *int) *UserCreate {
	if i != nil {
		uc.SetModelMode(*i)
	}
	return uc
}

// SetModelCritical sets the "model_critical" field.
func (uc *UserCreate) SetModelCritical(i int) *UserCreate {
	uc.mutation.SetModelCritical(i)
	return uc
}

// SetNillableModelCritical sets the "model_critical" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelCritical(i *int) *UserCreate {
	if i != nil {
		uc.SetModelCritical(*i)
	}
	return uc
}

// SetModelCriticalD sets the "model_critical_d" field.
func (uc *UserCreate) SetModelCriticalD(i int) *UserCreate {
	uc.mutation.SetModelCriticalD(i)
	return uc
}

// SetNillableModelCriticalD sets the "model_critical_d" field if the given value is not nil.
func (uc *UserCreate) SetNillableModelCriticalD(i *int) *UserCreate {
	if i != nil {
		uc.SetModelCriticalD(*i)
	}
	return uc
}

// SetGame sets the "game" field.
func (uc *UserCreate) SetGame(b bool) *UserCreate {
	uc.mutation.SetGame(b)
	return uc
}

// SetNillableGame sets the "game" field if the given value is not nil.
func (uc *UserCreate) SetNillableGame(b *bool) *UserCreate {
	if b != nil {
		uc.SetGame(*b)
	}
	return uc
}

// SetGameTest sets the "game_test" field.
func (uc *UserCreate) SetGameTest(b bool) *UserCreate {
	uc.mutation.SetGameTest(b)
	return uc
}

// SetNillableGameTest sets the "game_test" field if the given value is not nil.
func (uc *UserCreate) SetNillableGameTest(b *bool) *UserCreate {
	if b != nil {
		uc.SetGameTest(*b)
	}
	return uc
}

// SetGameEnd sets the "game_end" field.
func (uc *UserCreate) SetGameEnd(b bool) *UserCreate {
	uc.mutation.SetGameEnd(b)
	return uc
}

// SetNillableGameEnd sets the "game_end" field if the given value is not nil.
func (uc *UserCreate) SetNillableGameEnd(b *bool) *UserCreate {
	if b != nil {
		uc.SetGameEnd(*b)
	}
	return uc
}

// SetGameAccount sets the "game_account" field.
func (uc *UserCreate) SetGameAccount(b bool) *UserCreate {
	uc.mutation.SetGameAccount(b)
	return uc
}

// SetNillableGameAccount sets the "game_account" field if the given value is not nil.
func (uc *UserCreate) SetNillableGameAccount(b *bool) *UserCreate {
	if b != nil {
		uc.SetGameAccount(*b)
	}
	return uc
}

// SetGameLv sets the "game_lv" field.
func (uc *UserCreate) SetGameLv(i int) *UserCreate {
	uc.mutation.SetGameLv(i)
	return uc
}

// SetNillableGameLv sets the "game_lv" field if the given value is not nil.
func (uc *UserCreate) SetNillableGameLv(i *int) *UserCreate {
	if i != nil {
		uc.SetGameLv(*i)
	}
	return uc
}

// SetCoin sets the "coin" field.
func (uc *UserCreate) SetCoin(i int) *UserCreate {
	uc.mutation.SetCoin(i)
	return uc
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (uc *UserCreate) SetNillableCoin(i *int) *UserCreate {
	if i != nil {
		uc.SetCoin(*i)
	}
	return uc
}

// SetCoinOpen sets the "coin_open" field.
func (uc *UserCreate) SetCoinOpen(b bool) *UserCreate {
	uc.mutation.SetCoinOpen(b)
	return uc
}

// SetNillableCoinOpen sets the "coin_open" field if the given value is not nil.
func (uc *UserCreate) SetNillableCoinOpen(b *bool) *UserCreate {
	if b != nil {
		uc.SetCoinOpen(*b)
	}
	return uc
}

// SetCoinAt sets the "coin_at" field.
func (uc *UserCreate) SetCoinAt(t time.Time) *UserCreate {
	uc.mutation.SetCoinAt(t)
	return uc
}

// SetNillableCoinAt sets the "coin_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCoinAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCoinAt(*t)
	}
	return uc
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (uc *UserCreate) AddCardIDs(ids ...int) *UserCreate {
	uc.mutation.AddCardIDs(ids...)
	return uc
}

// AddCard adds the "card" edges to the Card entity.
func (uc *UserCreate) AddCard(c ...*Card) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCardIDs(ids...)
}

// AddUeIDs adds the "ue" edge to the Ue entity by IDs.
func (uc *UserCreate) AddUeIDs(ids ...int) *UserCreate {
	uc.mutation.AddUeIDs(ids...)
	return uc
}

// AddUe adds the "ue" edges to the Ue entity.
func (uc *UserCreate) AddUe(u ...*Ue) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUeIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Member(); !ok {
		v := user.DefaultMember
		uc.mutation.SetMember(v)
	}
	if _, ok := uc.mutation.Book(); !ok {
		v := user.DefaultBook
		uc.mutation.SetBook(v)
	}
	if _, ok := uc.mutation.Manga(); !ok {
		v := user.DefaultManga
		uc.mutation.SetManga(v)
	}
	if _, ok := uc.mutation.Badge(); !ok {
		v := user.DefaultBadge
		uc.mutation.SetBadge(v)
	}
	if _, ok := uc.mutation.Bsky(); !ok {
		v := user.DefaultBsky
		uc.mutation.SetBsky(v)
	}
	if _, ok := uc.mutation.Mastodon(); !ok {
		v := user.DefaultMastodon
		uc.mutation.SetMastodon(v)
	}
	if _, ok := uc.mutation.Delete(); !ok {
		v := user.DefaultDelete
		uc.mutation.SetDelete(v)
	}
	if _, ok := uc.mutation.Handle(); !ok {
		v := user.DefaultHandle
		uc.mutation.SetHandle(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.RaidAt(); !ok {
		v := user.DefaultRaidAt()
		uc.mutation.SetRaidAt(v)
	}
	if _, ok := uc.mutation.ServerAt(); !ok {
		v := user.DefaultServerAt()
		uc.mutation.SetServerAt(v)
	}
	if _, ok := uc.mutation.EggAt(); !ok {
		v := user.DefaultEggAt()
		uc.mutation.SetEggAt(v)
	}
	if _, ok := uc.mutation.LuckAt(); !ok {
		v := user.DefaultLuckAt()
		uc.mutation.SetLuckAt(v)
	}
	if _, ok := uc.mutation.LikeAt(); !ok {
		v := user.DefaultLikeAt()
		uc.mutation.SetLikeAt(v)
	}
	if _, ok := uc.mutation.TenAt(); !ok {
		v := user.DefaultTenAt()
		uc.mutation.SetTenAt(v)
	}
	if _, ok := uc.mutation.Next(); !ok {
		v := user.DefaultNext
		uc.mutation.SetNext(v)
	}
	if _, ok := uc.mutation.ModelAt(); !ok {
		v := user.DefaultModelAt()
		uc.mutation.SetModelAt(v)
	}
	if _, ok := uc.mutation.Game(); !ok {
		v := user.DefaultGame
		uc.mutation.SetGame(v)
	}
	if _, ok := uc.mutation.GameTest(); !ok {
		v := user.DefaultGameTest
		uc.mutation.SetGameTest(v)
	}
	if _, ok := uc.mutation.GameEnd(); !ok {
		v := user.DefaultGameEnd
		uc.mutation.SetGameEnd(v)
	}
	if _, ok := uc.mutation.GameAccount(); !ok {
		v := user.DefaultGameAccount
		uc.mutation.SetGameAccount(v)
	}
	if _, ok := uc.mutation.CoinOpen(); !ok {
		v := user.DefaultCoinOpen
		uc.mutation.SetCoinOpen(v)
	}
	if _, ok := uc.mutation.CoinAt(); !ok {
		v := user.DefaultCoinAt()
		uc.mutation.SetCoinAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Did(); ok {
		_spec.SetField(user.FieldDid, field.TypeString, value)
		_node.Did = value
	}
	if value, ok := uc.mutation.Member(); ok {
		_spec.SetField(user.FieldMember, field.TypeBool, value)
		_node.Member = value
	}
	if value, ok := uc.mutation.Book(); ok {
		_spec.SetField(user.FieldBook, field.TypeBool, value)
		_node.Book = value
	}
	if value, ok := uc.mutation.Manga(); ok {
		_spec.SetField(user.FieldManga, field.TypeBool, value)
		_node.Manga = value
	}
	if value, ok := uc.mutation.Badge(); ok {
		_spec.SetField(user.FieldBadge, field.TypeBool, value)
		_node.Badge = value
	}
	if value, ok := uc.mutation.Bsky(); ok {
		_spec.SetField(user.FieldBsky, field.TypeBool, value)
		_node.Bsky = value
	}
	if value, ok := uc.mutation.Mastodon(); ok {
		_spec.SetField(user.FieldMastodon, field.TypeBool, value)
		_node.Mastodon = value
	}
	if value, ok := uc.mutation.Delete(); ok {
		_spec.SetField(user.FieldDelete, field.TypeBool, value)
		_node.Delete = value
	}
	if value, ok := uc.mutation.Handle(); ok {
		_spec.SetField(user.FieldHandle, field.TypeBool, value)
		_node.Handle = value
	}
	if value, ok := uc.mutation.Token(); ok {
		_spec.SetField(user.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.RaidAt(); ok {
		_spec.SetField(user.FieldRaidAt, field.TypeTime, value)
		_node.RaidAt = value
	}
	if value, ok := uc.mutation.ServerAt(); ok {
		_spec.SetField(user.FieldServerAt, field.TypeTime, value)
		_node.ServerAt = value
	}
	if value, ok := uc.mutation.EggAt(); ok {
		_spec.SetField(user.FieldEggAt, field.TypeTime, value)
		_node.EggAt = value
	}
	if value, ok := uc.mutation.Luck(); ok {
		_spec.SetField(user.FieldLuck, field.TypeInt, value)
		_node.Luck = value
	}
	if value, ok := uc.mutation.LuckAt(); ok {
		_spec.SetField(user.FieldLuckAt, field.TypeTime, value)
		_node.LuckAt = value
	}
	if value, ok := uc.mutation.Like(); ok {
		_spec.SetField(user.FieldLike, field.TypeInt, value)
		_node.Like = value
	}
	if value, ok := uc.mutation.LikeRank(); ok {
		_spec.SetField(user.FieldLikeRank, field.TypeInt, value)
		_node.LikeRank = value
	}
	if value, ok := uc.mutation.LikeAt(); ok {
		_spec.SetField(user.FieldLikeAt, field.TypeTime, value)
		_node.LikeAt = value
	}
	if value, ok := uc.mutation.Fav(); ok {
		_spec.SetField(user.FieldFav, field.TypeInt, value)
		_node.Fav = value
	}
	if value, ok := uc.mutation.Ten(); ok {
		_spec.SetField(user.FieldTen, field.TypeBool, value)
		_node.Ten = value
	}
	if value, ok := uc.mutation.TenSu(); ok {
		_spec.SetField(user.FieldTenSu, field.TypeInt, value)
		_node.TenSu = value
	}
	if value, ok := uc.mutation.TenKai(); ok {
		_spec.SetField(user.FieldTenKai, field.TypeInt, value)
		_node.TenKai = value
	}
	if value, ok := uc.mutation.Aiten(); ok {
		_spec.SetField(user.FieldAiten, field.TypeInt, value)
		_node.Aiten = value
	}
	if value, ok := uc.mutation.TenCard(); ok {
		_spec.SetField(user.FieldTenCard, field.TypeString, value)
		_node.TenCard = value
	}
	if value, ok := uc.mutation.TenDelete(); ok {
		_spec.SetField(user.FieldTenDelete, field.TypeString, value)
		_node.TenDelete = value
	}
	if value, ok := uc.mutation.TenPost(); ok {
		_spec.SetField(user.FieldTenPost, field.TypeString, value)
		_node.TenPost = value
	}
	if value, ok := uc.mutation.TenGet(); ok {
		_spec.SetField(user.FieldTenGet, field.TypeString, value)
		_node.TenGet = value
	}
	if value, ok := uc.mutation.TenAt(); ok {
		_spec.SetField(user.FieldTenAt, field.TypeTime, value)
		_node.TenAt = value
	}
	if value, ok := uc.mutation.Next(); ok {
		_spec.SetField(user.FieldNext, field.TypeString, value)
		_node.Next = value
	}
	if value, ok := uc.mutation.Room(); ok {
		_spec.SetField(user.FieldRoom, field.TypeInt, value)
		_node.Room = value
	}
	if value, ok := uc.mutation.Model(); ok {
		_spec.SetField(user.FieldModel, field.TypeBool, value)
		_node.Model = value
	}
	if value, ok := uc.mutation.ModelAt(); ok {
		_spec.SetField(user.FieldModelAt, field.TypeTime, value)
		_node.ModelAt = value
	}
	if value, ok := uc.mutation.ModelAttack(); ok {
		_spec.SetField(user.FieldModelAttack, field.TypeInt, value)
		_node.ModelAttack = value
	}
	if value, ok := uc.mutation.ModelLimit(); ok {
		_spec.SetField(user.FieldModelLimit, field.TypeInt, value)
		_node.ModelLimit = value
	}
	if value, ok := uc.mutation.ModelSkill(); ok {
		_spec.SetField(user.FieldModelSkill, field.TypeInt, value)
		_node.ModelSkill = value
	}
	if value, ok := uc.mutation.ModelMode(); ok {
		_spec.SetField(user.FieldModelMode, field.TypeInt, value)
		_node.ModelMode = value
	}
	if value, ok := uc.mutation.ModelCritical(); ok {
		_spec.SetField(user.FieldModelCritical, field.TypeInt, value)
		_node.ModelCritical = value
	}
	if value, ok := uc.mutation.ModelCriticalD(); ok {
		_spec.SetField(user.FieldModelCriticalD, field.TypeInt, value)
		_node.ModelCriticalD = value
	}
	if value, ok := uc.mutation.Game(); ok {
		_spec.SetField(user.FieldGame, field.TypeBool, value)
		_node.Game = value
	}
	if value, ok := uc.mutation.GameTest(); ok {
		_spec.SetField(user.FieldGameTest, field.TypeBool, value)
		_node.GameTest = value
	}
	if value, ok := uc.mutation.GameEnd(); ok {
		_spec.SetField(user.FieldGameEnd, field.TypeBool, value)
		_node.GameEnd = value
	}
	if value, ok := uc.mutation.GameAccount(); ok {
		_spec.SetField(user.FieldGameAccount, field.TypeBool, value)
		_node.GameAccount = value
	}
	if value, ok := uc.mutation.GameLv(); ok {
		_spec.SetField(user.FieldGameLv, field.TypeInt, value)
		_node.GameLv = value
	}
	if value, ok := uc.mutation.Coin(); ok {
		_spec.SetField(user.FieldCoin, field.TypeInt, value)
		_node.Coin = value
	}
	if value, ok := uc.mutation.CoinOpen(); ok {
		_spec.SetField(user.FieldCoinOpen, field.TypeBool, value)
		_node.CoinOpen = value
	}
	if value, ok := uc.mutation.CoinAt(); ok {
		_spec.SetField(user.FieldCoinAt, field.TypeTime, value)
		_node.CoinAt = value
	}
	if nodes := uc.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UeTable,
			Columns: []string{user.UeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
