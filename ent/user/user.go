// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldDid holds the string denoting the did field in the database.
	FieldDid = "did"
	// FieldMember holds the string denoting the member field in the database.
	FieldMember = "member"
	// FieldBook holds the string denoting the book field in the database.
	FieldBook = "book"
	// FieldManga holds the string denoting the manga field in the database.
	FieldManga = "manga"
	// FieldBadge holds the string denoting the badge field in the database.
	FieldBadge = "badge"
	// FieldBsky holds the string denoting the bsky field in the database.
	FieldBsky = "bsky"
	// FieldMastodon holds the string denoting the mastodon field in the database.
	FieldMastodon = "mastodon"
	// FieldDelete holds the string denoting the delete field in the database.
	FieldDelete = "delete"
	// FieldHandle holds the string denoting the handle field in the database.
	FieldHandle = "handle"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRaidAt holds the string denoting the raid_at field in the database.
	FieldRaidAt = "raid_at"
	// FieldServerAt holds the string denoting the server_at field in the database.
	FieldServerAt = "server_at"
	// FieldEggAt holds the string denoting the egg_at field in the database.
	FieldEggAt = "egg_at"
	// FieldLuck holds the string denoting the luck field in the database.
	FieldLuck = "luck"
	// FieldLuckAt holds the string denoting the luck_at field in the database.
	FieldLuckAt = "luck_at"
	// FieldLike holds the string denoting the like field in the database.
	FieldLike = "like"
	// FieldLikeRank holds the string denoting the like_rank field in the database.
	FieldLikeRank = "like_rank"
	// FieldLikeAt holds the string denoting the like_at field in the database.
	FieldLikeAt = "like_at"
	// FieldFav holds the string denoting the fav field in the database.
	FieldFav = "fav"
	// FieldTen holds the string denoting the ten field in the database.
	FieldTen = "ten"
	// FieldTenSu holds the string denoting the ten_su field in the database.
	FieldTenSu = "ten_su"
	// FieldTenKai holds the string denoting the ten_kai field in the database.
	FieldTenKai = "ten_kai"
	// FieldAiten holds the string denoting the aiten field in the database.
	FieldAiten = "aiten"
	// FieldTenCard holds the string denoting the ten_card field in the database.
	FieldTenCard = "ten_card"
	// FieldTenDelete holds the string denoting the ten_delete field in the database.
	FieldTenDelete = "ten_delete"
	// FieldTenPost holds the string denoting the ten_post field in the database.
	FieldTenPost = "ten_post"
	// FieldTenGet holds the string denoting the ten_get field in the database.
	FieldTenGet = "ten_get"
	// FieldTenAt holds the string denoting the ten_at field in the database.
	FieldTenAt = "ten_at"
	// FieldNext holds the string denoting the next field in the database.
	FieldNext = "next"
	// FieldRoom holds the string denoting the room field in the database.
	FieldRoom = "room"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldModelAt holds the string denoting the model_at field in the database.
	FieldModelAt = "model_at"
	// FieldModelAttack holds the string denoting the model_attack field in the database.
	FieldModelAttack = "model_attack"
	// FieldModelLimit holds the string denoting the model_limit field in the database.
	FieldModelLimit = "model_limit"
	// FieldModelSkill holds the string denoting the model_skill field in the database.
	FieldModelSkill = "model_skill"
	// FieldModelMode holds the string denoting the model_mode field in the database.
	FieldModelMode = "model_mode"
	// FieldModelCritical holds the string denoting the model_critical field in the database.
	FieldModelCritical = "model_critical"
	// FieldModelCriticalD holds the string denoting the model_critical_d field in the database.
	FieldModelCriticalD = "model_critical_d"
	// FieldGame holds the string denoting the game field in the database.
	FieldGame = "game"
	// FieldGameTest holds the string denoting the game_test field in the database.
	FieldGameTest = "game_test"
	// FieldGameEnd holds the string denoting the game_end field in the database.
	FieldGameEnd = "game_end"
	// FieldGameAccount holds the string denoting the game_account field in the database.
	FieldGameAccount = "game_account"
	// FieldGameLv holds the string denoting the game_lv field in the database.
	FieldGameLv = "game_lv"
	// FieldCoin holds the string denoting the coin field in the database.
	FieldCoin = "coin"
	// FieldCoinOpen holds the string denoting the coin_open field in the database.
	FieldCoinOpen = "coin_open"
	// FieldCoinAt holds the string denoting the coin_at field in the database.
	FieldCoinAt = "coin_at"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// EdgeUe holds the string denoting the ue edge name in mutations.
	EdgeUe = "ue"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CardTable is the table that holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "user_card"
	// UeTable is the table that holds the ue relation/edge.
	UeTable = "ues"
	// UeInverseTable is the table name for the Ue entity.
	// It exists in this package in order to avoid circular dependency with the "ue" package.
	UeInverseTable = "ues"
	// UeColumn is the table column denoting the ue relation/edge.
	UeColumn = "user_ue"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldDid,
	FieldMember,
	FieldBook,
	FieldManga,
	FieldBadge,
	FieldBsky,
	FieldMastodon,
	FieldDelete,
	FieldHandle,
	FieldToken,
	FieldPassword,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRaidAt,
	FieldServerAt,
	FieldEggAt,
	FieldLuck,
	FieldLuckAt,
	FieldLike,
	FieldLikeRank,
	FieldLikeAt,
	FieldFav,
	FieldTen,
	FieldTenSu,
	FieldTenKai,
	FieldAiten,
	FieldTenCard,
	FieldTenDelete,
	FieldTenPost,
	FieldTenGet,
	FieldTenAt,
	FieldNext,
	FieldRoom,
	FieldModel,
	FieldModelAt,
	FieldModelAttack,
	FieldModelLimit,
	FieldModelSkill,
	FieldModelMode,
	FieldModelCritical,
	FieldModelCriticalD,
	FieldGame,
	FieldGameTest,
	FieldGameEnd,
	FieldGameAccount,
	FieldGameLv,
	FieldCoin,
	FieldCoinOpen,
	FieldCoinAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_users",
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultMember holds the default value on creation for the "member" field.
	DefaultMember bool
	// DefaultBook holds the default value on creation for the "book" field.
	DefaultBook bool
	// DefaultManga holds the default value on creation for the "manga" field.
	DefaultManga bool
	// DefaultBadge holds the default value on creation for the "badge" field.
	DefaultBadge bool
	// DefaultBsky holds the default value on creation for the "bsky" field.
	DefaultBsky bool
	// DefaultMastodon holds the default value on creation for the "mastodon" field.
	DefaultMastodon bool
	// DefaultDelete holds the default value on creation for the "delete" field.
	DefaultDelete bool
	// DefaultHandle holds the default value on creation for the "handle" field.
	DefaultHandle bool
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultRaidAt holds the default value on creation for the "raid_at" field.
	DefaultRaidAt func() time.Time
	// DefaultServerAt holds the default value on creation for the "server_at" field.
	DefaultServerAt func() time.Time
	// DefaultEggAt holds the default value on creation for the "egg_at" field.
	DefaultEggAt func() time.Time
	// DefaultLuckAt holds the default value on creation for the "luck_at" field.
	DefaultLuckAt func() time.Time
	// DefaultLikeAt holds the default value on creation for the "like_at" field.
	DefaultLikeAt func() time.Time
	// DefaultTenAt holds the default value on creation for the "ten_at" field.
	DefaultTenAt func() time.Time
	// DefaultNext holds the default value on creation for the "next" field.
	DefaultNext string
	// DefaultModelAt holds the default value on creation for the "model_at" field.
	DefaultModelAt func() time.Time
	// DefaultGame holds the default value on creation for the "game" field.
	DefaultGame bool
	// DefaultGameTest holds the default value on creation for the "game_test" field.
	DefaultGameTest bool
	// DefaultGameEnd holds the default value on creation for the "game_end" field.
	DefaultGameEnd bool
	// DefaultGameAccount holds the default value on creation for the "game_account" field.
	DefaultGameAccount bool
	// DefaultCoinOpen holds the default value on creation for the "coin_open" field.
	DefaultCoinOpen bool
	// DefaultCoinAt holds the default value on creation for the "coin_at" field.
	DefaultCoinAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByDid orders the results by the did field.
func ByDid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDid, opts...).ToFunc()
}

// ByMember orders the results by the member field.
func ByMember(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMember, opts...).ToFunc()
}

// ByBook orders the results by the book field.
func ByBook(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBook, opts...).ToFunc()
}

// ByManga orders the results by the manga field.
func ByManga(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManga, opts...).ToFunc()
}

// ByBadge orders the results by the badge field.
func ByBadge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBadge, opts...).ToFunc()
}

// ByBsky orders the results by the bsky field.
func ByBsky(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBsky, opts...).ToFunc()
}

// ByMastodon orders the results by the mastodon field.
func ByMastodon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMastodon, opts...).ToFunc()
}

// ByDelete orders the results by the delete field.
func ByDelete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelete, opts...).ToFunc()
}

// ByHandle orders the results by the handle field.
func ByHandle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHandle, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByRaidAt orders the results by the raid_at field.
func ByRaidAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRaidAt, opts...).ToFunc()
}

// ByServerAt orders the results by the server_at field.
func ByServerAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServerAt, opts...).ToFunc()
}

// ByEggAt orders the results by the egg_at field.
func ByEggAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEggAt, opts...).ToFunc()
}

// ByLuck orders the results by the luck field.
func ByLuck(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLuck, opts...).ToFunc()
}

// ByLuckAt orders the results by the luck_at field.
func ByLuckAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLuckAt, opts...).ToFunc()
}

// ByLike orders the results by the like field.
func ByLike(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLike, opts...).ToFunc()
}

// ByLikeRank orders the results by the like_rank field.
func ByLikeRank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikeRank, opts...).ToFunc()
}

// ByLikeAt orders the results by the like_at field.
func ByLikeAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikeAt, opts...).ToFunc()
}

// ByFav orders the results by the fav field.
func ByFav(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFav, opts...).ToFunc()
}

// ByTen orders the results by the ten field.
func ByTen(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTen, opts...).ToFunc()
}

// ByTenSu orders the results by the ten_su field.
func ByTenSu(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenSu, opts...).ToFunc()
}

// ByTenKai orders the results by the ten_kai field.
func ByTenKai(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenKai, opts...).ToFunc()
}

// ByAiten orders the results by the aiten field.
func ByAiten(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAiten, opts...).ToFunc()
}

// ByTenCard orders the results by the ten_card field.
func ByTenCard(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenCard, opts...).ToFunc()
}

// ByTenDelete orders the results by the ten_delete field.
func ByTenDelete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenDelete, opts...).ToFunc()
}

// ByTenPost orders the results by the ten_post field.
func ByTenPost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenPost, opts...).ToFunc()
}

// ByTenGet orders the results by the ten_get field.
func ByTenGet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenGet, opts...).ToFunc()
}

// ByTenAt orders the results by the ten_at field.
func ByTenAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenAt, opts...).ToFunc()
}

// ByNext orders the results by the next field.
func ByNext(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNext, opts...).ToFunc()
}

// ByRoom orders the results by the room field.
func ByRoom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoom, opts...).ToFunc()
}

// ByModel orders the results by the model field.
func ByModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModel, opts...).ToFunc()
}

// ByModelAt orders the results by the model_at field.
func ByModelAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelAt, opts...).ToFunc()
}

// ByModelAttack orders the results by the model_attack field.
func ByModelAttack(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelAttack, opts...).ToFunc()
}

// ByModelLimit orders the results by the model_limit field.
func ByModelLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelLimit, opts...).ToFunc()
}

// ByModelSkill orders the results by the model_skill field.
func ByModelSkill(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelSkill, opts...).ToFunc()
}

// ByModelMode orders the results by the model_mode field.
func ByModelMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelMode, opts...).ToFunc()
}

// ByModelCritical orders the results by the model_critical field.
func ByModelCritical(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelCritical, opts...).ToFunc()
}

// ByModelCriticalD orders the results by the model_critical_d field.
func ByModelCriticalD(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelCriticalD, opts...).ToFunc()
}

// ByGame orders the results by the game field.
func ByGame(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGame, opts...).ToFunc()
}

// ByGameTest orders the results by the game_test field.
func ByGameTest(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameTest, opts...).ToFunc()
}

// ByGameEnd orders the results by the game_end field.
func ByGameEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameEnd, opts...).ToFunc()
}

// ByGameAccount orders the results by the game_account field.
func ByGameAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameAccount, opts...).ToFunc()
}

// ByGameLv orders the results by the game_lv field.
func ByGameLv(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameLv, opts...).ToFunc()
}

// ByCoin orders the results by the coin field.
func ByCoin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoin, opts...).ToFunc()
}

// ByCoinOpen orders the results by the coin_open field.
func ByCoinOpen(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinOpen, opts...).ToFunc()
}

// ByCoinAt orders the results by the coin_at field.
func ByCoinAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinAt, opts...).ToFunc()
}

// ByCardCount orders the results by card count.
func ByCardCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCardStep(), opts...)
	}
}

// ByCard orders the results by card terms.
func ByCard(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCardStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUeCount orders the results by ue count.
func ByUeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUeStep(), opts...)
	}
}

// ByUe orders the results by ue terms.
func ByUe(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCardStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CardInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CardTable, CardColumn),
	)
}
func newUeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UeTable, UeColumn),
	)
}
