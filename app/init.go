package app

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Sirupsen/logrus"

	"github.com/gorilla/context"
)

// Init
// func Init(logger *logrus.Logger) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		now := time.Now()

// 		rc.Response = &access.LogResponseWriter{rc.Response, http.StatusOK, 0}

// 		ac := newRequestScope(now, logger, rc.Request)
// 		rc.Set("Context", ac)

// 		fault.Recovery(ac.Errorf, convertError)(rc)
// 		logAccess(rc, ac.Infof, ac.Now())
// 	})
// }

func Init(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		logger := logrus.New()
		ac := newRequestScope(now, logger, r)
		ac.SetDB()
		b, _ := ioutil.ReadAll(r.Body)
		ac.SetBody(b)
		context.Set(r, "Context", ac)
		ac.SetParams(mux.Vars(r))
		defer ac.DB().Session.Close()
		next.ServeHTTP(w, r)
		return
	})
}

// GetRequestScope returns the RequestScope of the current request.
func GetRequestScope(r *http.Request) RequestScope {
	return context.Get(r, "Context").(RequestScope)
}
