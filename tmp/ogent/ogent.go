// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"t/ent"
	"t/ent/card"
	"t/ent/group"
	"t/ent/user"

	"github.com/go-faster/jx"
	"os"
)

// origin-config
var password = os.Getenv("PASS")
var token = os.Getenv("TOKEN")

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateCard handles POST /cards requests.
func (h *OgentHandler) CreateCard(ctx context.Context, req *CreateCardReq) (CreateCardRes, error) {
	b := h.client.Card.Create()
	// Add all fields.
	b.SetPassword(req.Password)
	if v, ok := req.Card.Get(); ok {
		b.SetCard(v)
	}
	if v, ok := req.Skill.Get(); ok {
		b.SetSkill(v)
	}
	if v, ok := req.Status.Get(); ok {
		b.SetStatus(v)
	}
	if v, ok := req.Cp.Get(); ok {
		b.SetCp(v)
	}
	if v, ok := req.URL.Get(); ok {
		b.SetURL(v)
	}
	if v, ok := req.CreatedAt.Get(); ok {
		b.SetCreatedAt(v)
	}
	// Add all edges.
	//b.SetOwnerID(req.Owner)
	// origin-config
	if req.Password == password {
		b.SetOwnerID(req.Owner)
	} else {
		b.SetOwnerID(0)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Card.Query().Where(card.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewCardCreate(e), nil
}

// ReadCard handles GET /cards/{id} requests.
func (h *OgentHandler) ReadCard(ctx context.Context, params ReadCardParams) (ReadCardRes, error) {
	q := h.client.Card.Query().Where(card.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewCardRead(e), nil
}

// UpdateCard handles PATCH /cards/{id} requests.
func (h *OgentHandler) UpdateCard(ctx context.Context, req *UpdateCardReq, params UpdateCardParams) (UpdateCardRes, error) {
	b := h.client.Card.UpdateOneID(params.ID)
	if v, ok := req.Token.Get(); ok {
		if v == token {
			b.SetToken(v)
			if v, ok := req.Skill.Get(); ok {
				b.SetSkill(v)
			}
			if v, ok := req.Status.Get(); ok {
				b.SetStatus(v)
			}
			if v, ok := req.Token.Get(); ok {
				b.SetToken(v)
			}
			if v, ok := req.Cp.Get(); ok {
				b.SetCp(v)
			}
			if v, ok := req.Owner.Get(); ok {
				b.SetOwnerID(v)
			}
		}
	}

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Card.Query().Where(card.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewCardUpdate(e), nil
}

// DeleteCard handles DELETE /cards/{id} requests.
func (h *OgentHandler) DeleteCard(ctx context.Context, params DeleteCardParams) (DeleteCardRes, error) {
	err := h.client.Card.DeleteOneID(0).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteCardNoContent), nil

}

// ListCard handles GET /cards requests.
func (h *OgentHandler) ListCard(ctx context.Context, params ListCardParams) (ListCardRes, error) {
	q := h.client.Card.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewCardLists(es)
	return (*ListCardOKApplicationJSON)(&r), nil
}

// ReadCardOwner handles GET /cards/{id}/owner requests.
func (h *OgentHandler) ReadCardOwner(ctx context.Context, params ReadCardOwnerParams) (ReadCardOwnerRes, error) {
	q := h.client.Card.Query().Where(card.IDEQ(params.ID)).QueryOwner()
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewCardOwnerRead(e), nil
}

// CreateGroup handles POST /groups requests.
func (h *OgentHandler) CreateGroup(ctx context.Context, req *CreateGroupReq) (CreateGroupRes, error) {
	b := h.client.Group.Create()
	// Add all fields.
	b.SetName("")
	b.SetPassword(req.Password)
	// Add all edges.
	b.AddUserIDs(req.Users...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Group.Query().Where(group.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewGroupCreate(e), nil
}

// ReadGroup handles GET /groups/{id} requests.
func (h *OgentHandler) ReadGroup(ctx context.Context, params ReadGroupParams) (ReadGroupRes, error) {
	q := h.client.Group.Query().Where(group.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewGroupRead(e), nil
}

// UpdateGroup handles PATCH /groups/{id} requests.
func (h *OgentHandler) UpdateGroup(ctx context.Context, req *UpdateGroupReq, params UpdateGroupParams) (UpdateGroupRes, error) {
	b := h.client.Group.UpdateOneID(0)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	// Add all edges.
	if req.Users != nil {
		b.ClearUsers().AddUserIDs(req.Users...)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Group.Query().Where(group.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewGroupUpdate(e), nil
}

// DeleteGroup handles DELETE /groups/{id} requests.
func (h *OgentHandler) DeleteGroup(ctx context.Context, params DeleteGroupParams) (DeleteGroupRes, error) {
	err := h.client.Group.DeleteOneID(0).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteGroupNoContent), nil

}

// ListGroup handles GET /groups requests.
func (h *OgentHandler) ListGroup(ctx context.Context, params ListGroupParams) (ListGroupRes, error) {
	q := h.client.Group.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewGroupLists(es)
	return (*ListGroupOKApplicationJSON)(&r), nil
}

// ListGroupUsers handles GET /groups/{id}/users requests.
func (h *OgentHandler) ListGroupUsers(ctx context.Context, params ListGroupUsersParams) (ListGroupUsersRes, error) {
	q := h.client.Group.Query().Where(group.IDEQ(params.ID)).QueryUsers()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewGroupUsersLists(es)
	return (*ListGroupUsersOKApplicationJSON)(&r), nil
}

// CreateUser handles POST /users requests.
func (h *OgentHandler) CreateUser(ctx context.Context, req *CreateUserReq) (CreateUserRes, error) {
	b := h.client.User.Create()

	// Add all fields.
	//b.SetUsername(req.Username)
	//origin-config
	if req.Password == password {
		b.SetUsername(req.Username)
	} else {
		b.SetUsername("")
	}

	b.SetPassword(req.Password)

	if v, ok := req.Fav.Get(); ok {
		b.SetFav(v)
	}
	if v, ok := req.Did.Get(); ok {
		b.SetDid(v)
	}
	if v, ok := req.Delete.Get(); ok {
		b.SetDelete(v)
	}
	if v, ok := req.Handle.Get(); ok {
		b.SetHandle(v)
	}
	if v, ok := req.Token.Get(); ok {
		b.SetToken(v)
	}
	if v, ok := req.CreatedAt.Get(); ok {
		b.SetCreatedAt(v)
	}
	if v, ok := req.UpdatedAt.Get(); ok {
		b.SetUpdatedAt(v)
	}
	if v, ok := req.RaidAt.Get(); ok {
		b.SetRaidAt(v)
	}
	if v, ok := req.Luck.Get(); ok {
		b.SetLuck(v)
	}
	if v, ok := req.Aiten.Get(); ok {
		b.SetAiten(v)
	}
	if v, ok := req.LuckAt.Get(); ok {
		b.SetLuckAt(v)
	}
	if v, ok := req.Like.Get(); ok {
		b.SetLike(v)
	}
	if v, ok := req.LikeRank.Get(); ok {
		b.SetLikeRank(v)
	}
	if v, ok := req.LikeAt.Get(); ok {
		b.SetLikeAt(v)
	}
	if v, ok := req.Ten.Get(); ok {
		b.SetTen(v)
	}
	if v, ok := req.TenSu.Get(); ok {
		b.SetTenSu(v)
	}
	if v, ok := req.TenKai.Get(); ok {
		b.SetTenKai(v)
	}
	if v, ok := req.TenCard.Get(); ok {
		b.SetTenCard(v)
	}
	if v, ok := req.TenDelete.Get(); ok {
		b.SetTenDelete(v)
	}
	if v, ok := req.TenPost.Get(); ok {
		b.SetTenPost(v)
	}
	if v, ok := req.TenGet.Get(); ok {
		b.SetTenGet(v)
	}
	if v, ok := req.TenAt.Get(); ok {
		b.SetTenAt(v)
	}
	if v, ok := req.Next.Get(); ok {
		b.SetNext(v)
	}
	// Add all edges.
	b.AddCardIDs(req.Card...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewUserCreate(e), nil
}

// ReadUser handles GET /users/{id} requests.
func (h *OgentHandler) ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewUserRead(e), nil
}

// UpdateUser handles PATCH /users/{id} requests.
func (h *OgentHandler) UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error) {
	b := h.client.User.UpdateOneID(params.ID)

	if v, ok := req.Token.Get(); ok {
		if v == token {
			b.SetToken(v)

			if v, ok := req.Fav.Get(); ok {
				b.SetFav(v)
			}
			if v, ok := req.Did.Get(); ok {
				b.SetDid(v)
			}
			if v, ok := req.Delete.Get(); ok {
				b.SetDelete(v)
			}
			if v, ok := req.Handle.Get(); ok {
				b.SetHandle(v)
			}
			if v, ok := req.UpdatedAt.Get(); ok {
				b.SetUpdatedAt(v)
			}
			if v, ok := req.RaidAt.Get(); ok {
				b.SetRaidAt(v)
			}
			if v, ok := req.Aiten.Get(); ok {
				b.SetAiten(v)
			}
			if v, ok := req.Luck.Get(); ok {
				b.SetLuck(v)
			}
			if v, ok := req.LuckAt.Get(); ok {
				b.SetLuckAt(v)
			}
			if v, ok := req.Like.Get(); ok {
				b.SetLike(v)
			}
			if v, ok := req.LikeRank.Get(); ok {
				b.SetLikeRank(v)
			}
			if v, ok := req.LikeAt.Get(); ok {
				b.SetLikeAt(v)
			}
			if v, ok := req.Ten.Get(); ok {
				b.SetTen(v)
			}
			if v, ok := req.TenSu.Get(); ok {
				b.SetTenSu(v)
			}
			if v, ok := req.TenKai.Get(); ok {
				b.SetTenKai(v)
			}
			if v, ok := req.TenCard.Get(); ok {
				b.SetTenCard(v)
			}
			if v, ok := req.TenDelete.Get(); ok {
				b.SetTenDelete(v)
			}
			if v, ok := req.TenPost.Get(); ok {
				b.SetTenPost(v)
			}
			if v, ok := req.TenGet.Get(); ok {
				b.SetTenGet(v)
			}
			if v, ok := req.TenAt.Get(); ok {
				b.SetTenAt(v)
			}
			if v, ok := req.Next.Get(); ok {
				b.SetNext(v)
			}
			// Add all edges.
			if req.Card != nil {
				b.ClearCard().AddCardIDs(req.Card...)
			}
		}
	}

	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewUserUpdate(e), nil
}

// DeleteUser handles DELETE /users/{id} requests.
func (h *OgentHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
	err := h.client.User.DeleteOneID(0).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteUserNoContent), nil

}

// ListUser handles GET /users requests.
func (h *OgentHandler) ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error) {
	q := h.client.User.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewUserLists(es)
	return (*ListUserOKApplicationJSON)(&r), nil
}

// ListUserCard handles GET /users/{id}/card requests.
func (h *OgentHandler) ListUserCard(ctx context.Context, params ListUserCardParams) (ListUserCardRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID)).QueryCard()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewUserCardLists(es)
	return (*ListUserCardOKApplicationJSON)(&r), nil
}
