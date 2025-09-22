package webapp

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/innoai-tech/infra/pkg/http/webapp"
	"github.com/innoai-tech/infra/pkg/http/webapp/appconfig"
	"github.com/innoai-tech/postgres-operator/internal/version"
	"github.com/octohelm/x/ptr"
)

//go:embed public
var content embed.FS

type WebUI struct {
	// [前端] api 请求调用需要跟随 bash href
	AllApiPrefixWithBaseHref *bool `flag:",omitzero"`

	h       http.Handler
	once    sync.Once
	appname string
}

func (w *WebUI) Use(name string) {
	w.appname = name
}

func (w *WebUI) SetDefaults() {
	if w.AllApiPrefixWithBaseHref == nil {
		w.AllApiPrefixWithBaseHref = ptr.Ptr(true)
	}
}

func (w *WebUI) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	w.once.Do(func() {
		appConfig := appconfig.AppConfig{}
		appConfig.LoadFromEnviron(os.Environ())

		appBaseHref := "/"

		appConfig["APP_BASE_HREF"] = appBaseHref

		if *w.AllApiPrefixWithBaseHref {
			appConfig["ALL_API_PREFIX_WITH_BASE_HREF"] = "enabled"
		}

		root, _ := fs.Sub(content, path.Join("public", w.appname))

		w.h = webapp.ServeFS(
			root,

			webapp.WithBaseHref(appBaseHref),
			webapp.WithAppConfig(appConfig),
			webapp.WithAppVersion(version.Version()),
		)
	})

	w.h.ServeHTTP(rw, r)
}
