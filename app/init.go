package app

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
