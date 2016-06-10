package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devinmcgloin/morph/src/api/SQL"
	"github.com/devinmcgloin/morph/src/api/auth"
	"github.com/devinmcgloin/morph/src/api/endpoint"
	"github.com/devinmcgloin/morph/src/views/editView"
	"github.com/devinmcgloin/morph/src/views/publicView"
	"github.com/devinmcgloin/morph/src/views/settings"
	"github.com/julienschmidt/httprouter"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func main() {

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "/auth/github/callback"),
	)

	flag := log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	log.SetFlags(flag)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port must be set")
	}

	router := httprouter.New()

	log.Printf("Serving at http://localhost:%s", port)

	// API POST ROUTES
	//TODO need to figure out api formatting and tokens.
	//TODO maybe these are the best routes for posting changes.

	router.POST("/api/v0/upload", endpoint.UploadHandler)
	router.POST("/api/v0/users/:user", endpoint.UserHandler)
	router.POST("/api/v0/photos/:p_id", endpoint.ImageHandler)

	//TODO these really should be done based on a semantic thing, not ID.
	//TODO consider phasing out numerical id's entirely for hex strings.

	// CONTENT VIEW ROUTES
	router.GET("/", publicView.MostRecentView)
	router.GET("/i/:IID", publicView.FeatureImgView)
	router.GET("/tag/:tag", publicView.CollectionTagView)
	router.GET("/tag/:tag/:IID", publicView.CollectionTagFeatureView)
	router.GET("/album/:AID", publicView.AlbumView)
	router.GET("/u/:UID", publicView.UserProfileView)
	router.GET("/loc/:LID", publicView.LocationView)
	router.GET("/search/*query", publicView.SearchView)

	// CONTENT EDIT ROUTES
	router.GET("/i/:IID/edit", editView.FeatureImgEditView)
	router.GET("/album/:AID/edit", editView.AlbumEditView)
	router.GET("/u/:UID/edit", editView.UserProfileEditView)
	router.GET("/upload", editView.UploadView)

	//TODO need to setup similar content post routes for sending back edits.

	// BACKEND MANAGE ROUTES
	router.GET("/login", publicView.UserLoginView)
	router.GET("/auth/:provider", auth.BeginAuthHandler)
	router.GET("/auth/:provider/callback", auth.UserLoginCallback)
	router.GET("/settings", settings.UserSettingsView)

	// ASSETS
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	//FIXME Temp to test locally.
	router.ServeFiles("/content/*filepath", http.Dir("content/"))

	SQL.SetDB()

	log.Fatal(http.ListenAndServe(":"+port, router))

}
