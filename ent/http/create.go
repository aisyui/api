// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strings"
	"t/ent"
	"t/ent/compartment"
	"t/ent/entry"
	"t/ent/fridge"
	"t/ent/item"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

// Create creates a new ent.Compartment and stores it in the database.
func (h CompartmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d CompartmentCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Compartment.Create()
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create compartment", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Compartment.Query().Where(compartment.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read compartment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("compartment rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment3324871446View(ret), w)
}

// Create creates a new ent.Entry and stores it in the database.
func (h EntryHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d EntryCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.User == nil {
		errs["user"] = `missing required field: "user"`
	} else if err := entry.UserValidator(*d.User); err != nil {
		errs["user"] = strings.TrimPrefix(err.Error(), "entry: ")
	}
	if d.First == nil {
		errs["first"] = `missing required field: "first"`
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.Entry.Create()
	if d.User != nil {
		b.SetUser(*d.User)
	}
	if d.First != nil {
		b.SetFirst(*d.First)
	}
	if d.CreatedAt != nil {
		b.SetCreatedAt(*d.CreatedAt)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create entry", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Entry.Query().Where(entry.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.String("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.String("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read entry", zap.Error(err), zap.String("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("entry rendered", zap.String("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewEntry2675665849View(ret), w)
}

// Create creates a new ent.Fridge and stores it in the database.
func (h FridgeHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d FridgeCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Fridge.Create()
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create fridge", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Fridge.Query().Where(fridge.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read fridge", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("fridge rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewFridge2211356377View(ret), w)
}

// Create creates a new ent.Item and stores it in the database.
func (h ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d ItemCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Item.Create()
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create item", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Item.Query().Where(item.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read item", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("item rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewItem1548468123View(ret), w)
}