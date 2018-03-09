package app

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() string
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id string)
	// RequestID returns the ID of the current request
	RequestID() string
	// DB returns the currently active database transaction that can be used for DB query purpose
	DB() *mgo.Database
	// DB returns the currently active database transaction that can be used for DB query purpose
	SetDB()
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
	// Set All params
	SetParams(params map[string]string)
	// Get All Params
	GetParams() map[string]string
	// SetBody
	SetBody(body []byte)
	// GetBody
	GetBody() []byte
}

type requestScope struct {
	Logger                      // the logger tagged with the current request information
	now       time.Time         // the time when the request is being processed
	requestID string            // an ID identifying one or multiple correlated HTTP requests
	userID    string            // an ID identifying the current user
	db        *mgo.Database     // Database object
	params    map[string]string // url Params Params
	body      []byte            // Body
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("UserID", id)
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func (rs *requestScope) DB() *mgo.Database {
	return rs.db
}

func (rs *requestScope) SetDB() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	rs.db = session.DB("myDb")
}

func (rs *requestScope) SetParams(params map[string]string) {
	rs.params = params
}

func (rs *requestScope) GetParams() map[string]string {
	return rs.params
}

func (rs *requestScope) SetBody(body []byte) {
	rs.body = body
}

func (rs *requestScope) GetBody() []byte {
	return rs.body
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}
