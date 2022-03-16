package s3api

import (
	"github.com/filedag-project/filedag-storage/http/objectstore/consts"
	"github.com/filedag-project/filedag-storage/http/objectstore/iam"
	"github.com/filedag-project/filedag-storage/http/objectstore/iam/set"
	"github.com/filedag-project/filedag-storage/http/objectstore/response"
	"github.com/filedag-project/filedag-storage/http/objectstore/store"
	"github.com/gorilla/mux"
	"net/http"
)

type s3ApiServer struct {
	authSys iam.AuthSys
	store   store.StorageSys
}

//registerS3Router Register S3Router
func (s3a *s3ApiServer) registerS3Router(router *mux.Router) {
	// API Router
	apiRouter := router.PathPrefix("/").Subrouter()
	apiRouter.Methods(http.MethodPost).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		ctypeOk := set.MatchSimple("application/x-www-form-urlencoded*", r.Header.Get(consts.ContentType))
		authOk := set.MatchSimple(consts.SignV4Algorithm+"*", r.Header.Get(consts.Authorization))
		noQueries := len(r.URL.RawQuery) == 0
		return ctypeOk && authOk && noQueries
	}).HandlerFunc(s3a.AssumeRole)
	// Readiness Probe
	apiRouter.Methods(http.MethodGet).Path("/status").HandlerFunc(s3a.StatusHandler)
	// NotFound
	apiRouter.NotFoundHandler = http.HandlerFunc(response.NotFoundHandler)
	// ListBuckets
	apiRouter.Methods(http.MethodGet).Path("/").HandlerFunc(s3a.ListBucketsHandler)
	var routers []*mux.Router
	routers = append(routers, apiRouter.PathPrefix("/{bucket}").Subrouter())

	for _, bucket := range routers {
		// PutObject
		bucket.Methods(http.MethodPut).Path("/{object:.+}").HandlerFunc(s3a.PutObjectHandler)
		// PutBucketPolicy
		bucket.Methods(http.MethodPut).HandlerFunc(s3a.PutBucketPolicyHandler).Queries("policy", "")
		// DeleteBucketPolicy
		bucket.Methods(http.MethodDelete).HandlerFunc(s3a.DeleteBucketPolicyHandler).Queries("policy", "")
		// GetBucketPolicy
		bucket.Methods(http.MethodGet).HandlerFunc(s3a.GetBucketPolicyHandler).Queries("policy", "")
		// PutBucket
		bucket.Methods(http.MethodPut).HandlerFunc(s3a.PutBucketHandler)
		// HeadBucket
		bucket.Methods(http.MethodHead).HandlerFunc(s3a.HeadBucketHandler)
		// DeleteBucket
		bucket.Methods(http.MethodDelete).HandlerFunc(s3a.DeleteBucketHandler)

	}
}

//NewS3Server Start a S3Server
func NewS3Server(router *mux.Router) {
	var s3server s3ApiServer
	s3server.authSys.Init()
	s3server.store.Init()
	s3server.registerS3Router(router)
}
