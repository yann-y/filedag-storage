package s3api

import (
	"encoding/json"
	"fmt"
	"github.com/filedag-project/filedag-storage/http/objectstore/iam/policy"
	"github.com/filedag-project/filedag-storage/http/objectstore/iamapi"
	"github.com/filedag-project/filedag-storage/http/objectstore/response"
	"github.com/filedag-project/filedag-storage/http/objectstore/uleveldb"
	"github.com/filedag-project/filedag-storage/http/objectstore/utils"
	"github.com/filedag-project/filedag-storage/http/objectstore/utils/testsign"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"
)

var w *httptest.ResponseRecorder
var router = mux.NewRouter()

func TestMain(m *testing.M) {
	dir, err := ioutil.ReadDir("./test")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"./test", d.Name()}...))
	}
	uleveldb.DBClient, err = uleveldb.OpenDb("./test")
	if err != nil {
		return
	}
	defer uleveldb.DBClient.Close()
	iamapi.NewIamApiServer(router)
	NewS3Server(router)
	os.Exit(m.Run())
}
func reqTest(r *http.Request) *httptest.ResponseRecorder {
	// mock a response logger
	w = httptest.NewRecorder()
	// Let the server process the mock request and record the returned response content
	router.ServeHTTP(w, r)
	return w
}
func TestS3ApiServer_BucketHandler(t *testing.T) {
	bucketName := "/testbucket"
	// test cases with inputs and expected result for Bucket.
	testCases := []struct {
		bucketName string
		accessKey  string
		secretKey  string
		// expected output.
		expectedRespStatus int // expected response status body.
	}{
		// Test case - 1.
		// Fetching the entire Bucket and validating its contents.
		{
			bucketName:         bucketName,
			accessKey:          DefaultTestAccessKey,
			secretKey:          DefaultTestSecretKey,
			expectedRespStatus: http.StatusOK,
		},
		// Test case - 2.
		// wrong accessKey.
		{
			bucketName:         bucketName,
			accessKey:          "1",
			secretKey:          "1",
			expectedRespStatus: http.StatusForbidden,
		},
	}
	// Iterating over the cases, fetching the object validating the response.
	for i, testCase := range testCases {
		// mock an HTTP request
		reqPutBucket := testsign.MustNewSignedV4Request(http.MethodPut, "/testbucket", 0, nil, "s3", testCase.accessKey, testCase.secretKey, t)
		result1 := reqTest(reqPutBucket)
		if result1.Code != testCase.expectedRespStatus {
			t.Fatalf("Case %d: Expected the response status to be `%d`, but instead found `%d`", i+1, testCase.expectedRespStatus, result1.Code)
		}

		reqHeadBucket := testsign.MustNewSignedV4Request(http.MethodHead, "/testbucket", 0, nil, "s3", testCase.accessKey, testCase.secretKey, t)
		result2 := reqTest(reqHeadBucket)
		if result2.Code != testCase.expectedRespStatus {
			t.Fatalf("Case %d: Expected the response status to be `%d`, but instead found `%d`", i+1, testCase.expectedRespStatus, result2.Code)
		}

		reqListBucket := testsign.MustNewSignedV4Request(http.MethodGet, "/", 0, nil, "s3", testCase.accessKey, testCase.secretKey, t)
		result3 := reqTest(reqListBucket)
		if result3.Code != testCase.expectedRespStatus {
			t.Fatalf("Case %d: Expected the response status to be `%d`, but instead found `%d`", i+1, testCase.expectedRespStatus, result3.Code)
		}
		var resp1 response.ListAllMyBucketsResult
		utils.XmlDecoder(result3.Body, &resp1, reqListBucket.ContentLength)
		fmt.Println("list:", resp1)

		reqDeleteBucket := testsign.MustNewSignedV4Request(http.MethodDelete, "/testbucket", 0,
			nil, "s3", testCase.accessKey, testCase.secretKey, t)
		result4 := reqTest(reqDeleteBucket)
		if result4.Code != testCase.expectedRespStatus {
			t.Fatalf("Case %d: Expected the response status to be `%d`, but instead found `%d`", i+1, testCase.expectedRespStatus, result4.Code)
		}

		resp2 := response.ListAllMyBucketsResult{}
		utils.XmlDecoder(reqTest(reqListBucket).Body, &resp2, reqListBucket.ContentLength)
		fmt.Println("list:", resp2)
	}

}

func TestS3ApiServer_BucketPolicyHandler(t *testing.T) {
	u := "/testbucket"
	reqPutBucket := testsign.MustNewSignedV4Request(http.MethodPut, u, 0, nil, "s3", DefaultTestAccessKey, DefaultTestSecretKey, t)
	fmt.Println("putbucket:", reqTest(reqPutBucket).Body.String())

	p := `{"Version":"2008-10-17","Id":"aaaa-bbbb-cccc-dddd","Statement":[{"Effect":"Allow","Sid":"1","Principal":{"AWS":["111122223333","444455556666"]},"Action":["s3:*"],"Resource":"arn:aws:s3:::testbucket/*"}]}`
	reqPut := testsign.MustNewSignedV4Request(http.MethodPut, u+"?policy", int64(len(p)), strings.NewReader(p),
		"s3", DefaultTestAccessKey, DefaultTestSecretKey, t)
	fmt.Println("put:", reqTest(reqPut).Body.String())

	reqGet := testsign.MustNewSignedV4Request(http.MethodGet, u+"?policy", 0, nil, "s3",
		DefaultTestAccessKey, DefaultTestSecretKey, t)
	resp1 := policy.Policy{}
	json.Unmarshal([]byte(reqTest(reqGet).Body.String()), &resp1)
	fmt.Println("get:", resp1)

	reqDel := testsign.MustNewSignedV4Request(http.MethodDelete, u+"?policy", 0, nil, "s3", DefaultTestAccessKey, DefaultTestSecretKey, t)
	fmt.Println("del:", reqTest(reqDel).Body.String())
}

//func TestS3ApiServer_GetBucketLocationHandler(t *testing.T) {
//	u := "http://127.0.0.1:9985/test22"
//	//req.Header.Set("Content-Type", "text/plain")
//	req := testsign.MustNewSignedV4Request(http.MethodHead, u+"?location", 0, nil, "s3", DefaultTestAccessKey, DefaultTestSecretKey, t)
//	client := &http.Client{}
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//	body, err := ioutil.ReadAll(res.Body)
//
//	fmt.Println(res)
//	fmt.Println(string(body))
//}
//
//func TestS3ApiServer_GetBucketAclHandler(t *testing.T) {
//	u := "http://127.0.0.1:9985/test"
//	req := testsign.MustNewSignedV4Request(http.MethodGet, u+"?acl=", 0, nil, "s3", DefaultTestAccessKey, DefaultTestSecretKey, t)
//	//req.Header.Set("Content-Type", "text/plain")
//	client := &http.Client{}
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//	body, err := ioutil.ReadAll(res.Body)
//
//	fmt.Println(res)
//	fmt.Println(string(body))
//}
//func TestS3ApiServer_PutBucketAclHandler(t *testing.T) {
//	u := "http://127.0.0.1:9985/test"
//	a := `<?xml version="1.0" encoding="UTF-8"?>
//<AccessControlPolicy xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
//  <Owner>
//    <ID>*** Owner-Canonical-User-ID ***</ID>
//    <DisplayName>owner-display-name</DisplayName>
//  </Owner>
//  <AccessControlList>
//    <Grant>
//      <Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//               xsi:type="Canonical User">
//        <ID>*** Owner-Canonical-User-ID ***</ID>
//        <DisplayName>display-name</DisplayName>
//      </Grantee>
//      <Permission>FULL_CONTROL</Permission>
//    </Grant>
//  </AccessControlList>
//</AccessControlPolicy>`
//	req := testsign.MustNewSignedV4Request(http.MethodPut, u+"?acl=", int64(len(a)), strings.NewReader(a), "s3", DefaultTestAccessKey, DefaultTestSecretKey, t)
//	//req.Header.Set("Content-Type", "text/plain")
//	client := &http.Client{}
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//	body, err := ioutil.ReadAll(res.Body)
//
//	fmt.Println(res)
//	fmt.Println(string(body))
//}
