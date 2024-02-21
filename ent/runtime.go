// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/card"
	"api/ent/group"
	"api/ent/schema"
	"api/ent/ue"
	"api/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescPassword is the schema descriptor for password field.
	cardDescPassword := cardFields[0].Descriptor()
	// card.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	card.PasswordValidator = cardDescPassword.Validators[0].(func(string) error)
	// cardDescCard is the schema descriptor for card field.
	cardDescCard := cardFields[1].Descriptor()
	// card.DefaultCard holds the default value on creation for the card field.
	card.DefaultCard = cardDescCard.Default.(func() int)
	// cardDescSkill is the schema descriptor for skill field.
	cardDescSkill := cardFields[2].Descriptor()
	// card.DefaultSkill holds the default value on creation for the skill field.
	card.DefaultSkill = cardDescSkill.Default.(func() string)
	// cardDescStatus is the schema descriptor for status field.
	cardDescStatus := cardFields[3].Descriptor()
	// card.DefaultStatus holds the default value on creation for the status field.
	card.DefaultStatus = cardDescStatus.Default.(func() string)
	// cardDescCp is the schema descriptor for cp field.
	cardDescCp := cardFields[5].Descriptor()
	// card.DefaultCp holds the default value on creation for the cp field.
	card.DefaultCp = cardDescCp.Default.(func() int)
	// cardDescURL is the schema descriptor for url field.
	cardDescURL := cardFields[6].Descriptor()
	// card.DefaultURL holds the default value on creation for the url field.
	card.DefaultURL = cardDescURL.Default.(string)
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardFields[9].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescPassword is the schema descriptor for password field.
	groupDescPassword := groupFields[1].Descriptor()
	// group.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	group.PasswordValidator = groupDescPassword.Validators[0].(func(string) error)
	ueFields := schema.Ue{}.Fields()
	_ = ueFields
	// ueDescLimit is the schema descriptor for limit field.
	ueDescLimit := ueFields[0].Descriptor()
	// ue.DefaultLimit holds the default value on creation for the limit field.
	ue.DefaultLimit = ueDescLimit.Default.(bool)
	// ueDescLimitBoss is the schema descriptor for limit_boss field.
	ueDescLimitBoss := ueFields[1].Descriptor()
	// ue.DefaultLimitBoss holds the default value on creation for the limit_boss field.
	ue.DefaultLimitBoss = ueDescLimitBoss.Default.(bool)
	// ueDescLimitItem is the schema descriptor for limit_item field.
	ueDescLimitItem := ueFields[2].Descriptor()
	// ue.DefaultLimitItem holds the default value on creation for the limit_item field.
	ue.DefaultLimitItem = ueDescLimitItem.Default.(bool)
	// ueDescPassword is the schema descriptor for password field.
	ueDescPassword := ueFields[3].Descriptor()
	// ue.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	ue.PasswordValidator = ueDescPassword.Validators[0].(func(string) error)
	// ueDescCreatedAt is the schema descriptor for created_at field.
	ueDescCreatedAt := ueFields[18].Descriptor()
	// ue.DefaultCreatedAt holds the default value on creation for the created_at field.
	ue.DefaultCreatedAt = ueDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescMember is the schema descriptor for member field.
	userDescMember := userFields[2].Descriptor()
	// user.DefaultMember holds the default value on creation for the member field.
	user.DefaultMember = userDescMember.Default.(bool)
	// userDescBook is the schema descriptor for book field.
	userDescBook := userFields[3].Descriptor()
	// user.DefaultBook holds the default value on creation for the book field.
	user.DefaultBook = userDescBook.Default.(bool)
	// userDescManga is the schema descriptor for manga field.
	userDescManga := userFields[4].Descriptor()
	// user.DefaultManga holds the default value on creation for the manga field.
	user.DefaultManga = userDescManga.Default.(bool)
	// userDescBadge is the schema descriptor for badge field.
	userDescBadge := userFields[5].Descriptor()
	// user.DefaultBadge holds the default value on creation for the badge field.
	user.DefaultBadge = userDescBadge.Default.(bool)
	// userDescBsky is the schema descriptor for bsky field.
	userDescBsky := userFields[6].Descriptor()
	// user.DefaultBsky holds the default value on creation for the bsky field.
	user.DefaultBsky = userDescBsky.Default.(bool)
	// userDescMastodon is the schema descriptor for mastodon field.
	userDescMastodon := userFields[7].Descriptor()
	// user.DefaultMastodon holds the default value on creation for the mastodon field.
	user.DefaultMastodon = userDescMastodon.Default.(bool)
	// userDescDelete is the schema descriptor for delete field.
	userDescDelete := userFields[8].Descriptor()
	// user.DefaultDelete holds the default value on creation for the delete field.
	user.DefaultDelete = userDescDelete.Default.(bool)
	// userDescHandle is the schema descriptor for handle field.
	userDescHandle := userFields[9].Descriptor()
	// user.DefaultHandle holds the default value on creation for the handle field.
	user.DefaultHandle = userDescHandle.Default.(bool)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[11].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[12].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[13].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescRaidAt is the schema descriptor for raid_at field.
	userDescRaidAt := userFields[14].Descriptor()
	// user.DefaultRaidAt holds the default value on creation for the raid_at field.
	user.DefaultRaidAt = userDescRaidAt.Default.(func() time.Time)
	// userDescServerAt is the schema descriptor for server_at field.
	userDescServerAt := userFields[15].Descriptor()
	// user.DefaultServerAt holds the default value on creation for the server_at field.
	user.DefaultServerAt = userDescServerAt.Default.(func() time.Time)
	// userDescEggAt is the schema descriptor for egg_at field.
	userDescEggAt := userFields[16].Descriptor()
	// user.DefaultEggAt holds the default value on creation for the egg_at field.
	user.DefaultEggAt = userDescEggAt.Default.(func() time.Time)
	// userDescLuckAt is the schema descriptor for luck_at field.
	userDescLuckAt := userFields[18].Descriptor()
	// user.DefaultLuckAt holds the default value on creation for the luck_at field.
	user.DefaultLuckAt = userDescLuckAt.Default.(func() time.Time)
	// userDescLikeAt is the schema descriptor for like_at field.
	userDescLikeAt := userFields[21].Descriptor()
	// user.DefaultLikeAt holds the default value on creation for the like_at field.
	user.DefaultLikeAt = userDescLikeAt.Default.(func() time.Time)
	// userDescTenAt is the schema descriptor for ten_at field.
	userDescTenAt := userFields[31].Descriptor()
	// user.DefaultTenAt holds the default value on creation for the ten_at field.
	user.DefaultTenAt = userDescTenAt.Default.(func() time.Time)
	// userDescNext is the schema descriptor for next field.
	userDescNext := userFields[32].Descriptor()
	// user.DefaultNext holds the default value on creation for the next field.
	user.DefaultNext = userDescNext.Default.(string)
	// userDescModelAt is the schema descriptor for model_at field.
	userDescModelAt := userFields[35].Descriptor()
	// user.DefaultModelAt holds the default value on creation for the model_at field.
	user.DefaultModelAt = userDescModelAt.Default.(func() time.Time)
	// userDescGame is the schema descriptor for game field.
	userDescGame := userFields[42].Descriptor()
	// user.DefaultGame holds the default value on creation for the game field.
	user.DefaultGame = userDescGame.Default.(bool)
	// userDescGameTest is the schema descriptor for game_test field.
	userDescGameTest := userFields[43].Descriptor()
	// user.DefaultGameTest holds the default value on creation for the game_test field.
	user.DefaultGameTest = userDescGameTest.Default.(bool)
	// userDescGameEnd is the schema descriptor for game_end field.
	userDescGameEnd := userFields[44].Descriptor()
	// user.DefaultGameEnd holds the default value on creation for the game_end field.
	user.DefaultGameEnd = userDescGameEnd.Default.(bool)
	// userDescGameAccount is the schema descriptor for game_account field.
	userDescGameAccount := userFields[45].Descriptor()
	// user.DefaultGameAccount holds the default value on creation for the game_account field.
	user.DefaultGameAccount = userDescGameAccount.Default.(bool)
}
