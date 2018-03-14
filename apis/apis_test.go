package apis


type apiTestCase struct {
	tag      string
	method   string
	url      string
	body     string
	status   int
	response string
}


func newRouter() *mux.Router {
	logger := logrus.New()
	logger.Level = logrus.PanicLevel

	router := mux.NewRouter()

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		app.Transactional(testdata.DB),
	)

	return router
}


