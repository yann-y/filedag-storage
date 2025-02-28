/*
 * The following code tries to reverse engineer the Amazon S3 APIs,
 * and is mostly copied from minio implementation.
 */

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package iam

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/filedag-project/filedag-storage/objectservice/apierrors"
	"github.com/filedag-project/filedag-storage/objectservice/consts"
	"github.com/filedag-project/filedag-storage/objectservice/iam/auth"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Whitelist resource list that will be used in query string for signature-V2 calculation.
//
// This list should be kept alphabetically sorted, do not hastily edit.
var resourceList = []string{
	"acl",
	"cors",
	"delete",
	"encryption",
	"legal-hold",
	"lifecycle",
	"location",
	"logging",
	"notification",
	"partNumber",
	"policy",
	"requestPayment",
	"response-cache-control",
	"response-content-disposition",
	"response-content-encoding",
	"response-content-language",
	"response-content-type",
	"response-expires",
	"retention",
	"select",
	"select-type",
	"tagging",
	"torrent",
	"uploadId",
	"uploads",
	"versionId",
	"versioning",
	"versions",
	"website",
}

// Signature and API related constants.
const (
	signV2Algorithm = "AWS"
)

// AWS S3 Signature V2 calculation rule is give here:
// http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html#RESTAuthenticationStringToSign
func (s *AuthSys) doesPolicySignatureV2Match(formValues http.Header) (auth.Credentials, apierrors.ErrorCode) {
	accessKey := formValues.Get(consts.AmzAccessKeyID)

	r := &http.Request{Header: formValues}
	cred, _, s3Err := s.checkKeyValid(r, accessKey)
	if s3Err != apierrors.ErrNone {
		return cred, s3Err
	}
	policy := formValues.Get("Policy")
	signature := formValues.Get(consts.AmzSignatureV2)
	if !compareSignatureV2(signature, calculateSignatureV2(policy, cred.SecretKey)) {
		return cred, apierrors.ErrSignatureDoesNotMatch
	}
	return cred, apierrors.ErrNone
}

// Escape encodedQuery string into unescaped list of query params, returns error
// if any while unescaping the values.
func unescapeQueries(encodedQuery string) (unescapedQueries []string, err error) {
	for _, query := range strings.Split(encodedQuery, "&") {
		var unescapedQuery string
		unescapedQuery, err = url.QueryUnescape(query)
		if err != nil {
			return nil, err
		}
		unescapedQueries = append(unescapedQueries, unescapedQuery)
	}
	return unescapedQueries, nil
}

// doesPresignV2SignatureMatch - Verify query headers with presigned signature
//     - http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html#RESTAuthenticationQueryStringAuth
// returns apierrors.ErrNone if matches. S3 errors otherwise.
func (s *AuthSys) doesPresignV2SignatureMatch(r *http.Request) apierrors.ErrorCode {
	// r.RequestURI will have raw encoded URI as sent by the client.
	tokens := strings.SplitN(r.RequestURI, "?", 2)
	encodedResource := tokens[0]
	encodedQuery := ""
	if len(tokens) == 2 {
		encodedQuery = tokens[1]
	}

	var (
		filteredQueries []string
		gotSignature    string
		expires         string
		accessKey       string
		err             error
	)

	var unescapedQueries []string
	unescapedQueries, err = unescapeQueries(encodedQuery)
	if err != nil {
		return apierrors.ErrInvalidQueryParams
	}

	// Extract the necessary values from presigned query, construct a list of new filtered queries.
	for _, query := range unescapedQueries {
		keyval := strings.SplitN(query, "=", 2)
		if len(keyval) != 2 {
			return apierrors.ErrInvalidQueryParams
		}
		switch keyval[0] {
		case consts.AmzAccessKeyID:
			accessKey = keyval[1]
		case consts.AmzSignatureV2:
			gotSignature = keyval[1]
		case consts.Expires:
			expires = keyval[1]
		default:
			filteredQueries = append(filteredQueries, query)
		}
	}

	// Invalid values returns error.
	if accessKey == "" || gotSignature == "" || expires == "" {
		return apierrors.ErrInvalidQueryParams
	}

	cred, _, s3Err := s.checkKeyValid(r, accessKey)
	if s3Err != apierrors.ErrNone {
		return s3Err
	}

	// Make sure the request has not expired.
	expiresInt, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return apierrors.ErrAuthorizationHeaderMalformed
	}

	// Check if the presigned URL has expired.
	if expiresInt < time.Now().UTC().Unix() {
		return apierrors.ErrExpiredPresignRequest
	}

	encodedResource, err = getResource(encodedResource, r.Host)
	if err != nil {
		return apierrors.ErrInvalidRequest
	}

	expectedSignature := preSignatureV2(cred, r.Method, encodedResource, strings.Join(filteredQueries, "&"), r.Header, expires)
	if !compareSignatureV2(gotSignature, expectedSignature) {
		return apierrors.ErrSignatureDoesNotMatch
	}

	r.Form.Del(consts.Expires)

	return apierrors.ErrNone
}

func (s *AuthSys) getReqAccessKeyV2(r *http.Request) (auth.Credentials, bool, apierrors.ErrorCode) {
	if accessKey := r.Form.Get(consts.AmzAccessKeyID); accessKey != "" {
		return s.checkKeyValid(r, accessKey)
	}

	// below is V2 Signed Auth header format, splitting on `space` (after the `AWS` string).
	// Authorization = "AWS" + " " + AWSAccessKeyId + ":" + Signature
	authFields := strings.Split(r.Header.Get(consts.Authorization), " ")
	if len(authFields) != 2 {
		return auth.Credentials{}, false, apierrors.ErrMissingFields
	}

	// Then will be splitting on ":", this will seprate `AWSAccessKeyId` and `Signature` string.
	keySignFields := strings.Split(strings.TrimSpace(authFields[1]), ":")
	if len(keySignFields) != 2 {
		return auth.Credentials{}, false, apierrors.ErrMissingFields
	}

	return s.checkKeyValid(r, keySignFields[0])
}

// Authorization = "AWS" + " " + AWSAccessKeyId + ":" + Signature;
// Signature = Base64( HMAC-SHA1( YourSecretKey, UTF-8-Encoding-Of( StringToSign ) ) );
//
// StringToSign = HTTP-Verb + "\n" +
//  	Content-Md5 + "\n" +
//  	Content-Type + "\n" +
//  	Date + "\n" +
//  	CanonicalizedProtocolHeaders +
//  	CanonicalizedResource;
//
// CanonicalizedResource = [ consts.SlashSeparator + Bucket ] +
//  	<HTTP-Request-URI, from the protocol name up to the query string> +
//  	[ subresource, if present. For example "?acl", "?location", "?logging", or "?torrent"];
//
// CanonicalizedProtocolHeaders = <described below>

// doesSignV2Match - Verify authorization header with calculated header in accordance with
//     - http://docs.aws.amazon.com/AmazonS3/latest/dev/auth-request-sig-v2.html
// returns true if matches, false otherwise. if error is not nil then it is always false

func (s *AuthSys) validateV2AuthHeader(r *http.Request) (auth.Credentials, apierrors.ErrorCode) {
	var cred auth.Credentials
	v2Auth := r.Header.Get(consts.Authorization)
	if v2Auth == "" {
		return cred, apierrors.ErrAuthHeaderEmpty
	}

	// Verify if the header algorithm is supported or not.
	if !strings.HasPrefix(v2Auth, signV2Algorithm) {
		return cred, apierrors.ErrSignatureVersionNotSupported
	}

	cred, _, apiErr := s.getReqAccessKeyV2(r)
	if apiErr != apierrors.ErrNone {
		return cred, apiErr
	}

	return cred, apierrors.ErrNone
}

func (s *AuthSys) doesSignV2Match(r *http.Request) apierrors.ErrorCode {
	v2Auth := r.Header.Get(consts.Authorization)
	cred, apiError := s.validateV2AuthHeader(r)
	if apiError != apierrors.ErrNone {
		return apiError
	}

	// r.RequestURI will have raw encoded URI as sent by the client.
	tokens := strings.SplitN(r.RequestURI, "?", 2)
	encodedResource := tokens[0]
	encodedQuery := ""
	if len(tokens) == 2 {
		encodedQuery = tokens[1]
	}

	unescapedQueries, err := unescapeQueries(encodedQuery)
	if err != nil {
		return apierrors.ErrInvalidQueryParams
	}

	encodedResource, err = getResource(encodedResource, r.Host)
	if err != nil {
		return apierrors.ErrInvalidRequest
	}

	prefix := fmt.Sprintf("%s %s:", signV2Algorithm, cred.AccessKey)
	if !strings.HasPrefix(v2Auth, prefix) {
		return apierrors.ErrSignatureDoesNotMatch
	}
	v2Auth = v2Auth[len(prefix):]
	expectedAuth := signatureV2(cred, r.Method, encodedResource, strings.Join(unescapedQueries, "&"), r.Header)
	if !compareSignatureV2(v2Auth, expectedAuth) {
		return apierrors.ErrSignatureDoesNotMatch
	}
	return apierrors.ErrNone
}

func calculateSignatureV2(stringToSign string, secret string) string {
	hm := hmac.New(sha1.New, []byte(secret))
	hm.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}

// Return signature-v2 for the presigned request.
func preSignatureV2(cred auth.Credentials, method string, encodedResource string, encodedQuery string, headers http.Header, expires string) string {
	stringToSign := getStringToSignV2(method, encodedResource, encodedQuery, headers, expires)
	return calculateSignatureV2(stringToSign, cred.SecretKey)
}

// Return the signature v2 of a given request.
func signatureV2(cred auth.Credentials, method string, encodedResource string, encodedQuery string, headers http.Header) string {
	stringToSign := getStringToSignV2(method, encodedResource, encodedQuery, headers, "")
	signature := calculateSignatureV2(stringToSign, cred.SecretKey)
	return signature
}

// compareSignatureV2 returns true if and only if both signatures
// are equal. The signatures are expected to be base64 encoded strings
// according to the AWS S3 signature V2 spec.
func compareSignatureV2(sig1, sig2 string) bool {
	// Decode signature string to binary byte-sequence representation is required
	// as Base64 encoding of a value is not unique:
	// For example "aGVsbG8=" and "aGVsbG8=\r" will result in the same byte slice.
	signature1, err := base64.StdEncoding.DecodeString(sig1)
	if err != nil {
		return false
	}
	signature2, err := base64.StdEncoding.DecodeString(sig2)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(signature1, signature2) == 1
}

// Return canonical headers.
func canonicalizedAmzHeadersV2(headers http.Header) string {
	var keys []string
	keyval := make(map[string]string, len(headers))
	for key := range headers {
		lkey := strings.ToLower(key)
		if !strings.HasPrefix(lkey, "x-amz-") {
			continue
		}
		keys = append(keys, lkey)
		keyval[lkey] = strings.Join(headers[key], ",")
	}
	sort.Strings(keys)
	var canonicalHeaders []string
	for _, key := range keys {
		canonicalHeaders = append(canonicalHeaders, key+":"+keyval[key])
	}
	return strings.Join(canonicalHeaders, "\n")
}

// Return canonical resource string.
func canonicalizedResourceV2(encodedResource, encodedQuery string) string {
	queries := strings.Split(encodedQuery, "&")
	keyval := make(map[string]string)
	for _, query := range queries {
		key := query
		val := ""
		index := strings.Index(query, "=")
		if index != -1 {
			key = query[:index]
			val = query[index+1:]
		}
		keyval[key] = val
	}

	var canonicalQueries []string
	for _, key := range resourceList {
		val, ok := keyval[key]
		if !ok {
			continue
		}
		if val == "" {
			canonicalQueries = append(canonicalQueries, key)
			continue
		}
		canonicalQueries = append(canonicalQueries, key+"="+val)
	}

	// The queries will be already sorted as resourceList is sorted, if canonicalQueries
	// is empty strings.Join returns empty.
	canonicalQuery := strings.Join(canonicalQueries, "&")
	if canonicalQuery != "" {
		return encodedResource + "?" + canonicalQuery
	}
	return encodedResource
}

// Return string to sign under two different conditions.
// - if expires string is set then string to sign includes date instead of the Date header.
// - if expires string is empty then string to sign includes date header instead.
func getStringToSignV2(method string, encodedResource, encodedQuery string, headers http.Header, expires string) string {
	canonicalHeaders := canonicalizedAmzHeadersV2(headers)
	if len(canonicalHeaders) > 0 {
		canonicalHeaders += "\n"
	}

	date := expires // Date is set to expires date for presign operations.
	if date == "" {
		// If expires date is empty then request header Date is used.
		date = headers.Get(consts.Date)
	}

	// From the Amazon docs:
	//
	// StringToSign = HTTP-Verb + "\n" +
	// 	 Content-Md5 + "\n" +
	//	 Content-Type + "\n" +
	//	 Date/Expires + "\n" +
	//	 CanonicalizedProtocolHeaders +
	//	 CanonicalizedResource;
	stringToSign := strings.Join([]string{
		method,
		headers.Get(consts.ContentMD5),
		headers.Get(consts.ContentType),
		date,
		canonicalHeaders,
	}, "\n")

	return stringToSign + canonicalizedResourceV2(encodedResource, encodedQuery)
}

// Returns "/bucketName/objectName" for path-style or virtual-host-style requests.
func getResource(path string, host string) (string, error) {

	// If virtual-host-style is enabled construct the "resource" properly.
	if strings.Contains(host, ":") {
		// In bucket.mydomain.com:9000, strip out :9000
		var err error
		if host, _, err = net.SplitHostPort(host); err != nil {
			return "", err
		}
	}
	return path, nil
}
