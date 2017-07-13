package routes

import (
	"github.com/devinmcgloin/fokal/pkg/create"
	"github.com/devinmcgloin/fokal/pkg/handler"
	"github.com/devinmcgloin/fokal/pkg/model"
	"github.com/devinmcgloin/fokal/pkg/security"
	"github.com/devinmcgloin/fokal/pkg/security/permissions"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func RegisterCreateRoutes(state *handler.State, api *mux.Router, chain alice.Chain) {
	post := api.Methods("POST").Subrouter()
	post.Handle("/i", chain.Append(handler.Middleware{
		State: state,
		M:     security.Authenticate}.Handler).Then(handler.Handler{State: state, H: create.ImageHandler}))
	post.Handle("/u", chain.Append(handler.Middleware{
		State: state,
		M:     security.Authenticate}.Handler).Then(handler.Handler{State: state, H: create.UserHandler}))

	put := api.Methods("PUT").Subrouter()
	put.Handle("/u/{ID}/avatar", chain.Append(handler.Middleware{
		State: state,
		M:     security.Authenticate,
	}.Handler,
		permissions.Middleware{State: state,
			T:          permissions.CanEdit,
			TargetType: model.Users,
			M:          permissions.PermissionMiddle,
		}.Handler).Then(handler.Handler{
		State: state,
		H:     create.AvatarHandler,
	}))
}